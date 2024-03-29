// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: user/proto/authentication.proto

package service

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

type AuthenticationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server  string   `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	UserId  int64    `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Actions []string `protobuf:"bytes,3,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *AuthenticationRequest) Reset() {
	*x = AuthenticationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_authentication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationRequest) ProtoMessage() {}

func (x *AuthenticationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_authentication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationRequest.ProtoReflect.Descriptor instead.
func (*AuthenticationRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_authentication_proto_rawDescGZIP(), []int{0}
}

func (x *AuthenticationRequest) GetServer() string {
	if x != nil {
		return x.Server
	}
	return ""
}

func (x *AuthenticationRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *AuthenticationRequest) GetActions() []string {
	if x != nil {
		return x.Actions
	}
	return nil
}

type AuthenticationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result []*AuthenticationInfo `protobuf:"bytes,1,rep,name=result,proto3" json:"result,omitempty"`
}

func (x *AuthenticationResponse) Reset() {
	*x = AuthenticationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_authentication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationResponse) ProtoMessage() {}

func (x *AuthenticationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_authentication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationResponse.ProtoReflect.Descriptor instead.
func (*AuthenticationResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_authentication_proto_rawDescGZIP(), []int{1}
}

func (x *AuthenticationResponse) GetResult() []*AuthenticationInfo {
	if x != nil {
		return x.Result
	}
	return nil
}

type AuthenticationInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action        string `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	HasPermission bool   `protobuf:"varint,2,opt,name=hasPermission,proto3" json:"hasPermission,omitempty"`
}

func (x *AuthenticationInfo) Reset() {
	*x = AuthenticationInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_authentication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticationInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticationInfo) ProtoMessage() {}

func (x *AuthenticationInfo) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_authentication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticationInfo.ProtoReflect.Descriptor instead.
func (*AuthenticationInfo) Descriptor() ([]byte, []int) {
	return file_user_proto_authentication_proto_rawDescGZIP(), []int{2}
}

func (x *AuthenticationInfo) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *AuthenticationInfo) GetHasPermission() bool {
	if x != nil {
		return x.HasPermission
	}
	return false
}

var File_user_proto_authentication_proto protoreflect.FileDescriptor

var file_user_proto_authentication_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x62, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x4a, 0x0a, 0x16, 0x41,
	0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x52, 0x0a, 0x12, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x61, 0x73, 0x50, 0x65, 0x72, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x68, 0x61,
	0x73, 0x50, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x1c, 0x5a, 0x1a, 0x61,
	0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_user_proto_authentication_proto_rawDescOnce sync.Once
	file_user_proto_authentication_proto_rawDescData = file_user_proto_authentication_proto_rawDesc
)

func file_user_proto_authentication_proto_rawDescGZIP() []byte {
	file_user_proto_authentication_proto_rawDescOnce.Do(func() {
		file_user_proto_authentication_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_authentication_proto_rawDescData)
	})
	return file_user_proto_authentication_proto_rawDescData
}

var file_user_proto_authentication_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_user_proto_authentication_proto_goTypes = []interface{}{
	(*AuthenticationRequest)(nil),  // 0: user.AuthenticationRequest
	(*AuthenticationResponse)(nil), // 1: user.AuthenticationResponse
	(*AuthenticationInfo)(nil),     // 2: user.AuthenticationInfo
}
var file_user_proto_authentication_proto_depIdxs = []int32{
	2, // 0: user.AuthenticationResponse.result:type_name -> user.AuthenticationInfo
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_user_proto_authentication_proto_init() }
func file_user_proto_authentication_proto_init() {
	if File_user_proto_authentication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_authentication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationRequest); i {
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
		file_user_proto_authentication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationResponse); i {
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
		file_user_proto_authentication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticationInfo); i {
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
			RawDescriptor: file_user_proto_authentication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_proto_authentication_proto_goTypes,
		DependencyIndexes: file_user_proto_authentication_proto_depIdxs,
		MessageInfos:      file_user_proto_authentication_proto_msgTypes,
	}.Build()
	File_user_proto_authentication_proto = out.File
	file_user_proto_authentication_proto_rawDesc = nil
	file_user_proto_authentication_proto_goTypes = nil
	file_user_proto_authentication_proto_depIdxs = nil
}
