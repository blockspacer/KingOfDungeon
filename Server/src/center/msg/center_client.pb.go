// Code generated by protoc-gen-go. DO NOT EDIT.
// source: center_client.proto

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
	return fileDescriptor_832f0ed572a10b2d, []int{0}
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
	return fileDescriptor_832f0ed572a10b2d, []int{1}
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
	return fileDescriptor_832f0ed572a10b2d, []int{2}
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
	MapId                uint32   `protobuf:"varint,5,opt,name=mapId,proto3" json:"mapId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRoleInfo) Reset()         { *m = LoginRoleInfo{} }
func (m *LoginRoleInfo) String() string { return proto.CompactTextString(m) }
func (*LoginRoleInfo) ProtoMessage()    {}
func (*LoginRoleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_832f0ed572a10b2d, []int{3}
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

func (m *LoginRoleInfo) GetMapId() uint32 {
	if m != nil {
		return m.MapId
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
	return fileDescriptor_832f0ed572a10b2d, []int{4}
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
	return fileDescriptor_832f0ed572a10b2d, []int{5}
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
	return fileDescriptor_832f0ed572a10b2d, []int{6}
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

//删除角色
type ReqDelRole struct {
	RoleId               string   `protobuf:"bytes,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqDelRole) Reset()         { *m = ReqDelRole{} }
func (m *ReqDelRole) String() string { return proto.CompactTextString(m) }
func (*ReqDelRole) ProtoMessage()    {}
func (*ReqDelRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_832f0ed572a10b2d, []int{7}
}

func (m *ReqDelRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqDelRole.Unmarshal(m, b)
}
func (m *ReqDelRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqDelRole.Marshal(b, m, deterministic)
}
func (m *ReqDelRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqDelRole.Merge(m, src)
}
func (m *ReqDelRole) XXX_Size() int {
	return xxx_messageInfo_ReqDelRole.Size(m)
}
func (m *ReqDelRole) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqDelRole.DiscardUnknown(m)
}

var xxx_messageInfo_ReqDelRole proto.InternalMessageInfo

func (m *ReqDelRole) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

type RspDelRole struct {
	RetCode              int32    `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	RoleId               string   `protobuf:"bytes,2,opt,name=roleId,proto3" json:"roleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RspDelRole) Reset()         { *m = RspDelRole{} }
func (m *RspDelRole) String() string { return proto.CompactTextString(m) }
func (*RspDelRole) ProtoMessage()    {}
func (*RspDelRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_832f0ed572a10b2d, []int{8}
}

func (m *RspDelRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RspDelRole.Unmarshal(m, b)
}
func (m *RspDelRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RspDelRole.Marshal(b, m, deterministic)
}
func (m *RspDelRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RspDelRole.Merge(m, src)
}
func (m *RspDelRole) XXX_Size() int {
	return xxx_messageInfo_RspDelRole.Size(m)
}
func (m *RspDelRole) XXX_DiscardUnknown() {
	xxx_messageInfo_RspDelRole.DiscardUnknown(m)
}

var xxx_messageInfo_RspDelRole proto.InternalMessageInfo

func (m *RspDelRole) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *RspDelRole) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

//选择角色进入游戏地图
type ReqSelectRole struct {
	RoleId               string   `protobuf:"bytes,1,opt,name=roleId,proto3" json:"roleId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqSelectRole) Reset()         { *m = ReqSelectRole{} }
func (m *ReqSelectRole) String() string { return proto.CompactTextString(m) }
func (*ReqSelectRole) ProtoMessage()    {}
func (*ReqSelectRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_832f0ed572a10b2d, []int{9}
}

func (m *ReqSelectRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqSelectRole.Unmarshal(m, b)
}
func (m *ReqSelectRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqSelectRole.Marshal(b, m, deterministic)
}
func (m *ReqSelectRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqSelectRole.Merge(m, src)
}
func (m *ReqSelectRole) XXX_Size() int {
	return xxx_messageInfo_ReqSelectRole.Size(m)
}
func (m *ReqSelectRole) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqSelectRole.DiscardUnknown(m)
}

var xxx_messageInfo_ReqSelectRole proto.InternalMessageInfo

func (m *ReqSelectRole) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

type RspSelectRole struct {
	RetCode              int32    `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	RoleId               string   `protobuf:"bytes,2,opt,name=roleId,proto3" json:"roleId,omitempty"`
	MapId                uint32   `protobuf:"varint,3,opt,name=mapId,proto3" json:"mapId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RspSelectRole) Reset()         { *m = RspSelectRole{} }
func (m *RspSelectRole) String() string { return proto.CompactTextString(m) }
func (*RspSelectRole) ProtoMessage()    {}
func (*RspSelectRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_832f0ed572a10b2d, []int{10}
}

func (m *RspSelectRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RspSelectRole.Unmarshal(m, b)
}
func (m *RspSelectRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RspSelectRole.Marshal(b, m, deterministic)
}
func (m *RspSelectRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RspSelectRole.Merge(m, src)
}
func (m *RspSelectRole) XXX_Size() int {
	return xxx_messageInfo_RspSelectRole.Size(m)
}
func (m *RspSelectRole) XXX_DiscardUnknown() {
	xxx_messageInfo_RspSelectRole.DiscardUnknown(m)
}

var xxx_messageInfo_RspSelectRole proto.InternalMessageInfo

func (m *RspSelectRole) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *RspSelectRole) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *RspSelectRole) GetMapId() uint32 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func init() {
	proto.RegisterType((*ReqLogin)(nil), "msg.req_login")
	proto.RegisterType((*RspLogin)(nil), "msg.rsp_login")
	proto.RegisterType((*ReqRolelist)(nil), "msg.req_rolelist")
	proto.RegisterType((*LoginRoleInfo)(nil), "msg.login_role_info")
	proto.RegisterType((*RspRolelist)(nil), "msg.rsp_rolelist")
	proto.RegisterType((*ReqCreateRole)(nil), "msg.req_create_role")
	proto.RegisterType((*RspCreateRole)(nil), "msg.rsp_create_role")
	proto.RegisterType((*ReqDelRole)(nil), "msg.req_del_role")
	proto.RegisterType((*RspDelRole)(nil), "msg.rsp_del_role")
	proto.RegisterType((*ReqSelectRole)(nil), "msg.req_select_role")
	proto.RegisterType((*RspSelectRole)(nil), "msg.rsp_select_role")
}

func init() { proto.RegisterFile("center_client.proto", fileDescriptor_832f0ed572a10b2d) }

var fileDescriptor_832f0ed572a10b2d = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcf, 0x4b, 0x33, 0x31,
	0x10, 0x65, 0xbb, 0xfd, 0xf1, 0xed, 0x7c, 0xad, 0x95, 0x58, 0x24, 0xde, 0x4a, 0x44, 0x59, 0x2f,
	0x3d, 0xd4, 0x83, 0x57, 0xc1, 0x53, 0x2f, 0x1e, 0x16, 0x3c, 0x78, 0x2a, 0x75, 0x77, 0x5a, 0x16,
	0xb2, 0x49, 0x4d, 0x52, 0x15, 0xfc, 0xe7, 0x25, 0x93, 0xb6, 0xbb, 0x0a, 0x2b, 0x78, 0xcb, 0x1b,
	0xde, 0xcc, 0x7b, 0x6f, 0x26, 0x70, 0x96, 0xa3, 0x72, 0x68, 0x96, 0xb9, 0x2c, 0x51, 0xb9, 0xd9,
	0xd6, 0x68, 0xa7, 0x59, 0x5c, 0xd9, 0x8d, 0xb8, 0x84, 0xc4, 0xe0, 0xeb, 0x52, 0xea, 0x4d, 0xa9,
	0xd8, 0x39, 0xf4, 0x77, 0x16, 0xcd, 0xa2, 0xe0, 0xd1, 0x34, 0x4a, 0x93, 0x6c, 0x8f, 0xc4, 0x15,
	0x24, 0xc6, 0x6e, 0xf7, 0x24, 0x0e, 0x03, 0x83, 0xee, 0x41, 0x17, 0x48, 0xac, 0x5e, 0x76, 0x80,
	0xe2, 0x04, 0x86, 0x7e, 0x96, 0xd1, 0x12, 0x65, 0x69, 0x9d, 0xf8, 0x84, 0x31, 0xb5, 0x50, 0x65,
	0x59, 0xaa, 0xb5, 0xf6, 0x0a, 0x1e, 0xd4, 0x0a, 0x01, 0x31, 0x06, 0x5d, 0xb5, 0xaa, 0x90, 0x77,
	0xa8, 0x4a, 0x6f, 0x36, 0x81, 0x9e, 0xc4, 0x37, 0x94, 0x3c, 0x9e, 0x46, 0xe9, 0x28, 0x0b, 0x80,
	0x9d, 0x42, 0x6c, 0xf1, 0x83, 0x77, 0xa9, 0xe6, 0x9f, 0x9e, 0x57, 0xad, 0xb6, 0x8b, 0x82, 0xf7,
	0x02, 0x8f, 0x80, 0x78, 0x87, 0xa1, 0xf7, 0x7c, 0x30, 0xd3, 0x6e, 0x9b, 0x5d, 0xc0, 0x3f, 0x32,
	0xa8, 0x76, 0x15, 0xe9, 0x8f, 0xb2, 0x81, 0xc7, 0x8f, 0xbb, 0x8a, 0xcd, 0x21, 0x21, 0x83, 0x6a,
	0xad, 0x2d, 0x8f, 0xa7, 0x71, 0xfa, 0x7f, 0x3e, 0x99, 0x55, 0x76, 0x33, 0xfb, 0x91, 0x2b, 0xab,
	0x69, 0xe2, 0x0e, 0xc6, 0x7e, 0x0b, 0xb9, 0xc1, 0x95, 0x43, 0xa2, 0x1c, 0xd3, 0x45, 0x8d, 0x74,
	0xfb, 0x1c, 0x9d, 0x63, 0x0e, 0xf1, 0x04, 0x63, 0xef, 0xb8, 0xd9, 0xd8, 0x6e, 0x3a, 0x85, 0xae,
	0x17, 0xa6, 0xfe, 0x36, 0x53, 0xc4, 0x10, 0xd7, 0xe1, 0x2a, 0x05, 0xca, 0x30, 0xb3, 0xe5, 0x04,
	0xe2, 0x3e, 0x2c, 0xec, 0xc8, 0x6b, 0xd7, 0xae, 0x27, 0x74, 0xbe, 0x4d, 0xb8, 0x09, 0xc9, 0x2d,
	0x4a, 0xcc, 0xdd, 0xef, 0x62, 0xcf, 0x21, 0x6b, 0x93, 0xfa, 0x67, 0xbd, 0xfa, 0xf0, 0x71, 0xe3,
	0xf0, 0x2f, 0x7d, 0xfa, 0xdd, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x72, 0x80, 0x99,
	0xf4, 0x02, 0x00, 0x00,
}
