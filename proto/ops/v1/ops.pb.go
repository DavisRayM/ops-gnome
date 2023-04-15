// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: ops/v1/ops.proto

package proto

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

type ListSupportedDeploymentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSupportedDeploymentsRequest) Reset() {
	*x = ListSupportedDeploymentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_v1_ops_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSupportedDeploymentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSupportedDeploymentsRequest) ProtoMessage() {}

func (x *ListSupportedDeploymentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ops_v1_ops_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSupportedDeploymentsRequest.ProtoReflect.Descriptor instead.
func (*ListSupportedDeploymentsRequest) Descriptor() ([]byte, []int) {
	return file_ops_v1_ops_proto_rawDescGZIP(), []int{0}
}

type ListSupportedDeploymentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Deployments []string `protobuf:"bytes,1,rep,name=deployments,proto3" json:"deployments,omitempty"`
}

func (x *ListSupportedDeploymentsResponse) Reset() {
	*x = ListSupportedDeploymentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_v1_ops_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSupportedDeploymentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSupportedDeploymentsResponse) ProtoMessage() {}

func (x *ListSupportedDeploymentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ops_v1_ops_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSupportedDeploymentsResponse.ProtoReflect.Descriptor instead.
func (*ListSupportedDeploymentsResponse) Descriptor() ([]byte, []int) {
	return file_ops_v1_ops_proto_rawDescGZIP(), []int{1}
}

func (x *ListSupportedDeploymentsResponse) GetDeployments() []string {
	if x != nil {
		return x.Deployments
	}
	return nil
}

type GetDeploymentStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Release   string `protobuf:"bytes,1,opt,name=release,proto3" json:"release,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
}

func (x *GetDeploymentStatusRequest) Reset() {
	*x = GetDeploymentStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_v1_ops_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeploymentStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentStatusRequest) ProtoMessage() {}

func (x *GetDeploymentStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ops_v1_ops_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentStatusRequest.ProtoReflect.Descriptor instead.
func (*GetDeploymentStatusRequest) Descriptor() ([]byte, []int) {
	return file_ops_v1_ops_proto_rawDescGZIP(), []int{2}
}

func (x *GetDeploymentStatusRequest) GetRelease() string {
	if x != nil {
		return x.Release
	}
	return ""
}

func (x *GetDeploymentStatusRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type GetDeploymentStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status            string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	FirstDeployed     string `protobuf:"bytes,2,opt,name=first_deployed,json=firstDeployed,proto3" json:"first_deployed,omitempty"`
	LastDeployed      string `protobuf:"bytes,3,opt,name=last_deployed,json=lastDeployed,proto3" json:"last_deployed,omitempty"`
	StatusDescription string `protobuf:"bytes,4,opt,name=status_description,json=statusDescription,proto3" json:"status_description,omitempty"`
}

func (x *GetDeploymentStatusResponse) Reset() {
	*x = GetDeploymentStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_v1_ops_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeploymentStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeploymentStatusResponse) ProtoMessage() {}

func (x *GetDeploymentStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ops_v1_ops_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeploymentStatusResponse.ProtoReflect.Descriptor instead.
func (*GetDeploymentStatusResponse) Descriptor() ([]byte, []int) {
	return file_ops_v1_ops_proto_rawDescGZIP(), []int{3}
}

func (x *GetDeploymentStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetDeploymentStatusResponse) GetFirstDeployed() string {
	if x != nil {
		return x.FirstDeployed
	}
	return ""
}

func (x *GetDeploymentStatusResponse) GetLastDeployed() string {
	if x != nil {
		return x.LastDeployed
	}
	return ""
}

func (x *GetDeploymentStatusResponse) GetStatusDescription() string {
	if x != nil {
		return x.StatusDescription
	}
	return ""
}

type Task struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Command    string `protobuf:"bytes,2,opt,name=command,proto3" json:"command,omitempty"`
	Repository string `protobuf:"bytes,3,opt,name=repository,proto3" json:"repository,omitempty"`
}

func (x *Task) Reset() {
	*x = Task{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_v1_ops_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Task) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Task) ProtoMessage() {}

func (x *Task) ProtoReflect() protoreflect.Message {
	mi := &file_ops_v1_ops_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Task.ProtoReflect.Descriptor instead.
func (*Task) Descriptor() ([]byte, []int) {
	return file_ops_v1_ops_proto_rawDescGZIP(), []int{4}
}

func (x *Task) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Task) GetCommand() string {
	if x != nil {
		return x.Command
	}
	return ""
}

func (x *Task) GetRepository() string {
	if x != nil {
		return x.Repository
	}
	return ""
}

type ListSupportedTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSupportedTasksRequest) Reset() {
	*x = ListSupportedTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_v1_ops_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSupportedTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSupportedTasksRequest) ProtoMessage() {}

func (x *ListSupportedTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_ops_v1_ops_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSupportedTasksRequest.ProtoReflect.Descriptor instead.
func (*ListSupportedTasksRequest) Descriptor() ([]byte, []int) {
	return file_ops_v1_ops_proto_rawDescGZIP(), []int{5}
}

type ListSupportedTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*Task `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *ListSupportedTasksResponse) Reset() {
	*x = ListSupportedTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ops_v1_ops_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSupportedTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSupportedTasksResponse) ProtoMessage() {}

func (x *ListSupportedTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_ops_v1_ops_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSupportedTasksResponse.ProtoReflect.Descriptor instead.
func (*ListSupportedTasksResponse) Descriptor() ([]byte, []int) {
	return file_ops_v1_ops_proto_rawDescGZIP(), []int{6}
}

func (x *ListSupportedTasksResponse) GetTasks() []*Task {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_ops_v1_ops_proto protoreflect.FileDescriptor

var file_ops_v1_ops_proto_rawDesc = []byte{
	0x0a, 0x10, 0x6f, 0x70, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x70, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x6f, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x22, 0x21, 0x0a, 0x1f, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x44, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x44, 0x0a,
	0x20, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x73, 0x22, 0x54, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x6e,
	0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x1b, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f,
	0x79, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x69, 0x72, 0x73, 0x74,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x6c, 0x61, 0x73, 0x74,
	0x5f, 0x64, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x6c, 0x61, 0x73, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x65, 0x64, 0x12, 0x2d, 0x0a,
	0x12, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x04,
	0x54, 0x61, 0x73, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x61, 0x6e, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f,
	0x72, 0x79, 0x22, 0x1b, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x40, 0x0a, 0x1a, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a,
	0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6f,
	0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b,
	0x73, 0x32, 0xbe, 0x02, 0x0a, 0x0a, 0x4f, 0x70, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x6f, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65,
	0x64, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x27, 0x2e, 0x6f,
	0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x44, 0x65, 0x70, 0x6c,
	0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x60, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65,
	0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x22, 0x2e, 0x6f, 0x70, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x6f,
	0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x5d, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x21, 0x2e, 0x6f, 0x70, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64,
	0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x6f,
	0x70, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x44, 0x61, 0x76, 0x69, 0x73, 0x52, 0x61, 0x79, 0x4d, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x67,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ops_v1_ops_proto_rawDescOnce sync.Once
	file_ops_v1_ops_proto_rawDescData = file_ops_v1_ops_proto_rawDesc
)

func file_ops_v1_ops_proto_rawDescGZIP() []byte {
	file_ops_v1_ops_proto_rawDescOnce.Do(func() {
		file_ops_v1_ops_proto_rawDescData = protoimpl.X.CompressGZIP(file_ops_v1_ops_proto_rawDescData)
	})
	return file_ops_v1_ops_proto_rawDescData
}

var file_ops_v1_ops_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_ops_v1_ops_proto_goTypes = []interface{}{
	(*ListSupportedDeploymentsRequest)(nil),  // 0: ops.v1.ListSupportedDeploymentsRequest
	(*ListSupportedDeploymentsResponse)(nil), // 1: ops.v1.ListSupportedDeploymentsResponse
	(*GetDeploymentStatusRequest)(nil),       // 2: ops.v1.GetDeploymentStatusRequest
	(*GetDeploymentStatusResponse)(nil),      // 3: ops.v1.GetDeploymentStatusResponse
	(*Task)(nil),                             // 4: ops.v1.Task
	(*ListSupportedTasksRequest)(nil),        // 5: ops.v1.ListSupportedTasksRequest
	(*ListSupportedTasksResponse)(nil),       // 6: ops.v1.ListSupportedTasksResponse
}
var file_ops_v1_ops_proto_depIdxs = []int32{
	4, // 0: ops.v1.ListSupportedTasksResponse.tasks:type_name -> ops.v1.Task
	0, // 1: ops.v1.OpsService.ListSupportedDeployments:input_type -> ops.v1.ListSupportedDeploymentsRequest
	2, // 2: ops.v1.OpsService.GetDeploymentStatus:input_type -> ops.v1.GetDeploymentStatusRequest
	5, // 3: ops.v1.OpsService.ListSupportedTasks:input_type -> ops.v1.ListSupportedTasksRequest
	1, // 4: ops.v1.OpsService.ListSupportedDeployments:output_type -> ops.v1.ListSupportedDeploymentsResponse
	3, // 5: ops.v1.OpsService.GetDeploymentStatus:output_type -> ops.v1.GetDeploymentStatusResponse
	6, // 6: ops.v1.OpsService.ListSupportedTasks:output_type -> ops.v1.ListSupportedTasksResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_ops_v1_ops_proto_init() }
func file_ops_v1_ops_proto_init() {
	if File_ops_v1_ops_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_ops_v1_ops_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSupportedDeploymentsRequest); i {
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
		file_ops_v1_ops_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSupportedDeploymentsResponse); i {
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
		file_ops_v1_ops_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeploymentStatusRequest); i {
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
		file_ops_v1_ops_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeploymentStatusResponse); i {
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
		file_ops_v1_ops_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Task); i {
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
		file_ops_v1_ops_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSupportedTasksRequest); i {
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
		file_ops_v1_ops_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSupportedTasksResponse); i {
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
			RawDescriptor: file_ops_v1_ops_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ops_v1_ops_proto_goTypes,
		DependencyIndexes: file_ops_v1_ops_proto_depIdxs,
		MessageInfos:      file_ops_v1_ops_proto_msgTypes,
	}.Build()
	File_ops_v1_ops_proto = out.File
	file_ops_v1_ops_proto_rawDesc = nil
	file_ops_v1_ops_proto_goTypes = nil
	file_ops_v1_ops_proto_depIdxs = nil
}
