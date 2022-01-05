// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.8.0
// source: ganymede/proto/server.proto

package ganymedeservice

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

type BlankParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *BlankParams) Reset() {
	*x = BlankParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ganymede_proto_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlankParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlankParams) ProtoMessage() {}

func (x *BlankParams) ProtoReflect() protoreflect.Message {
	mi := &file_ganymede_proto_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlankParams.ProtoReflect.Descriptor instead.
func (*BlankParams) Descriptor() ([]byte, []int) {
	return file_ganymede_proto_server_proto_rawDescGZIP(), []int{0}
}

var File_ganymede_proto_server_proto protoreflect.FileDescriptor

var file_ganymede_proto_server_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x65, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67,
	0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x1a, 0x1b, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x65, 0x64, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x78, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x65, 0x64, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x19, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x65, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6a, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x21, 0x67, 0x61, 0x6e, 0x79,
	0x6d, 0x65, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x72, 0x79, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x26, 0x67,
	0x61, 0x6e, 0x79, 0x6d, 0x65, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x65, 0x64, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x23, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x65, 0x64, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f,
	0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x65,
	0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x61, 0x6e, 0x79,
	0x6d, 0x65, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x67, 0x61, 0x6e,
	0x79, 0x6d, 0x65, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x73, 0x5f, 0x6c,
	0x6f, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x61, 0x6e, 0x79,
	0x6d, 0x65, 0x64, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x42, 0x6c, 0x61, 0x6e, 0x6b,
	0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x32, 0xdd, 0x06, 0x0a, 0x0f, 0x47, 0x61, 0x6e, 0x79, 0x6d,
	0x65, 0x64, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x04, 0x50, 0x6f,
	0x6e, 0x67, 0x12, 0x14, 0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x2e, 0x42, 0x6c, 0x61,
	0x6e, 0x6b, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x73, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d,
	0x64, 0x65, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3a, 0x0a, 0x07, 0x52, 0x65, 0x67, 0x65, 0x78, 0x70, 0x73, 0x12, 0x14, 0x2e, 0x67,
	0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x2e, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x73, 0x1a, 0x17, 0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x67,
	0x65, 0x78, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x15, 0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65,
	0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12,
	0x14, 0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x2e,
	0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40,
	0x0a, 0x0c, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x73, 0x12, 0x14,
	0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x2e, 0x42, 0x6c, 0x61, 0x6e, 0x6b, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x18, 0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x2e, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00,
	0x12, 0x5c, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64,
	0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x61, 0x6e, 0x79,
	0x6d, 0x64, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x68,
	0x0a, 0x15, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x25, 0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64,
	0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f,
	0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26,
	0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x6f, 0x64, 0x65, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x0d, 0x52, 0x65, 0x73, 0x65,
	0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x2e, 0x67, 0x61, 0x6e, 0x79,
	0x6d, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d,
	0x64, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x0a, 0x4f, 0x41,
	0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d,
	0x64, 0x65, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x2e, 0x4f,
	0x41, 0x75, 0x74, 0x68, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x44, 0x0a, 0x09, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x19, 0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x61,
	0x6e, 0x79, 0x6d, 0x64, 0x65, 0x2e, 0x4f, 0x41, 0x75, 0x74, 0x68, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3e, 0x0a, 0x07, 0x49, 0x73, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x12, 0x17, 0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x2e, 0x49,
	0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e,
	0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x2e, 0x49, 0x73, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x06, 0x4c, 0x6f, 0x67,
	0x4f, 0x75, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x64, 0x65, 0x2e, 0x4c, 0x6f,
	0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x67, 0x61,
	0x6e, 0x79, 0x6d, 0x64, 0x65, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x28, 0x5a, 0x26, 0x61, 0x6d, 0x75, 0x73, 0x69, 0x6e,
	0x67, 0x2d, 0x78, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x65, 0x64, 0x65,
	0x2f, 0x67, 0x61, 0x6e, 0x79, 0x6d, 0x65, 0x64, 0x65, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_ganymede_proto_server_proto_rawDescOnce sync.Once
	file_ganymede_proto_server_proto_rawDescData = file_ganymede_proto_server_proto_rawDesc
)

func file_ganymede_proto_server_proto_rawDescGZIP() []byte {
	file_ganymede_proto_server_proto_rawDescOnce.Do(func() {
		file_ganymede_proto_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_ganymede_proto_server_proto_rawDescData)
	})
	return file_ganymede_proto_server_proto_rawDescData
}

var file_ganymede_proto_server_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_ganymede_proto_server_proto_goTypes = []interface{}{
	(*BlankParams)(nil),                   // 0: ganymde.BlankParams
	(*LoginRequest)(nil),                  // 1: ganymde.LoginRequest
	(*JoinRequest)(nil),                   // 2: ganymde.JoinRequest
	(*VerificationCodeRequest)(nil),       // 3: ganymde.VerificationCodeRequest
	(*VerificationCodeCheckRequest)(nil),  // 4: ganymde.VerificationCodeCheckRequest
	(*ResetPasswordRequest)(nil),          // 5: ganymde.ResetPasswordRequest
	(*OAuthLoginRequest)(nil),             // 6: ganymde.OAuthLoginRequest
	(*OAuthInfoRequest)(nil),              // 7: ganymde.OAuthInfoRequest
	(*IsLoginRequest)(nil),                // 8: ganymde.IsLoginRequest
	(*LogoutRequest)(nil),                 // 9: ganymde.LogoutRequest
	(*PongResponse)(nil),                  // 10: ganymde.PongResponse
	(*RegexpResponse)(nil),                // 11: ganymde.RegexpResponse
	(*LoginResponse)(nil),                 // 12: ganymde.LoginResponse
	(*JoinResponse)(nil),                  // 13: ganymde.JoinResponse
	(*CountryCodeList)(nil),               // 14: ganymde.CountryCodeList
	(*VerificationCodeResponse)(nil),      // 15: ganymde.VerificationCodeResponse
	(*VerificationCodeCheckResponse)(nil), // 16: ganymde.VerificationCodeCheckResponse
	(*ResetPasswordResponse)(nil),         // 17: ganymde.ResetPasswordResponse
	(*OAuthLoginResponse)(nil),            // 18: ganymde.OAuthLoginResponse
	(*OAuthInfoResponse)(nil),             // 19: ganymde.OAuthInfoResponse
	(*IsLoginResponse)(nil),               // 20: ganymde.IsLoginResponse
	(*LogoutResponse)(nil),                // 21: ganymde.LogoutResponse
}
var file_ganymede_proto_server_proto_depIdxs = []int32{
	0,  // 0: ganymde.GanymedeService.Pong:input_type -> ganymde.BlankParams
	0,  // 1: ganymde.GanymedeService.Regexps:input_type -> ganymde.BlankParams
	1,  // 2: ganymde.GanymedeService.Login:input_type -> ganymde.LoginRequest
	2,  // 3: ganymde.GanymedeService.Join:input_type -> ganymde.JoinRequest
	0,  // 4: ganymde.GanymedeService.CountryCodes:input_type -> ganymde.BlankParams
	3,  // 5: ganymde.GanymedeService.GetVerificationCode:input_type -> ganymde.VerificationCodeRequest
	4,  // 6: ganymde.GanymedeService.VerificationCodeCheck:input_type -> ganymde.VerificationCodeCheckRequest
	5,  // 7: ganymde.GanymedeService.ResetPassword:input_type -> ganymde.ResetPasswordRequest
	6,  // 8: ganymde.GanymedeService.OAuthLogin:input_type -> ganymde.OAuthLoginRequest
	7,  // 9: ganymde.GanymedeService.OAuthInfo:input_type -> ganymde.OAuthInfoRequest
	8,  // 10: ganymde.GanymedeService.IsLogin:input_type -> ganymde.IsLoginRequest
	9,  // 11: ganymde.GanymedeService.LogOut:input_type -> ganymde.LogoutRequest
	10, // 12: ganymde.GanymedeService.Pong:output_type -> ganymde.PongResponse
	11, // 13: ganymde.GanymedeService.Regexps:output_type -> ganymde.RegexpResponse
	12, // 14: ganymde.GanymedeService.Login:output_type -> ganymde.LoginResponse
	13, // 15: ganymde.GanymedeService.Join:output_type -> ganymde.JoinResponse
	14, // 16: ganymde.GanymedeService.CountryCodes:output_type -> ganymde.CountryCodeList
	15, // 17: ganymde.GanymedeService.GetVerificationCode:output_type -> ganymde.VerificationCodeResponse
	16, // 18: ganymde.GanymedeService.VerificationCodeCheck:output_type -> ganymde.VerificationCodeCheckResponse
	17, // 19: ganymde.GanymedeService.ResetPassword:output_type -> ganymde.ResetPasswordResponse
	18, // 20: ganymde.GanymedeService.OAuthLogin:output_type -> ganymde.OAuthLoginResponse
	19, // 21: ganymde.GanymedeService.OAuthInfo:output_type -> ganymde.OAuthInfoResponse
	20, // 22: ganymde.GanymedeService.IsLogin:output_type -> ganymde.IsLoginResponse
	21, // 23: ganymde.GanymedeService.LogOut:output_type -> ganymde.LogoutResponse
	12, // [12:24] is the sub-list for method output_type
	0,  // [0:12] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_ganymede_proto_server_proto_init() }
func file_ganymede_proto_server_proto_init() {
	if File_ganymede_proto_server_proto != nil {
		return
	}
	file_ganymede_proto_regexp_proto_init()
	file_ganymede_proto_login_proto_init()
	file_ganymede_proto_join_proto_init()
	file_ganymede_proto_country_code_proto_init()
	file_ganymede_proto_verification_code_proto_init()
	file_ganymede_proto_verification_code_check_proto_init()
	file_ganymede_proto_reset_password_proto_init()
	file_ganymede_proto_oauth_login_proto_init()
	file_ganymede_proto_oauth_info_proto_init()
	file_ganymede_proto_is_login_proto_init()
	file_ganymede_proto_logout_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ganymede_proto_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlankParams); i {
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
			RawDescriptor: file_ganymede_proto_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_ganymede_proto_server_proto_goTypes,
		DependencyIndexes: file_ganymede_proto_server_proto_depIdxs,
		MessageInfos:      file_ganymede_proto_server_proto_msgTypes,
	}.Build()
	File_ganymede_proto_server_proto = out.File
	file_ganymede_proto_server_proto_rawDesc = nil
	file_ganymede_proto_server_proto_goTypes = nil
	file_ganymede_proto_server_proto_depIdxs = nil
}