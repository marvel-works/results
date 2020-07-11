// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.12.1
// source: api.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type TaskRunResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskRun   *TaskRun `protobuf:"bytes,1,opt,name=task_run,json=taskRun,proto3" json:"task_run,omitempty"`
	ResultsId string   `protobuf:"bytes,2,opt,name=results_id,json=resultsId,proto3" json:"results_id,omitempty"`
}

func (x *TaskRunResult) Reset() {
	*x = TaskRunResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskRunResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskRunResult) ProtoMessage() {}

func (x *TaskRunResult) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskRunResult.ProtoReflect.Descriptor instead.
func (*TaskRunResult) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *TaskRunResult) GetTaskRun() *TaskRun {
	if x != nil {
		return x.TaskRun
	}
	return nil
}

func (x *TaskRunResult) GetResultsId() string {
	if x != nil {
		return x.ResultsId
	}
	return ""
}

type CreateTaskRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskRun *TaskRun `protobuf:"bytes,1,opt,name=task_run,json=taskRun,proto3" json:"task_run,omitempty"` // TODO: Accept some unique cluster identifier?
}

func (x *CreateTaskRunRequest) Reset() {
	*x = CreateTaskRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTaskRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTaskRunRequest) ProtoMessage() {}

func (x *CreateTaskRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTaskRunRequest.ProtoReflect.Descriptor instead.
func (*CreateTaskRunRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateTaskRunRequest) GetTaskRun() *TaskRun {
	if x != nil {
		return x.TaskRun
	}
	return nil
}

type DeleteTaskRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultsId string `protobuf:"bytes,1,opt,name=results_id,json=resultsId,proto3" json:"results_id,omitempty"`
}

func (x *DeleteTaskRunRequest) Reset() {
	*x = DeleteTaskRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTaskRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTaskRunRequest) ProtoMessage() {}

func (x *DeleteTaskRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTaskRunRequest.ProtoReflect.Descriptor instead.
func (*DeleteTaskRunRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteTaskRunRequest) GetResultsId() string {
	if x != nil {
		return x.ResultsId
	}
	return ""
}

type UpdateTaskRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskRun   *TaskRun `protobuf:"bytes,1,opt,name=task_run,json=taskRun,proto3" json:"task_run,omitempty"`
	ResultsId string   `protobuf:"bytes,2,opt,name=results_id,json=resultsId,proto3" json:"results_id,omitempty"`
}

func (x *UpdateTaskRunRequest) Reset() {
	*x = UpdateTaskRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTaskRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTaskRunRequest) ProtoMessage() {}

func (x *UpdateTaskRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTaskRunRequest.ProtoReflect.Descriptor instead.
func (*UpdateTaskRunRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateTaskRunRequest) GetTaskRun() *TaskRun {
	if x != nil {
		return x.TaskRun
	}
	return nil
}

func (x *UpdateTaskRunRequest) GetResultsId() string {
	if x != nil {
		return x.ResultsId
	}
	return ""
}

type GetTaskRunRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResultsId string `protobuf:"bytes,1,opt,name=results_id,json=resultsId,proto3" json:"results_id,omitempty"`
}

func (x *GetTaskRunRequest) Reset() {
	*x = GetTaskRunRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTaskRunRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTaskRunRequest) ProtoMessage() {}

func (x *GetTaskRunRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTaskRunRequest.ProtoReflect.Descriptor instead.
func (*GetTaskRunRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetTaskRunRequest) GetResultsId() string {
	if x != nil {
		return x.ResultsId
	}
	return ""
}

type ListTaskRunsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO: document query format.
	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"` // TODO: page_token, etc.
}

func (x *ListTaskRunsRequest) Reset() {
	*x = ListTaskRunsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTaskRunsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTaskRunsRequest) ProtoMessage() {}

func (x *ListTaskRunsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTaskRunsRequest.ProtoReflect.Descriptor instead.
func (*ListTaskRunsRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *ListTaskRunsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type ListTaskRunsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*TaskRun `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"` // TODO: next_page_token, etc.
}

func (x *ListTaskRunsResponse) Reset() {
	*x = ListTaskRunsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListTaskRunsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListTaskRunsResponse) ProtoMessage() {}

func (x *ListTaskRunsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListTaskRunsResponse.ProtoReflect.Descriptor instead.
func (*ListTaskRunsResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *ListTaskRunsResponse) GetItems() []*TaskRun {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x65, 0x6b,
	0x74, 0x6f, 0x6e, 0x1a, 0x0d, 0x74, 0x61, 0x73, 0x6b, 0x72, 0x75, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x5a, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x2a, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x65, 0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x75, 0x6e, 0x52, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x12, 0x1d, 0x0a, 0x0a,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x14, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x75, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x65, 0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x52, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x22,
	0x35, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x73, 0x49, 0x64, 0x22, 0x61, 0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2a,
	0x0a, 0x08, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x72, 0x75, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x74, 0x65, 0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75,
	0x6e, 0x52, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x49, 0x64, 0x22, 0x32, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x49, 0x64, 0x22, 0x2b, 0x0a,
	0x13, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0x3d, 0x0a, 0x14, 0x4c, 0x69,
	0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x74, 0x65, 0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52,
	0x75, 0x6e, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0x8f, 0x03, 0x0a, 0x07, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x12, 0x4c, 0x0a, 0x13, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54,
	0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x2e, 0x74,
	0x65, 0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x75, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x65, 0x6b,
	0x74, 0x6f, 0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73,
	0x6b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x2e, 0x74, 0x65, 0x6b,
	0x74, 0x6f, 0x6e, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x65, 0x6b, 0x74, 0x6f,
	0x6e, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x4c, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x19, 0x2e, 0x74, 0x65,
	0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x74, 0x65, 0x6b, 0x74, 0x6f, 0x6e, 0x2e,
	0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12,
	0x4d, 0x0a, 0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1c, 0x2e, 0x74, 0x65, 0x6b, 0x74, 0x6f, 0x6e, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x4b,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e, 0x73, 0x12, 0x1b,
	0x2e, 0x74, 0x65, 0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x75, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x65,
	0x6b, 0x74, 0x6f, 0x6e, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x75, 0x6e,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_proto_goTypes = []interface{}{
	(*TaskRunResult)(nil),        // 0: tekton.TaskRunResult
	(*CreateTaskRunRequest)(nil), // 1: tekton.CreateTaskRunRequest
	(*DeleteTaskRunRequest)(nil), // 2: tekton.DeleteTaskRunRequest
	(*UpdateTaskRunRequest)(nil), // 3: tekton.UpdateTaskRunRequest
	(*GetTaskRunRequest)(nil),    // 4: tekton.GetTaskRunRequest
	(*ListTaskRunsRequest)(nil),  // 5: tekton.ListTaskRunsRequest
	(*ListTaskRunsResponse)(nil), // 6: tekton.ListTaskRunsResponse
	(*TaskRun)(nil),              // 7: tekton.TaskRun
	(*empty.Empty)(nil),          // 8: google.protobuf.Empty
}
var file_api_proto_depIdxs = []int32{
	7, // 0: tekton.TaskRunResult.task_run:type_name -> tekton.TaskRun
	7, // 1: tekton.CreateTaskRunRequest.task_run:type_name -> tekton.TaskRun
	7, // 2: tekton.UpdateTaskRunRequest.task_run:type_name -> tekton.TaskRun
	7, // 3: tekton.ListTaskRunsResponse.items:type_name -> tekton.TaskRun
	1, // 4: tekton.Results.CreateTaskRunResult:input_type -> tekton.CreateTaskRunRequest
	3, // 5: tekton.Results.UpdateTaskRunResult:input_type -> tekton.UpdateTaskRunRequest
	4, // 6: tekton.Results.GetTaskRunResultResult:input_type -> tekton.GetTaskRunRequest
	2, // 7: tekton.Results.DeleteTaskRunResult:input_type -> tekton.DeleteTaskRunRequest
	5, // 8: tekton.Results.ListTaskRuns:input_type -> tekton.ListTaskRunsRequest
	0, // 9: tekton.Results.CreateTaskRunResult:output_type -> tekton.TaskRunResult
	0, // 10: tekton.Results.UpdateTaskRunResult:output_type -> tekton.TaskRunResult
	0, // 11: tekton.Results.GetTaskRunResultResult:output_type -> tekton.TaskRunResult
	8, // 12: tekton.Results.DeleteTaskRunResult:output_type -> google.protobuf.Empty
	6, // 13: tekton.Results.ListTaskRuns:output_type -> tekton.ListTaskRunsResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	file_taskrun_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TaskRunResult); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTaskRunRequest); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteTaskRunRequest); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateTaskRunRequest); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTaskRunRequest); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTaskRunsRequest); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListTaskRunsResponse); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ResultsClient is the client API for Results service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResultsClient interface {
	CreateTaskRunResult(ctx context.Context, in *CreateTaskRunRequest, opts ...grpc.CallOption) (*TaskRunResult, error)
	UpdateTaskRunResult(ctx context.Context, in *UpdateTaskRunRequest, opts ...grpc.CallOption) (*TaskRunResult, error)
	GetTaskRunResultResult(ctx context.Context, in *GetTaskRunRequest, opts ...grpc.CallOption) (*TaskRunResult, error)
	DeleteTaskRunResult(ctx context.Context, in *DeleteTaskRunRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	ListTaskRuns(ctx context.Context, in *ListTaskRunsRequest, opts ...grpc.CallOption) (*ListTaskRunsResponse, error)
}

type resultsClient struct {
	cc grpc.ClientConnInterface
}

func NewResultsClient(cc grpc.ClientConnInterface) ResultsClient {
	return &resultsClient{cc}
}

func (c *resultsClient) CreateTaskRunResult(ctx context.Context, in *CreateTaskRunRequest, opts ...grpc.CallOption) (*TaskRunResult, error) {
	out := new(TaskRunResult)
	err := c.cc.Invoke(ctx, "/tekton.Results/CreateTaskRunResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultsClient) UpdateTaskRunResult(ctx context.Context, in *UpdateTaskRunRequest, opts ...grpc.CallOption) (*TaskRunResult, error) {
	out := new(TaskRunResult)
	err := c.cc.Invoke(ctx, "/tekton.Results/UpdateTaskRunResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultsClient) GetTaskRunResultResult(ctx context.Context, in *GetTaskRunRequest, opts ...grpc.CallOption) (*TaskRunResult, error) {
	out := new(TaskRunResult)
	err := c.cc.Invoke(ctx, "/tekton.Results/GetTaskRunResultResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultsClient) DeleteTaskRunResult(ctx context.Context, in *DeleteTaskRunRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/tekton.Results/DeleteTaskRunResult", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resultsClient) ListTaskRuns(ctx context.Context, in *ListTaskRunsRequest, opts ...grpc.CallOption) (*ListTaskRunsResponse, error) {
	out := new(ListTaskRunsResponse)
	err := c.cc.Invoke(ctx, "/tekton.Results/ListTaskRuns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResultsServer is the server API for Results service.
type ResultsServer interface {
	CreateTaskRunResult(context.Context, *CreateTaskRunRequest) (*TaskRunResult, error)
	UpdateTaskRunResult(context.Context, *UpdateTaskRunRequest) (*TaskRunResult, error)
	GetTaskRunResultResult(context.Context, *GetTaskRunRequest) (*TaskRunResult, error)
	DeleteTaskRunResult(context.Context, *DeleteTaskRunRequest) (*empty.Empty, error)
	ListTaskRuns(context.Context, *ListTaskRunsRequest) (*ListTaskRunsResponse, error)
}

// UnimplementedResultsServer can be embedded to have forward compatible implementations.
type UnimplementedResultsServer struct {
}

func (*UnimplementedResultsServer) CreateTaskRunResult(context.Context, *CreateTaskRunRequest) (*TaskRunResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTaskRunResult not implemented")
}
func (*UnimplementedResultsServer) UpdateTaskRunResult(context.Context, *UpdateTaskRunRequest) (*TaskRunResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTaskRunResult not implemented")
}
func (*UnimplementedResultsServer) GetTaskRunResultResult(context.Context, *GetTaskRunRequest) (*TaskRunResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTaskRunResultResult not implemented")
}
func (*UnimplementedResultsServer) DeleteTaskRunResult(context.Context, *DeleteTaskRunRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteTaskRunResult not implemented")
}
func (*UnimplementedResultsServer) ListTaskRuns(context.Context, *ListTaskRunsRequest) (*ListTaskRunsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTaskRuns not implemented")
}

func RegisterResultsServer(s *grpc.Server, srv ResultsServer) {
	s.RegisterService(&_Results_serviceDesc, srv)
}

func _Results_CreateTaskRunResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTaskRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultsServer).CreateTaskRunResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tekton.Results/CreateTaskRunResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultsServer).CreateTaskRunResult(ctx, req.(*CreateTaskRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Results_UpdateTaskRunResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTaskRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultsServer).UpdateTaskRunResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tekton.Results/UpdateTaskRunResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultsServer).UpdateTaskRunResult(ctx, req.(*UpdateTaskRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Results_GetTaskRunResultResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTaskRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultsServer).GetTaskRunResultResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tekton.Results/GetTaskRunResultResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultsServer).GetTaskRunResultResult(ctx, req.(*GetTaskRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Results_DeleteTaskRunResult_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteTaskRunRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultsServer).DeleteTaskRunResult(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tekton.Results/DeleteTaskRunResult",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultsServer).DeleteTaskRunResult(ctx, req.(*DeleteTaskRunRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Results_ListTaskRuns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTaskRunsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResultsServer).ListTaskRuns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tekton.Results/ListTaskRuns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResultsServer).ListTaskRuns(ctx, req.(*ListTaskRunsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Results_serviceDesc = grpc.ServiceDesc{
	ServiceName: "tekton.Results",
	HandlerType: (*ResultsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateTaskRunResult",
			Handler:    _Results_CreateTaskRunResult_Handler,
		},
		{
			MethodName: "UpdateTaskRunResult",
			Handler:    _Results_UpdateTaskRunResult_Handler,
		},
		{
			MethodName: "GetTaskRunResultResult",
			Handler:    _Results_GetTaskRunResultResult_Handler,
		},
		{
			MethodName: "DeleteTaskRunResult",
			Handler:    _Results_DeleteTaskRunResult_Handler,
		},
		{
			MethodName: "ListTaskRuns",
			Handler:    _Results_ListTaskRuns_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
