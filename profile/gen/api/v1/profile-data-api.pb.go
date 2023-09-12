// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: api/v1/profile-data-api.proto

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

type ProfileData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address string `protobuf:"bytes,2,opt,name=Address,proto3" json:"Address,omitempty"`
	Name    string `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *ProfileData) Reset() {
	*x = ProfileData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_profile_data_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileData) ProtoMessage() {}

func (x *ProfileData) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_profile_data_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileData.ProtoReflect.Descriptor instead.
func (*ProfileData) Descriptor() ([]byte, []int) {
	return file_api_v1_profile_data_api_proto_rawDescGZIP(), []int{0}
}

func (x *ProfileData) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ProfileData) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *ProfileData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type CreateProfileDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileData *ProfileData `protobuf:"bytes,1,opt,name=profileData,proto3" json:"profileData,omitempty"`
}

func (x *CreateProfileDataRequest) Reset() {
	*x = CreateProfileDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_profile_data_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProfileDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfileDataRequest) ProtoMessage() {}

func (x *CreateProfileDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_profile_data_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfileDataRequest.ProtoReflect.Descriptor instead.
func (*CreateProfileDataRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_profile_data_api_proto_rawDescGZIP(), []int{1}
}

func (x *CreateProfileDataRequest) GetProfileData() *ProfileData {
	if x != nil {
		return x.ProfileData
	}
	return nil
}

type CreateProfileDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *CreateProfileDataResponse) Reset() {
	*x = CreateProfileDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_profile_data_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProfileDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProfileDataResponse) ProtoMessage() {}

func (x *CreateProfileDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_profile_data_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProfileDataResponse.ProtoReflect.Descriptor instead.
func (*CreateProfileDataResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_profile_data_api_proto_rawDescGZIP(), []int{2}
}

func (x *CreateProfileDataResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetProfileDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetProfileDataRequest) Reset() {
	*x = GetProfileDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_profile_data_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileDataRequest) ProtoMessage() {}

func (x *GetProfileDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_profile_data_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileDataRequest.ProtoReflect.Descriptor instead.
func (*GetProfileDataRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_profile_data_api_proto_rawDescGZIP(), []int{3}
}

func (x *GetProfileDataRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetProfileDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileData *ProfileData `protobuf:"bytes,1,opt,name=profileData,proto3" json:"profileData,omitempty"`
}

func (x *GetProfileDataResponse) Reset() {
	*x = GetProfileDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_profile_data_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProfileDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProfileDataResponse) ProtoMessage() {}

func (x *GetProfileDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_profile_data_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProfileDataResponse.ProtoReflect.Descriptor instead.
func (*GetProfileDataResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_profile_data_api_proto_rawDescGZIP(), []int{4}
}

func (x *GetProfileDataResponse) GetProfileData() *ProfileData {
	if x != nil {
		return x.ProfileData
	}
	return nil
}

type UpdateProfileDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ProfileData *ProfileData `protobuf:"bytes,2,opt,name=profileData,proto3" json:"profileData,omitempty"`
}

func (x *UpdateProfileDataRequest) Reset() {
	*x = UpdateProfileDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_profile_data_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProfileDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileDataRequest) ProtoMessage() {}

func (x *UpdateProfileDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_profile_data_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileDataRequest.ProtoReflect.Descriptor instead.
func (*UpdateProfileDataRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_profile_data_api_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateProfileDataRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateProfileDataRequest) GetProfileData() *ProfileData {
	if x != nil {
		return x.ProfileData
	}
	return nil
}

type UpdateProfileDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UpdateProfileDataResponse) Reset() {
	*x = UpdateProfileDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_profile_data_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateProfileDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateProfileDataResponse) ProtoMessage() {}

func (x *UpdateProfileDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_profile_data_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateProfileDataResponse.ProtoReflect.Descriptor instead.
func (*UpdateProfileDataResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_profile_data_api_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateProfileDataResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DeleteProfileDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteProfileDataRequest) Reset() {
	*x = DeleteProfileDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_profile_data_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProfileDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProfileDataRequest) ProtoMessage() {}

func (x *DeleteProfileDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_profile_data_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProfileDataRequest.ProtoReflect.Descriptor instead.
func (*DeleteProfileDataRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_profile_data_api_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteProfileDataRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteProfileDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *DeleteProfileDataResponse) Reset() {
	*x = DeleteProfileDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_profile_data_api_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteProfileDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteProfileDataResponse) ProtoMessage() {}

func (x *DeleteProfileDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_profile_data_api_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteProfileDataResponse.ProtoReflect.Descriptor instead.
func (*DeleteProfileDataResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_profile_data_api_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteProfileDataResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type ListProfileDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListProfileDataRequest) Reset() {
	*x = ListProfileDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_profile_data_api_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProfileDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProfileDataRequest) ProtoMessage() {}

func (x *ListProfileDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_profile_data_api_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProfileDataRequest.ProtoReflect.Descriptor instead.
func (*ListProfileDataRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_profile_data_api_proto_rawDescGZIP(), []int{9}
}

type ListProfileDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProfileData []*ProfileData `protobuf:"bytes,1,rep,name=profileData,proto3" json:"profileData,omitempty"`
}

func (x *ListProfileDataResponse) Reset() {
	*x = ListProfileDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_profile_data_api_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListProfileDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListProfileDataResponse) ProtoMessage() {}

func (x *ListProfileDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_profile_data_api_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListProfileDataResponse.ProtoReflect.Descriptor instead.
func (*ListProfileDataResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_profile_data_api_proto_rawDescGZIP(), []int{10}
}

func (x *ListProfileDataResponse) GetProfileData() []*ProfileData {
	if x != nil {
		return x.ProfileData
	}
	return nil
}

type ProfileDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ProfileDataRequest) Reset() {
	*x = ProfileDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_profile_data_api_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileDataRequest) ProtoMessage() {}

func (x *ProfileDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_profile_data_api_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileDataRequest.ProtoReflect.Descriptor instead.
func (*ProfileDataRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_profile_data_api_proto_rawDescGZIP(), []int{11}
}

type ProfileDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ProfileDataResponse) Reset() {
	*x = ProfileDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_v1_profile_data_api_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProfileDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProfileDataResponse) ProtoMessage() {}

func (x *ProfileDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_profile_data_api_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProfileDataResponse.ProtoReflect.Descriptor instead.
func (*ProfileDataResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_profile_data_api_proto_rawDescGZIP(), []int{12}
}

func (x *ProfileDataResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_api_v1_profile_data_api_proto protoreflect.FileDescriptor

var file_api_v1_profile_data_api_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x2d, 0x64, 0x61, 0x74, 0x61, 0x2d, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x02, 0x70, 0x62, 0x22, 0x4b, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x4d, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x31, 0x0a, 0x0b,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22,
	0x35, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x4b, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f,
	0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x5d, 0x0a, 0x18,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x31, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0b,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x35, 0x0a, 0x19, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x2a, 0x0a, 0x18, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x35,
	0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x18, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x4c, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0b, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x22, 0x14, 0x0a,
	0x12, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x2f, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x32, 0xe4, 0x03, 0x0a, 0x12, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x12, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x2e, 0x70,
	0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c, 0x2e, 0x70, 0x62, 0x2e,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x11, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1c,
	0x2e, 0x70, 0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70,
	0x62, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44,
	0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a,
	0x0f, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x12, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0d, 0x5a, 0x0b, 0x2e,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_api_v1_profile_data_api_proto_rawDescOnce sync.Once
	file_api_v1_profile_data_api_proto_rawDescData = file_api_v1_profile_data_api_proto_rawDesc
)

func file_api_v1_profile_data_api_proto_rawDescGZIP() []byte {
	file_api_v1_profile_data_api_proto_rawDescOnce.Do(func() {
		file_api_v1_profile_data_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_v1_profile_data_api_proto_rawDescData)
	})
	return file_api_v1_profile_data_api_proto_rawDescData
}

var file_api_v1_profile_data_api_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_api_v1_profile_data_api_proto_goTypes = []interface{}{
	(*ProfileData)(nil),               // 0: pb.ProfileData
	(*CreateProfileDataRequest)(nil),  // 1: pb.CreateProfileDataRequest
	(*CreateProfileDataResponse)(nil), // 2: pb.CreateProfileDataResponse
	(*GetProfileDataRequest)(nil),     // 3: pb.GetProfileDataRequest
	(*GetProfileDataResponse)(nil),    // 4: pb.GetProfileDataResponse
	(*UpdateProfileDataRequest)(nil),  // 5: pb.UpdateProfileDataRequest
	(*UpdateProfileDataResponse)(nil), // 6: pb.UpdateProfileDataResponse
	(*DeleteProfileDataRequest)(nil),  // 7: pb.DeleteProfileDataRequest
	(*DeleteProfileDataResponse)(nil), // 8: pb.DeleteProfileDataResponse
	(*ListProfileDataRequest)(nil),    // 9: pb.ListProfileDataRequest
	(*ListProfileDataResponse)(nil),   // 10: pb.ListProfileDataResponse
	(*ProfileDataRequest)(nil),        // 11: pb.ProfileDataRequest
	(*ProfileDataResponse)(nil),       // 12: pb.ProfileDataResponse
}
var file_api_v1_profile_data_api_proto_depIdxs = []int32{
	0,  // 0: pb.CreateProfileDataRequest.profileData:type_name -> pb.ProfileData
	0,  // 1: pb.GetProfileDataResponse.profileData:type_name -> pb.ProfileData
	0,  // 2: pb.UpdateProfileDataRequest.profileData:type_name -> pb.ProfileData
	0,  // 3: pb.ListProfileDataResponse.profileData:type_name -> pb.ProfileData
	11, // 4: pb.ProfileDataService.Ping:input_type -> pb.ProfileDataRequest
	1,  // 5: pb.ProfileDataService.CreateProfileData:input_type -> pb.CreateProfileDataRequest
	3,  // 6: pb.ProfileDataService.GetProfileData:input_type -> pb.GetProfileDataRequest
	5,  // 7: pb.ProfileDataService.UpdateProfileData:input_type -> pb.UpdateProfileDataRequest
	7,  // 8: pb.ProfileDataService.DeleteProfileData:input_type -> pb.DeleteProfileDataRequest
	9,  // 9: pb.ProfileDataService.ListProfileData:input_type -> pb.ListProfileDataRequest
	12, // 10: pb.ProfileDataService.Ping:output_type -> pb.ProfileDataResponse
	2,  // 11: pb.ProfileDataService.CreateProfileData:output_type -> pb.CreateProfileDataResponse
	4,  // 12: pb.ProfileDataService.GetProfileData:output_type -> pb.GetProfileDataResponse
	6,  // 13: pb.ProfileDataService.UpdateProfileData:output_type -> pb.UpdateProfileDataResponse
	8,  // 14: pb.ProfileDataService.DeleteProfileData:output_type -> pb.DeleteProfileDataResponse
	10, // 15: pb.ProfileDataService.ListProfileData:output_type -> pb.ListProfileDataResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_v1_profile_data_api_proto_init() }
func file_api_v1_profile_data_api_proto_init() {
	if File_api_v1_profile_data_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_v1_profile_data_api_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileData); i {
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
		file_api_v1_profile_data_api_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProfileDataRequest); i {
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
		file_api_v1_profile_data_api_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProfileDataResponse); i {
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
		file_api_v1_profile_data_api_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileDataRequest); i {
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
		file_api_v1_profile_data_api_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProfileDataResponse); i {
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
		file_api_v1_profile_data_api_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProfileDataRequest); i {
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
		file_api_v1_profile_data_api_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateProfileDataResponse); i {
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
		file_api_v1_profile_data_api_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProfileDataRequest); i {
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
		file_api_v1_profile_data_api_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteProfileDataResponse); i {
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
		file_api_v1_profile_data_api_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProfileDataRequest); i {
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
		file_api_v1_profile_data_api_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListProfileDataResponse); i {
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
		file_api_v1_profile_data_api_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileDataRequest); i {
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
		file_api_v1_profile_data_api_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProfileDataResponse); i {
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
			RawDescriptor: file_api_v1_profile_data_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_profile_data_api_proto_goTypes,
		DependencyIndexes: file_api_v1_profile_data_api_proto_depIdxs,
		MessageInfos:      file_api_v1_profile_data_api_proto_msgTypes,
	}.Build()
	File_api_v1_profile_data_api_proto = out.File
	file_api_v1_profile_data_api_proto_rawDesc = nil
	file_api_v1_profile_data_api_proto_goTypes = nil
	file_api_v1_profile_data_api_proto_depIdxs = nil
}
