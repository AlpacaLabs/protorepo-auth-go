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
type RevokeSessionByIDRequest struct {
	// session id
	SessionId            string   `protobuf:"bytes,1,opt,name=session_id,json=sessionId,proto3" json:"session_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeSessionByIDRequest) Reset()         { *m = RevokeSessionByIDRequest{} }
func (m *RevokeSessionByIDRequest) String() string { return proto.CompactTextString(m) }
func (*RevokeSessionByIDRequest) ProtoMessage()    {}
func (*RevokeSessionByIDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d08be2f88de0f4b1, []int{0}
}

func (m *RevokeSessionByIDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeSessionByIDRequest.Unmarshal(m, b)
}
func (m *RevokeSessionByIDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeSessionByIDRequest.Marshal(b, m, deterministic)
}
func (m *RevokeSessionByIDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeSessionByIDRequest.Merge(m, src)
}
func (m *RevokeSessionByIDRequest) XXX_Size() int {
	return xxx_messageInfo_RevokeSessionByIDRequest.Size(m)
}
func (m *RevokeSessionByIDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeSessionByIDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeSessionByIDRequest proto.InternalMessageInfo

func (m *RevokeSessionByIDRequest) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

// Response from revoking a session.
type RevokeSessionByIDResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeSessionByIDResponse) Reset()         { *m = RevokeSessionByIDResponse{} }
func (m *RevokeSessionByIDResponse) String() string { return proto.CompactTextString(m) }
func (*RevokeSessionByIDResponse) ProtoMessage()    {}
func (*RevokeSessionByIDResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d08be2f88de0f4b1, []int{1}
}

func (m *RevokeSessionByIDResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeSessionByIDResponse.Unmarshal(m, b)
}
func (m *RevokeSessionByIDResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeSessionByIDResponse.Marshal(b, m, deterministic)
}
func (m *RevokeSessionByIDResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeSessionByIDResponse.Merge(m, src)
}
func (m *RevokeSessionByIDResponse) XXX_Size() int {
	return xxx_messageInfo_RevokeSessionByIDResponse.Size(m)
}
func (m *RevokeSessionByIDResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeSessionByIDResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeSessionByIDResponse proto.InternalMessageInfo

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
	proto.RegisterType((*RevokeSessionByIDRequest)(nil), "alpacalabs.auth.v1.RevokeSessionByIDRequest")
	proto.RegisterType((*RevokeSessionByIDResponse)(nil), "alpacalabs.auth.v1.RevokeSessionByIDResponse")
	proto.RegisterType((*RevokeSessionsForAccountRequest)(nil), "alpacalabs.auth.v1.RevokeSessionsForAccountRequest")
	proto.RegisterType((*RevokeSessionsForAccountResponse)(nil), "alpacalabs.auth.v1.RevokeSessionsForAccountResponse")
}

func init() {
	proto.RegisterFile("alpacalabs/auth/v1/revoke_session.proto", fileDescriptor_d08be2f88de0f4b1)
}

var fileDescriptor_d08be2f88de0f4b1 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4f, 0xcc, 0x29, 0x48,
	0x4c, 0x4e, 0xcc, 0x49, 0x4c, 0x2a, 0xd6, 0x4f, 0x2c, 0x2d, 0xc9, 0xd0, 0x2f, 0x33, 0xd4, 0x2f,
	0x4a, 0x2d, 0xcb, 0xcf, 0x4e, 0x8d, 0x2f, 0x4e, 0x2d, 0x2e, 0xce, 0xcc, 0xcf, 0xd3, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0x42, 0x28, 0xd4, 0x03, 0x29, 0xd4, 0x2b, 0x33, 0x54, 0xb2, 0xe4,
	0x92, 0x08, 0x02, 0xab, 0x0d, 0x86, 0x28, 0x75, 0xaa, 0xf4, 0x74, 0x09, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x11, 0x92, 0xe5, 0xe2, 0x82, 0x1a, 0x10, 0x9f, 0x99, 0x22, 0xc1, 0xa8, 0xc0, 0xa8,
	0xc1, 0x19, 0xc4, 0x09, 0x15, 0xf1, 0x4c, 0x51, 0x92, 0xe6, 0x92, 0xc4, 0xa2, 0xb5, 0xb8, 0x20,
	0x3f, 0xaf, 0x38, 0x55, 0xc9, 0x81, 0x4b, 0x1e, 0x45, 0xb2, 0xd8, 0x2d, 0xbf, 0xc8, 0x31, 0x39,
	0x39, 0xbf, 0x34, 0xaf, 0x04, 0xc9, 0xf8, 0x44, 0x88, 0x08, 0x92, 0xf1, 0x50, 0x11, 0xcf, 0x14,
	0x25, 0x25, 0x2e, 0x05, 0xdc, 0x26, 0x40, 0x6c, 0x31, 0xea, 0x65, 0xe2, 0xe2, 0x85, 0x29, 0x2a,
	0x2a, 0xcb, 0x4c, 0x4e, 0x15, 0x2a, 0xe2, 0x12, 0xc4, 0x70, 0x94, 0x90, 0x8e, 0x1e, 0xa6, 0xcf,
	0xf5, 0x70, 0x79, 0x5b, 0x4a, 0x97, 0x48, 0xd5, 0x50, 0x9f, 0x32, 0x08, 0xb5, 0x33, 0xa2, 0x05,
	0x22, 0x92, 0x53, 0x85, 0x8c, 0x09, 0x9a, 0x86, 0x19, 0x34, 0x52, 0x26, 0xa4, 0x69, 0x82, 0xb9,
	0xc4, 0x69, 0x22, 0x23, 0x97, 0x58, 0x72, 0x7e, 0x2e, 0x16, 0xdd, 0x4e, 0x9c, 0x8e, 0xa5, 0x25,
	0x19, 0x01, 0xa0, 0x74, 0x10, 0xc0, 0x18, 0x65, 0x95, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97,
	0x9c, 0x9f, 0xab, 0xef, 0x08, 0x56, 0xeb, 0x03, 0x4a, 0x3d, 0xe0, 0x64, 0x52, 0x94, 0x5a, 0x90,
	0xaf, 0x0b, 0xd2, 0xa5, 0x9b, 0x9e, 0xaf, 0x8f, 0x99, 0xb4, 0x16, 0x31, 0x31, 0x3b, 0xfa, 0x38,
	0xae, 0x62, 0x12, 0x42, 0xe8, 0xd3, 0x03, 0x19, 0xad, 0x17, 0x66, 0x78, 0x0a, 0x59, 0x30, 0x06,
	0x24, 0x18, 0x13, 0x66, 0x98, 0xc4, 0x06, 0x36, 0xd5, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x27,
	0xac, 0xb8, 0x36, 0xa7, 0x02, 0x00, 0x00,
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
	RevokeSessionByID(ctx context.Context, in *RevokeSessionByIDRequest, opts ...grpc.CallOption) (*RevokeSessionByIDResponse, error)
	// Revoke your sessions (except for the one you're currently using, of course).
	RevokeSessionsForAccount(ctx context.Context, in *RevokeSessionsForAccountRequest, opts ...grpc.CallOption) (*RevokeSessionsForAccountResponse, error)
}

type revokeServiceClient struct {
	cc *grpc.ClientConn
}

func NewRevokeServiceClient(cc *grpc.ClientConn) RevokeServiceClient {
	return &revokeServiceClient{cc}
}

func (c *revokeServiceClient) RevokeSessionByID(ctx context.Context, in *RevokeSessionByIDRequest, opts ...grpc.CallOption) (*RevokeSessionByIDResponse, error) {
	out := new(RevokeSessionByIDResponse)
	err := c.cc.Invoke(ctx, "/alpacalabs.auth.v1.RevokeService/RevokeSessionByID", in, out, opts...)
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
	RevokeSessionByID(context.Context, *RevokeSessionByIDRequest) (*RevokeSessionByIDResponse, error)
	// Revoke your sessions (except for the one you're currently using, of course).
	RevokeSessionsForAccount(context.Context, *RevokeSessionsForAccountRequest) (*RevokeSessionsForAccountResponse, error)
}

func RegisterRevokeServiceServer(s *grpc.Server, srv RevokeServiceServer) {
	s.RegisterService(&_RevokeService_serviceDesc, srv)
}

func _RevokeService_RevokeSessionByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeSessionByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RevokeServiceServer).RevokeSessionByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alpacalabs.auth.v1.RevokeService/RevokeSessionByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RevokeServiceServer).RevokeSessionByID(ctx, req.(*RevokeSessionByIDRequest))
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
			MethodName: "RevokeSessionByID",
			Handler:    _RevokeService_RevokeSessionByID_Handler,
		},
		{
			MethodName: "RevokeSessionsForAccount",
			Handler:    _RevokeService_RevokeSessionsForAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alpacalabs/auth/v1/revoke_session.proto",
}
