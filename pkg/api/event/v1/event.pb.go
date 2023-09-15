// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: event/v1/event.proto

package eventv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Action int32

const (
	Action_ACTION_UNSPECIFIED Action = 0
	Action_ACTION_PUT         Action = 1
	Action_ACTION_DELETE      Action = 2
	Action_ACTION_START       Action = 3
	Action_ACTION_STOP        Action = 4
)

// Enum value maps for Action.
var (
	Action_name = map[int32]string{
		0: "ACTION_UNSPECIFIED",
		1: "ACTION_PUT",
		2: "ACTION_DELETE",
		3: "ACTION_START",
		4: "ACTION_STOP",
	}
	Action_value = map[string]int32{
		"ACTION_UNSPECIFIED": 0,
		"ACTION_PUT":         1,
		"ACTION_DELETE":      2,
		"ACTION_START":       3,
		"ACTION_STOP":        4,
	}
)

func (x Action) Enum() *Action {
	p := new(Action)
	*p = x
	return p
}

func (x Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Action) Descriptor() protoreflect.EnumDescriptor {
	return file_event_v1_event_proto_enumTypes[0].Descriptor()
}

func (Action) Type() protoreflect.EnumType {
	return &file_event_v1_event_proto_enumTypes[0]
}

func (x Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Action.Descriptor instead.
func (Action) EnumDescriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{0}
}

// EventPayload defines an event payload.
type EventPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// repository is required for all actions
	Repository string `protobuf:"bytes,1,opt,name=repository,proto3" json:"repository,omitempty"`
	Digest     string `protobuf:"bytes,2,opt,name=digest,proto3" json:"digest,omitempty"`
	Mediatype  string `protobuf:"bytes,3,opt,name=mediatype,proto3" json:"mediatype,omitempty"`
	Payload    []byte `protobuf:"bytes,4,opt,name=payload,proto3" json:"payload,omitempty"`
	Action     Action `protobuf:"varint,5,opt,name=action,proto3,enum=beskar.api.event.v1.Action" json:"action,omitempty"`
}

func (x *EventPayload) Reset() {
	*x = EventPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_v1_event_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventPayload) ProtoMessage() {}

func (x *EventPayload) ProtoReflect() protoreflect.Message {
	mi := &file_event_v1_event_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventPayload.ProtoReflect.Descriptor instead.
func (*EventPayload) Descriptor() ([]byte, []int) {
	return file_event_v1_event_proto_rawDescGZIP(), []int{0}
}

func (x *EventPayload) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

func (x *EventPayload) GetDigest() string {
	if x != nil {
		return x.Digest
	}
	return ""
}

func (x *EventPayload) GetMediatype() string {
	if x != nil {
		return x.Mediatype
	}
	return ""
}

func (x *EventPayload) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *EventPayload) GetAction() Action {
	if x != nil {
		return x.Action
	}
	return Action_ACTION_UNSPECIFIED
}

var File_event_v1_event_proto protoreflect.FileDescriptor

var file_event_v1_event_proto_rawDesc = []byte{
	0x0a, 0x14, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x62, 0x65, 0x73, 0x6b, 0x61, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x22, 0xb3, 0x01, 0x0a, 0x0c,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x1e, 0x0a, 0x0a,
	0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x69,
	0x67, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x79, 0x70,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6d, 0x65, 0x64, 0x69, 0x61, 0x74, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x33, 0x0a, 0x06,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x62,
	0x65, 0x73, 0x6b, 0x61, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2a, 0x66, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x12, 0x41,
	0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x50, 0x55,
	0x54, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x44, 0x45,
	0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x12, 0x10, 0x0a, 0x0c, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e,
	0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x03, 0x12, 0x0f, 0x0a, 0x0b, 0x41, 0x43, 0x54, 0x49,
	0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x4f, 0x50, 0x10, 0x04, 0x42, 0x2c, 0x5a, 0x2a, 0x67, 0x6f, 0x2e,
	0x63, 0x69, 0x71, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x62, 0x65, 0x73, 0x6b, 0x61, 0x72, 0x2f, 0x70,
	0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x3b,
	0x65, 0x76, 0x65, 0x6e, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_v1_event_proto_rawDescOnce sync.Once
	file_event_v1_event_proto_rawDescData = file_event_v1_event_proto_rawDesc
)

func file_event_v1_event_proto_rawDescGZIP() []byte {
	file_event_v1_event_proto_rawDescOnce.Do(func() {
		file_event_v1_event_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_v1_event_proto_rawDescData)
	})
	return file_event_v1_event_proto_rawDescData
}

var file_event_v1_event_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_event_v1_event_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_event_v1_event_proto_goTypes = []interface{}{
	(Action)(0),          // 0: beskar.api.event.v1.Action
	(*EventPayload)(nil), // 1: beskar.api.event.v1.EventPayload
}
var file_event_v1_event_proto_depIdxs = []int32{
	0, // 0: beskar.api.event.v1.EventPayload.action:type_name -> beskar.api.event.v1.Action
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_event_v1_event_proto_init() }
func file_event_v1_event_proto_init() {
	if File_event_v1_event_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_v1_event_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventPayload); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_event_v1_event_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_v1_event_proto_goTypes,
		DependencyIndexes: file_event_v1_event_proto_depIdxs,
		EnumInfos:         file_event_v1_event_proto_enumTypes,
		MessageInfos:      file_event_v1_event_proto_msgTypes,
	}.Build()
	File_event_v1_event_proto = out.File
	file_event_v1_event_proto_rawDesc = nil
	file_event_v1_event_proto_goTypes = nil
	file_event_v1_event_proto_depIdxs = nil
}
