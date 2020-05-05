// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alpacalabs/auth/v1/session.proto

package v1

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

// An auth session.
type Session struct {
	// the session's globally unique ID.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// the session's secret.
	// if one secret gets compromised, it only compromises this one session
	// and not all sessions in the database. great security feature to have!
	Secret string `protobuf:"bytes,2,opt,name=secret,proto3" json:"secret,omitempty"`
	// the account to which this session belongs.
	AccountId            string   `protobuf:"bytes,3,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_9c03b0f379635ac9, []int{0}
}

func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (m *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(m, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Session) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Session) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func init() {
	proto.RegisterType((*Session)(nil), "alpacalabs.auth.v1.Session")
}

func init() { proto.RegisterFile("alpacalabs/auth/v1/session.proto", fileDescriptor_9c03b0f379635ac9) }

var fileDescriptor_9c03b0f379635ac9 = []byte{
	// 218 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xcc, 0x29, 0x48,
	0x4c, 0x4e, 0xcc, 0x49, 0x4c, 0x2a, 0xd6, 0x4f, 0x2c, 0x2d, 0xc9, 0xd0, 0x2f, 0x33, 0xd4, 0x2f,
	0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x42, 0xa8,
	0xd0, 0x03, 0xa9, 0xd0, 0x2b, 0x33, 0x54, 0x0a, 0xe0, 0x62, 0x0f, 0x86, 0x28, 0x12, 0xe2, 0xe3,
	0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe3,
	0x62, 0x2b, 0x4e, 0x4d, 0x2e, 0x4a, 0x2d, 0x91, 0x60, 0x02, 0x8b, 0x41, 0x79, 0x42, 0xb2, 0x5c,
	0x5c, 0x89, 0xc9, 0xc9, 0xf9, 0xa5, 0x79, 0x25, 0xf1, 0x99, 0x29, 0x12, 0xcc, 0x60, 0x39, 0x4e,
	0xa8, 0x88, 0x67, 0x8a, 0xd3, 0x44, 0x46, 0x2e, 0xb1, 0xe4, 0xfc, 0x5c, 0x3d, 0x4c, 0xcb, 0x9c,
	0x38, 0x1d, 0x4b, 0x4b, 0x32, 0x02, 0x40, 0x6e, 0x09, 0x60, 0x8c, 0xb2, 0x4a, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x77, 0x04, 0xab, 0xf5, 0x01, 0x39, 0x1d, 0xec, 0xd4,
	0xa2, 0xd4, 0x82, 0x7c, 0x5d, 0x90, 0x2e, 0xdd, 0xf4, 0x7c, 0x7d, 0x4c, 0x7f, 0x2d, 0x62, 0x62,
	0x76, 0xf4, 0x71, 0x5c, 0xc5, 0x24, 0x84, 0xd0, 0xa7, 0x07, 0x32, 0x5a, 0x2f, 0xcc, 0xf0, 0x14,
	0xb2, 0x60, 0x0c, 0x48, 0x30, 0x26, 0xcc, 0x30, 0x89, 0x0d, 0x6c, 0xaa, 0x31, 0x20, 0x00, 0x00,
	0xff, 0xff, 0xd7, 0x09, 0x86, 0x43, 0x24, 0x01, 0x00, 0x00,
}
