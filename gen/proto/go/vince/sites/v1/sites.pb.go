// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: vince/sites/v1/sites.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	v11 "github.com/vinceanalytics/vince/gen/proto/go/vince/blocks/v1"
	_ "github.com/vinceanalytics/vince/gen/proto/go/vince/config/v1"
	v1 "github.com/vinceanalytics/vince/gen/proto/go/vince/goals/v1"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/emptypb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Site struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain      string              `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Goals       map[string]*v1.Goal `protobuf:"bytes,2,rep,name=goals,proto3" json:"goals,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Description string              `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	BaseStats   *v11.BaseStats      `protobuf:"bytes,4,opt,name=base_stats,json=baseStats,proto3" json:"base_stats,omitempty"`
}

func (x *Site) Reset() {
	*x = Site{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_sites_v1_sites_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Site) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Site) ProtoMessage() {}

func (x *Site) ProtoReflect() protoreflect.Message {
	mi := &file_vince_sites_v1_sites_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Site.ProtoReflect.Descriptor instead.
func (*Site) Descriptor() ([]byte, []int) {
	return file_vince_sites_v1_sites_proto_rawDescGZIP(), []int{0}
}

func (x *Site) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Site) GetGoals() map[string]*v1.Goal {
	if x != nil {
		return x.Goals
	}
	return nil
}

func (x *Site) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Site) GetBaseStats() *v11.BaseStats {
	if x != nil {
		return x.BaseStats
	}
	return nil
}

type CreateSiteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain      string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *CreateSiteRequest) Reset() {
	*x = CreateSiteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_sites_v1_sites_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSiteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSiteRequest) ProtoMessage() {}

func (x *CreateSiteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vince_sites_v1_sites_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSiteRequest.ProtoReflect.Descriptor instead.
func (*CreateSiteRequest) Descriptor() ([]byte, []int) {
	return file_vince_sites_v1_sites_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSiteRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *CreateSiteRequest) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CreateSiteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Site *Site `protobuf:"bytes,1,opt,name=site,proto3" json:"site,omitempty"`
}

func (x *CreateSiteResponse) Reset() {
	*x = CreateSiteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_sites_v1_sites_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSiteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSiteResponse) ProtoMessage() {}

func (x *CreateSiteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vince_sites_v1_sites_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSiteResponse.ProtoReflect.Descriptor instead.
func (*CreateSiteResponse) Descriptor() ([]byte, []int) {
	return file_vince_sites_v1_sites_proto_rawDescGZIP(), []int{2}
}

func (x *CreateSiteResponse) GetSite() *Site {
	if x != nil {
		return x.Site
	}
	return nil
}

type GetSiteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *GetSiteRequest) Reset() {
	*x = GetSiteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_sites_v1_sites_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSiteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSiteRequest) ProtoMessage() {}

func (x *GetSiteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vince_sites_v1_sites_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSiteRequest.ProtoReflect.Descriptor instead.
func (*GetSiteRequest) Descriptor() ([]byte, []int) {
	return file_vince_sites_v1_sites_proto_rawDescGZIP(), []int{3}
}

func (x *GetSiteRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type ListSitesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListSitesRequest) Reset() {
	*x = ListSitesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_sites_v1_sites_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSitesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSitesRequest) ProtoMessage() {}

func (x *ListSitesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vince_sites_v1_sites_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSitesRequest.ProtoReflect.Descriptor instead.
func (*ListSitesRequest) Descriptor() ([]byte, []int) {
	return file_vince_sites_v1_sites_proto_rawDescGZIP(), []int{4}
}

type ListSitesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Site `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListSitesResponse) Reset() {
	*x = ListSitesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_sites_v1_sites_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSitesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSitesResponse) ProtoMessage() {}

func (x *ListSitesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vince_sites_v1_sites_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSitesResponse.ProtoReflect.Descriptor instead.
func (*ListSitesResponse) Descriptor() ([]byte, []int) {
	return file_vince_sites_v1_sites_proto_rawDescGZIP(), []int{5}
}

func (x *ListSitesResponse) GetList() []*Site {
	if x != nil {
		return x.List
	}
	return nil
}

type DeleteSiteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Domain string `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
}

func (x *DeleteSiteRequest) Reset() {
	*x = DeleteSiteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_sites_v1_sites_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSiteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSiteRequest) ProtoMessage() {}

func (x *DeleteSiteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vince_sites_v1_sites_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSiteRequest.ProtoReflect.Descriptor instead.
func (*DeleteSiteRequest) Descriptor() ([]byte, []int) {
	return file_vince_sites_v1_sites_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteSiteRequest) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

type DeleteSiteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSiteResponse) Reset() {
	*x = DeleteSiteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vince_sites_v1_sites_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSiteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSiteResponse) ProtoMessage() {}

func (x *DeleteSiteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_vince_sites_v1_sites_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSiteResponse.ProtoReflect.Descriptor instead.
func (*DeleteSiteResponse) Descriptor() ([]byte, []int) {
	return file_vince_sites_v1_sites_proto_rawDescGZIP(), []int{7}
}

var File_vince_sites_v1_sites_proto protoreflect.FileDescriptor

var file_vince_sites_v1_sites_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x69, 0x74, 0x65, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x69, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31,
	0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x1c, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a,
	0x76, 0x69, 0x6e, 0x63, 0x65, 0x2f, 0x67, 0x6f, 0x61, 0x6c, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x67,
	0x6f, 0x61, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x2f, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xdd, 0x01, 0x0a, 0x04, 0x53, 0x69, 0x74, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61,
	0x69, 0x6e, 0x12, 0x29, 0x0a, 0x05, 0x67, 0x6f, 0x61, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x13, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x74, 0x65, 0x2e, 0x47, 0x6f, 0x61, 0x6c,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x67, 0x6f, 0x61, 0x6c, 0x73, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2c, 0x0a, 0x0a, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x76, 0x31, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x73, 0x52, 0x09, 0x62, 0x61, 0x73, 0x65, 0x53, 0x74, 0x61, 0x74, 0x73, 0x1a, 0x42, 0x0a,
	0x0a, 0x47, 0x6f, 0x61, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x1e, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x76,
	0x31, 0x2e, 0x47, 0x6f, 0x61, 0x6c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x59, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0xc8, 0x01, 0x01, 0x72, 0x02,
	0x68, 0x01, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65,
	0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1c, 0x0a, 0x04, 0x73, 0x69, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x74, 0x65, 0x52, 0x04, 0x73, 0x69, 0x74, 0x65,
	0x22, 0x34, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0xc8, 0x01, 0x01, 0x72, 0x02, 0x68, 0x01, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x12, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69,
	0x74, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x31, 0x0a, 0x11, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x1c, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x08, 0x2e,
	0x76, 0x31, 0x2e, 0x53, 0x69, 0x74, 0x65, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x22, 0x37, 0x0a,
	0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x22, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0xc8, 0x01, 0x01, 0x72, 0x02, 0x68, 0x01, 0x52, 0x06,
	0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xae, 0x02, 0x0a,
	0x05, 0x53, 0x69, 0x74, 0x65, 0x73, 0x12, 0x4e, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x69, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x22, 0x09, 0x2f, 0x76, 0x31,
	0x2f, 0x73, 0x69, 0x74, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x53, 0x69, 0x74,
	0x65, 0x12, 0x12, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x69, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x74, 0x65, 0x22,
	0x11, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x74,
	0x65, 0x73, 0x12, 0x4a, 0x0a, 0x09, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x74, 0x65, 0x73, 0x12,
	0x14, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x69, 0x74, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53,
	0x69, 0x74, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3,
	0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x74, 0x65, 0x12, 0x4d,
	0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x69, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0a, 0x2a, 0x08, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69, 0x74, 0x65, 0x42, 0x79, 0x0a,
	0x06, 0x63, 0x6f, 0x6d, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x53, 0x69, 0x74, 0x65, 0x73, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x2f, 0x73, 0x69, 0x74, 0x65, 0x73, 0x2f,
	0x76, 0x31, 0xa2, 0x02, 0x03, 0x56, 0x58, 0x58, 0xaa, 0x02, 0x02, 0x56, 0x31, 0xca, 0x02, 0x02,
	0x56, 0x31, 0xe2, 0x02, 0x0e, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x02, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vince_sites_v1_sites_proto_rawDescOnce sync.Once
	file_vince_sites_v1_sites_proto_rawDescData = file_vince_sites_v1_sites_proto_rawDesc
)

func file_vince_sites_v1_sites_proto_rawDescGZIP() []byte {
	file_vince_sites_v1_sites_proto_rawDescOnce.Do(func() {
		file_vince_sites_v1_sites_proto_rawDescData = protoimpl.X.CompressGZIP(file_vince_sites_v1_sites_proto_rawDescData)
	})
	return file_vince_sites_v1_sites_proto_rawDescData
}

var file_vince_sites_v1_sites_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_vince_sites_v1_sites_proto_goTypes = []interface{}{
	(*Site)(nil),               // 0: v1.Site
	(*CreateSiteRequest)(nil),  // 1: v1.CreateSiteRequest
	(*CreateSiteResponse)(nil), // 2: v1.CreateSiteResponse
	(*GetSiteRequest)(nil),     // 3: v1.GetSiteRequest
	(*ListSitesRequest)(nil),   // 4: v1.ListSitesRequest
	(*ListSitesResponse)(nil),  // 5: v1.ListSitesResponse
	(*DeleteSiteRequest)(nil),  // 6: v1.DeleteSiteRequest
	(*DeleteSiteResponse)(nil), // 7: v1.DeleteSiteResponse
	nil,                        // 8: v1.Site.GoalsEntry
	(*v11.BaseStats)(nil),      // 9: v1.BaseStats
	(*v1.Goal)(nil),            // 10: v1.Goal
}
var file_vince_sites_v1_sites_proto_depIdxs = []int32{
	8,  // 0: v1.Site.goals:type_name -> v1.Site.GoalsEntry
	9,  // 1: v1.Site.base_stats:type_name -> v1.BaseStats
	0,  // 2: v1.CreateSiteResponse.site:type_name -> v1.Site
	0,  // 3: v1.ListSitesResponse.list:type_name -> v1.Site
	10, // 4: v1.Site.GoalsEntry.value:type_name -> v1.Goal
	1,  // 5: v1.Sites.CreateSite:input_type -> v1.CreateSiteRequest
	3,  // 6: v1.Sites.GetSite:input_type -> v1.GetSiteRequest
	4,  // 7: v1.Sites.ListSites:input_type -> v1.ListSitesRequest
	6,  // 8: v1.Sites.DeleteSite:input_type -> v1.DeleteSiteRequest
	2,  // 9: v1.Sites.CreateSite:output_type -> v1.CreateSiteResponse
	0,  // 10: v1.Sites.GetSite:output_type -> v1.Site
	5,  // 11: v1.Sites.ListSites:output_type -> v1.ListSitesResponse
	7,  // 12: v1.Sites.DeleteSite:output_type -> v1.DeleteSiteResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_vince_sites_v1_sites_proto_init() }
func file_vince_sites_v1_sites_proto_init() {
	if File_vince_sites_v1_sites_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vince_sites_v1_sites_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Site); i {
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
		file_vince_sites_v1_sites_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSiteRequest); i {
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
		file_vince_sites_v1_sites_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSiteResponse); i {
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
		file_vince_sites_v1_sites_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSiteRequest); i {
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
		file_vince_sites_v1_sites_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSitesRequest); i {
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
		file_vince_sites_v1_sites_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSitesResponse); i {
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
		file_vince_sites_v1_sites_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSiteRequest); i {
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
		file_vince_sites_v1_sites_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSiteResponse); i {
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
			RawDescriptor: file_vince_sites_v1_sites_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_vince_sites_v1_sites_proto_goTypes,
		DependencyIndexes: file_vince_sites_v1_sites_proto_depIdxs,
		MessageInfos:      file_vince_sites_v1_sites_proto_msgTypes,
	}.Build()
	File_vince_sites_v1_sites_proto = out.File
	file_vince_sites_v1_sites_proto_rawDesc = nil
	file_vince_sites_v1_sites_proto_goTypes = nil
	file_vince_sites_v1_sites_proto_depIdxs = nil
}
