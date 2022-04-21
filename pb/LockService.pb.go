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
// source: LockService.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type LockType int32

const (
	LockType_EXCLUSIVE LockType = 1
	LockType_SHARED    LockType = 2
)

// Enum value maps for LockType.
var (
	LockType_name = map[int32]string{
		1: "EXCLUSIVE",
		2: "SHARED",
	}
	LockType_value = map[string]int32{
		"EXCLUSIVE": 1,
		"SHARED":    2,
	}
)

func (x LockType) Enum() *LockType {
	p := new(LockType)
	*p = x
	return p
}

func (x LockType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LockType) Descriptor() protoreflect.EnumDescriptor {
	return file_LockService_proto_enumTypes[0].Descriptor()
}

func (LockType) Type() protoreflect.EnumType {
	return &file_LockService_proto_enumTypes[0]
}

func (x LockType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *LockType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = LockType(num)
	return nil
}

// Deprecated: Use LockType.Descriptor instead.
func (LockType) EnumDescriptor() ([]byte, []int) {
	return file_LockService_proto_rawDescGZIP(), []int{0}
}

type LockedResourceType int32

const (
	LockedResourceType_SERVER    LockedResourceType = 1
	LockedResourceType_NAMESPACE LockedResourceType = 2
	LockedResourceType_TABLE     LockedResourceType = 3
	LockedResourceType_REGION    LockedResourceType = 4
	LockedResourceType_PEER      LockedResourceType = 5
)

// Enum value maps for LockedResourceType.
var (
	LockedResourceType_name = map[int32]string{
		1: "SERVER",
		2: "NAMESPACE",
		3: "TABLE",
		4: "REGION",
		5: "PEER",
	}
	LockedResourceType_value = map[string]int32{
		"SERVER":    1,
		"NAMESPACE": 2,
		"TABLE":     3,
		"REGION":    4,
		"PEER":      5,
	}
)

func (x LockedResourceType) Enum() *LockedResourceType {
	p := new(LockedResourceType)
	*p = x
	return p
}

func (x LockedResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LockedResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_LockService_proto_enumTypes[1].Descriptor()
}

func (LockedResourceType) Type() protoreflect.EnumType {
	return &file_LockService_proto_enumTypes[1]
}

func (x LockedResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *LockedResourceType) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = LockedResourceType(num)
	return nil
}

// Deprecated: Use LockedResourceType.Descriptor instead.
func (LockedResourceType) EnumDescriptor() ([]byte, []int) {
	return file_LockService_proto_rawDescGZIP(), []int{1}
}

type LockHeartbeatResponse_LockStatus int32

const (
	LockHeartbeatResponse_UNLOCKED LockHeartbeatResponse_LockStatus = 1
	LockHeartbeatResponse_LOCKED   LockHeartbeatResponse_LockStatus = 2
)

// Enum value maps for LockHeartbeatResponse_LockStatus.
var (
	LockHeartbeatResponse_LockStatus_name = map[int32]string{
		1: "UNLOCKED",
		2: "LOCKED",
	}
	LockHeartbeatResponse_LockStatus_value = map[string]int32{
		"UNLOCKED": 1,
		"LOCKED":   2,
	}
)

func (x LockHeartbeatResponse_LockStatus) Enum() *LockHeartbeatResponse_LockStatus {
	p := new(LockHeartbeatResponse_LockStatus)
	*p = x
	return p
}

func (x LockHeartbeatResponse_LockStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LockHeartbeatResponse_LockStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_LockService_proto_enumTypes[2].Descriptor()
}

func (LockHeartbeatResponse_LockStatus) Type() protoreflect.EnumType {
	return &file_LockService_proto_enumTypes[2]
}

func (x LockHeartbeatResponse_LockStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *LockHeartbeatResponse_LockStatus) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = LockHeartbeatResponse_LockStatus(num)
	return nil
}

// Deprecated: Use LockHeartbeatResponse_LockStatus.Descriptor instead.
func (LockHeartbeatResponse_LockStatus) EnumDescriptor() ([]byte, []int) {
	return file_LockService_proto_rawDescGZIP(), []int{3, 0}
}

type LockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LockType    *LockType     `protobuf:"varint,1,req,name=lock_type,json=lockType,enum=hbase.pb.LockType" json:"lock_type,omitempty"`
	Namespace   *string       `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	TableName   *TableName    `protobuf:"bytes,3,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	RegionInfo  []*RegionInfo `protobuf:"bytes,4,rep,name=region_info,json=regionInfo" json:"region_info,omitempty"`
	Description *string       `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	NonceGroup  *uint64       `protobuf:"varint,6,opt,name=nonce_group,json=nonceGroup,def=0" json:"nonce_group,omitempty"`
	Nonce       *uint64       `protobuf:"varint,7,opt,name=nonce,def=0" json:"nonce,omitempty"`
}

// Default values for LockRequest fields.
const (
	Default_LockRequest_NonceGroup = uint64(0)
	Default_LockRequest_Nonce      = uint64(0)
)

func (x *LockRequest) Reset() {
	*x = LockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LockService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockRequest) ProtoMessage() {}

func (x *LockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_LockService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockRequest.ProtoReflect.Descriptor instead.
func (*LockRequest) Descriptor() ([]byte, []int) {
	return file_LockService_proto_rawDescGZIP(), []int{0}
}

func (x *LockRequest) GetLockType() LockType {
	if x != nil && x.LockType != nil {
		return *x.LockType
	}
	return LockType_EXCLUSIVE
}

func (x *LockRequest) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *LockRequest) GetTableName() *TableName {
	if x != nil {
		return x.TableName
	}
	return nil
}

func (x *LockRequest) GetRegionInfo() []*RegionInfo {
	if x != nil {
		return x.RegionInfo
	}
	return nil
}

func (x *LockRequest) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *LockRequest) GetNonceGroup() uint64 {
	if x != nil && x.NonceGroup != nil {
		return *x.NonceGroup
	}
	return Default_LockRequest_NonceGroup
}

func (x *LockRequest) GetNonce() uint64 {
	if x != nil && x.Nonce != nil {
		return *x.Nonce
	}
	return Default_LockRequest_Nonce
}

type LockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcId *uint64 `protobuf:"varint,1,req,name=proc_id,json=procId" json:"proc_id,omitempty"`
}

func (x *LockResponse) Reset() {
	*x = LockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LockService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockResponse) ProtoMessage() {}

func (x *LockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_LockService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockResponse.ProtoReflect.Descriptor instead.
func (*LockResponse) Descriptor() ([]byte, []int) {
	return file_LockService_proto_rawDescGZIP(), []int{1}
}

func (x *LockResponse) GetProcId() uint64 {
	if x != nil && x.ProcId != nil {
		return *x.ProcId
	}
	return 0
}

type LockHeartbeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProcId    *uint64 `protobuf:"varint,1,req,name=proc_id,json=procId" json:"proc_id,omitempty"`
	KeepAlive *bool   `protobuf:"varint,2,opt,name=keep_alive,json=keepAlive,def=1" json:"keep_alive,omitempty"`
}

// Default values for LockHeartbeatRequest fields.
const (
	Default_LockHeartbeatRequest_KeepAlive = bool(true)
)

func (x *LockHeartbeatRequest) Reset() {
	*x = LockHeartbeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LockService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockHeartbeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockHeartbeatRequest) ProtoMessage() {}

func (x *LockHeartbeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_LockService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockHeartbeatRequest.ProtoReflect.Descriptor instead.
func (*LockHeartbeatRequest) Descriptor() ([]byte, []int) {
	return file_LockService_proto_rawDescGZIP(), []int{2}
}

func (x *LockHeartbeatRequest) GetProcId() uint64 {
	if x != nil && x.ProcId != nil {
		return *x.ProcId
	}
	return 0
}

func (x *LockHeartbeatRequest) GetKeepAlive() bool {
	if x != nil && x.KeepAlive != nil {
		return *x.KeepAlive
	}
	return Default_LockHeartbeatRequest_KeepAlive
}

type LockHeartbeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LockStatus *LockHeartbeatResponse_LockStatus `protobuf:"varint,1,req,name=lock_status,json=lockStatus,enum=hbase.pb.LockHeartbeatResponse_LockStatus" json:"lock_status,omitempty"`
	// Timeout of lock (if locked).
	TimeoutMs *uint32 `protobuf:"varint,2,opt,name=timeout_ms,json=timeoutMs" json:"timeout_ms,omitempty"`
}

func (x *LockHeartbeatResponse) Reset() {
	*x = LockHeartbeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LockService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockHeartbeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockHeartbeatResponse) ProtoMessage() {}

func (x *LockHeartbeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_LockService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockHeartbeatResponse.ProtoReflect.Descriptor instead.
func (*LockHeartbeatResponse) Descriptor() ([]byte, []int) {
	return file_LockService_proto_rawDescGZIP(), []int{3}
}

func (x *LockHeartbeatResponse) GetLockStatus() LockHeartbeatResponse_LockStatus {
	if x != nil && x.LockStatus != nil {
		return *x.LockStatus
	}
	return LockHeartbeatResponse_UNLOCKED
}

func (x *LockHeartbeatResponse) GetTimeoutMs() uint32 {
	if x != nil && x.TimeoutMs != nil {
		return *x.TimeoutMs
	}
	return 0
}

type LockProcedureData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LockType     *LockType     `protobuf:"varint,1,req,name=lock_type,json=lockType,enum=hbase.pb.LockType" json:"lock_type,omitempty"`
	Namespace    *string       `protobuf:"bytes,2,opt,name=namespace" json:"namespace,omitempty"`
	TableName    *TableName    `protobuf:"bytes,3,opt,name=table_name,json=tableName" json:"table_name,omitempty"`
	RegionInfo   []*RegionInfo `protobuf:"bytes,4,rep,name=region_info,json=regionInfo" json:"region_info,omitempty"`
	Description  *string       `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	IsMasterLock *bool         `protobuf:"varint,6,opt,name=is_master_lock,json=isMasterLock,def=0" json:"is_master_lock,omitempty"`
}

// Default values for LockProcedureData fields.
const (
	Default_LockProcedureData_IsMasterLock = bool(false)
)

func (x *LockProcedureData) Reset() {
	*x = LockProcedureData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LockService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockProcedureData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockProcedureData) ProtoMessage() {}

func (x *LockProcedureData) ProtoReflect() protoreflect.Message {
	mi := &file_LockService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockProcedureData.ProtoReflect.Descriptor instead.
func (*LockProcedureData) Descriptor() ([]byte, []int) {
	return file_LockService_proto_rawDescGZIP(), []int{4}
}

func (x *LockProcedureData) GetLockType() LockType {
	if x != nil && x.LockType != nil {
		return *x.LockType
	}
	return LockType_EXCLUSIVE
}

func (x *LockProcedureData) GetNamespace() string {
	if x != nil && x.Namespace != nil {
		return *x.Namespace
	}
	return ""
}

func (x *LockProcedureData) GetTableName() *TableName {
	if x != nil {
		return x.TableName
	}
	return nil
}

func (x *LockProcedureData) GetRegionInfo() []*RegionInfo {
	if x != nil {
		return x.RegionInfo
	}
	return nil
}

func (x *LockProcedureData) GetDescription() string {
	if x != nil && x.Description != nil {
		return *x.Description
	}
	return ""
}

func (x *LockProcedureData) GetIsMasterLock() bool {
	if x != nil && x.IsMasterLock != nil {
		return *x.IsMasterLock
	}
	return Default_LockProcedureData_IsMasterLock
}

type LockedResource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResourceType                *LockedResourceType `protobuf:"varint,1,req,name=resource_type,json=resourceType,enum=hbase.pb.LockedResourceType" json:"resource_type,omitempty"`
	ResourceName                *string             `protobuf:"bytes,2,opt,name=resource_name,json=resourceName" json:"resource_name,omitempty"`
	LockType                    *LockType           `protobuf:"varint,3,req,name=lock_type,json=lockType,enum=hbase.pb.LockType" json:"lock_type,omitempty"`
	ExclusiveLockOwnerProcedure *Procedure          `protobuf:"bytes,4,opt,name=exclusive_lock_owner_procedure,json=exclusiveLockOwnerProcedure" json:"exclusive_lock_owner_procedure,omitempty"`
	SharedLockCount             *int32              `protobuf:"varint,5,opt,name=shared_lock_count,json=sharedLockCount" json:"shared_lock_count,omitempty"`
	WaitingProcedures           []*Procedure        `protobuf:"bytes,6,rep,name=waitingProcedures" json:"waitingProcedures,omitempty"`
}

func (x *LockedResource) Reset() {
	*x = LockedResource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_LockService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockedResource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockedResource) ProtoMessage() {}

func (x *LockedResource) ProtoReflect() protoreflect.Message {
	mi := &file_LockService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockedResource.ProtoReflect.Descriptor instead.
func (*LockedResource) Descriptor() ([]byte, []int) {
	return file_LockService_proto_rawDescGZIP(), []int{5}
}

func (x *LockedResource) GetResourceType() LockedResourceType {
	if x != nil && x.ResourceType != nil {
		return *x.ResourceType
	}
	return LockedResourceType_SERVER
}

func (x *LockedResource) GetResourceName() string {
	if x != nil && x.ResourceName != nil {
		return *x.ResourceName
	}
	return ""
}

func (x *LockedResource) GetLockType() LockType {
	if x != nil && x.LockType != nil {
		return *x.LockType
	}
	return LockType_EXCLUSIVE
}

func (x *LockedResource) GetExclusiveLockOwnerProcedure() *Procedure {
	if x != nil {
		return x.ExclusiveLockOwnerProcedure
	}
	return nil
}

func (x *LockedResource) GetSharedLockCount() int32 {
	if x != nil && x.SharedLockCount != nil {
		return *x.SharedLockCount
	}
	return 0
}

func (x *LockedResource) GetWaitingProcedures() []*Procedure {
	if x != nil {
		return x.WaitingProcedures
	}
	return nil
}

var File_LockService_proto protoreflect.FileDescriptor

var file_LockService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x08, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x1a, 0x0b, 0x48,
	0x42, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0f, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x64, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa6, 0x02, 0x0a, 0x0b,
	0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2f, 0x0a, 0x09, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x12,
	0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x74, 0x61,
	0x62, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13,
	0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x35,
	0x0a, 0x0b, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0b, 0x6e, 0x6f, 0x6e, 0x63, 0x65,
	0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x3a, 0x01, 0x30, 0x52,
	0x0a, 0x6e, 0x6f, 0x6e, 0x63, 0x65, 0x47, 0x72, 0x6f, 0x75, 0x70, 0x12, 0x17, 0x0a, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x3a, 0x01, 0x30, 0x52, 0x05, 0x6e,
	0x6f, 0x6e, 0x63, 0x65, 0x22, 0x27, 0x0a, 0x0c, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x02, 0x28, 0x04, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x63, 0x49, 0x64, 0x22, 0x54, 0x0a,
	0x14, 0x4c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x63, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x02, 0x28, 0x04, 0x52, 0x06, 0x70, 0x72, 0x6f, 0x63, 0x49, 0x64, 0x12, 0x23,
	0x0a, 0x0a, 0x6b, 0x65, 0x65, 0x70, 0x5f, 0x61, 0x6c, 0x69, 0x76, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x3a, 0x04, 0x74, 0x72, 0x75, 0x65, 0x52, 0x09, 0x6b, 0x65, 0x65, 0x70, 0x41, 0x6c,
	0x69, 0x76, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x15, 0x4c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x72,
	0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a,
	0x0b, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x0e, 0x32, 0x2a, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f,
	0x63, 0x6b, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x0a,
	0x6c, 0x6f, 0x63, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x74, 0x69,
	0x6d, 0x65, 0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x4d, 0x73, 0x22, 0x26, 0x0a, 0x0a, 0x4c, 0x6f, 0x63,
	0x6b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x4c, 0x4f, 0x43,
	0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10,
	0x02, 0x22, 0x9c, 0x02, 0x0a, 0x11, 0x4c, 0x6f, 0x63, 0x6b, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64,
	0x75, 0x72, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x2f, 0x0a, 0x09, 0x6c, 0x6f, 0x63, 0x6b, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x68, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x52, 0x08,
	0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d,
	0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x32, 0x0a, 0x0a, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x52,
	0x09, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x0b, 0x72, 0x65,
	0x67, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0a, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2b, 0x0a, 0x0e, 0x69, 0x73, 0x5f, 0x6d, 0x61, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x3a, 0x05, 0x66, 0x61, 0x6c,
	0x73, 0x65, 0x52, 0x0c, 0x69, 0x73, 0x4d, 0x61, 0x73, 0x74, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x6b,
	0x22, 0xf2, 0x02, 0x0a, 0x0e, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x68, 0x62, 0x61,
	0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2f, 0x0a, 0x09, 0x6c,
	0x6f, 0x63, 0x6b, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0e, 0x32, 0x12,
	0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x54, 0x79,
	0x70, 0x65, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70, 0x65, 0x12, 0x58, 0x0a, 0x1e,
	0x65, 0x78, 0x63, 0x6c, 0x75, 0x73, 0x69, 0x76, 0x65, 0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e,
	0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x52, 0x1b, 0x65, 0x78, 0x63, 0x6c, 0x75,
	0x73, 0x69, 0x76, 0x65, 0x4c, 0x6f, 0x63, 0x6b, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x12, 0x2a, 0x0a, 0x11, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x5f, 0x6c, 0x6f, 0x63, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x4c, 0x6f, 0x63, 0x6b, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x41, 0x0a, 0x11, 0x77, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f,
	0x63, 0x65, 0x64, 0x75, 0x72, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x64, 0x75,
	0x72, 0x65, 0x52, 0x11, 0x77, 0x61, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x64, 0x75, 0x72, 0x65, 0x73, 0x2a, 0x25, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x6b, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x0d, 0x0a, 0x09, 0x45, 0x58, 0x43, 0x4c, 0x55, 0x53, 0x49, 0x56, 0x45, 0x10, 0x01,
	0x12, 0x0a, 0x0a, 0x06, 0x53, 0x48, 0x41, 0x52, 0x45, 0x44, 0x10, 0x02, 0x2a, 0x50, 0x0a, 0x12,
	0x4c, 0x6f, 0x63, 0x6b, 0x65, 0x64, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52, 0x10, 0x01, 0x12, 0x0d,
	0x0a, 0x09, 0x4e, 0x41, 0x4d, 0x45, 0x53, 0x50, 0x41, 0x43, 0x45, 0x10, 0x02, 0x12, 0x09, 0x0a,
	0x05, 0x54, 0x41, 0x42, 0x4c, 0x45, 0x10, 0x03, 0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x47, 0x49,
	0x4f, 0x4e, 0x10, 0x04, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x45, 0x45, 0x52, 0x10, 0x05, 0x32, 0x9d,
	0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3c,
	0x0a, 0x0b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4c, 0x6f, 0x63, 0x6b, 0x12, 0x15, 0x2e,
	0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e,
	0x4c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0d,
	0x4c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x12, 0x1e, 0x2e,
	0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e,
	0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x48, 0x65, 0x61,
	0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x55,
	0x0a, 0x31, 0x6f, 0x72, 0x67, 0x2e, 0x61, 0x70, 0x61, 0x63, 0x68, 0x65, 0x2e, 0x68, 0x61, 0x64,
	0x6f, 0x6f, 0x70, 0x2e, 0x68, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x73, 0x68, 0x61, 0x64, 0x65, 0x64,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x42, 0x11, 0x4c, 0x6f, 0x63, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x48, 0x01, 0x5a, 0x05, 0x2e, 0x2e, 0x2f, 0x70, 0x62, 0x88,
	0x01, 0x01, 0xa0, 0x01, 0x01,
}

var (
	file_LockService_proto_rawDescOnce sync.Once
	file_LockService_proto_rawDescData = file_LockService_proto_rawDesc
)

func file_LockService_proto_rawDescGZIP() []byte {
	file_LockService_proto_rawDescOnce.Do(func() {
		file_LockService_proto_rawDescData = protoimpl.X.CompressGZIP(file_LockService_proto_rawDescData)
	})
	return file_LockService_proto_rawDescData
}

var file_LockService_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_LockService_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_LockService_proto_goTypes = []interface{}{
	(LockType)(0),                         // 0: hbase.pb.LockType
	(LockedResourceType)(0),               // 1: hbase.pb.LockedResourceType
	(LockHeartbeatResponse_LockStatus)(0), // 2: hbase.pb.LockHeartbeatResponse.LockStatus
	(*LockRequest)(nil),                   // 3: hbase.pb.LockRequest
	(*LockResponse)(nil),                  // 4: hbase.pb.LockResponse
	(*LockHeartbeatRequest)(nil),          // 5: hbase.pb.LockHeartbeatRequest
	(*LockHeartbeatResponse)(nil),         // 6: hbase.pb.LockHeartbeatResponse
	(*LockProcedureData)(nil),             // 7: hbase.pb.LockProcedureData
	(*LockedResource)(nil),                // 8: hbase.pb.LockedResource
	(*TableName)(nil),                     // 9: hbase.pb.TableName
	(*RegionInfo)(nil),                    // 10: hbase.pb.RegionInfo
	(*Procedure)(nil),                     // 11: hbase.pb.Procedure
}
var file_LockService_proto_depIdxs = []int32{
	0,  // 0: hbase.pb.LockRequest.lock_type:type_name -> hbase.pb.LockType
	9,  // 1: hbase.pb.LockRequest.table_name:type_name -> hbase.pb.TableName
	10, // 2: hbase.pb.LockRequest.region_info:type_name -> hbase.pb.RegionInfo
	2,  // 3: hbase.pb.LockHeartbeatResponse.lock_status:type_name -> hbase.pb.LockHeartbeatResponse.LockStatus
	0,  // 4: hbase.pb.LockProcedureData.lock_type:type_name -> hbase.pb.LockType
	9,  // 5: hbase.pb.LockProcedureData.table_name:type_name -> hbase.pb.TableName
	10, // 6: hbase.pb.LockProcedureData.region_info:type_name -> hbase.pb.RegionInfo
	1,  // 7: hbase.pb.LockedResource.resource_type:type_name -> hbase.pb.LockedResourceType
	0,  // 8: hbase.pb.LockedResource.lock_type:type_name -> hbase.pb.LockType
	11, // 9: hbase.pb.LockedResource.exclusive_lock_owner_procedure:type_name -> hbase.pb.Procedure
	11, // 10: hbase.pb.LockedResource.waitingProcedures:type_name -> hbase.pb.Procedure
	3,  // 11: hbase.pb.LockService.RequestLock:input_type -> hbase.pb.LockRequest
	5,  // 12: hbase.pb.LockService.LockHeartbeat:input_type -> hbase.pb.LockHeartbeatRequest
	4,  // 13: hbase.pb.LockService.RequestLock:output_type -> hbase.pb.LockResponse
	6,  // 14: hbase.pb.LockService.LockHeartbeat:output_type -> hbase.pb.LockHeartbeatResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_LockService_proto_init() }
func file_LockService_proto_init() {
	if File_LockService_proto != nil {
		return
	}
	file_HBase_proto_init()
	file_Procedure_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_LockService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockRequest); i {
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
		file_LockService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockResponse); i {
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
		file_LockService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockHeartbeatRequest); i {
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
		file_LockService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockHeartbeatResponse); i {
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
		file_LockService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockProcedureData); i {
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
		file_LockService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockedResource); i {
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
			RawDescriptor: file_LockService_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_LockService_proto_goTypes,
		DependencyIndexes: file_LockService_proto_depIdxs,
		EnumInfos:         file_LockService_proto_enumTypes,
		MessageInfos:      file_LockService_proto_msgTypes,
	}.Build()
	File_LockService_proto = out.File
	file_LockService_proto_rawDesc = nil
	file_LockService_proto_goTypes = nil
	file_LockService_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LockServiceClient is the client API for LockService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LockServiceClient interface {
	//* Acquire lock on namespace/table/region
	RequestLock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockResponse, error)
	//* Keep alive (or not) a previously acquired lock
	LockHeartbeat(ctx context.Context, in *LockHeartbeatRequest, opts ...grpc.CallOption) (*LockHeartbeatResponse, error)
}

type lockServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLockServiceClient(cc grpc.ClientConnInterface) LockServiceClient {
	return &lockServiceClient{cc}
}

func (c *lockServiceClient) RequestLock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockResponse, error) {
	out := new(LockResponse)
	err := c.cc.Invoke(ctx, "/hbase.pb.LockService/RequestLock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lockServiceClient) LockHeartbeat(ctx context.Context, in *LockHeartbeatRequest, opts ...grpc.CallOption) (*LockHeartbeatResponse, error) {
	out := new(LockHeartbeatResponse)
	err := c.cc.Invoke(ctx, "/hbase.pb.LockService/LockHeartbeat", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LockServiceServer is the server API for LockService service.
type LockServiceServer interface {
	//* Acquire lock on namespace/table/region
	RequestLock(context.Context, *LockRequest) (*LockResponse, error)
	//* Keep alive (or not) a previously acquired lock
	LockHeartbeat(context.Context, *LockHeartbeatRequest) (*LockHeartbeatResponse, error)
}

// UnimplementedLockServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLockServiceServer struct {
}

func (*UnimplementedLockServiceServer) RequestLock(context.Context, *LockRequest) (*LockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestLock not implemented")
}
func (*UnimplementedLockServiceServer) LockHeartbeat(context.Context, *LockHeartbeatRequest) (*LockHeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LockHeartbeat not implemented")
}

func RegisterLockServiceServer(s *grpc.Server, srv LockServiceServer) {
	s.RegisterService(&_LockService_serviceDesc, srv)
}

func _LockService_RequestLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServiceServer).RequestLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hbase.pb.LockService/RequestLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServiceServer).RequestLock(ctx, req.(*LockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LockService_LockHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockHeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LockServiceServer).LockHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hbase.pb.LockService/LockHeartbeat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LockServiceServer).LockHeartbeat(ctx, req.(*LockHeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LockService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hbase.pb.LockService",
	HandlerType: (*LockServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestLock",
			Handler:    _LockService_RequestLock_Handler,
		},
		{
			MethodName: "LockHeartbeat",
			Handler:    _LockService_LockHeartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "LockService.proto",
}