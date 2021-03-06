// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.9.2
// source: services.proto

package services

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

type BusinessService struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessServiceID   int64   `protobuf:"varint,1,opt,name=businessServiceID,proto3" json:"businessServiceID,omitempty"`
	BusinessServiceName string  `protobuf:"bytes,2,opt,name=businessServiceName,proto3" json:"businessServiceName,omitempty"`
	SubCategories       []int64 `protobuf:"varint,3,rep,packed,name=subCategories,proto3" json:"subCategories,omitempty"`
}

func (x *BusinessService) Reset() {
	*x = BusinessService{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BusinessService) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BusinessService) ProtoMessage() {}

func (x *BusinessService) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BusinessService.ProtoReflect.Descriptor instead.
func (*BusinessService) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{0}
}

func (x *BusinessService) GetBusinessServiceID() int64 {
	if x != nil {
		return x.BusinessServiceID
	}
	return 0
}

func (x *BusinessService) GetBusinessServiceName() string {
	if x != nil {
		return x.BusinessServiceName
	}
	return ""
}

func (x *BusinessService) GetSubCategories() []int64 {
	if x != nil {
		return x.SubCategories
	}
	return nil
}

type GetBusinessServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessServiceID int64 `protobuf:"varint,1,opt,name=businessServiceID,proto3" json:"businessServiceID,omitempty"`
}

func (x *GetBusinessServiceRequest) Reset() {
	*x = GetBusinessServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBusinessServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBusinessServiceRequest) ProtoMessage() {}

func (x *GetBusinessServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBusinessServiceRequest.ProtoReflect.Descriptor instead.
func (*GetBusinessServiceRequest) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{1}
}

func (x *GetBusinessServiceRequest) GetBusinessServiceID() int64 {
	if x != nil {
		return x.BusinessServiceID
	}
	return 0
}

type GetBusinessServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessService *BusinessService `protobuf:"bytes,1,opt,name=businessService,proto3" json:"businessService,omitempty"`
}

func (x *GetBusinessServiceResponse) Reset() {
	*x = GetBusinessServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBusinessServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBusinessServiceResponse) ProtoMessage() {}

func (x *GetBusinessServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBusinessServiceResponse.ProtoReflect.Descriptor instead.
func (*GetBusinessServiceResponse) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{2}
}

func (x *GetBusinessServiceResponse) GetBusinessService() *BusinessService {
	if x != nil {
		return x.BusinessService
	}
	return nil
}

type GetBusinessServicesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessServices []*BusinessService `protobuf:"bytes,1,rep,name=businessServices,proto3" json:"businessServices,omitempty"`
}

func (x *GetBusinessServicesResponse) Reset() {
	*x = GetBusinessServicesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBusinessServicesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBusinessServicesResponse) ProtoMessage() {}

func (x *GetBusinessServicesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBusinessServicesResponse.ProtoReflect.Descriptor instead.
func (*GetBusinessServicesResponse) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{3}
}

func (x *GetBusinessServicesResponse) GetBusinessServices() []*BusinessService {
	if x != nil {
		return x.BusinessServices
	}
	return nil
}

type CreateBusinessServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessServiceName string  `protobuf:"bytes,1,opt,name=businessServiceName,proto3" json:"businessServiceName,omitempty"`
	SubCategories       []int64 `protobuf:"varint,2,rep,packed,name=subCategories,proto3" json:"subCategories,omitempty"`
}

func (x *CreateBusinessServiceRequest) Reset() {
	*x = CreateBusinessServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBusinessServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBusinessServiceRequest) ProtoMessage() {}

func (x *CreateBusinessServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBusinessServiceRequest.ProtoReflect.Descriptor instead.
func (*CreateBusinessServiceRequest) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{4}
}

func (x *CreateBusinessServiceRequest) GetBusinessServiceName() string {
	if x != nil {
		return x.BusinessServiceName
	}
	return ""
}

func (x *CreateBusinessServiceRequest) GetSubCategories() []int64 {
	if x != nil {
		return x.SubCategories
	}
	return nil
}

type CreateBusinessServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessService *BusinessService `protobuf:"bytes,1,opt,name=businessService,proto3" json:"businessService,omitempty"`
}

func (x *CreateBusinessServiceResponse) Reset() {
	*x = CreateBusinessServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBusinessServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBusinessServiceResponse) ProtoMessage() {}

func (x *CreateBusinessServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBusinessServiceResponse.ProtoReflect.Descriptor instead.
func (*CreateBusinessServiceResponse) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{5}
}

func (x *CreateBusinessServiceResponse) GetBusinessService() *BusinessService {
	if x != nil {
		return x.BusinessService
	}
	return nil
}

type GetServicesUnderSubCategoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SubCategoryID int64 `protobuf:"varint,1,opt,name=SubCategoryID,proto3" json:"SubCategoryID,omitempty"`
}

func (x *GetServicesUnderSubCategoryRequest) Reset() {
	*x = GetServicesUnderSubCategoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServicesUnderSubCategoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicesUnderSubCategoryRequest) ProtoMessage() {}

func (x *GetServicesUnderSubCategoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicesUnderSubCategoryRequest.ProtoReflect.Descriptor instead.
func (*GetServicesUnderSubCategoryRequest) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{6}
}

func (x *GetServicesUnderSubCategoryRequest) GetSubCategoryID() int64 {
	if x != nil {
		return x.SubCategoryID
	}
	return 0
}

type GetServicesUnderSubCategoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BusinessServices []*BusinessService `protobuf:"bytes,1,rep,name=businessServices,proto3" json:"businessServices,omitempty"`
}

func (x *GetServicesUnderSubCategoryResponse) Reset() {
	*x = GetServicesUnderSubCategoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_services_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetServicesUnderSubCategoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetServicesUnderSubCategoryResponse) ProtoMessage() {}

func (x *GetServicesUnderSubCategoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_services_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetServicesUnderSubCategoryResponse.ProtoReflect.Descriptor instead.
func (*GetServicesUnderSubCategoryResponse) Descriptor() ([]byte, []int) {
	return file_services_proto_rawDescGZIP(), []int{7}
}

func (x *GetServicesUnderSubCategoryResponse) GetBusinessServices() []*BusinessService {
	if x != nil {
		return x.BusinessServices
	}
	return nil
}

var File_services_proto protoreflect.FileDescriptor

var file_services_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x97, 0x01, 0x0a, 0x0f, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x11, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x11,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x44, 0x12, 0x30, 0x0a, 0x13, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x43,
	0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x49, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x11, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x11, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x44, 0x22, 0x69, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x75,
	0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x0f,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22,
	0x6c, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d,
	0x0a, 0x10, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e,
	0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x10, 0x62, 0x75, 0x73,
	0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x22, 0x76, 0x0a,
	0x1c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a,
	0x13, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x62, 0x75, 0x73, 0x69,
	0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12,
	0x24, 0x0a, 0x0d, 0x73, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x6c, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x52, 0x0f, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x4a, 0x0a, 0x22, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x53, 0x75, 0x62,
	0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x53, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x49, 0x44, 0x22,
	0x74, 0x0a, 0x23, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x55, 0x6e,
	0x64, 0x65, 0x72, 0x53, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a, 0x10, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x21, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x10, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x32, 0xf0, 0x03, 0x0a, 0x10, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65,
	0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x71, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x2b, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x5e, 0x0a,
	0x13, 0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x2d, 0x2e, 0x62,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x47, 0x65, 0x74, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x8c, 0x01,
	0x0a, 0x1b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x55, 0x6e, 0x64,
	0x65, 0x72, 0x53, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x34, 0x2e,
	0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x55, 0x6e, 0x64, 0x65,
	0x72, 0x53, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x73, 0x55, 0x6e, 0x64, 0x65, 0x72, 0x53, 0x75, 0x62, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7a, 0x0a, 0x15,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x62, 0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42,
	0x75, 0x73, 0x69, 0x6e, 0x65, 0x73, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x2e, 0x3b, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_services_proto_rawDescOnce sync.Once
	file_services_proto_rawDescData = file_services_proto_rawDesc
)

func file_services_proto_rawDescGZIP() []byte {
	file_services_proto_rawDescOnce.Do(func() {
		file_services_proto_rawDescData = protoimpl.X.CompressGZIP(file_services_proto_rawDescData)
	})
	return file_services_proto_rawDescData
}

var file_services_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_services_proto_goTypes = []interface{}{
	(*BusinessService)(nil),                     // 0: businessServices.BusinessService
	(*GetBusinessServiceRequest)(nil),           // 1: businessServices.GetBusinessServiceRequest
	(*GetBusinessServiceResponse)(nil),          // 2: businessServices.GetBusinessServiceResponse
	(*GetBusinessServicesResponse)(nil),         // 3: businessServices.GetBusinessServicesResponse
	(*CreateBusinessServiceRequest)(nil),        // 4: businessServices.CreateBusinessServiceRequest
	(*CreateBusinessServiceResponse)(nil),       // 5: businessServices.CreateBusinessServiceResponse
	(*GetServicesUnderSubCategoryRequest)(nil),  // 6: businessServices.GetServicesUnderSubCategoryRequest
	(*GetServicesUnderSubCategoryResponse)(nil), // 7: businessServices.GetServicesUnderSubCategoryResponse
	(*empty.Empty)(nil),                         // 8: google.protobuf.Empty
}
var file_services_proto_depIdxs = []int32{
	0, // 0: businessServices.GetBusinessServiceResponse.businessService:type_name -> businessServices.BusinessService
	0, // 1: businessServices.GetBusinessServicesResponse.businessServices:type_name -> businessServices.BusinessService
	0, // 2: businessServices.CreateBusinessServiceResponse.businessService:type_name -> businessServices.BusinessService
	0, // 3: businessServices.GetServicesUnderSubCategoryResponse.businessServices:type_name -> businessServices.BusinessService
	1, // 4: businessServices.BusinessServices.GetBusinessService:input_type -> businessServices.GetBusinessServiceRequest
	8, // 5: businessServices.BusinessServices.GetBusinessServices:input_type -> google.protobuf.Empty
	6, // 6: businessServices.BusinessServices.GetServicesUnderSubCategory:input_type -> businessServices.GetServicesUnderSubCategoryRequest
	4, // 7: businessServices.BusinessServices.CreateBusinessService:input_type -> businessServices.CreateBusinessServiceRequest
	2, // 8: businessServices.BusinessServices.GetBusinessService:output_type -> businessServices.GetBusinessServiceResponse
	3, // 9: businessServices.BusinessServices.GetBusinessServices:output_type -> businessServices.GetBusinessServicesResponse
	7, // 10: businessServices.BusinessServices.GetServicesUnderSubCategory:output_type -> businessServices.GetServicesUnderSubCategoryResponse
	5, // 11: businessServices.BusinessServices.CreateBusinessService:output_type -> businessServices.CreateBusinessServiceResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_services_proto_init() }
func file_services_proto_init() {
	if File_services_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_services_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BusinessService); i {
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
		file_services_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBusinessServiceRequest); i {
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
		file_services_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBusinessServiceResponse); i {
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
		file_services_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBusinessServicesResponse); i {
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
		file_services_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBusinessServiceRequest); i {
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
		file_services_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBusinessServiceResponse); i {
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
		file_services_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServicesUnderSubCategoryRequest); i {
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
		file_services_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetServicesUnderSubCategoryResponse); i {
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
			RawDescriptor: file_services_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_services_proto_goTypes,
		DependencyIndexes: file_services_proto_depIdxs,
		MessageInfos:      file_services_proto_msgTypes,
	}.Build()
	File_services_proto = out.File
	file_services_proto_rawDesc = nil
	file_services_proto_goTypes = nil
	file_services_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// BusinessServicesClient is the client API for BusinessServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BusinessServicesClient interface {
	GetBusinessService(ctx context.Context, in *GetBusinessServiceRequest, opts ...grpc.CallOption) (*GetBusinessServiceResponse, error)
	GetBusinessServices(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetBusinessServicesResponse, error)
	GetServicesUnderSubCategory(ctx context.Context, in *GetServicesUnderSubCategoryRequest, opts ...grpc.CallOption) (*GetServicesUnderSubCategoryResponse, error)
	CreateBusinessService(ctx context.Context, in *CreateBusinessServiceRequest, opts ...grpc.CallOption) (*CreateBusinessServiceResponse, error)
}

type businessServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewBusinessServicesClient(cc grpc.ClientConnInterface) BusinessServicesClient {
	return &businessServicesClient{cc}
}

func (c *businessServicesClient) GetBusinessService(ctx context.Context, in *GetBusinessServiceRequest, opts ...grpc.CallOption) (*GetBusinessServiceResponse, error) {
	out := new(GetBusinessServiceResponse)
	err := c.cc.Invoke(ctx, "/businessServices.BusinessServices/GetBusinessService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessServicesClient) GetBusinessServices(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*GetBusinessServicesResponse, error) {
	out := new(GetBusinessServicesResponse)
	err := c.cc.Invoke(ctx, "/businessServices.BusinessServices/GetBusinessServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessServicesClient) GetServicesUnderSubCategory(ctx context.Context, in *GetServicesUnderSubCategoryRequest, opts ...grpc.CallOption) (*GetServicesUnderSubCategoryResponse, error) {
	out := new(GetServicesUnderSubCategoryResponse)
	err := c.cc.Invoke(ctx, "/businessServices.BusinessServices/GetServicesUnderSubCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessServicesClient) CreateBusinessService(ctx context.Context, in *CreateBusinessServiceRequest, opts ...grpc.CallOption) (*CreateBusinessServiceResponse, error) {
	out := new(CreateBusinessServiceResponse)
	err := c.cc.Invoke(ctx, "/businessServices.BusinessServices/CreateBusinessService", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BusinessServicesServer is the server API for BusinessServices service.
type BusinessServicesServer interface {
	GetBusinessService(context.Context, *GetBusinessServiceRequest) (*GetBusinessServiceResponse, error)
	GetBusinessServices(context.Context, *empty.Empty) (*GetBusinessServicesResponse, error)
	GetServicesUnderSubCategory(context.Context, *GetServicesUnderSubCategoryRequest) (*GetServicesUnderSubCategoryResponse, error)
	CreateBusinessService(context.Context, *CreateBusinessServiceRequest) (*CreateBusinessServiceResponse, error)
}

// UnimplementedBusinessServicesServer can be embedded to have forward compatible implementations.
type UnimplementedBusinessServicesServer struct {
}

func (*UnimplementedBusinessServicesServer) GetBusinessService(context.Context, *GetBusinessServiceRequest) (*GetBusinessServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusinessService not implemented")
}
func (*UnimplementedBusinessServicesServer) GetBusinessServices(context.Context, *empty.Empty) (*GetBusinessServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusinessServices not implemented")
}
func (*UnimplementedBusinessServicesServer) GetServicesUnderSubCategory(context.Context, *GetServicesUnderSubCategoryRequest) (*GetServicesUnderSubCategoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServicesUnderSubCategory not implemented")
}
func (*UnimplementedBusinessServicesServer) CreateBusinessService(context.Context, *CreateBusinessServiceRequest) (*CreateBusinessServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBusinessService not implemented")
}

func RegisterBusinessServicesServer(s *grpc.Server, srv BusinessServicesServer) {
	s.RegisterService(&_BusinessServices_serviceDesc, srv)
}

func _BusinessServices_GetBusinessService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBusinessServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServicesServer).GetBusinessService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/businessServices.BusinessServices/GetBusinessService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServicesServer).GetBusinessService(ctx, req.(*GetBusinessServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessServices_GetBusinessServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServicesServer).GetBusinessServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/businessServices.BusinessServices/GetBusinessServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServicesServer).GetBusinessServices(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessServices_GetServicesUnderSubCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServicesUnderSubCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServicesServer).GetServicesUnderSubCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/businessServices.BusinessServices/GetServicesUnderSubCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServicesServer).GetServicesUnderSubCategory(ctx, req.(*GetServicesUnderSubCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessServices_CreateBusinessService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBusinessServiceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServicesServer).CreateBusinessService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/businessServices.BusinessServices/CreateBusinessService",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServicesServer).CreateBusinessService(ctx, req.(*CreateBusinessServiceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _BusinessServices_serviceDesc = grpc.ServiceDesc{
	ServiceName: "businessServices.BusinessServices",
	HandlerType: (*BusinessServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBusinessService",
			Handler:    _BusinessServices_GetBusinessService_Handler,
		},
		{
			MethodName: "GetBusinessServices",
			Handler:    _BusinessServices_GetBusinessServices_Handler,
		},
		{
			MethodName: "GetServicesUnderSubCategory",
			Handler:    _BusinessServices_GetServicesUnderSubCategory_Handler,
		},
		{
			MethodName: "CreateBusinessService",
			Handler:    _BusinessServices_CreateBusinessService_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "services.proto",
}
