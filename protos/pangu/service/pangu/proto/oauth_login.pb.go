// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: pangu/proto/oauth_login.proto

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

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Login string   `protobuf:"bytes,3,opt,name=login,proto3" json:"login,omitempty"`
	Email string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Roles []string `protobuf:"bytes,5,rep,name=roles,proto3" json:"roles,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pangu_proto_oauth_login_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pangu_proto_oauth_login_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_pangu_proto_oauth_login_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

type SessionInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionID string `protobuf:"bytes,1,opt,name=SessionID,proto3" json:"SessionID,omitempty"`
	MaxAge    int64  `protobuf:"varint,2,opt,name=MaxAge,proto3" json:"MaxAge,omitempty"` // second
}

func (x *SessionInfo) Reset() {
	*x = SessionInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pangu_proto_oauth_login_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SessionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SessionInfo) ProtoMessage() {}

func (x *SessionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pangu_proto_oauth_login_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SessionInfo.ProtoReflect.Descriptor instead.
func (*SessionInfo) Descriptor() ([]byte, []int) {
	return file_pangu_proto_oauth_login_proto_rawDescGZIP(), []int{1}
}

func (x *SessionInfo) GetSessionID() string {
	if x != nil {
		return x.SessionID
	}
	return ""
}

func (x *SessionInfo) GetMaxAge() int64 {
	if x != nil {
		return x.MaxAge
	}
	return 0
}

type LoginInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SessionId   string       `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	UserInfo    *UserInfo    `protobuf:"bytes,2,opt,name=user_info,json=userInfo,proto3" json:"user_info,omitempty"`
	SessionInfo *SessionInfo `protobuf:"bytes,3,opt,name=session_info,json=sessionInfo,proto3" json:"session_info,omitempty"`
}

func (x *LoginInfo) Reset() {
	*x = LoginInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pangu_proto_oauth_login_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginInfo) ProtoMessage() {}

func (x *LoginInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pangu_proto_oauth_login_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginInfo.ProtoReflect.Descriptor instead.
func (*LoginInfo) Descriptor() ([]byte, []int) {
	return file_pangu_proto_oauth_login_proto_rawDescGZIP(), []int{2}
}

func (x *LoginInfo) GetSessionId() string {
	if x != nil {
		return x.SessionId
	}
	return ""
}

func (x *LoginInfo) GetUserInfo() *UserInfo {
	if x != nil {
		return x.UserInfo
	}
	return nil
}

func (x *LoginInfo) GetSessionInfo() *SessionInfo {
	if x != nil {
		return x.SessionInfo
	}
	return nil
}

type OAuthLoginRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider string `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
	Code     string `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Service  string `protobuf:"bytes,3,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *OAuthLoginRequest) Reset() {
	*x = OAuthLoginRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pangu_proto_oauth_login_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthLoginRequest) ProtoMessage() {}

func (x *OAuthLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pangu_proto_oauth_login_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthLoginRequest.ProtoReflect.Descriptor instead.
func (*OAuthLoginRequest) Descriptor() ([]byte, []int) {
	return file_pangu_proto_oauth_login_proto_rawDescGZIP(), []int{3}
}

func (x *OAuthLoginRequest) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *OAuthLoginRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *OAuthLoginRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type OAuthLoginResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Succeed bool       `protobuf:"varint,1,opt,name=succeed,proto3" json:"succeed,omitempty"`
	Result  *LoginInfo `protobuf:"bytes,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *OAuthLoginResponse) Reset() {
	*x = OAuthLoginResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pangu_proto_oauth_login_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OAuthLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OAuthLoginResponse) ProtoMessage() {}

func (x *OAuthLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_pangu_proto_oauth_login_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OAuthLoginResponse.ProtoReflect.Descriptor instead.
func (*OAuthLoginResponse) Descriptor() ([]byte, []int) {
	return file_pangu_proto_oauth_login_proto_rawDescGZIP(), []int{4}
}

func (x *OAuthLoginResponse) GetSucceed() bool {
	if x != nil {
		return x.Succeed
	}
	return false
}

func (x *OAuthLoginResponse) GetResult() *LoginInfo {
	if x != nil {
		return x.Result
	}
	return nil
}

var File_pangu_proto_oauth_login_proto protoreflect.FileDescriptor

var file_pangu_proto_oauth_login_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0c, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x70, 0x0a,
	0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x22,
	0x43, 0x0a, 0x0b, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1c,
	0x0a, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x53, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x16, 0x0a, 0x06,
	0x4d, 0x61, 0x78, 0x41, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x4d, 0x61,
	0x78, 0x41, 0x67, 0x65, 0x22, 0x9d, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x49,
	0x64, 0x12, 0x33, 0x0a, 0x09, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x3c, 0x0a, 0x0c, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70,
	0x61, 0x6e, 0x67, 0x75, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0b, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x66, 0x6f, 0x22, 0x5d, 0x0a, 0x11, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x72, 0x6f,
	0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x22, 0x5f, 0x0a, 0x12, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x65, 0x64, 0x12, 0x2f, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x42, 0x22, 0x5a, 0x20, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d,
	0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x6e, 0x67, 0x75, 0x2f, 0x70, 0x61, 0x6e, 0x67,
	0x75, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pangu_proto_oauth_login_proto_rawDescOnce sync.Once
	file_pangu_proto_oauth_login_proto_rawDescData = file_pangu_proto_oauth_login_proto_rawDesc
)

func file_pangu_proto_oauth_login_proto_rawDescGZIP() []byte {
	file_pangu_proto_oauth_login_proto_rawDescOnce.Do(func() {
		file_pangu_proto_oauth_login_proto_rawDescData = protoimpl.X.CompressGZIP(file_pangu_proto_oauth_login_proto_rawDescData)
	})
	return file_pangu_proto_oauth_login_proto_rawDescData
}

var file_pangu_proto_oauth_login_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_pangu_proto_oauth_login_proto_goTypes = []interface{}{
	(*UserInfo)(nil),           // 0: panguservice.UserInfo
	(*SessionInfo)(nil),        // 1: panguservice.SessionInfo
	(*LoginInfo)(nil),          // 2: panguservice.LoginInfo
	(*OAuthLoginRequest)(nil),  // 3: panguservice.OAuthLoginRequest
	(*OAuthLoginResponse)(nil), // 4: panguservice.OAuthLoginResponse
}
var file_pangu_proto_oauth_login_proto_depIdxs = []int32{
	0, // 0: panguservice.LoginInfo.user_info:type_name -> panguservice.UserInfo
	1, // 1: panguservice.LoginInfo.session_info:type_name -> panguservice.SessionInfo
	2, // 2: panguservice.OAuthLoginResponse.result:type_name -> panguservice.LoginInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_pangu_proto_oauth_login_proto_init() }
func file_pangu_proto_oauth_login_proto_init() {
	if File_pangu_proto_oauth_login_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pangu_proto_oauth_login_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_pangu_proto_oauth_login_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SessionInfo); i {
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
		file_pangu_proto_oauth_login_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginInfo); i {
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
		file_pangu_proto_oauth_login_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthLoginRequest); i {
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
		file_pangu_proto_oauth_login_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OAuthLoginResponse); i {
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
			RawDescriptor: file_pangu_proto_oauth_login_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pangu_proto_oauth_login_proto_goTypes,
		DependencyIndexes: file_pangu_proto_oauth_login_proto_depIdxs,
		MessageInfos:      file_pangu_proto_oauth_login_proto_msgTypes,
	}.Build()
	File_pangu_proto_oauth_login_proto = out.File
	file_pangu_proto_oauth_login_proto_rawDesc = nil
	file_pangu_proto_oauth_login_proto_goTypes = nil
	file_pangu_proto_oauth_login_proto_depIdxs = nil
}
