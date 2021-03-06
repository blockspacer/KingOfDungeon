// Code generated by protoc-gen-go. DO NOT EDIT.
// source: center_server.proto

package proto

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

//游戏服务器通知中心服务器初始化完成
type NotifyServerInited struct {
	LoadedMaps           []uint32 `protobuf:"varint,1,rep,packed,name=loadedMaps,proto3" json:"loadedMaps,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyServerInited) Reset()         { *m = NotifyServerInited{} }
func (m *NotifyServerInited) String() string { return proto.CompactTextString(m) }
func (*NotifyServerInited) ProtoMessage()    {}
func (*NotifyServerInited) Descriptor() ([]byte, []int) {
	return fileDescriptor_91fb75f26ebeb9d0, []int{0}
}

func (m *NotifyServerInited) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyServerInited.Unmarshal(m, b)
}
func (m *NotifyServerInited) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyServerInited.Marshal(b, m, deterministic)
}
func (m *NotifyServerInited) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyServerInited.Merge(m, src)
}
func (m *NotifyServerInited) XXX_Size() int {
	return xxx_messageInfo_NotifyServerInited.Size(m)
}
func (m *NotifyServerInited) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyServerInited.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyServerInited proto.InternalMessageInfo

func (m *NotifyServerInited) GetLoadedMaps() []uint32 {
	if m != nil {
		return m.LoadedMaps
	}
	return nil
}

//中心服务器通知游戏服务器角色进入地图
type NotifyRoleEnter struct {
	UserId               string   `protobuf:"bytes,1,opt,name=userId,proto3" json:"userId,omitempty"`
	RoleId               string   `protobuf:"bytes,2,opt,name=roleId,proto3" json:"roleId,omitempty"`
	MapId                uint32   `protobuf:"varint,3,opt,name=mapId,proto3" json:"mapId,omitempty"`
	Token                string   `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyRoleEnter) Reset()         { *m = NotifyRoleEnter{} }
func (m *NotifyRoleEnter) String() string { return proto.CompactTextString(m) }
func (*NotifyRoleEnter) ProtoMessage()    {}
func (*NotifyRoleEnter) Descriptor() ([]byte, []int) {
	return fileDescriptor_91fb75f26ebeb9d0, []int{1}
}

func (m *NotifyRoleEnter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyRoleEnter.Unmarshal(m, b)
}
func (m *NotifyRoleEnter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyRoleEnter.Marshal(b, m, deterministic)
}
func (m *NotifyRoleEnter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyRoleEnter.Merge(m, src)
}
func (m *NotifyRoleEnter) XXX_Size() int {
	return xxx_messageInfo_NotifyRoleEnter.Size(m)
}
func (m *NotifyRoleEnter) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyRoleEnter.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyRoleEnter proto.InternalMessageInfo

func (m *NotifyRoleEnter) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *NotifyRoleEnter) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *NotifyRoleEnter) GetMapId() uint32 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *NotifyRoleEnter) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

//游戏服务器通知中心服务器角色已经可以进入地图了
type NotifyRoleEnteredReady struct {
	RetCode              int32    `protobuf:"varint,1,opt,name=retCode,proto3" json:"retCode,omitempty"`
	UserId               string   `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
	RoleId               string   `protobuf:"bytes,3,opt,name=roleId,proto3" json:"roleId,omitempty"`
	MapId                uint32   `protobuf:"varint,4,opt,name=mapId,proto3" json:"mapId,omitempty"`
	Token                string   `protobuf:"bytes,5,opt,name=token,proto3" json:"token,omitempty"`
	ServerIp             string   `protobuf:"bytes,6,opt,name=server_ip,json=serverIp,proto3" json:"server_ip,omitempty"`
	ServerPort           string   `protobuf:"bytes,7,opt,name=server_port,json=serverPort,proto3" json:"server_port,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NotifyRoleEnteredReady) Reset()         { *m = NotifyRoleEnteredReady{} }
func (m *NotifyRoleEnteredReady) String() string { return proto.CompactTextString(m) }
func (*NotifyRoleEnteredReady) ProtoMessage()    {}
func (*NotifyRoleEnteredReady) Descriptor() ([]byte, []int) {
	return fileDescriptor_91fb75f26ebeb9d0, []int{2}
}

func (m *NotifyRoleEnteredReady) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NotifyRoleEnteredReady.Unmarshal(m, b)
}
func (m *NotifyRoleEnteredReady) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NotifyRoleEnteredReady.Marshal(b, m, deterministic)
}
func (m *NotifyRoleEnteredReady) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NotifyRoleEnteredReady.Merge(m, src)
}
func (m *NotifyRoleEnteredReady) XXX_Size() int {
	return xxx_messageInfo_NotifyRoleEnteredReady.Size(m)
}
func (m *NotifyRoleEnteredReady) XXX_DiscardUnknown() {
	xxx_messageInfo_NotifyRoleEnteredReady.DiscardUnknown(m)
}

var xxx_messageInfo_NotifyRoleEnteredReady proto.InternalMessageInfo

func (m *NotifyRoleEnteredReady) GetRetCode() int32 {
	if m != nil {
		return m.RetCode
	}
	return 0
}

func (m *NotifyRoleEnteredReady) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *NotifyRoleEnteredReady) GetRoleId() string {
	if m != nil {
		return m.RoleId
	}
	return ""
}

func (m *NotifyRoleEnteredReady) GetMapId() uint32 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *NotifyRoleEnteredReady) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *NotifyRoleEnteredReady) GetServerIp() string {
	if m != nil {
		return m.ServerIp
	}
	return ""
}

func (m *NotifyRoleEnteredReady) GetServerPort() string {
	if m != nil {
		return m.ServerPort
	}
	return ""
}

func init() {
	proto.RegisterType((*NotifyServerInited)(nil), "proto.notify_server_inited")
	proto.RegisterType((*NotifyRoleEnter)(nil), "proto.notify_role_enter")
	proto.RegisterType((*NotifyRoleEnteredReady)(nil), "proto.notify_role_entered_ready")
}

func init() { proto.RegisterFile("center_server.proto", fileDescriptor_91fb75f26ebeb9d0) }

var fileDescriptor_91fb75f26ebeb9d0 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xc9, 0x76, 0xdb, 0x75, 0x47, 0xf6, 0x60, 0x5c, 0x24, 0x22, 0x68, 0xe9, 0xa9, 0x27,
	0x2f, 0x82, 0x2f, 0xe0, 0xa9, 0x07, 0x41, 0xfa, 0x02, 0xa5, 0x3a, 0x23, 0x14, 0xd7, 0x4e, 0x98,
	0x8e, 0xc2, 0x3e, 0xa5, 0xaf, 0x24, 0x4d, 0xba, 0x50, 0xb0, 0x9e, 0xc2, 0xff, 0xfd, 0x13, 0xe6,
	0x23, 0x81, 0xcb, 0x37, 0xea, 0x95, 0xa4, 0x19, 0x48, 0xbe, 0x49, 0xee, 0xbd, 0xb0, 0xb2, 0x4d,
	0xc3, 0x51, 0x3c, 0xc2, 0xbe, 0x67, 0xed, 0xde, 0x8f, 0x53, 0xdb, 0x74, 0x7d, 0xa7, 0x84, 0xf6,
	0x16, 0xe0, 0xc0, 0x2d, 0x12, 0x3e, 0xb7, 0x7e, 0x70, 0x26, 0x4f, 0xca, 0x5d, 0x3d, 0x23, 0x05,
	0xc3, 0xc5, 0x74, 0x4f, 0xf8, 0x40, 0x4d, 0x58, 0x60, 0xaf, 0x20, 0xfb, 0x1a, 0x48, 0x2a, 0x74,
	0x26, 0x37, 0xe5, 0xb6, 0x9e, 0xd2, 0xc8, 0xc7, 0xa9, 0x0a, 0xdd, 0x2a, 0xf2, 0x98, 0xec, 0x1e,
	0xd2, 0xcf, 0xd6, 0x57, 0xe8, 0x92, 0xdc, 0x94, 0xbb, 0x3a, 0x86, 0x91, 0x2a, 0x7f, 0x50, 0xef,
	0xd6, 0x61, 0x38, 0x86, 0xe2, 0xc7, 0xc0, 0xf5, 0x9f, 0x8d, 0x84, 0x8d, 0x50, 0x8b, 0x47, 0xeb,
	0x60, 0x23, 0xa4, 0x4f, 0x8c, 0x14, 0x56, 0xa7, 0xf5, 0x29, 0xce, 0x9c, 0x56, 0xff, 0x38, 0x25,
	0xcb, 0x4e, 0xeb, 0x45, 0xa7, 0x74, 0xe6, 0x64, 0x6f, 0x60, 0x7b, 0x7a, 0x35, 0xef, 0xb2, 0xd0,
	0x9c, 0x45, 0x50, 0x79, 0x7b, 0x07, 0xe7, 0x53, 0xe9, 0x59, 0xd4, 0x6d, 0x42, 0x0d, 0x11, 0xbd,
	0xb0, 0xe8, 0x6b, 0x16, 0x7e, 0xe0, 0xe1, 0x37, 0x00, 0x00, 0xff, 0xff, 0x65, 0x9f, 0x66, 0xda,
	0x9f, 0x01, 0x00, 0x00,
}
