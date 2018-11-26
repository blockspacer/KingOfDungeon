// Code generated by protoc-gen-go. DO NOT EDIT.
// source: server_client.proto

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

//客户端请求角色信息
type ReqGetRole struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReqGetRole) Reset()         { *m = ReqGetRole{} }
func (m *ReqGetRole) String() string { return proto.CompactTextString(m) }
func (*ReqGetRole) ProtoMessage()    {}
func (*ReqGetRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{0}
}

func (m *ReqGetRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReqGetRole.Unmarshal(m, b)
}
func (m *ReqGetRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReqGetRole.Marshal(b, m, deterministic)
}
func (m *ReqGetRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReqGetRole.Merge(m, src)
}
func (m *ReqGetRole) XXX_Size() int {
	return xxx_messageInfo_ReqGetRole.Size(m)
}
func (m *ReqGetRole) XXX_DiscardUnknown() {
	xxx_messageInfo_ReqGetRole.DiscardUnknown(m)
}

var xxx_messageInfo_ReqGetRole proto.InternalMessageInfo

//角色位置
type RolePos struct {
	X                    float32  `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    float32  `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z                    float32  `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RolePos) Reset()         { *m = RolePos{} }
func (m *RolePos) String() string { return proto.CompactTextString(m) }
func (*RolePos) ProtoMessage()    {}
func (*RolePos) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{1}
}

func (m *RolePos) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RolePos.Unmarshal(m, b)
}
func (m *RolePos) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RolePos.Marshal(b, m, deterministic)
}
func (m *RolePos) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RolePos.Merge(m, src)
}
func (m *RolePos) XXX_Size() int {
	return xxx_messageInfo_RolePos.Size(m)
}
func (m *RolePos) XXX_DiscardUnknown() {
	xxx_messageInfo_RolePos.DiscardUnknown(m)
}

var xxx_messageInfo_RolePos proto.InternalMessageInfo

func (m *RolePos) GetX() float32 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *RolePos) GetY() float32 {
	if m != nil {
		return m.Y
	}
	return 0
}

func (m *RolePos) GetZ() float32 {
	if m != nil {
		return m.Z
	}
	return 0
}

//角色角度
type RoleAngle struct {
	Angle                uint32   `protobuf:"varint,1,opt,name=angle,proto3" json:"angle,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleAngle) Reset()         { *m = RoleAngle{} }
func (m *RoleAngle) String() string { return proto.CompactTextString(m) }
func (*RoleAngle) ProtoMessage()    {}
func (*RoleAngle) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{2}
}

func (m *RoleAngle) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleAngle.Unmarshal(m, b)
}
func (m *RoleAngle) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleAngle.Marshal(b, m, deterministic)
}
func (m *RoleAngle) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleAngle.Merge(m, src)
}
func (m *RoleAngle) XXX_Size() int {
	return xxx_messageInfo_RoleAngle.Size(m)
}
func (m *RoleAngle) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleAngle.DiscardUnknown(m)
}

var xxx_messageInfo_RoleAngle proto.InternalMessageInfo

func (m *RoleAngle) GetAngle() uint32 {
	if m != nil {
		return m.Angle
	}
	return 0
}

//角色信息
type RoleInfo struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Level                uint32   `protobuf:"varint,2,opt,name=level,proto3" json:"level,omitempty"`
	MapId                uint32   `protobuf:"varint,3,opt,name=map_id,json=mapId,proto3" json:"map_id,omitempty"`
	Exp                  uint32   `protobuf:"varint,4,opt,name=exp,proto3" json:"exp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RoleInfo) Reset()         { *m = RoleInfo{} }
func (m *RoleInfo) String() string { return proto.CompactTextString(m) }
func (*RoleInfo) ProtoMessage()    {}
func (*RoleInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{3}
}

func (m *RoleInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoleInfo.Unmarshal(m, b)
}
func (m *RoleInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoleInfo.Marshal(b, m, deterministic)
}
func (m *RoleInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoleInfo.Merge(m, src)
}
func (m *RoleInfo) XXX_Size() int {
	return xxx_messageInfo_RoleInfo.Size(m)
}
func (m *RoleInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoleInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoleInfo proto.InternalMessageInfo

func (m *RoleInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RoleInfo) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *RoleInfo) GetMapId() uint32 {
	if m != nil {
		return m.MapId
	}
	return 0
}

func (m *RoleInfo) GetExp() uint32 {
	if m != nil {
		return m.Exp
	}
	return 0
}

//服务端响应角色信息
type RspGetRole struct {
	Id                   uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Info                 *RoleInfo `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RspGetRole) Reset()         { *m = RspGetRole{} }
func (m *RspGetRole) String() string { return proto.CompactTextString(m) }
func (*RspGetRole) ProtoMessage()    {}
func (*RspGetRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{4}
}

func (m *RspGetRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RspGetRole.Unmarshal(m, b)
}
func (m *RspGetRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RspGetRole.Marshal(b, m, deterministic)
}
func (m *RspGetRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RspGetRole.Merge(m, src)
}
func (m *RspGetRole) XXX_Size() int {
	return xxx_messageInfo_RspGetRole.Size(m)
}
func (m *RspGetRole) XXX_DiscardUnknown() {
	xxx_messageInfo_RspGetRole.DiscardUnknown(m)
}

var xxx_messageInfo_RspGetRole proto.InternalMessageInfo

func (m *RspGetRole) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *RspGetRole) GetInfo() *RoleInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

//服务端推送角色信息改变
type PushRole struct {
	Info                 *RoleInfo `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *PushRole) Reset()         { *m = PushRole{} }
func (m *PushRole) String() string { return proto.CompactTextString(m) }
func (*PushRole) ProtoMessage()    {}
func (*PushRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_55ccd995fd75bbbc, []int{5}
}

func (m *PushRole) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PushRole.Unmarshal(m, b)
}
func (m *PushRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PushRole.Marshal(b, m, deterministic)
}
func (m *PushRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PushRole.Merge(m, src)
}
func (m *PushRole) XXX_Size() int {
	return xxx_messageInfo_PushRole.Size(m)
}
func (m *PushRole) XXX_DiscardUnknown() {
	xxx_messageInfo_PushRole.DiscardUnknown(m)
}

var xxx_messageInfo_PushRole proto.InternalMessageInfo

func (m *PushRole) GetInfo() *RoleInfo {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*ReqGetRole)(nil), "msg.req_get_role")
	proto.RegisterType((*RolePos)(nil), "msg.role_pos")
	proto.RegisterType((*RoleAngle)(nil), "msg.role_angle")
	proto.RegisterType((*RoleInfo)(nil), "msg.role_info")
	proto.RegisterType((*RspGetRole)(nil), "msg.rsp_get_role")
	proto.RegisterType((*PushRole)(nil), "msg.push_role")
}

func init() { proto.RegisterFile("server_client.proto", fileDescriptor_55ccd995fd75bbbc) }

var fileDescriptor_55ccd995fd75bbbc = []byte{
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4e, 0xc4, 0x20,
	0x10, 0x86, 0x43, 0xdb, 0xdd, 0xd8, 0xb1, 0x6d, 0x0c, 0x6a, 0xd2, 0xa3, 0xe1, 0xe4, 0xa9, 0x26,
	0xea, 0x13, 0x78, 0xf3, 0xca, 0x03, 0x48, 0xaa, 0x1d, 0x2b, 0x09, 0x14, 0x84, 0xba, 0xe9, 0xee,
	0xd3, 0x1b, 0x86, 0xb8, 0x37, 0x4f, 0xcc, 0x07, 0xf3, 0xcd, 0x0f, 0xc0, 0x75, 0xc4, 0x70, 0xc0,
	0xa0, 0x3e, 0x8c, 0xc6, 0x65, 0x1d, 0x7c, 0x70, 0xab, 0xe3, 0xa5, 0x8d, 0xb3, 0xe8, 0xa0, 0x09,
	0xf8, 0xad, 0x66, 0x5c, 0x55, 0x70, 0x06, 0xc5, 0x33, 0x5c, 0xa4, 0x55, 0x79, 0x17, 0x79, 0x03,
	0x6c, 0xeb, 0xd9, 0x1d, 0xbb, 0x2f, 0x24, 0xdb, 0x12, 0x1d, 0xfb, 0x22, 0xd3, 0x31, 0xd1, 0xa9,
	0x2f, 0x33, 0x9d, 0x84, 0x00, 0x20, 0x6b, 0x5c, 0x66, 0x83, 0xfc, 0x06, 0x76, 0x54, 0x90, 0xdb,
	0xca, 0x0c, 0xe2, 0x0d, 0x6a, 0xea, 0xd1, 0xcb, 0xa7, 0xe3, 0x1c, 0xaa, 0x65, 0xb4, 0xb9, 0xa3,
	0x96, 0x54, 0x27, 0xcd, 0xe0, 0x01, 0x0d, 0x85, 0xb4, 0x32, 0x03, 0xbf, 0x85, 0xbd, 0x1d, 0xbd,
	0xd2, 0x13, 0xa5, 0xb5, 0x72, 0x67, 0x47, 0xff, 0x3a, 0xf1, 0x2b, 0x28, 0x71, 0xf3, 0x7d, 0x45,
	0x7b, 0xa9, 0x14, 0x2f, 0xd0, 0x84, 0xe8, 0xcf, 0x2f, 0xe1, 0x1d, 0x14, 0x7a, 0xa2, 0x80, 0x4a,
	0x16, 0x7a, 0xe2, 0x02, 0xaa, 0x14, 0x4d, 0xd3, 0x2f, 0x1f, 0xbb, 0xc1, 0xc6, 0x79, 0x38, 0x5f,
	0x48, 0xd2, 0x99, 0x78, 0x80, 0xda, 0xff, 0xc4, 0xaf, 0x3c, 0xe0, 0x4f, 0x60, 0xff, 0x0b, 0xef,
	0x7b, 0xfa, 0xca, 0xa7, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcd, 0xa8, 0x3c, 0xc1, 0x61, 0x01,
	0x00, 0x00,
}
