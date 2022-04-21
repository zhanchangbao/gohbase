//*
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
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
// 	protoc-gen-go v1.26.0-rc.1
// 	protoc        v3.17.3
// source: Comparator.proto

// This file contains protocol buffers that are used for filters

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

type BitComparator_BitwiseOp int32

const (
	BitComparator_AND BitComparator_BitwiseOp = 1
	BitComparator_OR  BitComparator_BitwiseOp = 2
	BitComparator_XOR BitComparator_BitwiseOp = 3
)

// Enum value maps for BitComparator_BitwiseOp.
var (
	BitComparator_BitwiseOp_name = map[int32]string{
		1: "AND",
		2: "OR",
		3: "XOR",
	}
	BitComparator_BitwiseOp_value = map[string]int32{
		"AND": 1,
		"OR":  2,
		"XOR": 3,
	}
)

func (x BitComparator_BitwiseOp) Enum() *BitComparator_BitwiseOp {
	p := new(BitComparator_BitwiseOp)
	*p = x
	return p
}

func (x BitComparator_BitwiseOp) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BitComparator_BitwiseOp) Descriptor() protoreflect.EnumDescriptor {
	return file_Comparator_proto_enumTypes[0].Descriptor()
}

func (BitComparator_BitwiseOp) Type() protoreflect.EnumType {
	return &file_Comparator_proto_enumTypes[0]
}

func (x BitComparator_BitwiseOp) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *BitComparator_BitwiseOp) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = BitComparator_BitwiseOp(num)
	return nil
}

// Deprecated: Use BitComparator_BitwiseOp.Descriptor instead.
func (BitComparator_BitwiseOp) EnumDescriptor() ([]byte, []int) {
	return file_Comparator_proto_rawDescGZIP(), []int{5, 0}
}

type Comparator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                 *string `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	SerializedComparator []byte  `protobuf:"bytes,2,opt,name=serialized_comparator,json=serializedComparator" json:"serialized_comparator,omitempty"`
}

func (x *Comparator) Reset() {
	*x = Comparator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Comparator_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comparator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comparator) ProtoMessage() {}

func (x *Comparator) ProtoReflect() protoreflect.Message {
	mi := &file_Comparator_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comparator.ProtoReflect.Descriptor instead.
func (*Comparator) Descriptor() ([]byte, []int) {
	return file_Comparator_proto_rawDescGZIP(), []int{0}
}

func (x *Comparator) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *Comparator) GetSerializedComparator() []byte {
	if x != nil {
		return x.SerializedComparator
	}
	return nil
}

type ByteArrayComparable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value []byte `protobuf:"bytes,1,opt,name=value" json:"value,omitempty"`
}

func (x *ByteArrayComparable) Reset() {
	*x = ByteArrayComparable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Comparator_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByteArrayComparable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByteArrayComparable) ProtoMessage() {}

func (x *ByteArrayComparable) ProtoReflect() protoreflect.Message {
	mi := &file_Comparator_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByteArrayComparable.ProtoReflect.Descriptor instead.
func (*ByteArrayComparable) Descriptor() ([]byte, []int) {
	return file_Comparator_proto_rawDescGZIP(), []int{1}
}

func (x *ByteArrayComparable) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type BinaryComparator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comparable *ByteArrayComparable `protobuf:"bytes,1,req,name=comparable" json:"comparable,omitempty"`
}

func (x *BinaryComparator) Reset() {
	*x = BinaryComparator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Comparator_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryComparator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryComparator) ProtoMessage() {}

func (x *BinaryComparator) ProtoReflect() protoreflect.Message {
	mi := &file_Comparator_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryComparator.ProtoReflect.Descriptor instead.
func (*BinaryComparator) Descriptor() ([]byte, []int) {
	return file_Comparator_proto_rawDescGZIP(), []int{2}
}

func (x *BinaryComparator) GetComparable() *ByteArrayComparable {
	if x != nil {
		return x.Comparable
	}
	return nil
}

type LongComparator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comparable *ByteArrayComparable `protobuf:"bytes,1,req,name=comparable" json:"comparable,omitempty"`
}

func (x *LongComparator) Reset() {
	*x = LongComparator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Comparator_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LongComparator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LongComparator) ProtoMessage() {}

func (x *LongComparator) ProtoReflect() protoreflect.Message {
	mi := &file_Comparator_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LongComparator.ProtoReflect.Descriptor instead.
func (*LongComparator) Descriptor() ([]byte, []int) {
	return file_Comparator_proto_rawDescGZIP(), []int{3}
}

func (x *LongComparator) GetComparable() *ByteArrayComparable {
	if x != nil {
		return x.Comparable
	}
	return nil
}

type BinaryPrefixComparator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comparable *ByteArrayComparable `protobuf:"bytes,1,req,name=comparable" json:"comparable,omitempty"`
}

func (x *BinaryPrefixComparator) Reset() {
	*x = BinaryPrefixComparator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Comparator_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BinaryPrefixComparator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BinaryPrefixComparator) ProtoMessage() {}

func (x *BinaryPrefixComparator) ProtoReflect() protoreflect.Message {
	mi := &file_Comparator_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BinaryPrefixComparator.ProtoReflect.Descriptor instead.
func (*BinaryPrefixComparator) Descriptor() ([]byte, []int) {
	return file_Comparator_proto_rawDescGZIP(), []int{4}
}

func (x *BinaryPrefixComparator) GetComparable() *ByteArrayComparable {
	if x != nil {
		return x.Comparable
	}
	return nil
}

type BitComparator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comparable *ByteArrayComparable     `protobuf:"bytes,1,req,name=comparable" json:"comparable,omitempty"`
	BitwiseOp  *BitComparator_BitwiseOp `protobuf:"varint,2,req,name=bitwise_op,json=bitwiseOp,enum=hbase.pb.BitComparator_BitwiseOp" json:"bitwise_op,omitempty"`
}

func (x *BitComparator) Reset() {
	*x = BitComparator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Comparator_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BitComparator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BitComparator) ProtoMessage() {}

func (x *BitComparator) ProtoReflect() protoreflect.Message {
	mi := &file_Comparator_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BitComparator.ProtoReflect.Descriptor instead.
func (*BitComparator) Descriptor() ([]byte, []int) {
	return file_Comparator_proto_rawDescGZIP(), []int{5}
}

func (x *BitComparator) GetComparable() *ByteArrayComparable {
	if x != nil {
		return x.Comparable
	}
	return nil
}

func (x *BitComparator) GetBitwiseOp() BitComparator_BitwiseOp {
	if x != nil && x.BitwiseOp != nil {
		return *x.BitwiseOp
	}
	return BitComparator_AND
}

type NullComparator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NullComparator) Reset() {
	*x = NullComparator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Comparator_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NullComparator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NullComparator) ProtoMessage() {}

func (x *NullComparator) ProtoReflect() protoreflect.Message {
	mi := &file_Comparator_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NullComparator.ProtoReflect.Descriptor instead.
func (*NullComparator) Descriptor() ([]byte, []int) {
	return file_Comparator_proto_rawDescGZIP(), []int{6}
}

type RegexStringComparator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pattern      *string `protobuf:"bytes,1,req,name=pattern" json:"pattern,omitempty"`
	PatternFlags *int32  `protobuf:"varint,2,req,name=pattern_flags,json=patternFlags" json:"pattern_flags,omitempty"`
	Charset      *string `protobuf:"bytes,3,req,name=charset" json:"charset,omitempty"`
	Engine       *string `protobuf:"bytes,4,opt,name=engine" json:"engine,omitempty"`
}

func (x *RegexStringComparator) Reset() {
	*x = RegexStringComparator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Comparator_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegexStringComparator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegexStringComparator) ProtoMessage() {}

func (x *RegexStringComparator) ProtoReflect() protoreflect.Message {
	mi := &file_Comparator_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegexStringComparator.ProtoReflect.Descriptor instead.
func (*RegexStringComparator) Descriptor() ([]byte, []int) {
	return file_Comparator_proto_rawDescGZIP(), []int{7}
}

func (x *RegexStringComparator) GetPattern() string {
	if x != nil && x.Pattern != nil {
		return *x.Pattern
	}
	return ""
}

func (x *RegexStringComparator) GetPatternFlags() int32 {
	if x != nil && x.PatternFlags != nil {
		return *x.PatternFlags
	}
	return 0
}

func (x *RegexStringComparator) GetCharset() string {
	if x != nil && x.Charset != nil {
		return *x.Charset
	}
	return ""
}

func (x *RegexStringComparator) GetEngine() string {
	if x != nil && x.Engine != nil {
		return *x.Engine
	}
	return ""
}

type SubstringComparator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Substr *string `protobuf:"bytes,1,req,name=substr" json:"substr,omitempty"`
}

func (x *SubstringComparator) Reset() {
	*x = SubstringComparator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Comparator_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubstringComparator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubstringComparator) ProtoMessage() {}

func (x *SubstringComparator) ProtoReflect() protoreflect.Message {
	mi := &file_Comparator_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubstringComparator.ProtoReflect.Descriptor instead.
func (*SubstringComparator) Descriptor() ([]byte, []int) {
	return file_Comparator_proto_rawDescGZIP(), []int{8}
}

func (x *SubstringComparator) GetSubstr() string {
	if x != nil && x.Substr != nil {
		return *x.Substr
	}
	return ""
}

type BigDecimalComparator struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comparable *ByteArrayComparable `protobuf:"bytes,1,req,name=comparable" json:"comparable,omitempty"`
}

func (x *BigDecimalComparator) Reset() {
	*x = BigDecimalComparator{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Comparator_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BigDecimalComparator) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BigDecimalComparator) ProtoMessage() {}

func (x *BigDecimalComparator) ProtoReflect() protoreflect.Message {
	mi := &file_Comparator_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BigDecimalComparator.ProtoReflect.Descriptor instead.
func (*BigDecimalComparator) Descriptor() ([]byte, []int) {
	return file_Comparator_proto_rawDescGZIP(), []int{9}
}

func (x *BigDecimalComparator) GetComparable() *ByteArrayComparable {
	if x != nil {
		return x.Comparable
	}
	return nil
}

var File_Comparator_proto protoreflect.FileDescriptor

var file_Comparator_proto_rawDesc = []byte{
	0x0a, 0x10, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x08, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x22, 0x55, 0x0a, 0x0a,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x33,
	0x0a, 0x15, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x5f, 0x63, 0x6f, 0x6d,
	0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x14, 0x73,
	0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x65, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x22, 0x2b, 0x0a, 0x13, 0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x22, 0x51, 0x0a, 0x10, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72,
	0x61, 0x74, 0x6f, 0x72, 0x12, 0x3d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x62,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x62, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79, 0x43, 0x6f, 0x6d,
	0x70, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61,
	0x62, 0x6c, 0x65, 0x22, 0x4f, 0x0a, 0x0e, 0x4c, 0x6f, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x3d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61,
	0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72,
	0x61, 0x62, 0x6c, 0x65, 0x22, 0x57, 0x0a, 0x16, 0x42, 0x69, 0x6e, 0x61, 0x72, 0x79, 0x50, 0x72,
	0x65, 0x66, 0x69, 0x78, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x3d,
	0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x79,
	0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x62, 0x6c,
	0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x22, 0xb7, 0x01,
	0x0a, 0x0d, 0x42, 0x69, 0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12,
	0x3d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20,
	0x02, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x42,
	0x79, 0x74, 0x65, 0x41, 0x72, 0x72, 0x61, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x62,
	0x6c, 0x65, 0x52, 0x0a, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x40,
	0x0a, 0x0a, 0x62, 0x69, 0x74, 0x77, 0x69, 0x73, 0x65, 0x5f, 0x6f, 0x70, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x0e, 0x32, 0x21, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x69,
	0x74, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x2e, 0x42, 0x69, 0x74, 0x77,
	0x69, 0x73, 0x65, 0x4f, 0x70, 0x52, 0x09, 0x62, 0x69, 0x74, 0x77, 0x69, 0x73, 0x65, 0x4f, 0x70,
	0x22, 0x25, 0x0a, 0x09, 0x42, 0x69, 0x74, 0x77, 0x69, 0x73, 0x65, 0x4f, 0x70, 0x12, 0x07, 0x0a,
	0x03, 0x41, 0x4e, 0x44, 0x10, 0x01, 0x12, 0x06, 0x0a, 0x02, 0x4f, 0x52, 0x10, 0x02, 0x12, 0x07,
	0x0a, 0x03, 0x58, 0x4f, 0x52, 0x10, 0x03, 0x22, 0x10, 0x0a, 0x0e, 0x4e, 0x75, 0x6c, 0x6c, 0x43,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x22, 0x88, 0x01, 0x0a, 0x15, 0x52, 0x65,
	0x67, 0x65, 0x78, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61,
	0x74, 0x6f, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x5f, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x02,
	0x20, 0x02, 0x28, 0x05, 0x52, 0x0c, 0x70, 0x61, 0x74, 0x74, 0x65, 0x72, 0x6e, 0x46, 0x6c, 0x61,
	0x67, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x18, 0x03, 0x20,
	0x02, 0x28, 0x09, 0x52, 0x07, 0x63, 0x68, 0x61, 0x72, 0x73, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x65, 0x6e,
	0x67, 0x69, 0x6e, 0x65, 0x22, 0x2d, 0x0a, 0x13, 0x53, 0x75, 0x62, 0x73, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x75, 0x62, 0x73, 0x74, 0x72, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62,
	0x73, 0x74, 0x72, 0x22, 0x55, 0x0a, 0x14, 0x42, 0x69, 0x67, 0x44, 0x65, 0x63, 0x69, 0x6d, 0x61,
	0x6c, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x3d, 0x0a, 0x0a, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x0a,
	0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x62, 0x6c, 0x65, 0x42, 0x54, 0x0a, 0x31, 0x6f, 0x72,
	0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e,
	0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42,
	0x10, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x48, 0x01, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x88, 0x01, 0x01, 0xa0, 0x01, 0x01,
}

var (
	file_Comparator_proto_rawDescOnce sync.Once
	file_Comparator_proto_rawDescData = file_Comparator_proto_rawDesc
)

func file_Comparator_proto_rawDescGZIP() []byte {
	file_Comparator_proto_rawDescOnce.Do(func() {
		file_Comparator_proto_rawDescData = protoimpl.X.CompressGZIP(file_Comparator_proto_rawDescData)
	})
	return file_Comparator_proto_rawDescData
}

var file_Comparator_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Comparator_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_Comparator_proto_goTypes = []interface{}{
	(BitComparator_BitwiseOp)(0),   // 0: hbase.pb.BitComparator.BitwiseOp
	(*Comparator)(nil),             // 1: hbase.pb.Comparator
	(*ByteArrayComparable)(nil),    // 2: hbase.pb.ByteArrayComparable
	(*BinaryComparator)(nil),       // 3: hbase.pb.BinaryComparator
	(*LongComparator)(nil),         // 4: hbase.pb.LongComparator
	(*BinaryPrefixComparator)(nil), // 5: hbase.pb.BinaryPrefixComparator
	(*BitComparator)(nil),          // 6: hbase.pb.BitComparator
	(*NullComparator)(nil),         // 7: hbase.pb.NullComparator
	(*RegexStringComparator)(nil),  // 8: hbase.pb.RegexStringComparator
	(*SubstringComparator)(nil),    // 9: hbase.pb.SubstringComparator
	(*BigDecimalComparator)(nil),   // 10: hbase.pb.BigDecimalComparator
}
var file_Comparator_proto_depIdxs = []int32{
	2, // 0: hbase.pb.BinaryComparator.comparable:type_name -> hbase.pb.ByteArrayComparable
	2, // 1: hbase.pb.LongComparator.comparable:type_name -> hbase.pb.ByteArrayComparable
	2, // 2: hbase.pb.BinaryPrefixComparator.comparable:type_name -> hbase.pb.ByteArrayComparable
	2, // 3: hbase.pb.BitComparator.comparable:type_name -> hbase.pb.ByteArrayComparable
	0, // 4: hbase.pb.BitComparator.bitwise_op:type_name -> hbase.pb.BitComparator.BitwiseOp
	2, // 5: hbase.pb.BigDecimalComparator.comparable:type_name -> hbase.pb.ByteArrayComparable
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_Comparator_proto_init() }
func file_Comparator_proto_init() {
	if File_Comparator_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Comparator_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comparator); i {
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
		file_Comparator_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByteArrayComparable); i {
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
		file_Comparator_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryComparator); i {
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
		file_Comparator_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LongComparator); i {
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
		file_Comparator_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BinaryPrefixComparator); i {
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
		file_Comparator_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BitComparator); i {
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
		file_Comparator_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NullComparator); i {
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
		file_Comparator_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegexStringComparator); i {
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
		file_Comparator_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubstringComparator); i {
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
		file_Comparator_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BigDecimalComparator); i {
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
			RawDescriptor: file_Comparator_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Comparator_proto_goTypes,
		DependencyIndexes: file_Comparator_proto_depIdxs,
		EnumInfos:         file_Comparator_proto_enumTypes,
		MessageInfos:      file_Comparator_proto_msgTypes,
	}.Build()
	File_Comparator_proto = out.File
	file_Comparator_proto_rawDesc = nil
	file_Comparator_proto_goTypes = nil
	file_Comparator_proto_depIdxs = nil
}
