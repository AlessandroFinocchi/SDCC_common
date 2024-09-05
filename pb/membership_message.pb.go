// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.20.3
// source: membership_message.proto

package pb

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

type MembershipMsgType int32

const (
	MembershipMsgType_REPLY   MembershipMsgType = 0
	MembershipMsgType_REQUEST MembershipMsgType = 1
)

// Enum value maps for MembershipMsgType.
var (
	MembershipMsgType_name = map[int32]string{
		0: "REPLY",
		1: "REQUEST",
	}
	MembershipMsgType_value = map[string]int32{
		"REPLY":   0,
		"REQUEST": 1,
	}
)

func (x MembershipMsgType) Enum() *MembershipMsgType {
	p := new(MembershipMsgType)
	*p = x
	return p
}

func (x MembershipMsgType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MembershipMsgType) Descriptor() protoreflect.EnumDescriptor {
	return file_membership_message_proto_enumTypes[0].Descriptor()
}

func (MembershipMsgType) Type() protoreflect.EnumType {
	return &file_membership_message_proto_enumTypes[0]
}

func (x MembershipMsgType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MembershipMsgType.Descriptor instead.
func (MembershipMsgType) EnumDescriptor() ([]byte, []int) {
	return file_membership_message_proto_rawDescGZIP(), []int{0}
}

type MembershipRequestMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes  []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	Source *Node   `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *MembershipRequestMessage) Reset() {
	*x = MembershipRequestMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_membership_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MembershipRequestMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembershipRequestMessage) ProtoMessage() {}

func (x *MembershipRequestMessage) ProtoReflect() protoreflect.Message {
	mi := &file_membership_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MembershipRequestMessage.ProtoReflect.Descriptor instead.
func (*MembershipRequestMessage) Descriptor() ([]byte, []int) {
	return file_membership_message_proto_rawDescGZIP(), []int{0}
}

func (x *MembershipRequestMessage) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

func (x *MembershipRequestMessage) GetSource() *Node {
	if x != nil {
		return x.Source
	}
	return nil
}

type MembershipReplyMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nodes []*Node `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
}

func (x *MembershipReplyMessage) Reset() {
	*x = MembershipReplyMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_membership_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MembershipReplyMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MembershipReplyMessage) ProtoMessage() {}

func (x *MembershipReplyMessage) ProtoReflect() protoreflect.Message {
	mi := &file_membership_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MembershipReplyMessage.ProtoReflect.Descriptor instead.
func (*MembershipReplyMessage) Descriptor() ([]byte, []int) {
	return file_membership_message_proto_rawDescGZIP(), []int{1}
}

func (x *MembershipReplyMessage) GetNodes() []*Node {
	if x != nil {
		return x.Nodes
	}
	return nil
}

var File_membership_message_proto protoreflect.FileDescriptor

var file_membership_message_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x5f, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x12, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x18, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73,
	0x68, 0x69, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x12, 0x21, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52, 0x05, 0x6e,
	0x6f, 0x64, 0x65, 0x73, 0x12, 0x23, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64,
	0x65, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x22, 0x3b, 0x0a, 0x16, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x73, 0x68, 0x69, 0x70, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x21, 0x0a, 0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x52,
	0x05, 0x6e, 0x6f, 0x64, 0x65, 0x73, 0x2a, 0x2b, 0x0a, 0x11, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x68, 0x69, 0x70, 0x4d, 0x73, 0x67, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x52,
	0x45, 0x50, 0x4c, 0x59, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53,
	0x54, 0x10, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_membership_message_proto_rawDescOnce sync.Once
	file_membership_message_proto_rawDescData = file_membership_message_proto_rawDesc
)

func file_membership_message_proto_rawDescGZIP() []byte {
	file_membership_message_proto_rawDescOnce.Do(func() {
		file_membership_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_membership_message_proto_rawDescData)
	})
	return file_membership_message_proto_rawDescData
}

var file_membership_message_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_membership_message_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_membership_message_proto_goTypes = []any{
	(MembershipMsgType)(0),           // 0: proto.MembershipMsgType
	(*MembershipRequestMessage)(nil), // 1: proto.MembershipRequestMessage
	(*MembershipReplyMessage)(nil),   // 2: proto.MembershipReplyMessage
	(*Node)(nil),                     // 3: proto.Node
}
var file_membership_message_proto_depIdxs = []int32{
	3, // 0: proto.MembershipRequestMessage.nodes:type_name -> proto.Node
	3, // 1: proto.MembershipRequestMessage.source:type_name -> proto.Node
	3, // 2: proto.MembershipReplyMessage.nodes:type_name -> proto.Node
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_membership_message_proto_init() }
func file_membership_message_proto_init() {
	if File_membership_message_proto != nil {
		return
	}
	file_node_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_membership_message_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*MembershipRequestMessage); i {
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
		file_membership_message_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*MembershipReplyMessage); i {
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
			RawDescriptor: file_membership_message_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_membership_message_proto_goTypes,
		DependencyIndexes: file_membership_message_proto_depIdxs,
		EnumInfos:         file_membership_message_proto_enumTypes,
		MessageInfos:      file_membership_message_proto_msgTypes,
	}.Build()
	File_membership_message_proto = out.File
	file_membership_message_proto_rawDesc = nil
	file_membership_message_proto_goTypes = nil
	file_membership_message_proto_depIdxs = nil
}
