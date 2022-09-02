// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: testdata.proto

package testdata

import (
	otherpackage "github.com/TheThingsIndustries/protoc-gen-fieldmask/testdata/otherpackage"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_testdata_proto_rawDescGZIP(), []int{0}
}

type Test struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A *Test_TestNested `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B *Test_TestNested `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	C *Test_TestNested `protobuf:"bytes,3,opt,name=c,proto3" json:"c,omitempty"`
	// Types that are assignable to TestOneof:
	//
	//	*Test_D
	//	*Test_E
	//	*Test_F
	//	*Test_K
	TestOneof isTest_TestOneof        `protobuf_oneof:"testOneof"`
	G         *Empty                  `protobuf:"bytes,7,opt,name=g,proto3" json:"g,omitempty"`
	H         *wrapperspb.StringValue `protobuf:"bytes,8,opt,name=h,proto3" json:"h,omitempty"`
	I         *wrapperspb.StringValue `protobuf:"bytes,9,opt,name=i,proto3" json:"i,omitempty"`
	J         *otherpackage.Embed     `protobuf:"bytes,11,opt,name=j,proto3" json:"j,omitempty"`
	L         string                  `protobuf:"bytes,12,opt,name=l,proto3" json:"l,omitempty"`
}

func (x *Test) Reset() {
	*x = Test{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test) ProtoMessage() {}

func (x *Test) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test.ProtoReflect.Descriptor instead.
func (*Test) Descriptor() ([]byte, []int) {
	return file_testdata_proto_rawDescGZIP(), []int{1}
}

func (x *Test) GetA() *Test_TestNested {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *Test) GetB() *Test_TestNested {
	if x != nil {
		return x.B
	}
	return nil
}

func (x *Test) GetC() *Test_TestNested {
	if x != nil {
		return x.C
	}
	return nil
}

func (m *Test) GetTestOneof() isTest_TestOneof {
	if m != nil {
		return m.TestOneof
	}
	return nil
}

func (x *Test) GetD() int32 {
	if x, ok := x.GetTestOneof().(*Test_D); ok {
		return x.D
	}
	return 0
}

func (x *Test) GetE() uint32 {
	if x, ok := x.GetTestOneof().(*Test_E); ok {
		return x.E
	}
	return 0
}

func (x *Test) GetF() []byte {
	if x, ok := x.GetTestOneof().(*Test_F); ok {
		return x.F
	}
	return nil
}

func (x *Test) GetK() *Test_TestNested {
	if x, ok := x.GetTestOneof().(*Test_K); ok {
		return x.K
	}
	return nil
}

func (x *Test) GetG() *Empty {
	if x != nil {
		return x.G
	}
	return nil
}

func (x *Test) GetH() *wrapperspb.StringValue {
	if x != nil {
		return x.H
	}
	return nil
}

func (x *Test) GetI() *wrapperspb.StringValue {
	if x != nil {
		return x.I
	}
	return nil
}

func (x *Test) GetJ() *otherpackage.Embed {
	if x != nil {
		return x.J
	}
	return nil
}

func (x *Test) GetL() string {
	if x != nil {
		return x.L
	}
	return ""
}

type isTest_TestOneof interface {
	isTest_TestOneof()
}

type Test_D struct {
	D int32 `protobuf:"varint,4,opt,name=d,proto3,oneof"`
}

type Test_E struct {
	E uint32 `protobuf:"varint,5,opt,name=e,proto3,oneof"`
}

type Test_F struct {
	F []byte `protobuf:"bytes,6,opt,name=f,proto3,oneof"`
}

type Test_K struct {
	K *Test_TestNested `protobuf:"bytes,10,opt,name=k,proto3,oneof"`
}

func (*Test_D) isTest_TestOneof() {}

func (*Test_E) isTest_TestOneof() {}

func (*Test_F) isTest_TestOneof() {}

func (*Test_K) isTest_TestOneof() {}

type Test_TestNested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A *Test_TestNested_TestNestedNested   `protobuf:"bytes,1,opt,name=a,proto3" json:"a,omitempty"`
	B []byte                              `protobuf:"bytes,2,opt,name=b,proto3" json:"b,omitempty"`
	C *durationpb.Duration                `protobuf:"bytes,3,opt,name=c,proto3" json:"c,omitempty"`
	D *timestamppb.Timestamp              `protobuf:"bytes,4,opt,name=d,proto3" json:"d,omitempty"`
	E string                              `protobuf:"bytes,5,opt,name=e,proto3" json:"e,omitempty"`
	F string                              `protobuf:"bytes,6,opt,name=f,proto3" json:"f,omitempty"`
	G []*Test_TestNested_TestNestedNested `protobuf:"bytes,7,rep,name=g,proto3" json:"g,omitempty"`
}

func (x *Test_TestNested) Reset() {
	*x = Test_TestNested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test_TestNested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test_TestNested) ProtoMessage() {}

func (x *Test_TestNested) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test_TestNested.ProtoReflect.Descriptor instead.
func (*Test_TestNested) Descriptor() ([]byte, []int) {
	return file_testdata_proto_rawDescGZIP(), []int{1, 0}
}

func (x *Test_TestNested) GetA() *Test_TestNested_TestNestedNested {
	if x != nil {
		return x.A
	}
	return nil
}

func (x *Test_TestNested) GetB() []byte {
	if x != nil {
		return x.B
	}
	return nil
}

func (x *Test_TestNested) GetC() *durationpb.Duration {
	if x != nil {
		return x.C
	}
	return nil
}

func (x *Test_TestNested) GetD() *timestamppb.Timestamp {
	if x != nil {
		return x.D
	}
	return nil
}

func (x *Test_TestNested) GetE() string {
	if x != nil {
		return x.E
	}
	return ""
}

func (x *Test_TestNested) GetF() string {
	if x != nil {
		return x.F
	}
	return ""
}

func (x *Test_TestNested) GetG() []*Test_TestNested_TestNestedNested {
	if x != nil {
		return x.G
	}
	return nil
}

type Test_TestNested_TestNestedNested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int32            `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int64            `protobuf:"fixed64,2,opt,name=b,proto3" json:"b,omitempty"`
	C []string         `protobuf:"bytes,3,rep,name=c,proto3" json:"c,omitempty"`
	D map[int32]uint32 `protobuf:"bytes,4,rep,name=d,proto3" json:"d,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	// Types that are assignable to TestNestedNestedOneOf:
	//
	//	*Test_TestNested_TestNestedNested_E
	//	*Test_TestNested_TestNestedNested_F
	//	*Test_TestNested_TestNestedNested_G
	TestNestedNestedOneOf isTest_TestNested_TestNestedNested_TestNestedNestedOneOf `protobuf_oneof:"testNestedNestedOneOf"`
	H                     *Test_TestNested_TestNestedNested_TestNestedNestedEmbed  `protobuf:"bytes,8,opt,name=h,proto3" json:"h,omitempty"`
	I                     *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2 `protobuf:"bytes,9,opt,name=i,proto3" json:"i,omitempty"`
}

func (x *Test_TestNested_TestNestedNested) Reset() {
	*x = Test_TestNested_TestNestedNested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test_TestNested_TestNestedNested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test_TestNested_TestNestedNested) ProtoMessage() {}

func (x *Test_TestNested_TestNestedNested) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test_TestNested_TestNestedNested.ProtoReflect.Descriptor instead.
func (*Test_TestNested_TestNestedNested) Descriptor() ([]byte, []int) {
	return file_testdata_proto_rawDescGZIP(), []int{1, 0, 0}
}

func (x *Test_TestNested_TestNestedNested) GetA() int32 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *Test_TestNested_TestNestedNested) GetB() int64 {
	if x != nil {
		return x.B
	}
	return 0
}

func (x *Test_TestNested_TestNestedNested) GetC() []string {
	if x != nil {
		return x.C
	}
	return nil
}

func (x *Test_TestNested_TestNestedNested) GetD() map[int32]uint32 {
	if x != nil {
		return x.D
	}
	return nil
}

func (m *Test_TestNested_TestNestedNested) GetTestNestedNestedOneOf() isTest_TestNested_TestNestedNested_TestNestedNestedOneOf {
	if m != nil {
		return m.TestNestedNestedOneOf
	}
	return nil
}

func (x *Test_TestNested_TestNestedNested) GetE() *Empty {
	if x, ok := x.GetTestNestedNestedOneOf().(*Test_TestNested_TestNestedNested_E); ok {
		return x.E
	}
	return nil
}

func (x *Test_TestNested_TestNestedNested) GetF() uint32 {
	if x, ok := x.GetTestNestedNestedOneOf().(*Test_TestNested_TestNestedNested_F); ok {
		return x.F
	}
	return 0
}

func (x *Test_TestNested_TestNestedNested) GetG() *wrapperspb.UInt64Value {
	if x, ok := x.GetTestNestedNestedOneOf().(*Test_TestNested_TestNestedNested_G); ok {
		return x.G
	}
	return nil
}

func (x *Test_TestNested_TestNestedNested) GetH() *Test_TestNested_TestNestedNested_TestNestedNestedEmbed {
	if x != nil {
		return x.H
	}
	return nil
}

func (x *Test_TestNested_TestNestedNested) GetI() *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2 {
	if x != nil {
		return x.I
	}
	return nil
}

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
	G *wrapperspb.UInt64Value `protobuf:"bytes,7,opt,name=g,proto3,oneof"`
}

func (*Test_TestNested_TestNestedNested_E) isTest_TestNested_TestNestedNested_TestNestedNestedOneOf() {
}

func (*Test_TestNested_TestNestedNested_F) isTest_TestNested_TestNestedNested_TestNestedNestedOneOf() {
}

func (*Test_TestNested_TestNestedNested_G) isTest_TestNested_TestNestedNested_TestNestedNestedOneOf() {
}

type Test_TestNested_TestNestedNested_TestNestedNestedEmbed struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NestedField int32 `protobuf:"varint,1,opt,name=nested_field,json=nestedField,proto3" json:"nested_field,omitempty"`
}

func (x *Test_TestNested_TestNestedNested_TestNestedNestedEmbed) Reset() {
	*x = Test_TestNested_TestNestedNested_TestNestedNestedEmbed{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test_TestNested_TestNestedNested_TestNestedNestedEmbed) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test_TestNested_TestNestedNested_TestNestedNestedEmbed) ProtoMessage() {}

func (x *Test_TestNested_TestNestedNested_TestNestedNestedEmbed) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test_TestNested_TestNestedNested_TestNestedNestedEmbed.ProtoReflect.Descriptor instead.
func (*Test_TestNested_TestNestedNested_TestNestedNestedEmbed) Descriptor() ([]byte, []int) {
	return file_testdata_proto_rawDescGZIP(), []int{1, 0, 0, 1}
}

func (x *Test_TestNested_TestNestedNested_TestNestedNestedEmbed) GetNestedField() int32 {
	if x != nil {
		return x.NestedField
	}
	return 0
}

type Test_TestNested_TestNestedNested_TestNestedNestedEmbed2 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NestedField_2 int32 `protobuf:"varint,1,opt,name=nested_field_2,json=nestedField2,proto3" json:"nested_field_2,omitempty"`
}

func (x *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) Reset() {
	*x = Test_TestNested_TestNestedNested_TestNestedNestedEmbed2{}
	if protoimpl.UnsafeEnabled {
		mi := &file_testdata_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) ProtoMessage() {}

func (x *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) ProtoReflect() protoreflect.Message {
	mi := &file_testdata_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Test_TestNested_TestNestedNested_TestNestedNestedEmbed2.ProtoReflect.Descriptor instead.
func (*Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) Descriptor() ([]byte, []int) {
	return file_testdata_proto_rawDescGZIP(), []int{1, 0, 0, 2}
}

func (x *Test_TestNested_TestNestedNested_TestNestedNestedEmbed2) GetNestedField_2() int32 {
	if x != nil {
		return x.NestedField_2
	}
	return 0
}

var File_testdata_proto protoreflect.FileDescriptor

var file_testdata_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x41, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x6e, 0x76, 0x6f, 0x79, 0x70, 0x72, 0x6f, 0x78,
	0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x6f, 0x74, 0x68, 0x65, 0x72, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2f, 0x6f, 0x74, 0x68,
	0x65, 0x72, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xc0, 0x0a, 0x0a, 0x04, 0x54, 0x65, 0x73,
	0x74, 0x12, 0x27, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74,
	0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73,
	0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x01, 0x61, 0x12, 0x27, 0x0a, 0x01, 0x62, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x52, 0x01, 0x62, 0x12, 0x31, 0x0a, 0x01, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x42, 0x08, 0xfa, 0x42, 0x05, 0x8a, 0x01,
	0x02, 0x10, 0x01, 0x52, 0x01, 0x63, 0x12, 0x17, 0x0a, 0x01, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02, 0x20, 0x05, 0x48, 0x00, 0x52, 0x01, 0x64, 0x12,
	0x17, 0x0a, 0x01, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a,
	0x02, 0x20, 0x05, 0x48, 0x00, 0x52, 0x01, 0x65, 0x12, 0x0e, 0x0a, 0x01, 0x66, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x48, 0x00, 0x52, 0x01, 0x66, 0x12, 0x29, 0x0a, 0x01, 0x6b, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54,
	0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x48, 0x00,
	0x52, 0x01, 0x6b, 0x12, 0x1d, 0x0a, 0x01, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52,
	0x01, 0x67, 0x12, 0x2a, 0x0a, 0x01, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x01, 0x68, 0x12, 0x2a,
	0x0a, 0x01, 0x69, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x01, 0x69, 0x12, 0x21, 0x0a, 0x01, 0x6a, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x70, 0x61, 0x63,
	0x6b, 0x61, 0x67, 0x65, 0x2e, 0x45, 0x6d, 0x62, 0x65, 0x64, 0x52, 0x01, 0x6a, 0x12, 0x19, 0x0a,
	0x01, 0x6c, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0xfa, 0x42, 0x08, 0x72, 0x06, 0x88,
	0x01, 0x01, 0xd0, 0x01, 0x01, 0x52, 0x01, 0x6c, 0x1a, 0xfe, 0x06, 0x0a, 0x0a, 0x54, 0x65, 0x73,
	0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x38, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x54, 0x65,
	0x73, 0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x52, 0x01,
	0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x01, 0x62, 0x12,
	0x33, 0x0a, 0x01, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0a, 0xfa, 0x42, 0x07, 0xaa, 0x01, 0x04, 0x32, 0x02, 0x08,
	0x2a, 0x52, 0x01, 0x63, 0x12, 0x28, 0x0a, 0x01, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x01, 0x64, 0x12, 0x0c,
	0x0a, 0x01, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x65, 0x12, 0x0c, 0x0a, 0x01,
	0x66, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x01, 0x66, 0x12, 0x38, 0x0a, 0x01, 0x67, 0x18,
	0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64,
	0x2e, 0x54, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4e, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x52, 0x01, 0x67, 0x1a, 0xf2, 0x04, 0x0a, 0x10, 0x54, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x73,
	0x74, 0x65, 0x64, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x17, 0x0a, 0x01, 0x61, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x1a, 0x04, 0x18, 0x2a, 0x20, 0x18, 0x52,
	0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x10, 0x52, 0x01, 0x62,
	0x12, 0x22, 0x0a, 0x01, 0x63, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x42, 0x14, 0xfa, 0x42, 0x05,
	0x92, 0x01, 0x02, 0x10, 0x09, 0xfa, 0x42, 0x09, 0x92, 0x01, 0x06, 0x22, 0x04, 0x72, 0x02, 0x18,
	0x40, 0x52, 0x01, 0x63, 0x12, 0x3f, 0x0a, 0x01, 0x64, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x31, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x44, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x52, 0x01, 0x64, 0x12, 0x1f, 0x0a, 0x01, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x48, 0x00, 0x52, 0x01, 0x65, 0x12, 0x0e, 0x0a, 0x01, 0x66, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0d, 0x48, 0x00, 0x52, 0x01, 0x66, 0x12, 0x2c, 0x0a, 0x01, 0x67, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x48,
	0x00, 0x52, 0x01, 0x67, 0x12, 0x4e, 0x0a, 0x01, 0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x40, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6d, 0x62, 0x65,
	0x64, 0x52, 0x01, 0x68, 0x12, 0x4f, 0x0a, 0x01, 0x69, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x41, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x2e,
	0x54, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x54, 0x65, 0x73, 0x74,
	0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6d, 0x62, 0x65,
	0x64, 0x32, 0x52, 0x01, 0x69, 0x1a, 0x34, 0x0a, 0x06, 0x44, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x15, 0x54,
	0x65, 0x73, 0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45,
	0x6d, 0x62, 0x65, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x66,
	0x69, 0x65, 0x6c, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x6e, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x1a, 0x47, 0x0a, 0x16, 0x54, 0x65, 0x73, 0x74, 0x4e,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x45, 0x6d, 0x62, 0x65, 0x64,
	0x32, 0x12, 0x2d, 0x0a, 0x0e, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x5f, 0x66, 0x69, 0x65, 0x6c,
	0x64, 0x5f, 0x32, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x1a, 0x02,
	0x08, 0x02, 0x52, 0x0c, 0x6e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x46, 0x69, 0x65, 0x6c, 0x64, 0x32,
	0x42, 0x17, 0x0a, 0x15, 0x74, 0x65, 0x73, 0x74, 0x4e, 0x65, 0x73, 0x74, 0x65, 0x64, 0x4e, 0x65,
	0x73, 0x74, 0x65, 0x64, 0x4f, 0x6e, 0x65, 0x4f, 0x66, 0x42, 0x10, 0x0a, 0x09, 0x74, 0x65, 0x73,
	0x74, 0x4f, 0x6e, 0x65, 0x6f, 0x66, 0x12, 0x03, 0xf8, 0x42, 0x01, 0x42, 0x3e, 0x5a, 0x3c, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x54, 0x68, 0x65, 0x54, 0x68, 0x69,
	0x6e, 0x67, 0x73, 0x49, 0x6e, 0x64, 0x75, 0x73, 0x74, 0x72, 0x69, 0x65, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x6d, 0x61,
	0x73, 0x6b, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_testdata_proto_rawDescOnce sync.Once
	file_testdata_proto_rawDescData = file_testdata_proto_rawDesc
)

func file_testdata_proto_rawDescGZIP() []byte {
	file_testdata_proto_rawDescOnce.Do(func() {
		file_testdata_proto_rawDescData = protoimpl.X.CompressGZIP(file_testdata_proto_rawDescData)
	})
	return file_testdata_proto_rawDescData
}

var file_testdata_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_testdata_proto_goTypes = []interface{}{
	(*Empty)(nil),                            // 0: testdata.Empty
	(*Test)(nil),                             // 1: testdata.Test
	(*Test_TestNested)(nil),                  // 2: testdata.Test.TestNested
	(*Test_TestNested_TestNestedNested)(nil), // 3: testdata.Test.TestNested.TestNestedNested
	nil,                                      // 4: testdata.Test.TestNested.TestNestedNested.DEntry
	(*Test_TestNested_TestNestedNested_TestNestedNestedEmbed)(nil),  // 5: testdata.Test.TestNested.TestNestedNested.TestNestedNestedEmbed
	(*Test_TestNested_TestNestedNested_TestNestedNestedEmbed2)(nil), // 6: testdata.Test.TestNested.TestNestedNested.TestNestedNestedEmbed2
	(*wrapperspb.StringValue)(nil),                                  // 7: google.protobuf.StringValue
	(*otherpackage.Embed)(nil),                                      // 8: otherpackage.Embed
	(*durationpb.Duration)(nil),                                     // 9: google.protobuf.Duration
	(*timestamppb.Timestamp)(nil),                                   // 10: google.protobuf.Timestamp
	(*wrapperspb.UInt64Value)(nil),                                  // 11: google.protobuf.UInt64Value
}
var file_testdata_proto_depIdxs = []int32{
	2,  // 0: testdata.Test.a:type_name -> testdata.Test.TestNested
	2,  // 1: testdata.Test.b:type_name -> testdata.Test.TestNested
	2,  // 2: testdata.Test.c:type_name -> testdata.Test.TestNested
	2,  // 3: testdata.Test.k:type_name -> testdata.Test.TestNested
	0,  // 4: testdata.Test.g:type_name -> testdata.Empty
	7,  // 5: testdata.Test.h:type_name -> google.protobuf.StringValue
	7,  // 6: testdata.Test.i:type_name -> google.protobuf.StringValue
	8,  // 7: testdata.Test.j:type_name -> otherpackage.Embed
	3,  // 8: testdata.Test.TestNested.a:type_name -> testdata.Test.TestNested.TestNestedNested
	9,  // 9: testdata.Test.TestNested.c:type_name -> google.protobuf.Duration
	10, // 10: testdata.Test.TestNested.d:type_name -> google.protobuf.Timestamp
	3,  // 11: testdata.Test.TestNested.g:type_name -> testdata.Test.TestNested.TestNestedNested
	4,  // 12: testdata.Test.TestNested.TestNestedNested.d:type_name -> testdata.Test.TestNested.TestNestedNested.DEntry
	0,  // 13: testdata.Test.TestNested.TestNestedNested.e:type_name -> testdata.Empty
	11, // 14: testdata.Test.TestNested.TestNestedNested.g:type_name -> google.protobuf.UInt64Value
	5,  // 15: testdata.Test.TestNested.TestNestedNested.h:type_name -> testdata.Test.TestNested.TestNestedNested.TestNestedNestedEmbed
	6,  // 16: testdata.Test.TestNested.TestNestedNested.i:type_name -> testdata.Test.TestNested.TestNestedNested.TestNestedNestedEmbed2
	17, // [17:17] is the sub-list for method output_type
	17, // [17:17] is the sub-list for method input_type
	17, // [17:17] is the sub-list for extension type_name
	17, // [17:17] is the sub-list for extension extendee
	0,  // [0:17] is the sub-list for field type_name
}

func init() { file_testdata_proto_init() }
func file_testdata_proto_init() {
	if File_testdata_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_testdata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_testdata_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test); i {
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
		file_testdata_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test_TestNested); i {
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
		file_testdata_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test_TestNested_TestNestedNested); i {
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
		file_testdata_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test_TestNested_TestNestedNested_TestNestedNestedEmbed); i {
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
		file_testdata_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Test_TestNested_TestNestedNested_TestNestedNestedEmbed2); i {
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
	file_testdata_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*Test_D)(nil),
		(*Test_E)(nil),
		(*Test_F)(nil),
		(*Test_K)(nil),
	}
	file_testdata_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*Test_TestNested_TestNestedNested_E)(nil),
		(*Test_TestNested_TestNestedNested_F)(nil),
		(*Test_TestNested_TestNestedNested_G)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_testdata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_testdata_proto_goTypes,
		DependencyIndexes: file_testdata_proto_depIdxs,
		MessageInfos:      file_testdata_proto_msgTypes,
	}.Build()
	File_testdata_proto = out.File
	file_testdata_proto_rawDesc = nil
	file_testdata_proto_goTypes = nil
	file_testdata_proto_depIdxs = nil
}
