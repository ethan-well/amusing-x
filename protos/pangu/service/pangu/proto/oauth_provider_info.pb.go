// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: pangu/proto/oauth_provider_info.proto

package panguservice

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

type OauthProviderInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Service  string `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *OauthProviderInfoRequest) Reset() {
	*x = OauthProviderInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pangu_proto_oauth_provider_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OauthProviderInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OauthProviderInfoRequest) ProtoMessage() {}

func (x *OauthProviderInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pangu_proto_oauth_provider_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OauthProviderInfoRequest.ProtoReflect.Descriptor instead.
func (*OauthProviderInfoRequest) Descriptor() ([]byte, []int) {
	return file_pangu_proto_oauth_provider_info_proto_rawDescGZIP(), []int{0}
}

func (x *OauthProviderInfoRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *OauthProviderInfoRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type OAuthProviderInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succeed bool               `protobuf:"varint,1,opt,name=succeed,proto3" json:"succeed,omitempty"`
	Result  *OAuthProviderInfo `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *OAuthProviderInfoResponse) Reset() {
	*x = OAuthProviderInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pangu_proto_oauth_provider_info_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthProviderInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthProviderInfoResponse) ProtoMessage() {}

func (x *OAuthProviderInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pangu_proto_oauth_provider_info_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthProviderInfoResponse.ProtoReflect.Descriptor instead.
func (*OAuthProviderInfoResponse) Descriptor() ([]byte, []int) {
	return file_pangu_proto_oauth_provider_info_proto_rawDescGZIP(), []int{1}
}

func (x *OAuthProviderInfoResponse) GetSucceed() bool {
	if x != nil {
		return x.Succeed
	}
	return false
}

func (x *OAuthProviderInfoResponse) GetResult() *OAuthProviderInfo {
	if x != nil {
		return x.Result
	}
	return nil
}

type OAuthProviderInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider     string `protobuf:"bytes,1,opt,name=Provider,json=provider,proto3" json:"Provider,omitempty"`
	OauthUrl     string `protobuf:"bytes,2,opt,name=OauthUrl,json=oauth_url,proto3" json:"OauthUrl,omitempty"`
	ClientID     string `protobuf:"bytes,3,opt,name=ClientID,json=client_id,proto3" json:"ClientID,omitempty"`
	Scope        string `protobuf:"bytes,4,opt,name=Scope,json=scope,proto3" json:"Scope,omitempty"`
	State        string `protobuf:"bytes,5,opt,name=State,json=state,proto3" json:"State,omitempty"`
	GrantType    string `protobuf:"bytes,6,opt,name=GrantType,json=grant_type,proto3" json:"GrantType,omitempty"`
	RedirectUrl  string `protobuf:"bytes,7,opt,name=RedirectUrl,json=redirect_url,proto3" json:"RedirectUrl,omitempty"`
	CompletePath string `protobuf:"bytes,8,opt,name=CompletePath,json=complete_path,proto3" json:"CompletePath,omitempty"`
}

func (x *OAuthProviderInfo) Reset() {
	*x = OAuthProviderInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pangu_proto_oauth_provider_info_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthProviderInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthProviderInfo) ProtoMessage() {}

func (x *OAuthProviderInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pangu_proto_oauth_provider_info_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthProviderInfo.ProtoReflect.Descriptor instead.
func (*OAuthProviderInfo) Descriptor() ([]byte, []int) {
	return file_pangu_proto_oauth_provider_info_proto_rawDescGZIP(), []int{2}
}

func (x *OAuthProviderInfo) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *OAuthProviderInfo) GetOauthUrl() string {
	if x != nil {
		return x.OauthUrl
	}
	return ""
}

func (x *OAuthProviderInfo) GetClientID() string {
	if x != nil {
		return x.ClientID
	}
	return ""
}

func (x *OAuthProviderInfo) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *OAuthProviderInfo) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *OAuthProviderInfo) GetGrantType() string {
	if x != nil {
		return x.GrantType
	}
	return ""
}

func (x *OAuthProviderInfo) GetRedirectUrl() string {
	if x != nil {
		return x.RedirectUrl
	}
	return ""
}

func (x *OAuthProviderInfo) GetCompletePath() string {
	if x != nil {
		return x.CompletePath
	}
	return ""
}

var File_pangu_proto_oauth_provider_info_proto protoreflect.FileDescriptor

var file_pangu_proto_oauth_provider_info_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x50, 0x0a, 0x18, 0x4f, 0x61, 0x75, 0x74, 0x68, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x18, 0x0a,
	0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x6e, 0x0a, 0x19, 0x4f, 0x41, 0x75, 0x74, 0x68,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x65, 0x64, 0x12, 0x37,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f,
	0x2e, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0xfc, 0x01, 0x0a, 0x11, 0x4f, 0x41, 0x75, 0x74,
	0x68, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x0a,
	0x08, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x08, 0x4f, 0x61, 0x75,
	0x74, 0x68, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x5f, 0x75, 0x72, 0x6c, 0x12, 0x1b, 0x0a, 0x08, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x53, 0x74, 0x61,
	0x74, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x1d, 0x0a, 0x09, 0x47, 0x72, 0x61, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x67, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x12, 0x21,
	0x0a, 0x0b, 0x52, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x55, 0x72, 0x6c, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x75, 0x72,
	0x6c, 0x12, 0x23, 0x0a, 0x0c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x74,
	0x68, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74,
	0x65, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x42, 0x22, 0x5a, 0x20, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e,
	0x67, 0x2d, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x70, 0x61,
	0x6e, 0x67, 0x75, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_pangu_proto_oauth_provider_info_proto_rawDescOnce sync.Once
	file_pangu_proto_oauth_provider_info_proto_rawDescData = file_pangu_proto_oauth_provider_info_proto_rawDesc
)

func file_pangu_proto_oauth_provider_info_proto_rawDescGZIP() []byte {
	file_pangu_proto_oauth_provider_info_proto_rawDescOnce.Do(func() {
		file_pangu_proto_oauth_provider_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_pangu_proto_oauth_provider_info_proto_rawDescData)
	})
	return file_pangu_proto_oauth_provider_info_proto_rawDescData
}

var file_pangu_proto_oauth_provider_info_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pangu_proto_oauth_provider_info_proto_goTypes = []interface{}{
	(*OauthProviderInfoRequest)(nil),  // 0: panguservice.OauthProviderInfoRequest
	(*OAuthProviderInfoResponse)(nil), // 1: panguservice.OAuthProviderInfoResponse
	(*OAuthProviderInfo)(nil),         // 2: panguservice.OAuthProviderInfo
}
var file_pangu_proto_oauth_provider_info_proto_depIdxs = []int32{
	2, // 0: panguservice.OAuthProviderInfoResponse.result:type_name -> panguservice.OAuthProviderInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_pangu_proto_oauth_provider_info_proto_init() }
func file_pangu_proto_oauth_provider_info_proto_init() {
	if File_pangu_proto_oauth_provider_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pangu_proto_oauth_provider_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OauthProviderInfoRequest); i {
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
		file_pangu_proto_oauth_provider_info_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthProviderInfoResponse); i {
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
		file_pangu_proto_oauth_provider_info_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthProviderInfo); i {
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
			RawDescriptor: file_pangu_proto_oauth_provider_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pangu_proto_oauth_provider_info_proto_goTypes,
		DependencyIndexes: file_pangu_proto_oauth_provider_info_proto_depIdxs,
		MessageInfos:      file_pangu_proto_oauth_provider_info_proto_msgTypes,
	}.Build()
	File_pangu_proto_oauth_provider_info_proto = out.File
	file_pangu_proto_oauth_provider_info_proto_rawDesc = nil
	file_pangu_proto_oauth_provider_info_proto_goTypes = nil
	file_pangu_proto_oauth_provider_info_proto_depIdxs = nil
}
