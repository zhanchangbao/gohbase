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
// source: Snapshot.proto

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

type SnapshotDescription_Type int32

const (
	SnapshotDescription_DISABLED  SnapshotDescription_Type = 0
	SnapshotDescription_FLUSH     SnapshotDescription_Type = 1
	SnapshotDescription_SKIPFLUSH SnapshotDescription_Type = 2
)

// Enum value maps for SnapshotDescription_Type.
var (
	SnapshotDescription_Type_name = map[int32]string{
		0: "DISABLED",
		1: "FLUSH",
		2: "SKIPFLUSH",
	}
	SnapshotDescription_Type_value = map[string]int32{
		"DISABLED":  0,
		"FLUSH":     1,
		"SKIPFLUSH": 2,
	}
)

func (x SnapshotDescription_Type) Enum() *SnapshotDescription_Type {
	p := new(SnapshotDescription_Type)
	*p = x
	return p
}

func (x SnapshotDescription_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SnapshotDescription_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_Snapshot_proto_enumTypes[0].Descriptor()
}

func (SnapshotDescription_Type) Type() protoreflect.EnumType {
	return &file_Snapshot_proto_enumTypes[0]
}

func (x SnapshotDescription_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SnapshotDescription_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SnapshotDescription_Type(num)
	return nil
}

// Deprecated: Use SnapshotDescription_Type.Descriptor instead.
func (SnapshotDescription_Type) EnumDescriptor() ([]byte, []int) {
	return file_Snapshot_proto_rawDescGZIP(), []int{0, 0}
}

type SnapshotFileInfo_Type int32

const (
	SnapshotFileInfo_HFILE SnapshotFileInfo_Type = 1
	SnapshotFileInfo_WAL   SnapshotFileInfo_Type = 2
)

// Enum value maps for SnapshotFileInfo_Type.
var (
	SnapshotFileInfo_Type_name = map[int32]string{
		1: "HFILE",
		2: "WAL",
	}
	SnapshotFileInfo_Type_value = map[string]int32{
		"HFILE": 1,
		"WAL":   2,
	}
)

func (x SnapshotFileInfo_Type) Enum() *SnapshotFileInfo_Type {
	p := new(SnapshotFileInfo_Type)
	*p = x
	return p
}

func (x SnapshotFileInfo_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SnapshotFileInfo_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_Snapshot_proto_enumTypes[1].Descriptor()
}

func (SnapshotFileInfo_Type) Type() protoreflect.EnumType {
	return &file_Snapshot_proto_enumTypes[1]
}

func (x SnapshotFileInfo_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *SnapshotFileInfo_Type) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = SnapshotFileInfo_Type(num)
	return nil
}

// Deprecated: Use SnapshotFileInfo_Type.Descriptor instead.
func (SnapshotFileInfo_Type) EnumDescriptor() ([]byte, []int) {
	return file_Snapshot_proto_rawDescGZIP(), []int{1, 0}
}

//*
// Description of the snapshot to take
type SnapshotDescription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name                *string                   `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Table               *string                   `protobuf:"bytes,2,opt,name=table" json:"table,omitempty"` // not needed for delete, but checked for in taking snapshot
	CreationTime        *int64                    `protobuf:"varint,3,opt,name=creation_time,json=creationTime,def=0" json:"creation_time,omitempty"`
	Type                *SnapshotDescription_Type `protobuf:"varint,4,opt,name=type,enum=hbase.pb.SnapshotDescription_Type,def=1" json:"type,omitempty"`
	Version             *int32                    `protobuf:"varint,5,opt,name=version" json:"version,omitempty"`
	Owner               *string                   `protobuf:"bytes,6,opt,name=owner" json:"owner,omitempty"`
	UsersAndPermissions *UsersAndPermissions      `protobuf:"bytes,7,opt,name=users_and_permissions,json=usersAndPermissions" json:"users_and_permissions,omitempty"`
}

// Default values for SnapshotDescription fields.
const (
	Default_SnapshotDescription_CreationTime = int64(0)
	Default_SnapshotDescription_Type         = SnapshotDescription_FLUSH
)

func (x *SnapshotDescription) Reset() {
	*x = SnapshotDescription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Snapshot_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotDescription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotDescription) ProtoMessage() {}

func (x *SnapshotDescription) ProtoReflect() protoreflect.Message {
	mi := &file_Snapshot_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotDescription.ProtoReflect.Descriptor instead.
func (*SnapshotDescription) Descriptor() ([]byte, []int) {
	return file_Snapshot_proto_rawDescGZIP(), []int{0}
}

func (x *SnapshotDescription) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *SnapshotDescription) GetTable() string {
	if x != nil && x.Table != nil {
		return *x.Table
	}
	return ""
}

func (x *SnapshotDescription) GetCreationTime() int64 {
	if x != nil && x.CreationTime != nil {
		return *x.CreationTime
	}
	return Default_SnapshotDescription_CreationTime
}

func (x *SnapshotDescription) GetType() SnapshotDescription_Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return Default_SnapshotDescription_Type
}

func (x *SnapshotDescription) GetVersion() int32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *SnapshotDescription) GetOwner() string {
	if x != nil && x.Owner != nil {
		return *x.Owner
	}
	return ""
}

func (x *SnapshotDescription) GetUsersAndPermissions() *UsersAndPermissions {
	if x != nil {
		return x.UsersAndPermissions
	}
	return nil
}

type SnapshotFileInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      *SnapshotFileInfo_Type `protobuf:"varint,1,req,name=type,enum=hbase.pb.SnapshotFileInfo_Type" json:"type,omitempty"`
	Hfile     *string                `protobuf:"bytes,3,opt,name=hfile" json:"hfile,omitempty"`
	WalServer *string                `protobuf:"bytes,4,opt,name=wal_server,json=walServer" json:"wal_server,omitempty"`
	WalName   *string                `protobuf:"bytes,5,opt,name=wal_name,json=walName" json:"wal_name,omitempty"`
}

func (x *SnapshotFileInfo) Reset() {
	*x = SnapshotFileInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Snapshot_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotFileInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotFileInfo) ProtoMessage() {}

func (x *SnapshotFileInfo) ProtoReflect() protoreflect.Message {
	mi := &file_Snapshot_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotFileInfo.ProtoReflect.Descriptor instead.
func (*SnapshotFileInfo) Descriptor() ([]byte, []int) {
	return file_Snapshot_proto_rawDescGZIP(), []int{1}
}

func (x *SnapshotFileInfo) GetType() SnapshotFileInfo_Type {
	if x != nil && x.Type != nil {
		return *x.Type
	}
	return SnapshotFileInfo_HFILE
}

func (x *SnapshotFileInfo) GetHfile() string {
	if x != nil && x.Hfile != nil {
		return *x.Hfile
	}
	return ""
}

func (x *SnapshotFileInfo) GetWalServer() string {
	if x != nil && x.WalServer != nil {
		return *x.WalServer
	}
	return ""
}

func (x *SnapshotFileInfo) GetWalName() string {
	if x != nil && x.WalName != nil {
		return *x.WalName
	}
	return ""
}

type SnapshotRegionManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     *int32                                `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	RegionInfo  *RegionInfo                           `protobuf:"bytes,2,req,name=region_info,json=regionInfo" json:"region_info,omitempty"`
	FamilyFiles []*SnapshotRegionManifest_FamilyFiles `protobuf:"bytes,3,rep,name=family_files,json=familyFiles" json:"family_files,omitempty"`
}

func (x *SnapshotRegionManifest) Reset() {
	*x = SnapshotRegionManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Snapshot_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotRegionManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotRegionManifest) ProtoMessage() {}

func (x *SnapshotRegionManifest) ProtoReflect() protoreflect.Message {
	mi := &file_Snapshot_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotRegionManifest.ProtoReflect.Descriptor instead.
func (*SnapshotRegionManifest) Descriptor() ([]byte, []int) {
	return file_Snapshot_proto_rawDescGZIP(), []int{2}
}

func (x *SnapshotRegionManifest) GetVersion() int32 {
	if x != nil && x.Version != nil {
		return *x.Version
	}
	return 0
}

func (x *SnapshotRegionManifest) GetRegionInfo() *RegionInfo {
	if x != nil {
		return x.RegionInfo
	}
	return nil
}

func (x *SnapshotRegionManifest) GetFamilyFiles() []*SnapshotRegionManifest_FamilyFiles {
	if x != nil {
		return x.FamilyFiles
	}
	return nil
}

type SnapshotDataManifest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TableSchema     *TableSchema              `protobuf:"bytes,1,req,name=table_schema,json=tableSchema" json:"table_schema,omitempty"`
	RegionManifests []*SnapshotRegionManifest `protobuf:"bytes,2,rep,name=region_manifests,json=regionManifests" json:"region_manifests,omitempty"`
}

func (x *SnapshotDataManifest) Reset() {
	*x = SnapshotDataManifest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Snapshot_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotDataManifest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotDataManifest) ProtoMessage() {}

func (x *SnapshotDataManifest) ProtoReflect() protoreflect.Message {
	mi := &file_Snapshot_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotDataManifest.ProtoReflect.Descriptor instead.
func (*SnapshotDataManifest) Descriptor() ([]byte, []int) {
	return file_Snapshot_proto_rawDescGZIP(), []int{3}
}

func (x *SnapshotDataManifest) GetTableSchema() *TableSchema {
	if x != nil {
		return x.TableSchema
	}
	return nil
}

func (x *SnapshotDataManifest) GetRegionManifests() []*SnapshotRegionManifest {
	if x != nil {
		return x.RegionManifests
	}
	return nil
}

type SnapshotRegionManifest_StoreFile struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      *string    `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Reference *Reference `protobuf:"bytes,2,opt,name=reference" json:"reference,omitempty"`
	// TODO: Add checksums or other fields to verify the file
	FileSize *uint64 `protobuf:"varint,3,opt,name=file_size,json=fileSize" json:"file_size,omitempty"`
}

func (x *SnapshotRegionManifest_StoreFile) Reset() {
	*x = SnapshotRegionManifest_StoreFile{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Snapshot_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotRegionManifest_StoreFile) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotRegionManifest_StoreFile) ProtoMessage() {}

func (x *SnapshotRegionManifest_StoreFile) ProtoReflect() protoreflect.Message {
	mi := &file_Snapshot_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotRegionManifest_StoreFile.ProtoReflect.Descriptor instead.
func (*SnapshotRegionManifest_StoreFile) Descriptor() ([]byte, []int) {
	return file_Snapshot_proto_rawDescGZIP(), []int{2, 0}
}

func (x *SnapshotRegionManifest_StoreFile) GetName() string {
	if x != nil && x.Name != nil {
		return *x.Name
	}
	return ""
}

func (x *SnapshotRegionManifest_StoreFile) GetReference() *Reference {
	if x != nil {
		return x.Reference
	}
	return nil
}

func (x *SnapshotRegionManifest_StoreFile) GetFileSize() uint64 {
	if x != nil && x.FileSize != nil {
		return *x.FileSize
	}
	return 0
}

type SnapshotRegionManifest_FamilyFiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FamilyName []byte                              `protobuf:"bytes,1,req,name=family_name,json=familyName" json:"family_name,omitempty"`
	StoreFiles []*SnapshotRegionManifest_StoreFile `protobuf:"bytes,2,rep,name=store_files,json=storeFiles" json:"store_files,omitempty"`
}

func (x *SnapshotRegionManifest_FamilyFiles) Reset() {
	*x = SnapshotRegionManifest_FamilyFiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Snapshot_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SnapshotRegionManifest_FamilyFiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SnapshotRegionManifest_FamilyFiles) ProtoMessage() {}

func (x *SnapshotRegionManifest_FamilyFiles) ProtoReflect() protoreflect.Message {
	mi := &file_Snapshot_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SnapshotRegionManifest_FamilyFiles.ProtoReflect.Descriptor instead.
func (*SnapshotRegionManifest_FamilyFiles) Descriptor() ([]byte, []int) {
	return file_Snapshot_proto_rawDescGZIP(), []int{2, 1}
}

func (x *SnapshotRegionManifest_FamilyFiles) GetFamilyName() []byte {
	if x != nil {
		return x.FamilyName
	}
	return nil
}

func (x *SnapshotRegionManifest_FamilyFiles) GetStoreFiles() []*SnapshotRegionManifest_StoreFile {
	if x != nil {
		return x.StoreFiles
	}
	return nil
}

var File_Snapshot_proto protoreflect.FileDescriptor

var file_Snapshot_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x1a, 0x13, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x08, 0x46, 0x53, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0b, 0x48, 0x42, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd9, 0x02, 0x0a, 0x13, 0x53, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x26, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x3a,
	0x01, 0x30, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x3d, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22,
	0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x79,
	0x70, 0x65, 0x3a, 0x05, 0x46, 0x4c, 0x55, 0x53, 0x48, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12,
	0x51, 0x0a, 0x15, 0x75, 0x73, 0x65, 0x72, 0x73, 0x5f, 0x61, 0x6e, 0x64, 0x5f, 0x70, 0x65, 0x72,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d,
	0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x41,
	0x6e, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x13, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x41, 0x6e, 0x64, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0x2e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49,
	0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x46, 0x4c, 0x55, 0x53,
	0x48, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x53, 0x4b, 0x49, 0x50, 0x46, 0x4c, 0x55, 0x53, 0x48,
	0x10, 0x02, 0x22, 0xb3, 0x01, 0x0a, 0x10, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x46,
	0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x33, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x1f, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62,
	0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x49, 0x6e, 0x66,
	0x6f, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x68, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x68, 0x66, 0x69,
	0x6c, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x61, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x77, 0x61, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x12, 0x19, 0x0a, 0x08, 0x77, 0x61, 0x6c, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x77, 0x61, 0x6c, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x1a, 0x0a, 0x04,
	0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x48, 0x46, 0x49, 0x4c, 0x45, 0x10, 0x01, 0x12,
	0x07, 0x0a, 0x03, 0x57, 0x41, 0x4c, 0x10, 0x02, 0x22, 0xa8, 0x03, 0x0a, 0x16, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x69, 0x66,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x35, 0x0a,
	0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x02,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x4f, 0x0a, 0x0c, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x68, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x46, 0x61, 0x6d,
	0x69, 0x6c, 0x79, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x52, 0x0b, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x1a, 0x6f, 0x0a, 0x09, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x31, 0x0a, 0x09, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x62, 0x61, 0x73,
	0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x52, 0x09,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x69, 0x6c,
	0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x66, 0x69,
	0x6c, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x1a, 0x7b, 0x0a, 0x0b, 0x46, 0x61, 0x6d, 0x69, 0x6c, 0x79,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x61, 0x6d, 0x69, 0x6c, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0c, 0x52, 0x0a, 0x66, 0x61, 0x6d, 0x69,
	0x6c, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x4b, 0x0a, 0x0b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x68, 0x62,
	0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x2e, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x52, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x22, 0x9d, 0x01, 0x0a, 0x14, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c,
	0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x52, 0x0b, 0x74, 0x61, 0x62, 0x6c, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x6d, 0x61, 0x12, 0x4b, 0x0a, 0x10, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e,
	0x5f, 0x6d, 0x61, 0x6e, 0x69, 0x66, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x52, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65,
	0x73, 0x74, 0x52, 0x0f, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6e, 0x69, 0x66, 0x65,
	0x73, 0x74, 0x73, 0x42, 0x52, 0x0a, 0x31, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68,
	0x65, 0x2e, 0x68, 0x61, 0x64, 0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73,
	0x68, 0x61, 0x64, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x42, 0x0e, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70,
	0x62, 0x88, 0x01, 0x01, 0xa0, 0x01, 0x01,
}

var (
	file_Snapshot_proto_rawDescOnce sync.Once
	file_Snapshot_proto_rawDescData = file_Snapshot_proto_rawDesc
)

func file_Snapshot_proto_rawDescGZIP() []byte {
	file_Snapshot_proto_rawDescOnce.Do(func() {
		file_Snapshot_proto_rawDescData = protoimpl.X.CompressGZIP(file_Snapshot_proto_rawDescData)
	})
	return file_Snapshot_proto_rawDescData
}

var file_Snapshot_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_Snapshot_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_Snapshot_proto_goTypes = []interface{}{
	(SnapshotDescription_Type)(0),              // 0: hbase.pb.SnapshotDescription.Type
	(SnapshotFileInfo_Type)(0),                 // 1: hbase.pb.SnapshotFileInfo.Type
	(*SnapshotDescription)(nil),                // 2: hbase.pb.SnapshotDescription
	(*SnapshotFileInfo)(nil),                   // 3: hbase.pb.SnapshotFileInfo
	(*SnapshotRegionManifest)(nil),             // 4: hbase.pb.SnapshotRegionManifest
	(*SnapshotDataManifest)(nil),               // 5: hbase.pb.SnapshotDataManifest
	(*SnapshotRegionManifest_StoreFile)(nil),   // 6: hbase.pb.SnapshotRegionManifest.StoreFile
	(*SnapshotRegionManifest_FamilyFiles)(nil), // 7: hbase.pb.SnapshotRegionManifest.FamilyFiles
	(*UsersAndPermissions)(nil),                // 8: hbase.pb.UsersAndPermissions
	(*RegionInfo)(nil),                         // 9: hbase.pb.RegionInfo
	(*TableSchema)(nil),                        // 10: hbase.pb.TableSchema
	(*Reference)(nil),                          // 11: hbase.pb.Reference
}
var file_Snapshot_proto_depIdxs = []int32{
	0,  // 0: hbase.pb.SnapshotDescription.type:type_name -> hbase.pb.SnapshotDescription.Type
	8,  // 1: hbase.pb.SnapshotDescription.users_and_permissions:type_name -> hbase.pb.UsersAndPermissions
	1,  // 2: hbase.pb.SnapshotFileInfo.type:type_name -> hbase.pb.SnapshotFileInfo.Type
	9,  // 3: hbase.pb.SnapshotRegionManifest.region_info:type_name -> hbase.pb.RegionInfo
	7,  // 4: hbase.pb.SnapshotRegionManifest.family_files:type_name -> hbase.pb.SnapshotRegionManifest.FamilyFiles
	10, // 5: hbase.pb.SnapshotDataManifest.table_schema:type_name -> hbase.pb.TableSchema
	4,  // 6: hbase.pb.SnapshotDataManifest.region_manifests:type_name -> hbase.pb.SnapshotRegionManifest
	11, // 7: hbase.pb.SnapshotRegionManifest.StoreFile.reference:type_name -> hbase.pb.Reference
	6,  // 8: hbase.pb.SnapshotRegionManifest.FamilyFiles.store_files:type_name -> hbase.pb.SnapshotRegionManifest.StoreFile
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_Snapshot_proto_init() }
func file_Snapshot_proto_init() {
	if File_Snapshot_proto != nil {
		return
	}
	file_AccessControl_proto_init()
	file_FS_proto_init()
	file_HBase_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_Snapshot_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotDescription); i {
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
		file_Snapshot_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotFileInfo); i {
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
		file_Snapshot_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotRegionManifest); i {
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
		file_Snapshot_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotDataManifest); i {
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
		file_Snapshot_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotRegionManifest_StoreFile); i {
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
		file_Snapshot_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SnapshotRegionManifest_FamilyFiles); i {
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
			RawDescriptor: file_Snapshot_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_Snapshot_proto_goTypes,
		DependencyIndexes: file_Snapshot_proto_depIdxs,
		EnumInfos:         file_Snapshot_proto_enumTypes,
		MessageInfos:      file_Snapshot_proto_msgTypes,
	}.Build()
	File_Snapshot_proto = out.File
	file_Snapshot_proto_rawDesc = nil
	file_Snapshot_proto_goTypes = nil
	file_Snapshot_proto_depIdxs = nil
}
