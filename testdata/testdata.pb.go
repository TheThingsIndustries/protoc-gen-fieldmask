// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: testdata.proto

package testdata

import (
	fmt "fmt"
	github_com_TheThingsIndustries_protoc_gen_fieldmask_testdata_testpackage "github.com/TheThingsIndustries/protoc-gen-fieldmask/testdata/testpackage"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/duration"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	math "math"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{0}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

type Test struct {
	A          *Test_TestNested `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	CustomName *Test_TestNested `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	C          Test_TestNested  `protobuf:"bytes,3,opt,name=c,proto3" json:"c"`
	// Types that are valid to be assigned to TestOneof:
	//	*Test_D
	//	*Test_CustomNameOneof
	//	*Test_F
	//	*Test_K
	TestOneof            isTest_TestOneof   `protobuf_oneof:"testOneof"`
	G                    *Empty             `protobuf:"bytes,7,opt,name=g,proto3" json:"g,omitempty"`
	H                    *types.StringValue `protobuf:"bytes,8,opt,name=h,proto3" json:"h,omitempty"`
	I                    types.StringValue  `protobuf:"bytes,9,opt,name=i,proto3" json:"i"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Test) Reset()         { *m = Test{} }
func (m *Test) String() string { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()    {}
func (*Test) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{1}
}
func (m *Test) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test.Unmarshal(m, b)
}
func (m *Test) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test.Marshal(b, m, deterministic)
}
func (m *Test) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test.Merge(m, src)
}
func (m *Test) XXX_Size() int {
	return xxx_messageInfo_Test.Size(m)
}
func (m *Test) XXX_DiscardUnknown() {
	xxx_messageInfo_Test.DiscardUnknown(m)
}

var xxx_messageInfo_Test proto.InternalMessageInfo

type isTest_TestOneof interface {
	isTest_TestOneof()
}

type Test_D struct {
	D int32 `protobuf:"varint,4,opt,name=d,proto3,oneof"`
}
type Test_CustomNameOneof struct {
	CustomNameOneof uint32 `protobuf:"varint,5,opt,name=e,proto3,oneof"`
}
type Test_F struct {
	F []byte `protobuf:"bytes,6,opt,name=f,proto3,oneof"`
}
type Test_K struct {
	K *Test_TestNested `protobuf:"bytes,10,opt,name=k,proto3,oneof"`
}

func (*Test_D) isTest_TestOneof()               {}
func (*Test_CustomNameOneof) isTest_TestOneof() {}
func (*Test_F) isTest_TestOneof()               {}
func (*Test_K) isTest_TestOneof()               {}

func (m *Test) GetTestOneof() isTest_TestOneof {
	if m != nil {
		return m.TestOneof
	}
	return nil
}

func (m *Test) GetA() *Test_TestNested {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *Test) GetCustomName() *Test_TestNested {
	if m != nil {
		return m.CustomName
	}
	return nil
}

func (m *Test) GetC() Test_TestNested {
	if m != nil {
		return m.C
	}
	return Test_TestNested{}
}

func (m *Test) GetD() int32 {
	if x, ok := m.GetTestOneof().(*Test_D); ok {
		return x.D
	}
	return 0
}

func (m *Test) GetCustomNameOneof() uint32 {
	if x, ok := m.GetTestOneof().(*Test_CustomNameOneof); ok {
		return x.CustomNameOneof
	}
	return 0
}

func (m *Test) GetF() []byte {
	if x, ok := m.GetTestOneof().(*Test_F); ok {
		return x.F
	}
	return nil
}

func (m *Test) GetK() *Test_TestNested {
	if x, ok := m.GetTestOneof().(*Test_K); ok {
		return x.K
	}
	return nil
}

func (m *Test) GetG() *Empty {
	if m != nil {
		return m.G
	}
	return nil
}

func (m *Test) GetH() *types.StringValue {
	if m != nil {
		return m.H
	}
	return nil
}

func (m *Test) GetI() types.StringValue {
	if m != nil {
		return m.I
	}
	return types.StringValue{}
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Test) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Test_D)(nil),
		(*Test_CustomNameOneof)(nil),
		(*Test_F)(nil),
		(*Test_K)(nil),
	}
}

type Test_TestNested struct {
	A                    *Test_TestNested_TestNestedNested                                                    `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B                    []byte                                                                               `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	C                    *time.Duration                                                                       `protobuf:"bytes,3,opt,name=c,proto3,stdduration" json:"c,omitempty"`
	D                    *time.Time                                                                           `protobuf:"bytes,4,opt,name=d,proto3,stdtime" json:"d,omitempty"`
	E                    *github_com_TheThingsIndustries_protoc_gen_fieldmask_testdata_testpackage.CustomType `protobuf:"bytes,5,opt,name=e,proto3,customtype=github.com/TheThingsIndustries/protoc-gen-fieldmask/testdata/testpackage.CustomType" json:"e,omitempty"`
	F                    github_com_TheThingsIndustries_protoc_gen_fieldmask_testdata_testpackage.CustomType  `protobuf:"bytes,6,opt,name=f,proto3,customtype=github.com/TheThingsIndustries/protoc-gen-fieldmask/testdata/testpackage.CustomType" json:"f"`
	G                    []*Test_TestNested_TestNestedNested                                                  `protobuf:"bytes,7,rep,name=g,proto3" json:"g,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                                                             `json:"-"`
	XXX_unrecognized     []byte                                                                               `json:"-"`
	XXX_sizecache        int32                                                                                `json:"-"`
}

func (m *Test_TestNested) Reset()         { *m = Test_TestNested{} }
func (m *Test_TestNested) String() string { return proto.CompactTextString(m) }
func (*Test_TestNested) ProtoMessage()    {}
func (*Test_TestNested) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{1, 0}
}
func (m *Test_TestNested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test_TestNested.Unmarshal(m, b)
}
func (m *Test_TestNested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test_TestNested.Marshal(b, m, deterministic)
}
func (m *Test_TestNested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test_TestNested.Merge(m, src)
}
func (m *Test_TestNested) XXX_Size() int {
	return xxx_messageInfo_Test_TestNested.Size(m)
}
func (m *Test_TestNested) XXX_DiscardUnknown() {
	xxx_messageInfo_Test_TestNested.DiscardUnknown(m)
}

var xxx_messageInfo_Test_TestNested proto.InternalMessageInfo

func (m *Test_TestNested) GetA() *Test_TestNested_TestNestedNested {
	if m != nil {
		return m.A
	}
	return nil
}

func (m *Test_TestNested) GetB() []byte {
	if m != nil {
		return m.B
	}
	return nil
}

func (m *Test_TestNested) GetC() *time.Duration {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *Test_TestNested) GetD() *time.Time {
	if m != nil {
		return m.D
	}
	return nil
}

func (m *Test_TestNested) GetG() []*Test_TestNested_TestNestedNested {
	if m != nil {
		return m.G
	}
	return nil
}

type Test_TestNested_TestNestedNested struct {
	A int32            `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int64            `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"`
	C [][]byte         `protobuf:"bytes,3,rep,name=c,proto3" json:"c,omitempty"`
	D map[int32]uint32 `protobuf:"bytes,4,rep,name=d,proto3" json:"d,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Types that are valid to be assigned to TestNestedNestedOneOf:
	//	*Test_TestNested_TestNestedNested_E
	//	*Test_TestNested_TestNestedNested_F
	//	*Test_TestNested_TestNestedNested_G
	TestNestedNestedOneOf                                   isTest_TestNested_TestNestedNested_TestNestedNestedOneOf `protobuf_oneof:"testNestedNestedOneOf"`
	*Test_TestNested_TestNestedNested_TestNestedNestedEmbed `protobuf:"bytes,8,opt,name=h,proto3,embedded=h" json:"h,omitempty"`
	Test_TestNested_TestNestedNested_TestNestedNestedEmbed2 `protobuf:"bytes,9,opt,name=i,proto3,embedded=i" json:"i"`
	XXX_NoUnkeyedLiteral                                    struct{} `json:"-"`
	XXX_unrecognized                                        []byte   `json:"-"`
	XXX_sizecache                                           int32    `json:"-"`
}

func (m *Test_TestNested_TestNestedNested) Reset()         { *m = Test_TestNested_TestNestedNested{} }
func (m *Test_TestNested_TestNestedNested) String() string { return proto.CompactTextString(m) }
func (*Test_TestNested_TestNestedNested) ProtoMessage()    {}
func (*Test_TestNested_TestNestedNested) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{1, 0, 0}
}
func (m *Test_TestNested_TestNestedNested) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test_TestNested_TestNestedNested.Unmarshal(m, b)
}
func (m *Test_TestNested_TestNestedNested) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test_TestNested_TestNestedNested.Marshal(b, m, deterministic)
}
func (m *Test_TestNested_TestNestedNested) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test_TestNested_TestNestedNested.Merge(m, src)
}
func (m *Test_TestNested_TestNestedNested) XXX_Size() int {
	return xxx_messageInfo_Test_TestNested_TestNestedNested.Size(m)
}
func (m *Test_TestNested_TestNestedNested) XXX_DiscardUnknown() {
	xxx_messageInfo_Test_TestNested_TestNestedNested.DiscardUnknown(m)
}

var xxx_messageInfo_Test_TestNested_TestNestedNested proto.InternalMessageInfo

type isTest_TestNested_TestNestedNested_TestNestedNestedOneOf interface {
	isTest_TestNested_TestNestedNested_TestNestedNestedOneOf()
}

type Test_TestNested_TestNestedNested_E struct {
	E *Empty `protobuf:"bytes,5,opt,name=e,proto3,oneof"`
}
type Test_TestNested_TestNestedNested_F struct {
	F uint32 `protobuf:"varint,6,opt,name=f,proto3,oneof"`
}
type Test_TestNested_TestNestedNested_G struct {
	G *types.UInt64Value `protobuf:"bytes,7,opt,name=g,proto3,oneof"`
}

func (*Test_TestNested_TestNestedNested_E) isTest_TestNested_TestNestedNested_TestNestedNestedOneOf() {
}
func (*Test_TestNested_TestNestedNested_F) isTest_TestNested_TestNestedNested_TestNestedNestedOneOf() {
}
func (*Test_TestNested_TestNestedNested_G) isTest_TestNested_TestNestedNested_TestNestedNestedOneOf() {
}

func (m *Test_TestNested_TestNestedNested) GetTestNestedNestedOneOf() isTest_TestNested_TestNestedNested_TestNestedNestedOneOf {
	if m != nil {
		return m.TestNestedNestedOneOf
	}
	return nil
}

func (m *Test_TestNested_TestNestedNested) GetA() int32 {
	if m != nil {
		return m.A
	}
	return 0
}

func (m *Test_TestNested_TestNestedNested) GetB() int64 {
	if m != nil {
		return m.B
	}
	return 0
}

func (m *Test_TestNested_TestNestedNested) GetC() [][]byte {
	if m != nil {
		return m.C
	}
	return nil
}

func (m *Test_TestNested_TestNestedNested) GetD() map[int32]uint32 {
	if m != nil {
		return m.D
	}
	return nil
}

func (m *Test_TestNested_TestNestedNested) GetE() *Empty {
	if x, ok := m.GetTestNestedNestedOneOf().(*Test_TestNested_TestNestedNested_E); ok {
		return x.E
	}
	return nil
}

func (m *Test_TestNested_TestNestedNested) GetF() uint32 {
	if x, ok := m.GetTestNestedNestedOneOf().(*Test_TestNested_TestNestedNested_F); ok {
		return x.F
	}
	return 0
}

func (m *Test_TestNested_TestNestedNested) GetG() *types.UInt64Value {
	if x, ok := m.GetTestNestedNestedOneOf().(*Test_TestNested_TestNestedNested_G); ok {
		return x.G
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Test_TestNested_TestNestedNested) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Test_TestNested_TestNestedNested_E)(nil),
		(*Test_TestNested_TestNestedNested_F)(nil),
		(*Test_TestNested_TestNestedNested_G)(nil),
	}
}

type Test_TestNested_TestNestedNested_TestNestedNestedEmbed struct {
	NestedField          int32    `protobuf:"varint,1,opt,name=nested_field,json=nestedField,proto3" json:"nested_field,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test_TestNested_TestNestedNested_TestNestedNestedEmbed) Reset() {
	*m = Test_TestNested_TestNestedNested_TestNestedNestedEmbed{}
}
func (m *Test_TestNested_TestNestedNested_TestNestedNestedEmbed) String() string {
	return proto.CompactTextString(m)
}
func (*Test_TestNested_TestNestedNested_TestNestedNestedEmbed) ProtoMessage() {}
func (*Test_TestNested_TestNestedNested_TestNestedNestedEmbed) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{1, 0, 0, 1}
}
func (m *Test_TestNested_TestNestedNested_TestNestedNestedEmbed) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test_TestNested_TestNestedNested_TestNestedNestedEmbed.Unmarshal(m, b)
}
func (m *Test_TestNested_TestNestedNested_TestNestedNestedEmbed) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test_TestNested_TestNestedNested_TestNestedNestedEmbed.Marshal(b, m, deterministic)
}
func (m *Test_TestNested_TestNestedNested_TestNestedNestedEmbed) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test_TestNested_TestNestedNested_TestNestedNestedEmbed.Merge(m, src)
}
func (m *Test_TestNested_TestNestedNested_TestNestedNestedEmbed) XXX_Size() int {
	return xxx_messageInfo_Test_TestNested_TestNestedNested_TestNestedNestedEmbed.Size(m)
}
func (m *Test_TestNested_TestNestedNested_TestNestedNestedEmbed) XXX_DiscardUnknown() {
	xxx_messageInfo_Test_TestNested_TestNestedNested_TestNestedNestedEmbed.DiscardUnknown(m)
}

var xxx_messageInfo_Test_TestNested_TestNestedNested_TestNestedNestedEmbed proto.InternalMessageInfo

func (m *Test_TestNested_TestNestedNested_TestNestedNestedEmbed) GetNestedField() int32 {
	if m != nil {
		return m.NestedField
	}
	return 0
}

type Test_TestNested_TestNestedNested_TestNestedNestedEmbed2 struct {
	NestedField_2        int32    `protobuf:"varint,1,opt,name=nested_field_2,json=nestedField2,proto3" json:"nested_field_2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) Reset() {
	*m = Test_TestNested_TestNestedNested_TestNestedNestedEmbed2{}
}
func (m *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) String() string {
	return proto.CompactTextString(m)
}
func (*Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) ProtoMessage() {}
func (*Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) Descriptor() ([]byte, []int) {
	return fileDescriptor_40c4782d007dfce9, []int{1, 0, 0, 2}
}
func (m *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Test_TestNested_TestNestedNested_TestNestedNestedEmbed2.Unmarshal(m, b)
}
func (m *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Test_TestNested_TestNestedNested_TestNestedNestedEmbed2.Marshal(b, m, deterministic)
}
func (m *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Test_TestNested_TestNestedNested_TestNestedNestedEmbed2.Merge(m, src)
}
func (m *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) XXX_Size() int {
	return xxx_messageInfo_Test_TestNested_TestNestedNested_TestNestedNestedEmbed2.Size(m)
}
func (m *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) XXX_DiscardUnknown() {
	xxx_messageInfo_Test_TestNested_TestNestedNested_TestNestedNestedEmbed2.DiscardUnknown(m)
}

var xxx_messageInfo_Test_TestNested_TestNestedNested_TestNestedNestedEmbed2 proto.InternalMessageInfo

func (m *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) GetNestedField_2() int32 {
	if m != nil {
		return m.NestedField_2
	}
	return 0
}

func init() {
	proto.RegisterType((*Empty)(nil), "testdata.Empty")
	proto.RegisterType((*Test)(nil), "testdata.Test")
	proto.RegisterType((*Test_TestNested)(nil), "testdata.Test.TestNested")
	proto.RegisterType((*Test_TestNested_TestNestedNested)(nil), "testdata.Test.TestNested.TestNestedNested")
	proto.RegisterMapType((map[int32]uint32)(nil), "testdata.Test.TestNested.TestNestedNested.DEntry")
	proto.RegisterType((*Test_TestNested_TestNestedNested_TestNestedNestedEmbed)(nil), "testdata.Test.TestNested.TestNestedNested.TestNestedNestedEmbed")
	proto.RegisterType((*Test_TestNested_TestNestedNested_TestNestedNestedEmbed2)(nil), "testdata.Test.TestNested.TestNestedNested.TestNestedNestedEmbed2")
}

func init() { proto.RegisterFile("testdata.proto", fileDescriptor_40c4782d007dfce9) }

var fileDescriptor_40c4782d007dfce9 = []byte{
	// 753 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xcd, 0x6e, 0xd3, 0x4a,
	0x14, 0xee, 0x24, 0xce, 0xdf, 0x24, 0x4d, 0xa3, 0xb9, 0xb7, 0xb7, 0xbe, 0x16, 0x90, 0x50, 0x16,
	0x84, 0x88, 0x38, 0x10, 0x0a, 0xad, 0x2a, 0x04, 0x74, 0x68, 0x81, 0x6e, 0x5a, 0xc9, 0x0d, 0x20,
	0x60, 0x51, 0x39, 0xf1, 0xc4, 0xb1, 0x12, 0xff, 0xc8, 0x9e, 0x14, 0xf2, 0x0a, 0xac, 0x58, 0xf2,
	0x0c, 0x3c, 0x01, 0x8f, 0xc0, 0x92, 0x2d, 0x2c, 0x82, 0xd4, 0x35, 0x0f, 0x80, 0xb2, 0x42, 0x33,
	0x13, 0xe3, 0x90, 0xb4, 0x0a, 0x55, 0xc5, 0x26, 0x9a, 0x33, 0xe7, 0x3b, 0x67, 0x8e, 0xbf, 0xef,
	0x3b, 0x81, 0x79, 0x4a, 0x02, 0x6a, 0xe8, 0x54, 0x57, 0x3d, 0xdf, 0xa5, 0x2e, 0x4a, 0x87, 0xb1,
	0x52, 0x35, 0x2d, 0xda, 0xe9, 0x37, 0xd5, 0x96, 0x6b, 0xd7, 0x4c, 0xd7, 0x74, 0x6b, 0x1c, 0xd0,
	0xec, 0xb7, 0x79, 0xc4, 0x03, 0x7e, 0x12, 0x85, 0xca, 0xd6, 0x04, 0x9c, 0x38, 0x47, 0xee, 0xc0,
	0xf3, 0xdd, 0x37, 0x03, 0x51, 0xd4, 0xaa, 0x9a, 0xc4, 0xa9, 0x1e, 0xe9, 0x3d, 0xcb, 0xd0, 0x29,
	0xa9, 0xcd, 0x1c, 0xc6, 0x2d, 0x2e, 0x99, 0xae, 0x6b, 0xf6, 0x48, 0xf4, 0x90, 0xd1, 0xf7, 0x75,
	0x6a, 0xb9, 0xce, 0x38, 0x5f, 0x9c, 0xce, 0x53, 0xcb, 0x26, 0x01, 0xd5, 0x6d, 0xef, 0xb4, 0x06,
	0xaf, 0x7d, 0xdd, 0xf3, 0x88, 0x1f, 0x88, 0xfc, 0x6a, 0x0a, 0x26, 0x76, 0x6c, 0x8f, 0x0e, 0x56,
	0xbf, 0x64, 0xa1, 0xd4, 0x20, 0x01, 0x45, 0x57, 0x21, 0xd0, 0x65, 0x50, 0x02, 0xe5, 0x6c, 0xfd,
	0x7f, 0xf5, 0x17, 0x15, 0x2c, 0xc5, 0x7f, 0xf6, 0x48, 0x40, 0x89, 0xa1, 0x01, 0x1d, 0xad, 0x43,
	0xd0, 0x94, 0x63, 0x73, 0x80, 0x38, 0x7f, 0x3c, 0x2c, 0xc2, 0x87, 0xfd, 0x80, 0xba, 0xf6, 0x9e,
	0x6e, 0x13, 0x0d, 0x34, 0xd1, 0x6d, 0x08, 0x5a, 0x72, 0x7c, 0x5e, 0x61, 0x6e, 0x84, 0x13, 0x6f,
	0x41, 0xac, 0x00, 0x3e, 0x0d, 0x8b, 0x0b, 0x1a, 0x68, 0xa1, 0x3c, 0x04, 0x86, 0x2c, 0x95, 0x40,
	0x39, 0xf1, 0x64, 0x41, 0x03, 0x06, 0xba, 0x02, 0x01, 0x91, 0x13, 0x25, 0x50, 0x5e, 0xc4, 0xff,
	0x1c, 0x0f, 0x8b, 0x4b, 0xd1, 0x23, 0xfb, 0x0e, 0x71, 0xdb, 0x0c, 0x44, 0x58, 0x51, 0x5b, 0x4e,
	0x96, 0x40, 0x39, 0xc7, 0xe2, 0x36, 0xba, 0x06, 0x41, 0x57, 0x86, 0x73, 0xde, 0x66, 0xd0, 0x2e,
	0xba, 0x08, 0x81, 0x29, 0xa7, 0x38, 0x74, 0x29, 0x82, 0x72, 0xb6, 0x34, 0x60, 0xa2, 0x0a, 0x04,
	0x1d, 0x39, 0xcd, 0xd3, 0x17, 0x54, 0xc1, 0xb2, 0x1a, 0xb2, 0xac, 0x1e, 0x50, 0xdf, 0x72, 0xcc,
	0x67, 0x7a, 0xaf, 0x4f, 0x34, 0xd0, 0x41, 0x37, 0x20, 0xb0, 0xe4, 0xcc, 0x7c, 0x2c, 0x96, 0xc4,
	0xc7, 0x5a, 0xca, 0xc7, 0x34, 0x84, 0xd1, 0x40, 0x68, 0x23, 0x12, 0xa5, 0x72, 0xea, 0xd8, 0x13,
	0xc7, 0x48, 0xa5, 0x5c, 0xa8, 0x52, 0x8e, 0x51, 0xbf, 0x3e, 0x49, 0xfd, 0xf4, 0x20, 0xdb, 0x63,
	0x6f, 0xe1, 0xfc, 0x08, 0xa7, 0x3e, 0x00, 0xa9, 0x1e, 0x4b, 0x57, 0xde, 0x7f, 0x2b, 0x02, 0x46,
	0xbe, 0x1a, 0x92, 0x9f, 0xad, 0x2b, 0x33, 0x85, 0x8d, 0xd0, 0x74, 0x58, 0x7a, 0xc7, 0xf1, 0x06,
	0x22, 0xa1, 0x38, 0x19, 0xfc, 0xfc, 0xeb, 0xb0, 0x78, 0x30, 0xb1, 0x0a, 0x8d, 0x0e, 0x69, 0x74,
	0x2c, 0xc7, 0x0c, 0x76, 0x1d, 0xa3, 0x1f, 0x50, 0xdf, 0x22, 0xc1, 0xe4, 0x4e, 0xb4, 0x2d, 0xd2,
	0x33, 0x6c, 0x3d, 0xe8, 0xd6, 0xc2, 0xef, 0xe4, 0x07, 0x4f, 0x6f, 0x75, 0x75, 0x93, 0xa8, 0x42,
	0xe5, 0xc6, 0xc0, 0x23, 0x4c, 0x5e, 0x2b, 0x94, 0x37, 0x83, 0x5f, 0x31, 0xea, 0xfe, 0xda, 0x53,
	0x6d, 0x26, 0x01, 0xb3, 0x43, 0xfc, 0xac, 0x12, 0x98, 0xca, 0x77, 0x09, 0x16, 0xa6, 0xef, 0xd1,
	0x4a, 0xa8, 0x68, 0x02, 0x67, 0x46, 0x38, 0xa9, 0x48, 0x72, 0xa5, 0x24, 0xff, 0x26, 0x58, 0x81,
	0x09, 0x96, 0x13, 0x82, 0xc5, 0x99, 0x7c, 0x2d, 0x74, 0x5f, 0xa8, 0xc0, 0x66, 0xb8, 0xf9, 0xe7,
	0x33, 0xa8, 0xdb, 0x3b, 0x0e, 0xf5, 0x07, 0x4c, 0x96, 0x62, 0x28, 0xcb, 0xac, 0xa7, 0xa7, 0xf6,
	0x65, 0x51, 0xec, 0xcb, 0xf5, 0x68, 0x09, 0x66, 0x9d, 0xfb, 0x74, 0xd7, 0xa1, 0x77, 0xd6, 0xb8,
	0x73, 0x19, 0xda, 0x44, 0x8d, 0x68, 0x27, 0x1e, 0x9c, 0x61, 0xbe, 0xe9, 0x8b, 0x1d, 0xbb, 0x49,
	0x0c, 0x2c, 0x7d, 0x1e, 0x32, 0x2f, 0x75, 0xd0, 0x8b, 0x68, 0x7b, 0xb6, 0xce, 0xdb, 0xb5, 0x8e,
	0xd3, 0xcc, 0x27, 0xa2, 0xb5, 0xa5, 0xac, 0xc1, 0xa4, 0x20, 0x07, 0x15, 0x60, 0xbc, 0x4b, 0x06,
	0x42, 0x11, 0x8d, 0x1d, 0xd1, 0xbf, 0x30, 0x71, 0xc4, 0x3e, 0x8d, 0x8b, 0xb1, 0xa8, 0x89, 0x60,
	0x33, 0xb6, 0x01, 0x94, 0x4d, 0xb8, 0x7c, 0x62, 0x73, 0x74, 0x19, 0xe6, 0x1c, 0x1e, 0x1e, 0x72,
	0x6b, 0x8d, 0xbb, 0x65, 0xc5, 0xdd, 0x23, 0x76, 0xa5, 0x3c, 0x86, 0xff, 0x9d, 0x3c, 0x18, 0xaa,
	0xc2, 0xfc, 0x64, 0xf1, 0x61, 0x7d, 0x6c, 0x8f, 0xd4, 0x08, 0x4b, 0x4a, 0x2c, 0x1d, 0xd3, 0x72,
	0x13, 0x7d, 0xea, 0x78, 0x05, 0x2e, 0xd3, 0xa9, 0x46, 0xfb, 0x0e, 0xd9, 0x6f, 0xe3, 0x02, 0xcc,
	0xb0, 0x04, 0xff, 0x13, 0x44, 0xf1, 0x1f, 0x18, 0xe0, 0x7b, 0x2f, 0xef, 0x9e, 0x67, 0x29, 0x9a,
	0x49, 0x9e, 0xbe, 0xf5, 0x33, 0x00, 0x00, 0xff, 0xff, 0xed, 0x1e, 0xbb, 0xe8, 0x1a, 0x07, 0x00,
	0x00,
}
