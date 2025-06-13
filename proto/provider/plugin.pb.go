// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: proto/provider/plugin.proto

package providerpb

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

type InitReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ClientId      string                 `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	ClientSecret  string                 `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
	RedirectUrl   string                 `protobuf:"bytes,3,opt,name=redirect_url,json=redirectUrl,proto3" json:"redirect_url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InitReq) Reset() {
	*x = InitReq{}
	mi := &file_proto_provider_plugin_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InitReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InitReq) ProtoMessage() {}

func (x *InitReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InitReq.ProtoReflect.Descriptor instead.
func (*InitReq) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *InitReq) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *InitReq) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

func (x *InitReq) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

type GetTokenReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTokenReq) Reset() {
	*x = GetTokenReq{}
	mi := &file_proto_provider_plugin_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTokenReq) ProtoMessage() {}

func (x *GetTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTokenReq.ProtoReflect.Descriptor instead.
func (*GetTokenReq) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *GetTokenReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type RefreshTokenReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RefreshToken  string                 `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RefreshTokenReq) Reset() {
	*x = RefreshTokenReq{}
	mi := &file_proto_provider_plugin_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RefreshTokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokenReq) ProtoMessage() {}

func (x *RefreshTokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokenReq.ProtoReflect.Descriptor instead.
func (*RefreshTokenReq) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{2}
}

func (x *RefreshTokenReq) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type ProviderResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProviderResp) Reset() {
	*x = ProviderResp{}
	mi := &file_proto_provider_plugin_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProviderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProviderResp) ProtoMessage() {}

func (x *ProviderResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProviderResp.ProtoReflect.Descriptor instead.
func (*ProviderResp) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{3}
}

func (x *ProviderResp) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type NewAuthURLReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	State         string                 `protobuf:"bytes,1,opt,name=state,proto3" json:"state,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NewAuthURLReq) Reset() {
	*x = NewAuthURLReq{}
	mi := &file_proto_provider_plugin_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewAuthURLReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAuthURLReq) ProtoMessage() {}

func (x *NewAuthURLReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAuthURLReq.ProtoReflect.Descriptor instead.
func (*NewAuthURLReq) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{4}
}

func (x *NewAuthURLReq) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

type NewAuthURLResp struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Url           string                 `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NewAuthURLResp) Reset() {
	*x = NewAuthURLResp{}
	mi := &file_proto_provider_plugin_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NewAuthURLResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewAuthURLResp) ProtoMessage() {}

func (x *NewAuthURLResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewAuthURLResp.ProtoReflect.Descriptor instead.
func (*NewAuthURLResp) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{5}
}

func (x *NewAuthURLResp) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type GetUserInfoReq struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUserInfoReq) Reset() {
	*x = GetUserInfoReq{}
	mi := &file_proto_provider_plugin_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoReq) ProtoMessage() {}

func (x *GetUserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoReq.ProtoReflect.Descriptor instead.
func (*GetUserInfoReq) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{6}
}

func (x *GetUserInfoReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GetUserInfoResp struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	Username       string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	ProviderUserId string                 `protobuf:"bytes,2,opt,name=provider_user_id,json=providerUserId,proto3" json:"provider_user_id,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *GetUserInfoResp) Reset() {
	*x = GetUserInfoResp{}
	mi := &file_proto_provider_plugin_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserInfoResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserInfoResp) ProtoMessage() {}

func (x *GetUserInfoResp) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserInfoResp.ProtoReflect.Descriptor instead.
func (*GetUserInfoResp) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserInfoResp) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetUserInfoResp) GetProviderUserId() string {
	if x != nil {
		return x.ProviderUserId
	}
	return ""
}

type Enpty struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Enpty) Reset() {
	*x = Enpty{}
	mi := &file_proto_provider_plugin_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Enpty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Enpty) ProtoMessage() {}

func (x *Enpty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_provider_plugin_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Enpty.ProtoReflect.Descriptor instead.
func (*Enpty) Descriptor() ([]byte, []int) {
	return file_proto_provider_plugin_proto_rawDescGZIP(), []int{8}
}

var File_proto_provider_plugin_proto protoreflect.FileDescriptor

var file_proto_provider_plugin_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x07, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x23, 0x0a, 0x0d,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63,
	0x74, 0x55, 0x72, 0x6c, 0x22, 0x21, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x36, 0x0a, 0x0f, 0x52, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x22, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x22, 0x25, 0x0a, 0x0d, 0x4e, 0x65, 0x77, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x22, 0x22, 0x0a, 0x0e, 0x4e, 0x65,
	0x77, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03,
	0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x24,
	0x0a, 0x0e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x22, 0x57, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x07, 0x0a,
	0x05, 0x45, 0x6e, 0x70, 0x74, 0x79, 0x32, 0xe4, 0x01, 0x0a, 0x0c, 0x4f, 0x61, 0x75, 0x74, 0x68,
	0x32, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x26, 0x0a, 0x04, 0x49, 0x6e, 0x69, 0x74, 0x12,
	0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x6e, 0x69, 0x74, 0x52, 0x65, 0x71, 0x1a,
	0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12,
	0x2f, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x0c, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x45, 0x6e, 0x70, 0x74, 0x79, 0x1a, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x12, 0x14,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77, 0x41, 0x75, 0x74, 0x68, 0x55, 0x52,
	0x4c, 0x52, 0x65, 0x71, 0x1a, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x65, 0x77,
	0x41, 0x75, 0x74, 0x68, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x12, 0x3e, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x15, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x1a, 0x16, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x22, 0x00, 0x42, 0x0e, 0x5a,
	0x0c, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_provider_plugin_proto_rawDescOnce sync.Once
	file_proto_provider_plugin_proto_rawDescData = file_proto_provider_plugin_proto_rawDesc
)

func file_proto_provider_plugin_proto_rawDescGZIP() []byte {
	file_proto_provider_plugin_proto_rawDescOnce.Do(func() {
		file_proto_provider_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_provider_plugin_proto_rawDescData)
	})
	return file_proto_provider_plugin_proto_rawDescData
}

var file_proto_provider_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_provider_plugin_proto_goTypes = []any{
	(*InitReq)(nil),         // 0: proto.InitReq
	(*GetTokenReq)(nil),     // 1: proto.GetTokenReq
	(*RefreshTokenReq)(nil), // 2: proto.RefreshTokenReq
	(*ProviderResp)(nil),    // 3: proto.ProviderResp
	(*NewAuthURLReq)(nil),   // 4: proto.NewAuthURLReq
	(*NewAuthURLResp)(nil),  // 5: proto.NewAuthURLResp
	(*GetUserInfoReq)(nil),  // 6: proto.GetUserInfoReq
	(*GetUserInfoResp)(nil), // 7: proto.GetUserInfoResp
	(*Enpty)(nil),           // 8: proto.Enpty
}
var file_proto_provider_plugin_proto_depIdxs = []int32{
	0, // 0: proto.Oauth2Plugin.Init:input_type -> proto.InitReq
	8, // 1: proto.Oauth2Plugin.Provider:input_type -> proto.Enpty
	4, // 2: proto.Oauth2Plugin.NewAuthURL:input_type -> proto.NewAuthURLReq
	6, // 3: proto.Oauth2Plugin.GetUserInfo:input_type -> proto.GetUserInfoReq
	8, // 4: proto.Oauth2Plugin.Init:output_type -> proto.Enpty
	3, // 5: proto.Oauth2Plugin.Provider:output_type -> proto.ProviderResp
	5, // 6: proto.Oauth2Plugin.NewAuthURL:output_type -> proto.NewAuthURLResp
	7, // 7: proto.Oauth2Plugin.GetUserInfo:output_type -> proto.GetUserInfoResp
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_provider_plugin_proto_init() }
func file_proto_provider_plugin_proto_init() {
	if File_proto_provider_plugin_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_provider_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_provider_plugin_proto_goTypes,
		DependencyIndexes: file_proto_provider_plugin_proto_depIdxs,
		MessageInfos:      file_proto_provider_plugin_proto_msgTypes,
	}.Build()
	File_proto_provider_plugin_proto = out.File
	file_proto_provider_plugin_proto_rawDesc = nil
	file_proto_provider_plugin_proto_goTypes = nil
	file_proto_provider_plugin_proto_depIdxs = nil
}
