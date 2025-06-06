// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.9.0
// source: user.proto

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

type UserLoginRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password1     string                 `protobuf:"bytes,2,opt,name=password1,proto3" json:"password1,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserLoginRequest) Reset() {
	*x = UserLoginRequest{}
	mi := &file_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserLoginRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginRequest) ProtoMessage() {}

func (x *UserLoginRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginRequest.ProtoReflect.Descriptor instead.
func (*UserLoginRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserLoginRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserLoginRequest) GetPassword1() string {
	if x != nil {
		return x.Password1
	}
	return ""
}

type UserLoginResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	UserId        uint32                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // 添加token字段
	Msg           string                 `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserLoginResponse) Reset() {
	*x = UserLoginResponse{}
	mi := &file_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserLoginResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginResponse) ProtoMessage() {}

func (x *UserLoginResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginResponse.ProtoReflect.Descriptor instead.
func (*UserLoginResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserLoginResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserLoginResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserLoginResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UserRegisterRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password1     string                 `protobuf:"bytes,2,opt,name=password1,proto3" json:"password1,omitempty"`
	Password2     string                 `protobuf:"bytes,3,opt,name=password2,proto3" json:"password2,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRegisterRequest) Reset() {
	*x = UserRegisterRequest{}
	mi := &file_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRegisterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterRequest) ProtoMessage() {}

func (x *UserRegisterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterRequest.ProtoReflect.Descriptor instead.
func (*UserRegisterRequest) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{2}
}

func (x *UserRegisterRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserRegisterRequest) GetPassword1() string {
	if x != nil {
		return x.Password1
	}
	return ""
}

func (x *UserRegisterRequest) GetPassword2() string {
	if x != nil {
		return x.Password2
	}
	return ""
}

type UserRegisterResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Username      string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	UserId        uint32                 `protobuf:"varint,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Msg           string                 `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *UserRegisterResponse) Reset() {
	*x = UserRegisterResponse{}
	mi := &file_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterResponse) ProtoMessage() {}

func (x *UserRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterResponse.ProtoReflect.Descriptor instead.
func (*UserRegisterResponse) Descriptor() ([]byte, []int) {
	return file_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserRegisterResponse) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserRegisterResponse) GetUserId() uint32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserRegisterResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_user_proto protoreflect.FileDescriptor

const file_user_proto_rawDesc = "" +
	"\n" +
	"\n" +
	"user.proto\x12\x04user\"L\n" +
	"\x10userLoginRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1c\n" +
	"\tpassword1\x18\x02 \x01(\tR\tpassword1\"Z\n" +
	"\x11userLoginResponse\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\rR\x06userId\x12\x10\n" +
	"\x03msg\x18\x03 \x01(\tR\x03msg\"m\n" +
	"\x13userRegisterRequest\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x1c\n" +
	"\tpassword1\x18\x02 \x01(\tR\tpassword1\x12\x1c\n" +
	"\tpassword2\x18\x03 \x01(\tR\tpassword2\"]\n" +
	"\x14userRegisterResponse\x12\x1a\n" +
	"\busername\x18\x01 \x01(\tR\busername\x12\x17\n" +
	"\auser_id\x18\x02 \x01(\rR\x06userId\x12\x10\n" +
	"\x03msg\x18\x03 \x01(\tR\x03msg2\x8b\x01\n" +
	"\x04user\x12<\n" +
	"\tuserLogin\x12\x16.user.userLoginRequest\x1a\x17.user.userLoginResponse\x12E\n" +
	"\fuserRegister\x12\x19.user.userRegisterRequest\x1a\x1a.user.userRegisterResponseB\bZ\x06./userb\x06proto3"

var (
	file_user_proto_rawDescOnce sync.Once
	file_user_proto_rawDescData []byte
)

func file_user_proto_rawDescGZIP() []byte {
	file_user_proto_rawDescOnce.Do(func() {
		file_user_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)))
	})
	return file_user_proto_rawDescData
}

var file_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_user_proto_goTypes = []any{
	(*UserLoginRequest)(nil),     // 0: user.userLoginRequest
	(*UserLoginResponse)(nil),    // 1: user.userLoginResponse
	(*UserRegisterRequest)(nil),  // 2: user.userRegisterRequest
	(*UserRegisterResponse)(nil), // 3: user.userRegisterResponse
}
var file_user_proto_depIdxs = []int32{
	0, // 0: user.user.userLogin:input_type -> user.userLoginRequest
	2, // 1: user.user.userRegister:input_type -> user.userRegisterRequest
	1, // 2: user.user.userLogin:output_type -> user.userLoginResponse
	3, // 3: user.user.userRegister:output_type -> user.userRegisterResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_proto_init() }
func file_user_proto_init() {
	if File_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_user_proto_rawDesc), len(file_user_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_user_proto_goTypes,
		DependencyIndexes: file_user_proto_depIdxs,
		MessageInfos:      file_user_proto_msgTypes,
	}.Build()
	File_user_proto = out.File
	file_user_proto_goTypes = nil
	file_user_proto_depIdxs = nil
}
