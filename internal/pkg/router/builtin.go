// SPDX-FileCopyrightText: Copyright (c) 2023, CIQ, Inc. All rights reserved
// SPDX-License-Identifier: Apache-2.0

package router

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/distribution/distribution/v3"
	"github.com/distribution/distribution/v3/reference"
	v1 "github.com/google/go-containerregistry/pkg/v1"
	regtypes "github.com/google/go-containerregistry/pkg/v1/types"
	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/types"
)

var funcContextKey uint8

type funcContext struct {
	req        *http.Request
	registry   distribution.Namespace
	builtinErr error
}

func newBuffer() interface{} {
	buffer := make([]byte, 8192)
	return &buffer
}

var bufferPool = sync.Pool{
	New: newBuffer,
}

type bufPoolCloser struct {
	io.Reader
	buf *[]byte
}

func (bpc bufPoolCloser) Close() error {
	bufferPool.Put(bpc.buf)
	return nil
}

var ociBlobDigestBuiltin = rego.Function2(
	&rego.Function{
		Name:             "oci.blob_digest",
		Decl:             types.NewFunction(types.Args(types.S, types.S), types.S),
		Nondeterministic: true,
	},
	func(bctx rego.BuiltinContext, a, b *ast.Term) (term *ast.Term, errFn error) {
		funcContext, ok := bctx.Context.Value(&funcContextKey).(*funcContext)
		if !ok {
			bctx.Cancel.Cancel()
			return nil, fmt.Errorf("bad context")
		}

		defer func() {
			if errFn != nil {
				funcContext.builtinErr = fmt.Errorf("%s builtin eval oci.blob_digest error: %w", bctx.Location, errFn)
				bctx.Cancel.Cancel()
			}
		}()

		astRef, ok := a.Value.(ast.String)
		if !ok {
			return nil, fmt.Errorf("oci reference is not a string")
		}
		astMediaType, ok := b.Value.(ast.String)
		if !ok {
			return nil, fmt.Errorf("oci layer mediatype is not a string")
		}

		ref := string(astRef)

		tagIndex := strings.LastIndexByte(ref, ':')
		if tagIndex < 0 {
			return nil, fmt.Errorf("reference without tag")
		}
		namedRef, err := reference.WithName(ref[:tagIndex])
		if err != nil {
			return nil, fmt.Errorf("bad reference name")
		}

		repository, err := funcContext.registry.Repository(bctx.Context, namedRef)
		if err != nil {
			return nil, fmt.Errorf("while getting repository %s: %w", namedRef, err)
		}
		tagDesc, err := repository.Tags(bctx.Context).Get(bctx.Context, ref[tagIndex+1:])
		if err != nil {
			var tagUnknown distribution.ErrTagUnknown
			if errors.As(err, &tagUnknown) {
				return ast.StringTerm(""), nil
			}
			return nil, fmt.Errorf("while getting tag %s: %w", ref[tagIndex+1:], err)
		}
		manifestService, err := repository.Manifests(bctx.Context)
		if err != nil {
			return nil, fmt.Errorf("while getting manifest service for %s: %w", namedRef, err)
		}
		registryManifest, err := manifestService.Get(bctx.Context, tagDesc.Digest)
		if err != nil {
			return nil, fmt.Errorf("while getting manifest for %s: %w", namedRef, err)
		}
		_, manifestPayload, err := registryManifest.Payload()
		if err != nil {
			return nil, err
		}
		manifest := new(v1.Manifest)
		if err := json.Unmarshal(manifestPayload, manifest); err != nil {
			return nil, err
		}

		mediaType := regtypes.MediaType(astMediaType)
		for _, layer := range manifest.Layers {
			if layer.MediaType != mediaType {
				continue
			}
			return ast.StringTerm(layer.Digest.Hex), nil
		}

		return ast.StringTerm(""), nil
	},
)

var requestBodyBuiltin = rego.FunctionDyn(
	&rego.Function{
		Name:             "request.body",
		Decl:             types.NewFunction(types.Args(), types.A),
		Nondeterministic: true,
	},
	func(bctx rego.BuiltinContext, _ []*ast.Term) (_ *ast.Term, errFn error) {
		funcContext, ok := bctx.Context.Value(&funcContextKey).(*funcContext)
		if !ok {
			bctx.Cancel.Cancel()
			return nil, fmt.Errorf("bad context")
		}

		defer func() {
			if errFn != nil {
				funcContext.builtinErr = fmt.Errorf("%s builtin eval request.body error: %w", bctx.Location, errFn)
				bctx.Cancel.Cancel()
			}
		}()

		if funcContext.req.Body != nil && funcContext.req.Body != http.NoBody {
			buf := bufferPool.Get().(*[]byte)

			n, err := io.ReadAtLeast(funcContext.req.Body, *buf, 1)
			if err != nil {
				return nil, fmt.Errorf("empty body request")
			}

			body := *buf
			bodyReader := bytes.NewReader(body[:n])

			v, err := ast.ValueFromReader(bodyReader)
			if err != nil {
				return nil, err
			}

			_, _ = bodyReader.Seek(0, io.SeekStart)

			funcContext.req.Body = &bufPoolCloser{bodyReader, buf}

			return ast.NewTerm(v), nil
		}

		v, err := ast.InterfaceToValue(nil)

		return ast.NewTerm(v), err
	},
)
