// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: internal/app/handler/grpc/user/user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetAllUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Page          int32                  `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	PerPage       int32                  `protobuf:"varint,5,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllUserRequest) Reset() {
	*x = GetAllUserRequest{}
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserRequest) ProtoMessage() {}

func (x *GetAllUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserRequest.ProtoReflect.Descriptor instead.
func (*GetAllUserRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_handler_grpc_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *GetAllUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetAllUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetAllUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetAllUserRequest) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetAllUserRequest) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

type GetUserDetailRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserDetailRequest) Reset() {
	*x = GetUserDetailRequest{}
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserDetailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDetailRequest) ProtoMessage() {}

func (x *GetUserDetailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDetailRequest.ProtoReflect.Descriptor instead.
func (*GetUserDetailRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_handler_grpc_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetUserDetailRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type CreateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	SubDistrictId int64                  `protobuf:"varint,4,opt,name=sub_district_id,json=subDistrictId,proto3" json:"sub_district_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_handler_grpc_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *CreateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *CreateUserRequest) GetSubDistrictId() int64 {
	if x != nil {
		return x.SubDistrictId
	}
	return 0
}

type UpdateUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Username      string                 `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	SubDistrictId int64                  `protobuf:"varint,5,opt,name=sub_district_id,json=subDistrictId,proto3" json:"sub_district_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UpdateUserRequest) Reset() {
	*x = UpdateUserRequest{}
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UpdateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUserRequest) ProtoMessage() {}

func (x *UpdateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUserRequest.ProtoReflect.Descriptor instead.
func (*UpdateUserRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_handler_grpc_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateUserRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UpdateUserRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UpdateUserRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UpdateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateUserRequest) GetSubDistrictId() int64 {
	if x != nil {
		return x.SubDistrictId
	}
	return 0
}

type DeleteUserRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteUserRequest) Reset() {
	*x = DeleteUserRequest{}
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteUserRequest) ProtoMessage() {}

func (x *DeleteUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteUserRequest.ProtoReflect.Descriptor instead.
func (*DeleteUserRequest) Descriptor() ([]byte, []int) {
	return file_internal_app_handler_grpc_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteUserRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type CreateUpdateDeleteResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateUpdateDeleteResponse) Reset() {
	*x = CreateUpdateDeleteResponse{}
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateUpdateDeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUpdateDeleteResponse) ProtoMessage() {}

func (x *CreateUpdateDeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUpdateDeleteResponse.ProtoReflect.Descriptor instead.
func (*CreateUpdateDeleteResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_handler_grpc_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUpdateDeleteResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetAllUserResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data          *GetAllUserData        `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllUserResponse) Reset() {
	*x = GetAllUserResponse{}
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserResponse) ProtoMessage() {}

func (x *GetAllUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserResponse.ProtoReflect.Descriptor instead.
func (*GetAllUserResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_handler_grpc_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllUserResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetAllUserResponse) GetData() *GetAllUserData {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetUserDetailResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       string                 `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Data          *UserData              `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserDetailResponse) Reset() {
	*x = GetUserDetailResponse{}
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserDetailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserDetailResponse) ProtoMessage() {}

func (x *GetUserDetailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserDetailResponse.ProtoReflect.Descriptor instead.
func (*GetUserDetailResponse) Descriptor() ([]byte, []int) {
	return file_internal_app_handler_grpc_user_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserDetailResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetUserDetailResponse) GetData() *UserData {
	if x != nil {
		return x.Data
	}
	return nil
}

type GetAllUserData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Users         []*UserData            `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
	Pagination    *PaginationData        `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetAllUserData) Reset() {
	*x = GetAllUserData{}
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAllUserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUserData) ProtoMessage() {}

func (x *GetAllUserData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUserData.ProtoReflect.Descriptor instead.
func (*GetAllUserData) Descriptor() ([]byte, []int) {
	return file_internal_app_handler_grpc_user_user_proto_rawDescGZIP(), []int{8}
}

func (x *GetAllUserData) GetUsers() []*UserData {
	if x != nil {
		return x.Users
	}
	return nil
}

func (x *GetAllUserData) GetPagination() *PaginationData {
	if x != nil {
		return x.Pagination
	}
	return nil
}

type UserData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Uuid          string                 `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Username      string                 `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
	Email         string                 `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	CreatedAt     string                 `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Location      *UserLocationData      `protobuf:"bytes,7,opt,name=location,proto3,oneof" json:"location,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserData) Reset() {
	*x = UserData{}
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_internal_app_handler_grpc_user_user_proto_rawDescGZIP(), []int{9}
}

func (x *UserData) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *UserData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserData) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserData) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserData) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *UserData) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *UserData) GetLocation() *UserLocationData {
	if x != nil {
		return x.Location
	}
	return nil
}

type UserLocationData struct {
	state           protoimpl.MessageState `protogen:"open.v1"`
	SubDistrictId   int64                  `protobuf:"varint,1,opt,name=sub_district_id,json=subDistrictId,proto3" json:"sub_district_id,omitempty"`
	SubDistrictName string                 `protobuf:"bytes,2,opt,name=sub_district_name,json=subDistrictName,proto3" json:"sub_district_name,omitempty"`
	DistrictId      int64                  `protobuf:"varint,3,opt,name=district_id,json=districtId,proto3" json:"district_id,omitempty"`
	DistrictName    string                 `protobuf:"bytes,4,opt,name=district_name,json=districtName,proto3" json:"district_name,omitempty"`
	CityId          int64                  `protobuf:"varint,5,opt,name=city_id,json=cityId,proto3" json:"city_id,omitempty"`
	CityName        string                 `protobuf:"bytes,6,opt,name=city_name,json=cityName,proto3" json:"city_name,omitempty"`
	ProvinceId      int64                  `protobuf:"varint,7,opt,name=province_id,json=provinceId,proto3" json:"province_id,omitempty"`
	ProvinceName    string                 `protobuf:"bytes,8,opt,name=province_name,json=provinceName,proto3" json:"province_name,omitempty"`
	unknownFields   protoimpl.UnknownFields
	sizeCache       protoimpl.SizeCache
}

func (x *UserLocationData) Reset() {
	*x = UserLocationData{}
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserLocationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLocationData) ProtoMessage() {}

func (x *UserLocationData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLocationData.ProtoReflect.Descriptor instead.
func (*UserLocationData) Descriptor() ([]byte, []int) {
	return file_internal_app_handler_grpc_user_user_proto_rawDescGZIP(), []int{10}
}

func (x *UserLocationData) GetSubDistrictId() int64 {
	if x != nil {
		return x.SubDistrictId
	}
	return 0
}

func (x *UserLocationData) GetSubDistrictName() string {
	if x != nil {
		return x.SubDistrictName
	}
	return ""
}

func (x *UserLocationData) GetDistrictId() int64 {
	if x != nil {
		return x.DistrictId
	}
	return 0
}

func (x *UserLocationData) GetDistrictName() string {
	if x != nil {
		return x.DistrictName
	}
	return ""
}

func (x *UserLocationData) GetCityId() int64 {
	if x != nil {
		return x.CityId
	}
	return 0
}

func (x *UserLocationData) GetCityName() string {
	if x != nil {
		return x.CityName
	}
	return ""
}

func (x *UserLocationData) GetProvinceId() int64 {
	if x != nil {
		return x.ProvinceId
	}
	return 0
}

func (x *UserLocationData) GetProvinceName() string {
	if x != nil {
		return x.ProvinceName
	}
	return ""
}

type PaginationData struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Page          int32                  `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	PerPage       int32                  `protobuf:"varint,2,opt,name=per_page,json=perPage,proto3" json:"per_page,omitempty"`
	PageCount     int32                  `protobuf:"varint,3,opt,name=page_count,json=pageCount,proto3" json:"page_count,omitempty"`
	TotalCount    int64                  `protobuf:"varint,4,opt,name=total_count,json=totalCount,proto3" json:"total_count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PaginationData) Reset() {
	*x = PaginationData{}
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PaginationData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginationData) ProtoMessage() {}

func (x *PaginationData) ProtoReflect() protoreflect.Message {
	mi := &file_internal_app_handler_grpc_user_user_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginationData.ProtoReflect.Descriptor instead.
func (*PaginationData) Descriptor() ([]byte, []int) {
	return file_internal_app_handler_grpc_user_user_proto_rawDescGZIP(), []int{11}
}

func (x *PaginationData) GetPage() int32 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginationData) GetPerPage() int32 {
	if x != nil {
		return x.PerPage
	}
	return 0
}

func (x *PaginationData) GetPageCount() int32 {
	if x != nil {
		return x.PageCount
	}
	return 0
}

func (x *PaginationData) GetTotalCount() int64 {
	if x != nil {
		return x.TotalCount
	}
	return 0
}

var File_internal_app_handler_grpc_user_user_proto protoreflect.FileDescriptor

var file_internal_app_handler_grpc_user_user_proto_rawDesc = string([]byte{
	0x0a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x68,
	0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65,
	0x72, 0x22, 0x88, 0x01, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75,
	0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x22, 0x2a, 0x0a, 0x14,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x81, 0x01, 0x0a, 0x11, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73,
	0x75, 0x62, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x64, 0x22, 0x95, 0x01, 0x0a,
	0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x26, 0x0a, 0x0f,
	0x73, 0x75, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x75, 0x62, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69,
	0x63, 0x74, 0x49, 0x64, 0x22, 0x27, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x36, 0x0a,
	0x1a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x58, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22,
	0x55, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61,
	0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x6c, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x24, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x34,
	0x0a, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x14, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x69, 0x6e, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x22, 0xe8, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65,
	0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1d, 0x0a, 0x0a, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x37, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x61, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88,
	0x01, 0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xa8, 0x02, 0x0a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x26, 0x0a, 0x0f, 0x73, 0x75, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73,
	0x75, 0x62, 0x44, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x64, 0x12, 0x2a, 0x0a, 0x11,
	0x73, 0x75, 0x62, 0x5f, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x73, 0x75, 0x62, 0x44, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x73, 0x74,
	0x72, 0x69, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64,
	0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x64, 0x69, 0x73,
	0x74, 0x72, 0x69, 0x63, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x64, 0x69, 0x73, 0x74, 0x72, 0x69, 0x63, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x17,
	0x0a, 0x07, 0x63, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x63, 0x69, 0x74, 0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x69, 0x74, 0x79, 0x5f,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x69, 0x74, 0x79,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69,
	0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x6e, 0x63,
	0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x6e, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x7f, 0x0a, 0x0e, 0x50, 0x61,
	0x67, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x07, 0x70, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x70, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0xdf, 0x02, 0x0a, 0x0b,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x06, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x1a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43,
	0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x17, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x12, 0x17, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x49, 0x5a,
	0x47, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x6e, 0x73, 0x61,
	0x6e, 0x65, 0x61, 0x64, 0x69, 0x6e, 0x65, 0x73, 0x69, 0x61, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x6f,
	0x69, 0x6c, 0x65, 0x72, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_internal_app_handler_grpc_user_user_proto_rawDescOnce sync.Once
	file_internal_app_handler_grpc_user_user_proto_rawDescData []byte
)

func file_internal_app_handler_grpc_user_user_proto_rawDescGZIP() []byte {
	file_internal_app_handler_grpc_user_user_proto_rawDescOnce.Do(func() {
		file_internal_app_handler_grpc_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_app_handler_grpc_user_user_proto_rawDesc), len(file_internal_app_handler_grpc_user_user_proto_rawDesc)))
	})
	return file_internal_app_handler_grpc_user_user_proto_rawDescData
}

var file_internal_app_handler_grpc_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_internal_app_handler_grpc_user_user_proto_goTypes = []any{
	(*GetAllUserRequest)(nil),          // 0: user.GetAllUserRequest
	(*GetUserDetailRequest)(nil),       // 1: user.GetUserDetailRequest
	(*CreateUserRequest)(nil),          // 2: user.CreateUserRequest
	(*UpdateUserRequest)(nil),          // 3: user.UpdateUserRequest
	(*DeleteUserRequest)(nil),          // 4: user.DeleteUserRequest
	(*CreateUpdateDeleteResponse)(nil), // 5: user.CreateUpdateDeleteResponse
	(*GetAllUserResponse)(nil),         // 6: user.GetAllUserResponse
	(*GetUserDetailResponse)(nil),      // 7: user.GetUserDetailResponse
	(*GetAllUserData)(nil),             // 8: user.GetAllUserData
	(*UserData)(nil),                   // 9: user.UserData
	(*UserLocationData)(nil),           // 10: user.UserLocationData
	(*PaginationData)(nil),             // 11: user.PaginationData
}
var file_internal_app_handler_grpc_user_user_proto_depIdxs = []int32{
	8,  // 0: user.GetAllUserResponse.data:type_name -> user.GetAllUserData
	9,  // 1: user.GetUserDetailResponse.data:type_name -> user.UserData
	9,  // 2: user.GetAllUserData.users:type_name -> user.UserData
	11, // 3: user.GetAllUserData.pagination:type_name -> user.PaginationData
	10, // 4: user.UserData.location:type_name -> user.UserLocationData
	0,  // 5: user.UserService.GetAll:input_type -> user.GetAllUserRequest
	1,  // 6: user.UserService.GetDetail:input_type -> user.GetUserDetailRequest
	2,  // 7: user.UserService.Create:input_type -> user.CreateUserRequest
	3,  // 8: user.UserService.Update:input_type -> user.UpdateUserRequest
	4,  // 9: user.UserService.Delete:input_type -> user.DeleteUserRequest
	6,  // 10: user.UserService.GetAll:output_type -> user.GetAllUserResponse
	7,  // 11: user.UserService.GetDetail:output_type -> user.GetUserDetailResponse
	5,  // 12: user.UserService.Create:output_type -> user.CreateUpdateDeleteResponse
	5,  // 13: user.UserService.Update:output_type -> user.CreateUpdateDeleteResponse
	5,  // 14: user.UserService.Delete:output_type -> user.CreateUpdateDeleteResponse
	10, // [10:15] is the sub-list for method output_type
	5,  // [5:10] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_internal_app_handler_grpc_user_user_proto_init() }
func file_internal_app_handler_grpc_user_user_proto_init() {
	if File_internal_app_handler_grpc_user_user_proto != nil {
		return
	}
	file_internal_app_handler_grpc_user_user_proto_msgTypes[9].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_app_handler_grpc_user_user_proto_rawDesc), len(file_internal_app_handler_grpc_user_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_app_handler_grpc_user_user_proto_goTypes,
		DependencyIndexes: file_internal_app_handler_grpc_user_user_proto_depIdxs,
		MessageInfos:      file_internal_app_handler_grpc_user_user_proto_msgTypes,
	}.Build()
	File_internal_app_handler_grpc_user_user_proto = out.File
	file_internal_app_handler_grpc_user_user_proto_goTypes = nil
	file_internal_app_handler_grpc_user_user_proto_depIdxs = nil
}
