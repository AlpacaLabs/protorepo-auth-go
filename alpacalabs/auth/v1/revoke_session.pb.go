// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alpacalabs/auth/v1/revoke_session.proto

package v1

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

// Request to revoke a session.
type RevokeSessionRequest struct {
	// session id
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeSessionRequest) Reset()         { *m = RevokeSessionRequest{} }
func (m *RevokeSessionRequest) String() string { return proto.CompactTextString(m) }
func (*RevokeSessionRequest) ProtoMessage()    {}
func (*RevokeSessionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d08be2f88de0f4b1, []int{0}
}

func (m *RevokeSessionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeSessionRequest.Unmarshal(m, b)
}
func (m *RevokeSessionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeSessionRequest.Marshal(b, m, deterministic)
}
func (m *RevokeSessionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeSessionRequest.Merge(m, src)
}
func (m *RevokeSessionRequest) XXX_Size() int {
	return xxx_messageInfo_RevokeSessionRequest.Size(m)
}
func (m *RevokeSessionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeSessionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeSessionRequest proto.InternalMessageInfo

func (m *RevokeSessionRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

// Response from revoking a session.
type RevokeSessionResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeSessionResponse) Reset()         { *m = RevokeSessionResponse{} }
func (m *RevokeSessionResponse) String() string { return proto.CompactTextString(m) }
func (*RevokeSessionResponse) ProtoMessage()    {}
func (*RevokeSessionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d08be2f88de0f4b1, []int{1}
}

func (m *RevokeSessionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeSessionResponse.Unmarshal(m, b)
}
func (m *RevokeSessionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeSessionResponse.Marshal(b, m, deterministic)
}
func (m *RevokeSessionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeSessionResponse.Merge(m, src)
}
func (m *RevokeSessionResponse) XXX_Size() int {
	return xxx_messageInfo_RevokeSessionResponse.Size(m)
}
func (m *RevokeSessionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeSessionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeSessionResponse proto.InternalMessageInfo

// Request to revoke your sessions (except for the one you're currently using, of course).
type RevokeSessionsForAccountRequest struct {
	// the account
	AccountId            string   `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeSessionsForAccountRequest) Reset()         { *m = RevokeSessionsForAccountRequest{} }
func (m *RevokeSessionsForAccountRequest) String() string { return proto.CompactTextString(m) }
func (*RevokeSessionsForAccountRequest) ProtoMessage()    {}
func (*RevokeSessionsForAccountRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d08be2f88de0f4b1, []int{2}
}

func (m *RevokeSessionsForAccountRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeSessionsForAccountRequest.Unmarshal(m, b)
}
func (m *RevokeSessionsForAccountRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeSessionsForAccountRequest.Marshal(b, m, deterministic)
}
func (m *RevokeSessionsForAccountRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeSessionsForAccountRequest.Merge(m, src)
}
func (m *RevokeSessionsForAccountRequest) XXX_Size() int {
	return xxx_messageInfo_RevokeSessionsForAccountRequest.Size(m)
}
func (m *RevokeSessionsForAccountRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeSessionsForAccountRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeSessionsForAccountRequest proto.InternalMessageInfo

func (m *RevokeSessionsForAccountRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

// Response from revoking your sessions.
type RevokeSessionsForAccountResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeSessionsForAccountResponse) Reset()         { *m = RevokeSessionsForAccountResponse{} }
func (m *RevokeSessionsForAccountResponse) String() string { return proto.CompactTextString(m) }
func (*RevokeSessionsForAccountResponse) ProtoMessage()    {}
func (*RevokeSessionsForAccountResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d08be2f88de0f4b1, []int{3}
}

func (m *RevokeSessionsForAccountResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeSessionsForAccountResponse.Unmarshal(m, b)
}
func (m *RevokeSessionsForAccountResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeSessionsForAccountResponse.Marshal(b, m, deterministic)
}
func (m *RevokeSessionsForAccountResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeSessionsForAccountResponse.Merge(m, src)
}
func (m *RevokeSessionsForAccountResponse) XXX_Size() int {
	return xxx_messageInfo_RevokeSessionsForAccountResponse.Size(m)
}
func (m *RevokeSessionsForAccountResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeSessionsForAccountResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeSessionsForAccountResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RevokeSessionRequest)(nil), "alpacalabs.auth.v1.RevokeSessionRequest")
	proto.RegisterType((*RevokeSessionResponse)(nil), "alpacalabs.auth.v1.RevokeSessionResponse")
	proto.RegisterType((*RevokeSessionsForAccountRequest)(nil), "alpacalabs.auth.v1.RevokeSessionsForAccountRequest")
	proto.RegisterType((*RevokeSessionsForAccountResponse)(nil), "alpacalabs.auth.v1.RevokeSessionsForAccountResponse")
}

func init() {
	proto.RegisterFile("alpacalabs/auth/v1/revoke_session.proto", fileDescriptor_d08be2f88de0f4b1)
}

var fileDescriptor_d08be2f88de0f4b1 = []byte{
	// 306 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xcc, 0x29, 0x48,
	0x4c, 0x4e, 0xcc, 0x49, 0x4c, 0x2a, 0xd6, 0x4f, 0x2c, 0x2d, 0xc9, 0xd0, 0x2f, 0x33, 0xd4, 0x2f,
	0x4a, 0x2d, 0xcb, 0xcf, 0x4e, 0x8d, 0x2f, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0x42, 0x28, 0xd4, 0x03, 0x29, 0xd4, 0x2b, 0x33, 0x54, 0x32, 0xe5,
	0x12, 0x09, 0x02, 0xab, 0x0d, 0x86, 0x28, 0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92,
	0xe5, 0xe2, 0x82, 0x6a, 0x8e, 0xcf, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0xe2, 0x84,
	0x8a, 0x78, 0xa6, 0x28, 0x89, 0x73, 0x89, 0xa2, 0x69, 0x2b, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x55,
	0x72, 0xe0, 0x92, 0x47, 0x91, 0x28, 0x76, 0xcb, 0x2f, 0x72, 0x4c, 0x4e, 0xce, 0x2f, 0xcd, 0x2b,
	0x41, 0x32, 0x3a, 0x11, 0x22, 0x82, 0x64, 0x34, 0x54, 0xc4, 0x33, 0x45, 0x49, 0x89, 0x4b, 0x01,
	0xb7, 0x09, 0x10, 0x5b, 0x8c, 0x1a, 0x99, 0xb8, 0x78, 0x61, 0x8a, 0x8a, 0xca, 0x32, 0x93, 0x53,
	0x85, 0xd2, 0x10, 0x02, 0x60, 0x5d, 0x42, 0x1a, 0x7a, 0x98, 0xbe, 0xd5, 0xc3, 0xe6, 0x55, 0x29,
	0x4d, 0x22, 0x54, 0x42, 0x7d, 0xc7, 0x20, 0xd4, 0xce, 0xc8, 0x25, 0x81, 0xcb, 0x79, 0x42, 0xc6,
	0x04, 0x4d, 0xc2, 0x0c, 0x0e, 0x29, 0x13, 0xd2, 0x34, 0xc1, 0x5c, 0xe2, 0x34, 0x91, 0x91, 0x4b,
	0x2c, 0x39, 0x3f, 0x17, 0x8b, 0x6e, 0x27, 0x4e, 0xc7, 0xd2, 0x92, 0x8c, 0x00, 0x50, 0x9c, 0x07,
	0x30, 0x46, 0x59, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x3b, 0x82,
	0xd5, 0xfa, 0x80, 0x52, 0x0a, 0x38, 0x49, 0x14, 0xa5, 0x16, 0xe4, 0xeb, 0x82, 0x74, 0xe9, 0xa6,
	0xe7, 0xeb, 0x63, 0x26, 0xa3, 0x45, 0x4c, 0xcc, 0x8e, 0x3e, 0x8e, 0xab, 0x98, 0x84, 0x10, 0xfa,
	0xf4, 0x40, 0x46, 0xeb, 0x85, 0x19, 0x9e, 0x42, 0x16, 0x8c, 0x01, 0x09, 0xc6, 0x84, 0x19, 0x26,
	0xb1, 0x81, 0x4d, 0x35, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x01, 0xd5, 0xec, 0x93, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RevokeServiceClient is the client API for RevokeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RevokeServiceClient interface {
	// Revoke a particular session.
	RevokeSession(ctx context.Context, in *RevokeSessionRequest, opts ...grpc.CallOption) (*RevokeSessionResponse, error)
	// Revoke your sessions (except for the one you're currently using, of course).
	RevokeSessionsForAccount(ctx context.Context, in *RevokeSessionsForAccountRequest, opts ...grpc.CallOption) (*RevokeSessionsForAccountResponse, error)
}

type revokeServiceClient struct {
	cc *grpc.ClientConn
}

func NewRevokeServiceClient(cc *grpc.ClientConn) RevokeServiceClient {
	return &revokeServiceClient{cc}
}

func (c *revokeServiceClient) RevokeSession(ctx context.Context, in *RevokeSessionRequest, opts ...grpc.CallOption) (*RevokeSessionResponse, error) {
	out := new(RevokeSessionResponse)
	err := c.cc.Invoke(ctx, "/alpacalabs.auth.v1.RevokeService/RevokeSession", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *revokeServiceClient) RevokeSessionsForAccount(ctx context.Context, in *RevokeSessionsForAccountRequest, opts ...grpc.CallOption) (*RevokeSessionsForAccountResponse, error) {
	out := new(RevokeSessionsForAccountResponse)
	err := c.cc.Invoke(ctx, "/alpacalabs.auth.v1.RevokeService/RevokeSessionsForAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RevokeServiceServer is the server API for RevokeService service.
type RevokeServiceServer interface {
	// Revoke a particular session.
	RevokeSession(context.Context, *RevokeSessionRequest) (*RevokeSessionResponse, error)
	// Revoke your sessions (except for the one you're currently using, of course).
	RevokeSessionsForAccount(context.Context, *RevokeSessionsForAccountRequest) (*RevokeSessionsForAccountResponse, error)
}

func RegisterRevokeServiceServer(s *grpc.Server, srv RevokeServiceServer) {
	s.RegisterService(&_RevokeService_serviceDesc, srv)
}

func _RevokeService_RevokeSession_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeSessionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RevokeServiceServer).RevokeSession(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alpacalabs.auth.v1.RevokeService/RevokeSession",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RevokeServiceServer).RevokeSession(ctx, req.(*RevokeSessionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RevokeService_RevokeSessionsForAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeSessionsForAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RevokeServiceServer).RevokeSessionsForAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alpacalabs.auth.v1.RevokeService/RevokeSessionsForAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RevokeServiceServer).RevokeSessionsForAccount(ctx, req.(*RevokeSessionsForAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RevokeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alpacalabs.auth.v1.RevokeService",
	HandlerType: (*RevokeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RevokeSession",
			Handler:    _RevokeService_RevokeSession_Handler,
		},
		{
			MethodName: "RevokeSessionsForAccount",
			Handler:    _RevokeService_RevokeSessionsForAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alpacalabs/auth/v1/revoke_session.proto",
}
