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
// source: MapReduce.proto

//This file includes protocol buffers used in MapReduce only.

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

type ScanMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metrics []*NameInt64Pair `protobuf:"bytes,1,rep,name=metrics" json:"metrics,omitempty"`
}

func (x *ScanMetrics) Reset() {
	*x = ScanMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MapReduce_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ScanMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ScanMetrics) ProtoMessage() {}

func (x *ScanMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_MapReduce_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ScanMetrics.ProtoReflect.Descriptor instead.
func (*ScanMetrics) Descriptor() ([]byte, []int) {
	return file_MapReduce_proto_rawDescGZIP(), []int{0}
}

func (x *ScanMetrics) GetMetrics() []*NameInt64Pair {
	if x != nil {
		return x.Metrics
	}
	return nil
}

type TableSnapshotRegionSplit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locations []string     `protobuf:"bytes,2,rep,name=locations" json:"locations,omitempty"`
	Table     *TableSchema `protobuf:"bytes,3,opt,name=table" json:"table,omitempty"`
	Region    *RegionInfo  `protobuf:"bytes,4,opt,name=region" json:"region,omitempty"`
}

func (x *TableSnapshotRegionSplit) Reset() {
	*x = TableSnapshotRegionSplit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_MapReduce_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TableSnapshotRegionSplit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TableSnapshotRegionSplit) ProtoMessage() {}

func (x *TableSnapshotRegionSplit) ProtoReflect() protoreflect.Message {
	mi := &file_MapReduce_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TableSnapshotRegionSplit.ProtoReflect.Descriptor instead.
func (*TableSnapshotRegionSplit) Descriptor() ([]byte, []int) {
	return file_MapReduce_proto_rawDescGZIP(), []int{1}
}

func (x *TableSnapshotRegionSplit) GetLocations() []string {
	if x != nil {
		return x.Locations
	}
	return nil
}

func (x *TableSnapshotRegionSplit) GetTable() *TableSchema {
	if x != nil {
		return x.Table
	}
	return nil
}

func (x *TableSnapshotRegionSplit) GetRegion() *RegionInfo {
	if x != nil {
		return x.Region
	}
	return nil
}

var File_MapReduce_proto protoreflect.FileDescriptor

var file_MapReduce_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x1a, 0x0b, 0x48, 0x42, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x40, 0x0a, 0x0b, 0x53, 0x63, 0x61, 0x6e,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x31, 0x0a, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x62, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x50, 0x61, 0x69,
	0x72, 0x52, 0x07, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x93, 0x01, 0x0a, 0x18, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x67, 0x69,
	0x6f, 0x6e, 0x53, 0x70, 0x6c, 0x69, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x2b, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x05, 0x74, 0x61, 0x62,
	0x6c, 0x65, 0x12, 0x2c, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x42, 0x50, 0x0a, 0x31, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x68,
	0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x68, 0x61, 0x64,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x0f, 0x4d, 0x61, 0x70, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0xa0,
	0x01, 0x01,
}

var (
	file_MapReduce_proto_rawDescOnce sync.Once
	file_MapReduce_proto_rawDescData = file_MapReduce_proto_rawDesc
)

func file_MapReduce_proto_rawDescGZIP() []byte {
	file_MapReduce_proto_rawDescOnce.Do(func() {
		file_MapReduce_proto_rawDescData = protoimpl.X.CompressGZIP(file_MapReduce_proto_rawDescData)
	})
	return file_MapReduce_proto_rawDescData
}

var file_MapReduce_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_MapReduce_proto_goTypes = []interface{}{
	(*ScanMetrics)(nil),              // 0: hbase.pb.ScanMetrics
	(*TableSnapshotRegionSplit)(nil), // 1: hbase.pb.TableSnapshotRegionSplit
	(*NameInt64Pair)(nil),            // 2: hbase.pb.NameInt64Pair
	(*TableSchema)(nil),              // 3: hbase.pb.TableSchema
	(*RegionInfo)(nil),               // 4: hbase.pb.RegionInfo
}
var file_MapReduce_proto_depIdxs = []int32{
	2, // 0: hbase.pb.ScanMetrics.metrics:type_name -> hbase.pb.NameInt64Pair
	3, // 1: hbase.pb.TableSnapshotRegionSplit.table:type_name -> hbase.pb.TableSchema
	4, // 2: hbase.pb.TableSnapshotRegionSplit.region:type_name -> hbase.pb.RegionInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_MapReduce_proto_init() }
func file_MapReduce_proto_init() {
	if File_MapReduce_proto != nil {
		return
	}
	file_HBase_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_MapReduce_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ScanMetrics); i {
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
		file_MapReduce_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TableSnapshotRegionSplit); i {
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
			RawDescriptor: file_MapReduce_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_MapReduce_proto_goTypes,
		DependencyIndexes: file_MapReduce_proto_depIdxs,
		MessageInfos:      file_MapReduce_proto_msgTypes,
	}.Build()
	File_MapReduce_proto = out.File
	file_MapReduce_proto_rawDesc = nil
	file_MapReduce_proto_goTypes = nil
	file_MapReduce_proto_depIdxs = nil
}
