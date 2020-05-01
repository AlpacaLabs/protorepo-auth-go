// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alpacalabs/auth/v1/login.proto

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

// Request for logging in.
type LoginRequest struct {
	// email
	EmailAddress string `protobuf:"bytes,1,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	// password
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	// device type
	DeviceType string `protobuf:"bytes,3,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	// device id
	DeviceId             string   `protobuf:"bytes,4,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginRequest) Reset()         { *m = LoginRequest{} }
func (m *LoginRequest) String() string { return proto.CompactTextString(m) }
func (*LoginRequest) ProtoMessage()    {}
func (*LoginRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_35de948650bf2aa6, []int{0}
}

func (m *LoginRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginRequest.Unmarshal(m, b)
}
func (m *LoginRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginRequest.Marshal(b, m, deterministic)
}
func (m *LoginRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginRequest.Merge(m, src)
}
func (m *LoginRequest) XXX_Size() int {
	return xxx_messageInfo_LoginRequest.Size(m)
}
func (m *LoginRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoginRequest proto.InternalMessageInfo

func (m *LoginRequest) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *LoginRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginRequest) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *LoginRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

// Response containing your new JWT.
type LoginResponse struct {
	// JWT (json web token)
	Jwt                  string   `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LoginResponse) Reset()         { *m = LoginResponse{} }
func (m *LoginResponse) String() string { return proto.CompactTextString(m) }
func (*LoginResponse) ProtoMessage()    {}
func (*LoginResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_35de948650bf2aa6, []int{1}
}

func (m *LoginResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginResponse.Unmarshal(m, b)
}
func (m *LoginResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginResponse.Marshal(b, m, deterministic)
}
func (m *LoginResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginResponse.Merge(m, src)
}
func (m *LoginResponse) XXX_Size() int {
	return xxx_messageInfo_LoginResponse.Size(m)
}
func (m *LoginResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoginResponse proto.InternalMessageInfo

func (m *LoginResponse) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

func init() {
	proto.RegisterType((*LoginRequest)(nil), "alpacalabs.auth.v1.LoginRequest")
	proto.RegisterType((*LoginResponse)(nil), "alpacalabs.auth.v1.LoginResponse")
}

func init() { proto.RegisterFile("alpacalabs/auth/v1/login.proto", fileDescriptor_35de948650bf2aa6) }

var fileDescriptor_35de948650bf2aa6 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0x4a, 0xfb, 0x40,
	0x10, 0xc6, 0xff, 0x69, 0xff, 0x8a, 0x59, 0x5b, 0x90, 0x3d, 0x48, 0xa8, 0xa0, 0x6d, 0xbc, 0x78,
	0xe9, 0x86, 0xe8, 0xcd, 0x5b, 0x7a, 0x13, 0x82, 0x94, 0x2a, 0x3d, 0x48, 0xb1, 0x6c, 0xb2, 0x43,
	0x12, 0x49, 0xba, 0x6b, 0x76, 0x93, 0xd2, 0x87, 0xf0, 0xe0, 0x2b, 0x78, 0xf4, 0x51, 0x7c, 0x2a,
	0xd9, 0x4d, 0xb0, 0x85, 0x8a, 0xb7, 0xcc, 0xef, 0xfb, 0xbe, 0xc9, 0xcc, 0x2c, 0x3a, 0xa7, 0xb9,
	0xa0, 0x31, 0xcd, 0x69, 0x24, 0x3d, 0x5a, 0xa9, 0xd4, 0xab, 0x7d, 0x2f, 0xe7, 0x49, 0xb6, 0x22,
	0xa2, 0xe4, 0x8a, 0x63, 0xbc, 0xd5, 0x89, 0xd6, 0x49, 0xed, 0xbb, 0x6f, 0x16, 0xea, 0x85, 0xda,
	0x33, 0x83, 0xd7, 0x0a, 0xa4, 0xc2, 0x97, 0xa8, 0x0f, 0x05, 0xcd, 0xf2, 0x25, 0x65, 0xac, 0x04,
	0x29, 0x1d, 0x6b, 0x68, 0x5d, 0xd9, 0xb3, 0x9e, 0x81, 0x41, 0xc3, 0xf0, 0x00, 0x1d, 0x09, 0x2a,
	0xe5, 0x9a, 0x97, 0xcc, 0xe9, 0x18, 0xfd, 0xa7, 0xc6, 0x17, 0xe8, 0x98, 0x41, 0x9d, 0xc5, 0xb0,
	0x54, 0x1b, 0x01, 0x4e, 0xd7, 0xc8, 0xa8, 0x41, 0x8f, 0x1b, 0x01, 0xf8, 0x0c, 0xd9, 0xad, 0x21,
	0x63, 0xce, 0xff, 0x26, 0xdd, 0x80, 0x3b, 0xe6, 0x8e, 0x50, 0xbf, 0x1d, 0x47, 0x0a, 0xbe, 0x92,
	0x80, 0x4f, 0x50, 0xf7, 0x65, 0xad, 0xda, 0x29, 0xf4, 0xe7, 0xf5, 0x73, 0x3b, 0xf1, 0x03, 0x94,
	0x3a, 0x84, 0xef, 0xd1, 0x81, 0xa9, 0xf1, 0x90, 0xec, 0x2f, 0x48, 0x76, 0x97, 0x1b, 0x8c, 0xfe,
	0x70, 0x34, 0xff, 0x73, 0xff, 0x4d, 0xde, 0x2d, 0x74, 0x1a, 0xf3, 0xe2, 0x17, 0xeb, 0xc4, 0x0e,
	0x2a, 0x95, 0x4e, 0xf5, 0x31, 0xa7, 0xd6, 0xd3, 0x6d, 0x92, 0xa9, 0xb4, 0x8a, 0x48, 0xcc, 0x0b,
	0x2f, 0x30, 0xde, 0x50, 0x5f, 0xde, 0xdc, 0xba, 0x04, 0xc1, 0xc7, 0x3a, 0x35, 0x4e, 0xb8, 0xb7,
	0xff, 0x2c, 0x1f, 0x9d, 0x6e, 0x10, 0x06, 0x9f, 0x1d, 0xbc, 0xcd, 0x11, 0xdd, 0x9a, 0xcc, 0xfd,
	0xaf, 0x5d, 0xb8, 0xd0, 0x70, 0x31, 0xf7, 0xa3, 0x43, 0xd3, 0xf5, 0xe6, 0x3b, 0x00, 0x00, 0xff,
	0xff, 0xc5, 0xda, 0x55, 0x58, 0xe3, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoginServiceClient interface {
	// Login.
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
}

type loginServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoginServiceClient(cc *grpc.ClientConn) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/alpacalabs.auth.v1.LoginService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LoginServiceServer is the server API for LoginService service.
type LoginServiceServer interface {
	// Login.
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
}

func RegisterLoginServiceServer(s *grpc.Server, srv LoginServiceServer) {
	s.RegisterService(&_LoginService_serviceDesc, srv)
}

func _LoginService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alpacalabs.auth.v1.LoginService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LoginService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alpacalabs.auth.v1.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _LoginService_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alpacalabs/auth/v1/login.proto",
}
