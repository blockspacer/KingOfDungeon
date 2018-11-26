// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login.proto

package msg

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//登录请求
type ReqLogin struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqLogin) Reset()         { *m = ReqLogin{} }
func (m *ReqLogin) String() string { return proto.CompactTextString(m) }
func (*ReqLogin) ProtoMessage()    {}
func (*ReqLogin) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{0}
}

func (m *ReqLogin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqLogin.Unmarshal(m, b)
}
func (m *ReqLogin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqLogin.Marshal(b, m, deterministic)
}
func (m *ReqLogin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqLogin.Merge(m, src)
}
func (m *ReqLogin) XXX_Size() int {
	return xxx_messageInfo_ReqLogin.Size(m)
}
func (m *ReqLogin) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqLogin.DiscardUnknown(m)
}

var xxx_messageInfo_ReqLogin proto.InternalMessageInfo

func (m *ReqLogin) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

//登录返回
type RspLogin struct {
	RetCode              int32    `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RspLogin) Reset()         { *m = RspLogin{} }
func (m *RspLogin) String() string { return proto.CompactTextString(m) }
func (*RspLogin) ProtoMessage()    {}
func (*RspLogin) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{1}
}

func (m *RspLogin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RspLogin.Unmarshal(m, b)
}
func (m *RspLogin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RspLogin.Marshal(b, m, deterministic)
}
func (m *RspLogin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RspLogin.Merge(m, src)
}
func (m *RspLogin) XXX_Size() int {
	return xxx_messageInfo_RspLogin.Size(m)
}
func (m *RspLogin) XXX_DiscardUnknown() {
	xxx_messageInfo_RspLogin.DiscardUnknown(m)
}

var xxx_messageInfo_RspLogin proto.InternalMessageInfo

func (m *RspLogin) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

//获取角色列表
type ReqRolelist struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqRolelist) Reset()         { *m = ReqRolelist{} }
func (m *ReqRolelist) String() string { return proto.CompactTextString(m) }
func (*ReqRolelist) ProtoMessage()    {}
func (*ReqRolelist) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{2}
}

func (m *ReqRolelist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqRolelist.Unmarshal(m, b)
}
func (m *ReqRolelist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqRolelist.Marshal(b, m, deterministic)
}
func (m *ReqRolelist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqRolelist.Merge(m, src)
}
func (m *ReqRolelist) XXX_Size() int {
	return xxx_messageInfo_ReqRolelist.Size(m)
}
func (m *ReqRolelist) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqRolelist.DiscardUnknown(m)
}

var xxx_messageInfo_ReqRolelist proto.InternalMessageInfo

//角色列表中的表项定义
type LoginRoleInfo struct {
	RoleId               string   `protobuf:"bytes,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Level                uint32   `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	Sex                  uint32   `protobuf:"varint,4,opt,name=sex,proto3" json:"sex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRoleInfo) Reset()         { *m = LoginRoleInfo{} }
func (m *LoginRoleInfo) String() string { return proto.CompactTextString(m) }
func (*LoginRoleInfo) ProtoMessage()    {}
func (*LoginRoleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{3}
}

func (m *LoginRoleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRoleInfo.Unmarshal(m, b)
}
func (m *LoginRoleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRoleInfo.Marshal(b, m, deterministic)
}
func (m *LoginRoleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRoleInfo.Merge(m, src)
}
func (m *LoginRoleInfo) XXX_Size() int {
	return xxx_messageInfo_LoginRoleInfo.Size(m)
}
func (m *LoginRoleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRoleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRoleInfo proto.InternalMessageInfo

func (m *LoginRoleInfo) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *LoginRoleInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginRoleInfo) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *LoginRoleInfo) GetSex() uint32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

//返回角色列表
type RspRolelist struct {
	RetCode              int32            `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	RoleNum              uint32           `protobuf:"varint,2,opt,name=role_num,json=roleNum,proto3" json:"role_num,omitempty"`
	RoleInfos            []*LoginRoleInfo `protobuf:"bytes,3,rep,name=roleInfos,proto3" json:"roleInfos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *RspRolelist) Reset()         { *m = RspRolelist{} }
func (m *RspRolelist) String() string { return proto.CompactTextString(m) }
func (*RspRolelist) ProtoMessage()    {}
func (*RspRolelist) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{4}
}

func (m *RspRolelist) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RspRolelist.Unmarshal(m, b)
}
func (m *RspRolelist) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RspRolelist.Marshal(b, m, deterministic)
}
func (m *RspRolelist) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RspRolelist.Merge(m, src)
}
func (m *RspRolelist) XXX_Size() int {
	return xxx_messageInfo_RspRolelist.Size(m)
}
func (m *RspRolelist) XXX_DiscardUnknown() {
	xxx_messageInfo_RspRolelist.DiscardUnknown(m)
}

var xxx_messageInfo_RspRolelist proto.InternalMessageInfo

func (m *RspRolelist) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *RspRolelist) GetRoleNum() uint32 {
	if m != nil {
		return m.RoleNum
	}
	return 0
}

func (m *RspRolelist) GetRoleInfos() []*LoginRoleInfo {
	if m != nil {
		return m.RoleInfos
	}
	return nil
}

//客户端请求创建角色
type ReqCreateRole struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Sex                  uint32   `protobuf:"varint,2,opt,name=sex,proto3" json:"sex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqCreateRole) Reset()         { *m = ReqCreateRole{} }
func (m *ReqCreateRole) String() string { return proto.CompactTextString(m) }
func (*ReqCreateRole) ProtoMessage()    {}
func (*ReqCreateRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{5}
}

func (m *ReqCreateRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqCreateRole.Unmarshal(m, b)
}
func (m *ReqCreateRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqCreateRole.Marshal(b, m, deterministic)
}
func (m *ReqCreateRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqCreateRole.Merge(m, src)
}
func (m *ReqCreateRole) XXX_Size() int {
	return xxx_messageInfo_ReqCreateRole.Size(m)
}
func (m *ReqCreateRole) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqCreateRole.DiscardUnknown(m)
}

var xxx_messageInfo_ReqCreateRole proto.InternalMessageInfo

func (m *ReqCreateRole) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ReqCreateRole) GetSex() uint32 {
	if m != nil {
		return m.Sex
	}
	return 0
}

//服务端相应创建角色
type RspCreateRole struct {
	RetCode              int32          `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	Info                 *LoginRoleInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *RspCreateRole) Reset()         { *m = RspCreateRole{} }
func (m *RspCreateRole) String() string { return proto.CompactTextString(m) }
func (*RspCreateRole) ProtoMessage()    {}
func (*RspCreateRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{6}
}

func (m *RspCreateRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RspCreateRole.Unmarshal(m, b)
}
func (m *RspCreateRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RspCreateRole.Marshal(b, m, deterministic)
}
func (m *RspCreateRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RspCreateRole.Merge(m, src)
}
func (m *RspCreateRole) XXX_Size() int {
	return xxx_messageInfo_RspCreateRole.Size(m)
}
func (m *RspCreateRole) XXX_DiscardUnknown() {
	xxx_messageInfo_RspCreateRole.DiscardUnknown(m)
}

var xxx_messageInfo_RspCreateRole proto.InternalMessageInfo

func (m *RspCreateRole) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *RspCreateRole) GetInfo() *LoginRoleInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*ReqLogin)(nil), "msg.req_login")
	proto.RegisterType((*RspLogin)(nil), "msg.rsp_login")
	proto.RegisterType((*ReqRolelist)(nil), "msg.req_rolelist")
	proto.RegisterType((*LoginRoleInfo)(nil), "msg.login_role_info")
	proto.RegisterType((*RspRolelist)(nil), "msg.rsp_rolelist")
	proto.RegisterType((*ReqCreateRole)(nil), "msg.req_create_role")
	proto.RegisterType((*RspCreateRole)(nil), "msg.rsp_create_role")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor_67c21677aa7f4e4f) }

var fileDescriptor_67c21677aa7f4e4f = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4b, 0xc3, 0x40,
	0x14, 0xc4, 0x49, 0x37, 0x6d, 0xcd, 0x6b, 0x6b, 0x64, 0x29, 0xb2, 0xde, 0xc2, 0x8a, 0x90, 0x53,
	0x0e, 0xf5, 0xe0, 0x07, 0xf0, 0xe4, 0xc5, 0xc3, 0x82, 0xe7, 0x52, 0xed, 0x6b, 0x08, 0xec, 0x9f,
	0xba, 0x9b, 0xa8, 0x1f, 0x5f, 0xf6, 0x35, 0x35, 0x45, 0xc8, 0x6d, 0x67, 0x98, 0x97, 0xf9, 0x0d,
	0x81, 0x85, 0x76, 0x75, 0x63, 0xab, 0xa3, 0x77, 0xad, 0xe3, 0xcc, 0x84, 0x5a, 0xde, 0x43, 0xe6,
	0xf1, 0x73, 0x4b, 0x3e, 0xbf, 0x85, 0x59, 0x17, 0xd0, 0xbf, 0xec, 0x45, 0x52, 0x24, 0x65, 0xa6,
	0x7a, 0x25, 0x1f, 0x20, 0xf3, 0xe1, 0xd8, 0x87, 0x04, 0xcc, 0x3d, 0xb6, 0xcf, 0x6e, 0x8f, 0x94,
	0x9a, 0xaa, 0xb3, 0x94, 0xd7, 0xb0, 0x8c, 0xdf, 0xf2, 0x4e, 0xa3, 0x6e, 0x42, 0x2b, 0x11, 0x72,
	0x3a, 0x21, 0x67, 0xdb, 0xd8, 0x83, 0x8b, 0x0d, 0x51, 0x0c, 0x0d, 0x27, 0xc5, 0x39, 0xa4, 0x76,
	0x67, 0x50, 0x4c, 0xc8, 0xa5, 0x37, 0x5f, 0xc3, 0x54, 0xe3, 0x17, 0x6a, 0xc1, 0x8a, 0xa4, 0x5c,
	0xa9, 0x93, 0xe0, 0x37, 0xc0, 0x02, 0xfe, 0x88, 0x94, 0xbc, 0xf8, 0x94, 0xdf, 0xb0, 0x8c, 0x74,
	0xe7, 0xda, 0x71, 0x40, 0x7e, 0x07, 0x57, 0x84, 0x62, 0x3b, 0x43, 0x4d, 0x2b, 0x35, 0x8f, 0xfa,
	0xb5, 0x33, 0x7c, 0x03, 0x19, 0xa1, 0xd8, 0x83, 0x0b, 0x82, 0x15, 0xac, 0x5c, 0x6c, 0xd6, 0x95,
	0x09, 0x75, 0xf5, 0x6f, 0x81, 0x1a, 0x62, 0xf2, 0x09, 0xf2, 0xb8, 0xf7, 0xc3, 0xe3, 0xae, 0x45,
	0x8a, 0xfc, 0xed, 0x48, 0x2e, 0x76, 0xf4, 0xc4, 0x93, 0x81, 0xf8, 0x0d, 0xf2, 0x48, 0x7c, 0x79,
	0x38, 0x0e, 0x5d, 0x42, 0x1a, 0x8b, 0xe9, 0x7e, 0x0c, 0x8a, 0x12, 0xef, 0x33, 0xfa, 0xaf, 0x8f,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x3b, 0x74, 0x6d, 0xed, 0xe6, 0x01, 0x00, 0x00,
}
