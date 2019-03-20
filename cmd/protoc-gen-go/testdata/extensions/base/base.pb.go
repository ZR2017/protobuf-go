// Code generated by protoc-gen-go. DO NOT EDIT.
// source: extensions/base/base.proto

package base

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	reflect "reflect"
)

type BaseMessage struct {
	Field                        *string  `protobuf:"bytes,1,opt,name=field" json:"field,omitempty"`
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *BaseMessage) ProtoReflect() protoreflect.Message {
	return xxx_File_extensions_base_base_proto_messageTypes[0].MessageOf(m)
}
func (m *BaseMessage) Reset()         { *m = BaseMessage{} }
func (m *BaseMessage) String() string { return proto.CompactTextString(m) }
func (*BaseMessage) ProtoMessage()    {}

// Deprecated: Use BaseMessage.ProtoReflect.Type instead.
func (*BaseMessage) Descriptor() ([]byte, []int) {
	return xxx_File_extensions_base_base_proto_rawdesc_gzipped, []int{0}
}

var extRange_BaseMessage = []proto.ExtensionRange{
	{Start: 4, End: 9},
	{Start: 16, End: 536870911},
}

// Deprecated: Use BaseMessage.ProtoReflect.Type.ExtensionRanges instead.
func (*BaseMessage) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_BaseMessage
}

func (m *BaseMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseMessage.Unmarshal(m, b)
}
func (m *BaseMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseMessage.Marshal(b, m, deterministic)
}
func (m *BaseMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseMessage.Merge(m, src)
}
func (m *BaseMessage) XXX_Size() int {
	return xxx_messageInfo_BaseMessage.Size(m)
}
func (m *BaseMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseMessage.DiscardUnknown(m)
}

var xxx_messageInfo_BaseMessage proto.InternalMessageInfo

func (m *BaseMessage) GetField() string {
	if m != nil && m.Field != nil {
		return *m.Field
	}
	return ""
}

type MessageSetWireFormatMessage struct {
	XXX_NoUnkeyedLiteral         struct{} `json:"-"`
	proto.XXX_InternalExtensions `protobuf_messageset:"1" json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *MessageSetWireFormatMessage) ProtoReflect() protoreflect.Message {
	return xxx_File_extensions_base_base_proto_messageTypes[1].MessageOf(m)
}
func (m *MessageSetWireFormatMessage) Reset()         { *m = MessageSetWireFormatMessage{} }
func (m *MessageSetWireFormatMessage) String() string { return proto.CompactTextString(m) }
func (*MessageSetWireFormatMessage) ProtoMessage()    {}

// Deprecated: Use MessageSetWireFormatMessage.ProtoReflect.Type instead.
func (*MessageSetWireFormatMessage) Descriptor() ([]byte, []int) {
	return xxx_File_extensions_base_base_proto_rawdesc_gzipped, []int{1}
}

var extRange_MessageSetWireFormatMessage = []proto.ExtensionRange{
	{Start: 100, End: 2147483646},
}

// Deprecated: Use MessageSetWireFormatMessage.ProtoReflect.Type.ExtensionRanges instead.
func (*MessageSetWireFormatMessage) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_MessageSetWireFormatMessage
}

func (m *MessageSetWireFormatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MessageSetWireFormatMessage.Unmarshal(m, b)
}
func (m *MessageSetWireFormatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MessageSetWireFormatMessage.Marshal(b, m, deterministic)
}
func (m *MessageSetWireFormatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MessageSetWireFormatMessage.Merge(m, src)
}
func (m *MessageSetWireFormatMessage) XXX_Size() int {
	return xxx_messageInfo_MessageSetWireFormatMessage.Size(m)
}
func (m *MessageSetWireFormatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_MessageSetWireFormatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_MessageSetWireFormatMessage proto.InternalMessageInfo

var xxx_File_extensions_base_base_proto_rawdesc = []byte{
	// 233 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x1a, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x62, 0x61, 0x73,
	0x65, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1d, 0x67, 0x6f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x22, 0x33, 0x0a, 0x0b, 0x42,
	0x61, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x69,
	0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x66, 0x69, 0x65, 0x6c, 0x64,
	0x2a, 0x04, 0x08, 0x04, 0x10, 0x0a, 0x2a, 0x08, 0x08, 0x10, 0x10, 0x80, 0x80, 0x80, 0x80, 0x02,
	0x22, 0x2b, 0x0a, 0x1b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x53, 0x65, 0x74, 0x57, 0x69,
	0x72, 0x65, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a,
	0x08, 0x08, 0x64, 0x10, 0xff, 0xff, 0xff, 0xff, 0x07, 0x3a, 0x02, 0x08, 0x01, 0x42, 0x4a, 0x5a,
	0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61,
	0x6e, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x63,
	0x6d, 0x64, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f,
	0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65,
}

var xxx_File_extensions_base_base_proto_rawdesc_gzipped = protoimpl.X.CompressGZIP(xxx_File_extensions_base_base_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_extensions_base_base_proto protoreflect.FileDescriptor

var xxx_File_extensions_base_base_proto_messageTypes = make([]protoimpl.MessageType, 2)
var xxx_File_extensions_base_base_proto_goTypes = []interface{}{
	(*BaseMessage)(nil),                 // 0: goproto.protoc.extension.base.BaseMessage
	(*MessageSetWireFormatMessage)(nil), // 1: goproto.protoc.extension.base.MessageSetWireFormatMessage
}
var xxx_File_extensions_base_base_proto_depIdxs = []int32{}

func init() { xxx_File_extensions_base_base_proto_init() }
func xxx_File_extensions_base_base_proto_init() {
	if File_extensions_base_base_proto != nil {
		return
	}
	messageTypes := make([]protoreflect.MessageType, 2)
	File_extensions_base_base_proto = protoimpl.FileBuilder{
		RawDescriptor:      xxx_File_extensions_base_base_proto_rawdesc,
		GoTypes:            xxx_File_extensions_base_base_proto_goTypes,
		DependencyIndexes:  xxx_File_extensions_base_base_proto_depIdxs,
		MessageOutputTypes: messageTypes,
	}.Init()
	messageGoTypes := xxx_File_extensions_base_base_proto_goTypes[0:][:2]
	for i, mt := range messageTypes {
		xxx_File_extensions_base_base_proto_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_File_extensions_base_base_proto_messageTypes[i].PBType = mt
	}
	proto.RegisterFile("extensions/base/base.proto", xxx_File_extensions_base_base_proto_rawdesc_gzipped)
	proto.RegisterType((*BaseMessage)(nil), "goproto.protoc.extension.base.BaseMessage")
	proto.RegisterType((*MessageSetWireFormatMessage)(nil), "goproto.protoc.extension.base.MessageSetWireFormatMessage")
	xxx_File_extensions_base_base_proto_goTypes = nil
	xxx_File_extensions_base_base_proto_depIdxs = nil
}
