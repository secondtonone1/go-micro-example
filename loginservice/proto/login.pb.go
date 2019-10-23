// Code generated by protoc-gen-go. DO NOT EDIT.
// source: login.proto

package loginproto

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

//protoc   --proto_path=.  --micro_out=.  --go_out=. ./*.proto
type LoginReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Passwd               string   `protobuf:"bytes,2,opt,name=passwd,proto3" json:"passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{0}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LoginReq) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type LoginRsp struct {
	Errid                int32    `protobuf:"varint,1,opt,name=errid,proto3" json:"errid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRsp) Reset()         { *m = LoginRsp{} }
func (m *LoginRsp) String() string { return proto.CompactTextString(m) }
func (*LoginRsp) ProtoMessage()    {}
func (*LoginRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{1}
}

func (m *LoginRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRsp.Unmarshal(m, b)
}
func (m *LoginRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRsp.Marshal(b, m, deterministic)
}
func (m *LoginRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRsp.Merge(m, src)
}
func (m *LoginRsp) XXX_Size() int {
	return xxx_messageInfo_LoginRsp.Size(m)
}
func (m *LoginRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRsp.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRsp proto.InternalMessageInfo

func (m *LoginRsp) GetErrid() int32 {
	if m != nil {
		return m.Errid
	}
	return 0
}

func (m *LoginRsp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RegUsrReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Passwd               string   `protobuf:"bytes,2,opt,name=passwd,proto3" json:"passwd,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegUsrReq) Reset()         { *m = RegUsrReq{} }
func (m *RegUsrReq) String() string { return proto.CompactTextString(m) }
func (*RegUsrReq) ProtoMessage()    {}
func (*RegUsrReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{2}
}

func (m *RegUsrReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegUsrReq.Unmarshal(m, b)
}
func (m *RegUsrReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegUsrReq.Marshal(b, m, deterministic)
}
func (m *RegUsrReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegUsrReq.Merge(m, src)
}
func (m *RegUsrReq) XXX_Size() int {
	return xxx_messageInfo_RegUsrReq.Size(m)
}
func (m *RegUsrReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RegUsrReq.DiscardUnknown(m)
}

var xxx_messageInfo_RegUsrReq proto.InternalMessageInfo

func (m *RegUsrReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegUsrReq) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

func (m *RegUsrReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type RegUsrRsp struct {
	Errid                int32    `protobuf:"varint,1,opt,name=errid,proto3" json:"errid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegUsrRsp) Reset()         { *m = RegUsrRsp{} }
func (m *RegUsrRsp) String() string { return proto.CompactTextString(m) }
func (*RegUsrRsp) ProtoMessage()    {}
func (*RegUsrRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{3}
}

func (m *RegUsrRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegUsrRsp.Unmarshal(m, b)
}
func (m *RegUsrRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegUsrRsp.Marshal(b, m, deterministic)
}
func (m *RegUsrRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegUsrRsp.Merge(m, src)
}
func (m *RegUsrRsp) XXX_Size() int {
	return xxx_messageInfo_RegUsrRsp.Size(m)
}
func (m *RegUsrRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_RegUsrRsp.DiscardUnknown(m)
}

var xxx_messageInfo_RegUsrRsp proto.InternalMessageInfo

func (m *RegUsrRsp) GetErrid() int32 {
	if m != nil {
		return m.Errid
	}
	return 0
}

func (m *RegUsrRsp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegUsrRsp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type ChangewdReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangewdReq) Reset()         { *m = ChangewdReq{} }
func (m *ChangewdReq) String() string { return proto.CompactTextString(m) }
func (*ChangewdReq) ProtoMessage()    {}
func (*ChangewdReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{4}
}

func (m *ChangewdReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangewdReq.Unmarshal(m, b)
}
func (m *ChangewdReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangewdReq.Marshal(b, m, deterministic)
}
func (m *ChangewdReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangewdReq.Merge(m, src)
}
func (m *ChangewdReq) XXX_Size() int {
	return xxx_messageInfo_ChangewdReq.Size(m)
}
func (m *ChangewdReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangewdReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChangewdReq proto.InternalMessageInfo

func (m *ChangewdReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChangewdReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type ChangewdRsp struct {
	Errid                int32    `protobuf:"varint,1,opt,name=errid,proto3" json:"errid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangewdRsp) Reset()         { *m = ChangewdRsp{} }
func (m *ChangewdRsp) String() string { return proto.CompactTextString(m) }
func (*ChangewdRsp) ProtoMessage()    {}
func (*ChangewdRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{5}
}

func (m *ChangewdRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangewdRsp.Unmarshal(m, b)
}
func (m *ChangewdRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangewdRsp.Marshal(b, m, deterministic)
}
func (m *ChangewdRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangewdRsp.Merge(m, src)
}
func (m *ChangewdRsp) XXX_Size() int {
	return xxx_messageInfo_ChangewdRsp.Size(m)
}
func (m *ChangewdRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangewdRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ChangewdRsp proto.InternalMessageInfo

func (m *ChangewdRsp) GetErrid() int32 {
	if m != nil {
		return m.Errid
	}
	return 0
}

func (m *ChangewdRsp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChangewdRsp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

type ChangewdCfmReq struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Passwd               string   `protobuf:"bytes,3,opt,name=passwd,proto3" json:"passwd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangewdCfmReq) Reset()         { *m = ChangewdCfmReq{} }
func (m *ChangewdCfmReq) String() string { return proto.CompactTextString(m) }
func (*ChangewdCfmReq) ProtoMessage()    {}
func (*ChangewdCfmReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{6}
}

func (m *ChangewdCfmReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangewdCfmReq.Unmarshal(m, b)
}
func (m *ChangewdCfmReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangewdCfmReq.Marshal(b, m, deterministic)
}
func (m *ChangewdCfmReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangewdCfmReq.Merge(m, src)
}
func (m *ChangewdCfmReq) XXX_Size() int {
	return xxx_messageInfo_ChangewdCfmReq.Size(m)
}
func (m *ChangewdCfmReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangewdCfmReq.DiscardUnknown(m)
}

var xxx_messageInfo_ChangewdCfmReq proto.InternalMessageInfo

func (m *ChangewdCfmReq) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChangewdCfmReq) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ChangewdCfmReq) GetPasswd() string {
	if m != nil {
		return m.Passwd
	}
	return ""
}

type ChangewdCfmRsp struct {
	Errid                int32    `protobuf:"varint,1,opt,name=errid,proto3" json:"errid,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email                string   `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChangewdCfmRsp) Reset()         { *m = ChangewdCfmRsp{} }
func (m *ChangewdCfmRsp) String() string { return proto.CompactTextString(m) }
func (*ChangewdCfmRsp) ProtoMessage()    {}
func (*ChangewdCfmRsp) Descriptor() ([]byte, []int) {
	return fileDescriptor_67c21677aa7f4e4f, []int{7}
}

func (m *ChangewdCfmRsp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChangewdCfmRsp.Unmarshal(m, b)
}
func (m *ChangewdCfmRsp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChangewdCfmRsp.Marshal(b, m, deterministic)
}
func (m *ChangewdCfmRsp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangewdCfmRsp.Merge(m, src)
}
func (m *ChangewdCfmRsp) XXX_Size() int {
	return xxx_messageInfo_ChangewdCfmRsp.Size(m)
}
func (m *ChangewdCfmRsp) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangewdCfmRsp.DiscardUnknown(m)
}

var xxx_messageInfo_ChangewdCfmRsp proto.InternalMessageInfo

func (m *ChangewdCfmRsp) GetErrid() int32 {
	if m != nil {
		return m.Errid
	}
	return 0
}

func (m *ChangewdCfmRsp) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ChangewdCfmRsp) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginReq)(nil), "loginproto.LoginReq")
	proto.RegisterType((*LoginRsp)(nil), "loginproto.LoginRsp")
	proto.RegisterType((*RegUsrReq)(nil), "loginproto.RegUsrReq")
	proto.RegisterType((*RegUsrRsp)(nil), "loginproto.RegUsrRsp")
	proto.RegisterType((*ChangewdReq)(nil), "loginproto.ChangewdReq")
	proto.RegisterType((*ChangewdRsp)(nil), "loginproto.ChangewdRsp")
	proto.RegisterType((*ChangewdCfmReq)(nil), "loginproto.ChangewdCfmReq")
	proto.RegisterType((*ChangewdCfmRsp)(nil), "loginproto.ChangewdCfmRsp")
}

func init() { proto.RegisterFile("login.proto", fileDescriptor_67c21677aa7f4e4f) }

var fileDescriptor_67c21677aa7f4e4f = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x3d, 0x4f, 0xc3, 0x30,
	0x14, 0xa4, 0x29, 0xa9, 0xda, 0x17, 0xc4, 0xf0, 0x28, 0x10, 0x79, 0x42, 0x9e, 0x98, 0x32, 0xf0,
	0x39, 0xb1, 0x90, 0x11, 0x2a, 0x2a, 0x4b, 0xf9, 0x01, 0x41, 0x71, 0x83, 0xa5, 0xe6, 0x03, 0xbb,
	0x52, 0x7f, 0x09, 0xff, 0x17, 0xf5, 0x99, 0xa4, 0xae, 0x94, 0x0c, 0x11, 0xdd, 0xde, 0x9d, 0x7d,
	0xe7, 0x4b, 0xde, 0x41, 0xb0, 0xae, 0x72, 0x55, 0x46, 0xb5, 0xae, 0x36, 0x15, 0x02, 0x01, 0x9a,
	0xf9, 0x13, 0x4c, 0xdf, 0x77, 0x48, 0xc8, 0x6f, 0x44, 0x38, 0x2d, 0xd3, 0x42, 0x86, 0xa3, 0x9b,
	0xd1, 0xed, 0x4c, 0xd0, 0x8c, 0x57, 0x30, 0xa9, 0x53, 0x63, 0xb6, 0x59, 0xe8, 0x11, 0xfb, 0x87,
	0xf8, 0x43, 0xa3, 0x33, 0x35, 0xce, 0xc1, 0x97, 0x5a, 0xab, 0x8c, 0x84, 0xbe, 0xb0, 0xa0, 0x75,
	0xf3, 0xf6, 0x6e, 0x7c, 0x01, 0x33, 0x21, 0xf3, 0xc4, 0xe8, 0x81, 0xcf, 0xd1, 0x13, 0x45, 0xaa,
	0xd6, 0xe1, 0x98, 0x68, 0x0b, 0xf8, 0x5b, 0x6b, 0x37, 0x24, 0x45, 0x8f, 0xd9, 0x33, 0x04, 0xf1,
	0x57, 0x5a, 0xe6, 0x72, 0x9b, 0xf5, 0xa5, 0x6b, 0x85, 0x9e, 0x2b, 0x5c, 0x38, 0xc2, 0x23, 0xe4,
	0x10, 0x70, 0xde, 0xd8, 0xc5, 0xab, 0x62, 0x50, 0x14, 0xe7, 0xf7, 0x8d, 0x0f, 0xb6, 0xb5, 0x3c,
	0xf4, 0xfc, 0x7f, 0xca, 0xbb, 0x1f, 0x0f, 0xa6, 0x89, 0xd1, 0xd4, 0x01, 0x7c, 0x04, 0xdf, 0x0e,
	0xf3, 0x68, 0x5f, 0xad, 0xa8, 0xe9, 0x15, 0xeb, 0x60, 0x4d, 0xcd, 0x4f, 0xf0, 0x05, 0x02, 0x21,
	0x73, 0x65, 0x36, 0x52, 0x27, 0x46, 0xe3, 0xa5, 0x7b, 0xad, 0xad, 0x09, 0xeb, 0xa2, 0x49, 0xfe,
	0x0a, 0x67, 0xf6, 0xa3, 0x96, 0xb6, 0x23, 0xd7, 0xee, 0x45, 0x67, 0x95, 0xac, 0xfb, 0x80, 0x3c,
	0x3e, 0xe0, 0xc2, 0xf5, 0x88, 0xab, 0x72, 0xa5, 0x74, 0x81, 0xac, 0x4b, 0x61, 0xb7, 0xc1, 0x7a,
	0xcf, 0x76, 0x86, 0x9f, 0x13, 0xe2, 0xef, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x34, 0x62, 0xbc,
	0x18, 0x71, 0x03, 0x00, 0x00,
}
