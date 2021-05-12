// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: cc.proto

package cc

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

type AppAddReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AppName string `protobuf:"bytes,1,opt,name=AppName,proto3" json:"AppName,omitempty"`
	AppKey  string `protobuf:"bytes,2,opt,name=AppKey,proto3" json:"AppKey,omitempty"`
	Desc    string `protobuf:"bytes,3,opt,name=Desc,proto3" json:"Desc,omitempty"`
}

func (x *AppAddReq) Reset() {
	*x = AppAddReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppAddReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppAddReq) ProtoMessage() {}

func (x *AppAddReq) ProtoReflect() protoreflect.Message {
	mi := &file_cc_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppAddReq.ProtoReflect.Descriptor instead.
func (*AppAddReq) Descriptor() ([]byte, []int) {
	return file_cc_proto_rawDescGZIP(), []int{0}
}

func (x *AppAddReq) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *AppAddReq) GetAppKey() string {
	if x != nil {
		return x.AppKey
	}
	return ""
}

func (x *AppAddReq) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type AppAddReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
	Id   int64  `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AppAddReply) Reset() {
	*x = AppAddReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppAddReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppAddReply) ProtoMessage() {}

func (x *AppAddReply) ProtoReflect() protoreflect.Message {
	mi := &file_cc_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppAddReply.ProtoReflect.Descriptor instead.
func (*AppAddReply) Descriptor() ([]byte, []int) {
	return file_cc_proto_rawDescGZIP(), []int{1}
}

func (x *AppAddReply) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

func (x *AppAddReply) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AppListReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Current  int64 `protobuf:"varint,1,opt,name=current,proto3" json:"current,omitempty"`
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
}

func (x *AppListReq) Reset() {
	*x = AppListReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppListReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppListReq) ProtoMessage() {}

func (x *AppListReq) ProtoReflect() protoreflect.Message {
	mi := &file_cc_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppListReq.ProtoReflect.Descriptor instead.
func (*AppListReq) Descriptor() ([]byte, []int) {
	return file_cc_proto_rawDescGZIP(), []int{2}
}

func (x *AppListReq) GetCurrent() int64 {
	if x != nil {
		return x.Current
	}
	return 0
}

func (x *AppListReq) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListAppData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AppKey     string `protobuf:"bytes,2,opt,name=appKey,proto3" json:"appKey,omitempty"`
	AppName    string `protobuf:"bytes,3,opt,name=appName,proto3" json:"appName,omitempty"`
	Desc       string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
	CreateTime string `protobuf:"bytes,5,opt,name=createTime,proto3" json:"createTime,omitempty"`
	UpdateTime string `protobuf:"bytes,6,opt,name=updateTime,proto3" json:"updateTime,omitempty"`
}

func (x *ListAppData) Reset() {
	*x = ListAppData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAppData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAppData) ProtoMessage() {}

func (x *ListAppData) ProtoReflect() protoreflect.Message {
	mi := &file_cc_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAppData.ProtoReflect.Descriptor instead.
func (*ListAppData) Descriptor() ([]byte, []int) {
	return file_cc_proto_rawDescGZIP(), []int{3}
}

func (x *ListAppData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ListAppData) GetAppKey() string {
	if x != nil {
		return x.AppKey
	}
	return ""
}

func (x *ListAppData) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *ListAppData) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *ListAppData) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *ListAppData) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

type AppListReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64          `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	List  []*ListAppData `protobuf:"bytes,2,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *AppListReply) Reset() {
	*x = AppListReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppListReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppListReply) ProtoMessage() {}

func (x *AppListReply) ProtoReflect() protoreflect.Message {
	mi := &file_cc_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppListReply.ProtoReflect.Descriptor instead.
func (*AppListReply) Descriptor() ([]byte, []int) {
	return file_cc_proto_rawDescGZIP(), []int{4}
}

func (x *AppListReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *AppListReply) GetList() []*ListAppData {
	if x != nil {
		return x.List
	}
	return nil
}

type AppUpdateReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AppKey  string `protobuf:"bytes,2,opt,name=appKey,proto3" json:"appKey,omitempty"`
	AppName string `protobuf:"bytes,3,opt,name=appName,proto3" json:"appName,omitempty"`
	Desc    string `protobuf:"bytes,4,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *AppUpdateReq) Reset() {
	*x = AppUpdateReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppUpdateReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppUpdateReq) ProtoMessage() {}

func (x *AppUpdateReq) ProtoReflect() protoreflect.Message {
	mi := &file_cc_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppUpdateReq.ProtoReflect.Descriptor instead.
func (*AppUpdateReq) Descriptor() ([]byte, []int) {
	return file_cc_proto_rawDescGZIP(), []int{5}
}

func (x *AppUpdateReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *AppUpdateReq) GetAppKey() string {
	if x != nil {
		return x.AppKey
	}
	return ""
}

func (x *AppUpdateReq) GetAppName() string {
	if x != nil {
		return x.AppName
	}
	return ""
}

func (x *AppUpdateReq) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type AppUpdateReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *AppUpdateReply) Reset() {
	*x = AppUpdateReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppUpdateReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppUpdateReply) ProtoMessage() {}

func (x *AppUpdateReply) ProtoReflect() protoreflect.Message {
	mi := &file_cc_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppUpdateReply.ProtoReflect.Descriptor instead.
func (*AppUpdateReply) Descriptor() ([]byte, []int) {
	return file_cc_proto_rawDescGZIP(), []int{6}
}

func (x *AppUpdateReply) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

type AppDelReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AppDelReq) Reset() {
	*x = AppDelReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppDelReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppDelReq) ProtoMessage() {}

func (x *AppDelReq) ProtoReflect() protoreflect.Message {
	mi := &file_cc_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppDelReq.ProtoReflect.Descriptor instead.
func (*AppDelReq) Descriptor() ([]byte, []int) {
	return file_cc_proto_rawDescGZIP(), []int{7}
}

func (x *AppDelReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type AppDelReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong string `protobuf:"bytes,1,opt,name=pong,proto3" json:"pong,omitempty"`
}

func (x *AppDelReply) Reset() {
	*x = AppDelReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_cc_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppDelReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppDelReply) ProtoMessage() {}

func (x *AppDelReply) ProtoReflect() protoreflect.Message {
	mi := &file_cc_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppDelReply.ProtoReflect.Descriptor instead.
func (*AppDelReply) Descriptor() ([]byte, []int) {
	return file_cc_proto_rawDescGZIP(), []int{8}
}

func (x *AppDelReply) GetPong() string {
	if x != nil {
		return x.Pong
	}
	return ""
}

var File_cc_proto protoreflect.FileDescriptor

var file_cc_proto_rawDesc = []byte{
	0x0a, 0x08, 0x63, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x63, 0x63, 0x22, 0x51,
	0x0a, 0x09, 0x41, 0x70, 0x70, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x41,
	0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x70,
	0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x70, 0x70, 0x4b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x41, 0x70, 0x70, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x44, 0x65, 0x73, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x44, 0x65, 0x73,
	0x63, 0x22, 0x31, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x41, 0x64, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x6f, 0x6e, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x42, 0x0a, 0x0a, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08,
	0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0xa3, 0x01, 0x0a, 0x0b, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x70, 0x70, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x4b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x70, 0x4b, 0x65, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65,
	0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x1e,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x49,
	0x0a, 0x0c, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74,
	0x6f, 0x74, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x63, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x70, 0x70, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x64, 0x0a, 0x0c, 0x41, 0x70, 0x70,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x70, 0x70,
	0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x70, 0x70, 0x4b, 0x65,
	0x79, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x70, 0x70, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x65, 0x73, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22,
	0x24, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x6f, 0x6e, 0x67, 0x22, 0x1b, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x44, 0x65, 0x6c, 0x52,
	0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x21, 0x0a, 0x0b, 0x41, 0x70, 0x70, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x6f, 0x6e, 0x67, 0x32, 0xb8, 0x01, 0x0a, 0x02, 0x63, 0x63, 0x12, 0x28, 0x0a, 0x06,
	0x41, 0x70, 0x70, 0x41, 0x64, 0x64, 0x12, 0x0d, 0x2e, 0x63, 0x63, 0x2e, 0x41, 0x70, 0x70, 0x41,
	0x64, 0x64, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x63, 0x63, 0x2e, 0x41, 0x70, 0x70, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2b, 0x0a, 0x07, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x0e, 0x2e, 0x63, 0x63, 0x2e, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x10, 0x2e, 0x63, 0x63, 0x2e, 0x41, 0x70, 0x70, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x12, 0x10, 0x2e, 0x63, 0x63, 0x2e, 0x41, 0x70, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x1a, 0x12, 0x2e, 0x63, 0x63, 0x2e, 0x41, 0x70, 0x70, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x28, 0x0a, 0x06, 0x41, 0x70, 0x70, 0x44, 0x65, 0x6c,
	0x12, 0x0d, 0x2e, 0x63, 0x63, 0x2e, 0x41, 0x70, 0x70, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x1a,
	0x0f, 0x2e, 0x63, 0x63, 0x2e, 0x41, 0x70, 0x70, 0x44, 0x65, 0x6c, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x42, 0x05, 0x5a, 0x03, 0x2f, 0x63, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_cc_proto_rawDescOnce sync.Once
	file_cc_proto_rawDescData = file_cc_proto_rawDesc
)

func file_cc_proto_rawDescGZIP() []byte {
	file_cc_proto_rawDescOnce.Do(func() {
		file_cc_proto_rawDescData = protoimpl.X.CompressGZIP(file_cc_proto_rawDescData)
	})
	return file_cc_proto_rawDescData
}

var file_cc_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_cc_proto_goTypes = []interface{}{
	(*AppAddReq)(nil),      // 0: cc.AppAddReq
	(*AppAddReply)(nil),    // 1: cc.AppAddReply
	(*AppListReq)(nil),     // 2: cc.AppListReq
	(*ListAppData)(nil),    // 3: cc.ListAppData
	(*AppListReply)(nil),   // 4: cc.AppListReply
	(*AppUpdateReq)(nil),   // 5: cc.AppUpdateReq
	(*AppUpdateReply)(nil), // 6: cc.AppUpdateReply
	(*AppDelReq)(nil),      // 7: cc.AppDelReq
	(*AppDelReply)(nil),    // 8: cc.AppDelReply
}
var file_cc_proto_depIdxs = []int32{
	3, // 0: cc.AppListReply.list:type_name -> cc.ListAppData
	0, // 1: cc.cc.AppAdd:input_type -> cc.AppAddReq
	2, // 2: cc.cc.AppList:input_type -> cc.AppListReq
	5, // 3: cc.cc.AppUpdate:input_type -> cc.AppUpdateReq
	7, // 4: cc.cc.AppDel:input_type -> cc.AppDelReq
	1, // 5: cc.cc.AppAdd:output_type -> cc.AppAddReply
	4, // 6: cc.cc.AppList:output_type -> cc.AppListReply
	6, // 7: cc.cc.AppUpdate:output_type -> cc.AppUpdateReply
	8, // 8: cc.cc.AppDel:output_type -> cc.AppDelReply
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_cc_proto_init() }
func file_cc_proto_init() {
	if File_cc_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_cc_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppAddReq); i {
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
		file_cc_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppAddReply); i {
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
		file_cc_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppListReq); i {
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
		file_cc_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAppData); i {
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
		file_cc_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppListReply); i {
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
		file_cc_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppUpdateReq); i {
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
		file_cc_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppUpdateReply); i {
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
		file_cc_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppDelReq); i {
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
		file_cc_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppDelReply); i {
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
			RawDescriptor: file_cc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_cc_proto_goTypes,
		DependencyIndexes: file_cc_proto_depIdxs,
		MessageInfos:      file_cc_proto_msgTypes,
	}.Build()
	File_cc_proto = out.File
	file_cc_proto_rawDesc = nil
	file_cc_proto_goTypes = nil
	file_cc_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// CcClient is the client API for Cc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CcClient interface {
	AppAdd(ctx context.Context, in *AppAddReq, opts ...grpc.CallOption) (*AppAddReply, error)
	AppList(ctx context.Context, in *AppListReq, opts ...grpc.CallOption) (*AppListReply, error)
	AppUpdate(ctx context.Context, in *AppUpdateReq, opts ...grpc.CallOption) (*AppUpdateReply, error)
	AppDel(ctx context.Context, in *AppDelReq, opts ...grpc.CallOption) (*AppDelReply, error)
}

type ccClient struct {
	cc grpc.ClientConnInterface
}

func NewCcClient(cc grpc.ClientConnInterface) CcClient {
	return &ccClient{cc}
}

func (c *ccClient) AppAdd(ctx context.Context, in *AppAddReq, opts ...grpc.CallOption) (*AppAddReply, error) {
	out := new(AppAddReply)
	err := c.cc.Invoke(ctx, "/cc.cc/AppAdd", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ccClient) AppList(ctx context.Context, in *AppListReq, opts ...grpc.CallOption) (*AppListReply, error) {
	out := new(AppListReply)
	err := c.cc.Invoke(ctx, "/cc.cc/AppList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ccClient) AppUpdate(ctx context.Context, in *AppUpdateReq, opts ...grpc.CallOption) (*AppUpdateReply, error) {
	out := new(AppUpdateReply)
	err := c.cc.Invoke(ctx, "/cc.cc/AppUpdate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ccClient) AppDel(ctx context.Context, in *AppDelReq, opts ...grpc.CallOption) (*AppDelReply, error) {
	out := new(AppDelReply)
	err := c.cc.Invoke(ctx, "/cc.cc/AppDel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CcServer is the server API for Cc service.
type CcServer interface {
	AppAdd(context.Context, *AppAddReq) (*AppAddReply, error)
	AppList(context.Context, *AppListReq) (*AppListReply, error)
	AppUpdate(context.Context, *AppUpdateReq) (*AppUpdateReply, error)
	AppDel(context.Context, *AppDelReq) (*AppDelReply, error)
}

// UnimplementedCcServer can be embedded to have forward compatible implementations.
type UnimplementedCcServer struct {
}

func (*UnimplementedCcServer) AppAdd(context.Context, *AppAddReq) (*AppAddReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppAdd not implemented")
}
func (*UnimplementedCcServer) AppList(context.Context, *AppListReq) (*AppListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppList not implemented")
}
func (*UnimplementedCcServer) AppUpdate(context.Context, *AppUpdateReq) (*AppUpdateReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppUpdate not implemented")
}
func (*UnimplementedCcServer) AppDel(context.Context, *AppDelReq) (*AppDelReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppDel not implemented")
}

func RegisterCcServer(s *grpc.Server, srv CcServer) {
	s.RegisterService(&_Cc_serviceDesc, srv)
}

func _Cc_AppAdd_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppAddReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CcServer).AppAdd(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cc.cc/AppAdd",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CcServer).AppAdd(ctx, req.(*AppAddReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cc_AppList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CcServer).AppList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cc.cc/AppList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CcServer).AppList(ctx, req.(*AppListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cc_AppUpdate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppUpdateReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CcServer).AppUpdate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cc.cc/AppUpdate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CcServer).AppUpdate(ctx, req.(*AppUpdateReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cc_AppDel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AppDelReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CcServer).AppDel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cc.cc/AppDel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CcServer).AppDel(ctx, req.(*AppDelReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cc_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cc.cc",
	HandlerType: (*CcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AppAdd",
			Handler:    _Cc_AppAdd_Handler,
		},
		{
			MethodName: "AppList",
			Handler:    _Cc_AppList_Handler,
		},
		{
			MethodName: "AppUpdate",
			Handler:    _Cc_AppUpdate_Handler,
		},
		{
			MethodName: "AppDel",
			Handler:    _Cc_AppDel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cc.proto",
}
