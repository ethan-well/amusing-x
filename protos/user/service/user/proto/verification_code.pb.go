// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: user/proto/verification_code.proto

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

type VerificationCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone    string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
	Action   string `protobuf:"bytes,2,opt,name=action,proto3" json:"action,omitempty"`
	AreaCode string `protobuf:"bytes,3,opt,name=area_code,json=areaCode,proto3" json:"area_code,omitempty"`
}

func (x *VerificationCodeRequest) Reset() {
	*x = VerificationCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_verification_code_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerificationCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerificationCodeRequest) ProtoMessage() {}

func (x *VerificationCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_verification_code_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerificationCodeRequest.ProtoReflect.Descriptor instead.
func (*VerificationCodeRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_verification_code_proto_rawDescGZIP(), []int{0}
}

func (x *VerificationCodeRequest) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *VerificationCodeRequest) GetAction() string {
	if x != nil {
		return x.Action
	}
	return ""
}

func (x *VerificationCodeRequest) GetAreaCode() string {
	if x != nil {
		return x.AreaCode
	}
	return ""
}

type VerificationCodeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *VerificationCodeResponse) Reset() {
	*x = VerificationCodeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_proto_verification_code_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerificationCodeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerificationCodeResponse) ProtoMessage() {}

func (x *VerificationCodeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_verification_code_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerificationCodeResponse.ProtoReflect.Descriptor instead.
func (*VerificationCodeResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_verification_code_proto_rawDescGZIP(), []int{1}
}

func (x *VerificationCodeResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

var File_user_proto_verification_code_proto protoreflect.FileDescriptor

var file_user_proto_verification_code_proto_rawDesc = []byte{
	0x0a, 0x22, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x64, 0x0a, 0x17, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x72, 0x65, 0x61, 0x5f, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x72, 0x65, 0x61, 0x43, 0x6f, 0x64, 0x65,
	0x22, 0x2e, 0x0a, 0x18, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x42, 0x1c, 0x5a, 0x1a, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e, 0x67, 0x2d, 0x78, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_proto_verification_code_proto_rawDescOnce sync.Once
	file_user_proto_verification_code_proto_rawDescData = file_user_proto_verification_code_proto_rawDesc
)

func file_user_proto_verification_code_proto_rawDescGZIP() []byte {
	file_user_proto_verification_code_proto_rawDescOnce.Do(func() {
		file_user_proto_verification_code_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_proto_verification_code_proto_rawDescData)
	})
	return file_user_proto_verification_code_proto_rawDescData
}

var file_user_proto_verification_code_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_proto_verification_code_proto_goTypes = []interface{}{
	(*VerificationCodeRequest)(nil),  // 0: user.VerificationCodeRequest
	(*VerificationCodeResponse)(nil), // 1: user.VerificationCodeResponse
}
var file_user_proto_verification_code_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_proto_verification_code_proto_init() }
func file_user_proto_verification_code_proto_init() {
	if File_user_proto_verification_code_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_proto_verification_code_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerificationCodeRequest); i {
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
		file_user_proto_verification_code_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerificationCodeResponse); i {
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
			RawDescriptor: file_user_proto_verification_code_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_proto_verification_code_proto_goTypes,
		DependencyIndexes: file_user_proto_verification_code_proto_depIdxs,
		MessageInfos:      file_user_proto_verification_code_proto_msgTypes,
	}.Build()
	File_user_proto_verification_code_proto = out.File
	file_user_proto_verification_code_proto_rawDesc = nil
	file_user_proto_verification_code_proto_goTypes = nil
	file_user_proto_verification_code_proto_depIdxs = nil
}
