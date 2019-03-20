// Code generated by protoc-gen-go. DO NOT EDIT.
// source: import_public/sub/a.proto

package sub

import (
	proto "github.com/golang/protobuf/proto"
	sub2 "github.com/golang/protobuf/v2/cmd/protoc-gen-go/testdata/import_public/sub2"
	protoreflect "github.com/golang/protobuf/v2/reflect/protoreflect"
	protoimpl "github.com/golang/protobuf/v2/runtime/protoimpl"
	math "math"
	reflect "reflect"
)

// Symbols defined in public import of import_public/sub2/a.proto

type Sub2Message = sub2.Sub2Message

type E int32

const (
	E_ZERO E = 0
)

func (e E) Type() protoreflect.EnumType {
	return xxx_File_import_public_sub_a_proto_enumTypes[0]
}
func (e E) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

// Deprecated: Use E.Type.Values instead.
var E_name = map[int32]string{
	0: "ZERO",
}

// Deprecated: Use E.Type.Values instead.
var E_value = map[string]int32{
	"ZERO": 0,
}

func (x E) Enum() *E {
	return &x
}

func (x E) String() string {
	return protoimpl.X.EnumStringOf(x.Type(), protoreflect.EnumNumber(x))
}

// Deprecated: Do not use.
func (x *E) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Type(), b)
	if err != nil {
		return err
	}
	*x = E(num)
	return nil
}

// Deprecated: Use E.Type instead.
func (E) EnumDescriptor() ([]byte, []int) {
	return xxx_File_import_public_sub_a_proto_rawdesc_gzipped, []int{0}
}

type M_Subenum int32

const (
	M_M_ZERO M_Subenum = 0
)

func (e M_Subenum) Type() protoreflect.EnumType {
	return xxx_File_import_public_sub_a_proto_enumTypes[1]
}
func (e M_Subenum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

// Deprecated: Use M_Subenum.Type.Values instead.
var M_Subenum_name = map[int32]string{
	0: "M_ZERO",
}

// Deprecated: Use M_Subenum.Type.Values instead.
var M_Subenum_value = map[string]int32{
	"M_ZERO": 0,
}

func (x M_Subenum) Enum() *M_Subenum {
	return &x
}

func (x M_Subenum) String() string {
	return protoimpl.X.EnumStringOf(x.Type(), protoreflect.EnumNumber(x))
}

// Deprecated: Do not use.
func (x *M_Subenum) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Type(), b)
	if err != nil {
		return err
	}
	*x = M_Subenum(num)
	return nil
}

// Deprecated: Use M_Subenum.Type instead.
func (M_Subenum) EnumDescriptor() ([]byte, []int) {
	return xxx_File_import_public_sub_a_proto_rawdesc_gzipped, []int{0, 0}
}

type M_Submessage_Submessage_Subenum int32

const (
	M_Submessage_M_SUBMESSAGE_ZERO M_Submessage_Submessage_Subenum = 0
)

func (e M_Submessage_Submessage_Subenum) Type() protoreflect.EnumType {
	return xxx_File_import_public_sub_a_proto_enumTypes[2]
}
func (e M_Submessage_Submessage_Subenum) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(e)
}

// Deprecated: Use M_Submessage_Submessage_Subenum.Type.Values instead.
var M_Submessage_Submessage_Subenum_name = map[int32]string{
	0: "M_SUBMESSAGE_ZERO",
}

// Deprecated: Use M_Submessage_Submessage_Subenum.Type.Values instead.
var M_Submessage_Submessage_Subenum_value = map[string]int32{
	"M_SUBMESSAGE_ZERO": 0,
}

func (x M_Submessage_Submessage_Subenum) Enum() *M_Submessage_Submessage_Subenum {
	return &x
}

func (x M_Submessage_Submessage_Subenum) String() string {
	return protoimpl.X.EnumStringOf(x.Type(), protoreflect.EnumNumber(x))
}

// Deprecated: Do not use.
func (x *M_Submessage_Submessage_Subenum) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Type(), b)
	if err != nil {
		return err
	}
	*x = M_Submessage_Submessage_Subenum(num)
	return nil
}

// Deprecated: Use M_Submessage_Submessage_Subenum.Type instead.
func (M_Submessage_Submessage_Subenum) EnumDescriptor() ([]byte, []int) {
	return xxx_File_import_public_sub_a_proto_rawdesc_gzipped, []int{0, 0, 0}
}

type M struct {
	// Field using a type in the same Go package, but a different source file.
	M2 *M2      `protobuf:"bytes,1,opt,name=m2" json:"m2,omitempty"`
	S  *string  `protobuf:"bytes,4,opt,name=s,def=default" json:"s,omitempty"`
	B  []byte   `protobuf:"bytes,5,opt,name=b,def=default" json:"b,omitempty"`
	F  *float64 `protobuf:"fixed64,6,opt,name=f,def=nan" json:"f,omitempty"`
	// Types that are valid to be assigned to OneofField:
	//	*M_OneofInt32
	//	*M_OneofInt64
	OneofField                   isM_OneofField `protobuf_oneof:"oneof_field"`
	XXX_NoUnkeyedLiteral         struct{}       `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *M) ProtoReflect() protoreflect.Message {
	return xxx_File_import_public_sub_a_proto_messageTypes[0].MessageOf(m)
}
func (m *M) Reset()         { *m = M{} }
func (m *M) String() string { return proto.CompactTextString(m) }
func (*M) ProtoMessage()    {}

// Deprecated: Use M.ProtoReflect.Type instead.
func (*M) Descriptor() ([]byte, []int) {
	return xxx_File_import_public_sub_a_proto_rawdesc_gzipped, []int{0}
}

var extRange_M = []proto.ExtensionRange{
	{Start: 100, End: 536870911},
}

// Deprecated: Use M.ProtoReflect.Type.ExtensionRanges instead.
func (*M) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_M
}

func (m *M) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M.Unmarshal(m, b)
}
func (m *M) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M.Marshal(b, m, deterministic)
}
func (m *M) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M.Merge(m, src)
}
func (m *M) XXX_Size() int {
	return xxx_messageInfo_M.Size(m)
}
func (m *M) XXX_DiscardUnknown() {
	xxx_messageInfo_M.DiscardUnknown(m)
}

var xxx_messageInfo_M proto.InternalMessageInfo

const Default_M_S string = "default"

var Default_M_B []byte = []byte("default")
var Default_M_F float64 = math.NaN()

func (m *M) GetM2() *M2 {
	if m != nil {
		return m.M2
	}
	return nil
}

func (m *M) GetS() string {
	if m != nil && m.S != nil {
		return *m.S
	}
	return Default_M_S
}

func (m *M) GetB() []byte {
	if m != nil && m.B != nil {
		return m.B
	}
	return append([]byte(nil), Default_M_B...)
}

func (m *M) GetF() float64 {
	if m != nil && m.F != nil {
		return *m.F
	}
	return Default_M_F
}

type isM_OneofField interface {
	isM_OneofField()
}

type M_OneofInt32 struct {
	OneofInt32 int32 `protobuf:"varint,2,opt,name=oneof_int32,json=oneofInt32,oneof"`
}

type M_OneofInt64 struct {
	OneofInt64 int64 `protobuf:"varint,3,opt,name=oneof_int64,json=oneofInt64,oneof"`
}

func (*M_OneofInt32) isM_OneofField() {}

func (*M_OneofInt64) isM_OneofField() {}

func (m *M) GetOneofField() isM_OneofField {
	if m != nil {
		return m.OneofField
	}
	return nil
}

func (m *M) GetOneofInt32() int32 {
	if x, ok := m.GetOneofField().(*M_OneofInt32); ok {
		return x.OneofInt32
	}
	return 0
}

func (m *M) GetOneofInt64() int64 {
	if x, ok := m.GetOneofField().(*M_OneofInt64); ok {
		return x.OneofInt64
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*M) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*M_OneofInt32)(nil),
		(*M_OneofInt64)(nil),
	}
}

type M_Submessage struct {
	// Types that are valid to be assigned to SubmessageOneofField:
	//	*M_Submessage_SubmessageOneofInt32
	//	*M_Submessage_SubmessageOneofInt64
	SubmessageOneofField isM_Submessage_SubmessageOneofField `protobuf_oneof:"submessage_oneof_field"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *M_Submessage) ProtoReflect() protoreflect.Message {
	return xxx_File_import_public_sub_a_proto_messageTypes[1].MessageOf(m)
}
func (m *M_Submessage) Reset()         { *m = M_Submessage{} }
func (m *M_Submessage) String() string { return proto.CompactTextString(m) }
func (*M_Submessage) ProtoMessage()    {}

// Deprecated: Use M_Submessage.ProtoReflect.Type instead.
func (*M_Submessage) Descriptor() ([]byte, []int) {
	return xxx_File_import_public_sub_a_proto_rawdesc_gzipped, []int{0, 0}
}

func (m *M_Submessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M_Submessage.Unmarshal(m, b)
}
func (m *M_Submessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M_Submessage.Marshal(b, m, deterministic)
}
func (m *M_Submessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M_Submessage.Merge(m, src)
}
func (m *M_Submessage) XXX_Size() int {
	return xxx_messageInfo_M_Submessage.Size(m)
}
func (m *M_Submessage) XXX_DiscardUnknown() {
	xxx_messageInfo_M_Submessage.DiscardUnknown(m)
}

var xxx_messageInfo_M_Submessage proto.InternalMessageInfo

type isM_Submessage_SubmessageOneofField interface {
	isM_Submessage_SubmessageOneofField()
}

type M_Submessage_SubmessageOneofInt32 struct {
	SubmessageOneofInt32 int32 `protobuf:"varint,1,opt,name=submessage_oneof_int32,json=submessageOneofInt32,oneof"`
}

type M_Submessage_SubmessageOneofInt64 struct {
	SubmessageOneofInt64 int64 `protobuf:"varint,2,opt,name=submessage_oneof_int64,json=submessageOneofInt64,oneof"`
}

func (*M_Submessage_SubmessageOneofInt32) isM_Submessage_SubmessageOneofField() {}

func (*M_Submessage_SubmessageOneofInt64) isM_Submessage_SubmessageOneofField() {}

func (m *M_Submessage) GetSubmessageOneofField() isM_Submessage_SubmessageOneofField {
	if m != nil {
		return m.SubmessageOneofField
	}
	return nil
}

func (m *M_Submessage) GetSubmessageOneofInt32() int32 {
	if x, ok := m.GetSubmessageOneofField().(*M_Submessage_SubmessageOneofInt32); ok {
		return x.SubmessageOneofInt32
	}
	return 0
}

func (m *M_Submessage) GetSubmessageOneofInt64() int64 {
	if x, ok := m.GetSubmessageOneofField().(*M_Submessage_SubmessageOneofInt64); ok {
		return x.SubmessageOneofInt64
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*M_Submessage) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*M_Submessage_SubmessageOneofInt32)(nil),
		(*M_Submessage_SubmessageOneofInt64)(nil),
	}
}

var xxx_File_import_public_sub_a_proto_extDescs = []proto.ExtensionDesc{
	{
		ExtendedType:  (*M)(nil),
		ExtensionType: (*string)(nil),
		Field:         100,
		Name:          "goproto.protoc.import_public.sub.extension_field",
		Tag:           "bytes,100,opt,name=extension_field",
		Filename:      "import_public/sub/a.proto",
	},
}
var (
	// extend goproto.protoc.import_public.sub.M { optional string extension_field = 100; }
	E_ExtensionField = &xxx_File_import_public_sub_a_proto_extDescs[0]
)
var xxx_File_import_public_sub_a_proto_rawdesc = []byte{
	// 730 bytes of the wire-encoded FileDescriptorProto
	0x0a, 0x19, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f,
	0x73, 0x75, 0x62, 0x2f, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x20, 0x67, 0x6f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x69, 0x6d, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x75, 0x62, 0x1a, 0x19, 0x69,
	0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x73, 0x75, 0x62,
	0x2f, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74,
	0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x73, 0x75, 0x62, 0x32, 0x2f, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb6, 0x03, 0x0a, 0x01, 0x4d, 0x12, 0x34, 0x0a, 0x02, 0x6d, 0x32,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x73, 0x75, 0x62, 0x2e, 0x4d, 0x32, 0x52, 0x02, 0x6d, 0x32,
	0x12, 0x15, 0x0a, 0x01, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x3a, 0x07, 0x64, 0x65, 0x66,
	0x61, 0x75, 0x6c, 0x74, 0x52, 0x01, 0x73, 0x12, 0x15, 0x0a, 0x01, 0x62, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x0c, 0x3a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x01, 0x62, 0x12, 0x11,
	0x0a, 0x01, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x3a, 0x03, 0x6e, 0x61, 0x6e, 0x52, 0x01,
	0x66, 0x12, 0x21, 0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x49,
	0x6e, 0x74, 0x33, 0x32, 0x12, 0x21, 0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x69, 0x6e,
	0x74, 0x36, 0x34, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0a, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x1a, 0xc3, 0x01, 0x0a, 0x0a, 0x53, 0x75, 0x62, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x16, 0x73, 0x75, 0x62, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x14, 0x73, 0x75, 0x62, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x49, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x36,
	0x0a, 0x16, 0x73, 0x75, 0x62, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x6f, 0x6e, 0x65,
	0x6f, 0x66, 0x5f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00,
	0x52, 0x14, 0x73, 0x75, 0x62, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4f, 0x6e, 0x65, 0x6f,
	0x66, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x22, 0x2b, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x5f, 0x53, 0x75, 0x62, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x15, 0x0a, 0x11,
	0x4d, 0x5f, 0x53, 0x55, 0x42, 0x4d, 0x45, 0x53, 0x53, 0x41, 0x47, 0x45, 0x5f, 0x5a, 0x45, 0x52,
	0x4f, 0x10, 0x00, 0x42, 0x18, 0x0a, 0x16, 0x73, 0x75, 0x62, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x22, 0x15, 0x0a,
	0x07, 0x53, 0x75, 0x62, 0x65, 0x6e, 0x75, 0x6d, 0x12, 0x0a, 0x0a, 0x06, 0x4d, 0x5f, 0x5a, 0x45,
	0x52, 0x4f, 0x10, 0x00, 0x2a, 0x08, 0x08, 0x64, 0x10, 0x80, 0x80, 0x80, 0x80, 0x02, 0x42, 0x0d,
	0x0a, 0x0b, 0x6f, 0x6e, 0x65, 0x6f, 0x66, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x2a, 0x0d, 0x0a,
	0x01, 0x45, 0x12, 0x08, 0x0a, 0x04, 0x5a, 0x45, 0x52, 0x4f, 0x10, 0x00, 0x3a, 0x4c, 0x0a, 0x0f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x12,
	0x23, 0x2e, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63,
	0x2e, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2e, 0x73,
	0x75, 0x62, 0x2e, 0x4d, 0x18, 0x64, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x65, 0x78, 0x74, 0x65,
	0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x42, 0x4c, 0x5a, 0x4a, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x32, 0x2f, 0x63, 0x6d, 0x64, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x74, 0x65,
	0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x69, 0x6d, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x2f, 0x73, 0x75, 0x62, 0x50, 0x01,
}

var xxx_File_import_public_sub_a_proto_rawdesc_gzipped = protoimpl.X.CompressGZIP(xxx_File_import_public_sub_a_proto_rawdesc)

const _ = protoimpl.EnforceVersion(protoimpl.Version - 0)

var File_import_public_sub_a_proto protoreflect.FileDescriptor

var xxx_File_import_public_sub_a_proto_enumTypes = make([]protoreflect.EnumType, 3)
var xxx_File_import_public_sub_a_proto_messageTypes = make([]protoimpl.MessageType, 2)
var xxx_File_import_public_sub_a_proto_goTypes = []interface{}{
	(E)(0),                               // 0: goproto.protoc.import_public.sub.E
	(M_Subenum)(0),                       // 1: goproto.protoc.import_public.sub.M.Subenum
	(M_Submessage_Submessage_Subenum)(0), // 2: goproto.protoc.import_public.sub.M.Submessage.Submessage_Subenum
	(*M)(nil),                            // 3: goproto.protoc.import_public.sub.M
	(*M_Submessage)(nil),                 // 4: goproto.protoc.import_public.sub.M.Submessage
	(*M2)(nil),                           // 5: goproto.protoc.import_public.sub.M2
}
var xxx_File_import_public_sub_a_proto_depIdxs = []int32{
	3, // goproto.protoc.import_public.sub.extension_field:extendee -> goproto.protoc.import_public.sub.M
	5, // goproto.protoc.import_public.sub.M.m2:type_name -> goproto.protoc.import_public.sub.M2
}

func init() { xxx_File_import_public_sub_a_proto_init() }
func xxx_File_import_public_sub_a_proto_init() {
	if File_import_public_sub_a_proto != nil {
		return
	}
	xxx_File_import_public_sub_b_proto_init()
	messageTypes := make([]protoreflect.MessageType, 2)
	extensionTypes := make([]protoreflect.ExtensionType, 1)
	File_import_public_sub_a_proto = protoimpl.FileBuilder{
		RawDescriptor:        xxx_File_import_public_sub_a_proto_rawdesc,
		GoTypes:              xxx_File_import_public_sub_a_proto_goTypes,
		DependencyIndexes:    xxx_File_import_public_sub_a_proto_depIdxs,
		LegacyExtensions:     xxx_File_import_public_sub_a_proto_extDescs,
		EnumOutputTypes:      xxx_File_import_public_sub_a_proto_enumTypes,
		MessageOutputTypes:   messageTypes,
		ExtensionOutputTypes: extensionTypes,
	}.Init()
	messageGoTypes := xxx_File_import_public_sub_a_proto_goTypes[3:][:2]
	for i, mt := range messageTypes {
		xxx_File_import_public_sub_a_proto_messageTypes[i].GoType = reflect.TypeOf(messageGoTypes[i])
		xxx_File_import_public_sub_a_proto_messageTypes[i].PBType = mt
	}
	proto.RegisterFile("import_public/sub/a.proto", xxx_File_import_public_sub_a_proto_rawdesc_gzipped)
	proto.RegisterEnum("goproto.protoc.import_public.sub.E", E_name, E_value)
	proto.RegisterEnum("goproto.protoc.import_public.sub.M_Subenum", M_Subenum_name, M_Subenum_value)
	proto.RegisterEnum("goproto.protoc.import_public.sub.M_Submessage_Submessage_Subenum", M_Submessage_Submessage_Subenum_name, M_Submessage_Submessage_Subenum_value)
	proto.RegisterType((*M)(nil), "goproto.protoc.import_public.sub.M")
	proto.RegisterType((*M_Submessage)(nil), "goproto.protoc.import_public.sub.M.Submessage")
	proto.RegisterExtension(E_ExtensionField)
	xxx_File_import_public_sub_a_proto_goTypes = nil
	xxx_File_import_public_sub_a_proto_depIdxs = nil
}
