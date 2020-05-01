// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alpacalabs/auth/v1/internal_get_token.proto

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

// internal-only request to create a token on behalf of a newly signed up user
type GetTokenRequest struct {
	// account ID
	AccountId string `protobuf:"bytes,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	// email
	EmailAddress string `protobuf:"bytes,2,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	// device type
	DeviceType string `protobuf:"bytes,3,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	// device id
	DeviceId             string   `protobuf:"bytes,4,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenRequest) Reset()         { *m = GetTokenRequest{} }
func (m *GetTokenRequest) String() string { return proto.CompactTextString(m) }
func (*GetTokenRequest) ProtoMessage()    {}
func (*GetTokenRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0344758a20033b17, []int{0}
}

func (m *GetTokenRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenRequest.Unmarshal(m, b)
}
func (m *GetTokenRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenRequest.Marshal(b, m, deterministic)
}
func (m *GetTokenRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenRequest.Merge(m, src)
}
func (m *GetTokenRequest) XXX_Size() int {
	return xxx_messageInfo_GetTokenRequest.Size(m)
}
func (m *GetTokenRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenRequest proto.InternalMessageInfo

func (m *GetTokenRequest) GetAccountId() string {
	if m != nil {
		return m.AccountId
	}
	return ""
}

func (m *GetTokenRequest) GetEmailAddress() string {
	if m != nil {
		return m.EmailAddress
	}
	return ""
}

func (m *GetTokenRequest) GetDeviceType() string {
	if m != nil {
		return m.DeviceType
	}
	return ""
}

func (m *GetTokenRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

// internal-only response containing a token for use by a newly signed up user
type GetTokenResponse struct {
	// jwt
	Jwt                  string   `protobuf:"bytes,1,opt,name=jwt,proto3" json:"jwt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetTokenResponse) Reset()         { *m = GetTokenResponse{} }
func (m *GetTokenResponse) String() string { return proto.CompactTextString(m) }
func (*GetTokenResponse) ProtoMessage()    {}
func (*GetTokenResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0344758a20033b17, []int{1}
}

func (m *GetTokenResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetTokenResponse.Unmarshal(m, b)
}
func (m *GetTokenResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetTokenResponse.Marshal(b, m, deterministic)
}
func (m *GetTokenResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetTokenResponse.Merge(m, src)
}
func (m *GetTokenResponse) XXX_Size() int {
	return xxx_messageInfo_GetTokenResponse.Size(m)
}
func (m *GetTokenResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetTokenResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetTokenResponse proto.InternalMessageInfo

func (m *GetTokenResponse) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

func init() {
	proto.RegisterType((*GetTokenRequest)(nil), "alpacalabs.auth.v1.GetTokenRequest")
	proto.RegisterType((*GetTokenResponse)(nil), "alpacalabs.auth.v1.GetTokenResponse")
}

func init() {
	proto.RegisterFile("alpacalabs/auth/v1/internal_get_token.proto", fileDescriptor_0344758a20033b17)
}

var fileDescriptor_0344758a20033b17 = []byte{
	// 327 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0xc3, 0x30,
	0x14, 0xc7, 0xed, 0x26, 0xb2, 0x46, 0xc5, 0x91, 0x83, 0x14, 0x45, 0x94, 0x6d, 0x07, 0x41, 0x96,
	0x52, 0xbd, 0x79, 0xeb, 0x2e, 0x32, 0xd8, 0x61, 0xcc, 0x31, 0x41, 0x06, 0x25, 0x6b, 0x1e, 0x5b,
	0x67, 0xd7, 0xd4, 0xe6, 0x75, 0xb2, 0x8f, 0x21, 0x7e, 0x03, 0x8f, 0x7e, 0x14, 0x3f, 0x95, 0x24,
	0xad, 0x54, 0x9c, 0x78, 0x0b, 0xbf, 0xfc, 0xde, 0xcb, 0x7b, 0xf9, 0x93, 0x2b, 0x1e, 0xa7, 0x3c,
	0xe4, 0x31, 0x9f, 0x29, 0x97, 0xe7, 0xb8, 0x70, 0xd7, 0x9e, 0x1b, 0x25, 0x08, 0x59, 0xc2, 0xe3,
	0x60, 0x0e, 0x18, 0xa0, 0x7c, 0x82, 0x84, 0xa5, 0x99, 0x44, 0x49, 0x69, 0x25, 0x33, 0x2d, 0xb3,
	0xb5, 0xd7, 0x7a, 0xb3, 0xc8, 0xd1, 0x1d, 0xe0, 0x58, 0x6b, 0x23, 0x78, 0xce, 0x41, 0x21, 0x3d,
	0x23, 0x84, 0x87, 0xa1, 0xcc, 0x13, 0x0c, 0x22, 0xe1, 0x58, 0x17, 0xd6, 0xa5, 0x3d, 0xb2, 0x4b,
	0xd2, 0x17, 0xb4, 0x4d, 0x0e, 0x61, 0xc5, 0xa3, 0x38, 0xe0, 0x42, 0x64, 0xa0, 0x94, 0x53, 0x33,
	0xc6, 0x81, 0x81, 0x7e, 0xc1, 0xe8, 0x39, 0xd9, 0x17, 0xb0, 0x8e, 0x42, 0x08, 0x70, 0x93, 0x82,
	0x53, 0x37, 0x0a, 0x29, 0xd0, 0x78, 0x93, 0x02, 0x3d, 0x25, 0x76, 0x29, 0x44, 0xc2, 0xd9, 0x35,
	0xd7, 0x8d, 0x02, 0xf4, 0x45, 0xab, 0x43, 0x9a, 0xd5, 0x50, 0x2a, 0x95, 0x89, 0x02, 0xda, 0x24,
	0xf5, 0xe5, 0x0b, 0x96, 0xe3, 0xe8, 0xe3, 0xf5, 0xb2, 0x1a, 0xfd, 0x1e, 0x32, 0x5d, 0x4a, 0x1f,
	0x48, 0xe3, 0x1b, 0xd1, 0x36, 0xdb, 0xde, 0x97, 0xfd, 0xda, 0xf5, 0xa4, 0xf3, 0xbf, 0x54, 0xbc,
	0xdd, 0xda, 0xe9, 0xbd, 0x5a, 0xe4, 0x38, 0x94, 0xab, 0x3f, 0xec, 0x9e, 0xed, 0xe7, 0xb8, 0x18,
	0xea, 0x1f, 0x1e, 0x5a, 0x8f, 0xb7, 0xf3, 0x08, 0x17, 0xf9, 0x8c, 0x85, 0x72, 0xe5, 0xfa, 0xc6,
	0x1d, 0xe8, 0x6c, 0x4c, 0x00, 0x19, 0xa4, 0xb2, 0xab, 0xab, 0xba, 0x73, 0xe9, 0x6e, 0x07, 0xf7,
	0x5e, 0xab, 0xfb, 0x03, 0xff, 0xa3, 0x46, 0xab, 0x3a, 0xa6, 0x5b, 0xb3, 0x89, 0xf7, 0xf9, 0x13,
	0x4e, 0x35, 0x9c, 0x4e, 0xbc, 0xd9, 0x9e, 0xe9, 0x7a, 0xf3, 0x15, 0x00, 0x00, 0xff, 0xff, 0x07,
	0xeb, 0x4e, 0xcd, 0x05, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GetTokenServiceClient is the client API for GetTokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GetTokenServiceClient interface {
	// GetToken requests a token for the current user.
	// This endpoint issues a token without requiring a password.
	// That is only safe to do because it's meant to be internal-only.
	GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error)
}

type getTokenServiceClient struct {
	cc *grpc.ClientConn
}

func NewGetTokenServiceClient(cc *grpc.ClientConn) GetTokenServiceClient {
	return &getTokenServiceClient{cc}
}

func (c *getTokenServiceClient) GetToken(ctx context.Context, in *GetTokenRequest, opts ...grpc.CallOption) (*GetTokenResponse, error) {
	out := new(GetTokenResponse)
	err := c.cc.Invoke(ctx, "/alpacalabs.auth.v1.GetTokenService/GetToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetTokenServiceServer is the server API for GetTokenService service.
type GetTokenServiceServer interface {
	// GetToken requests a token for the current user.
	// This endpoint issues a token without requiring a password.
	// That is only safe to do because it's meant to be internal-only.
	GetToken(context.Context, *GetTokenRequest) (*GetTokenResponse, error)
}

func RegisterGetTokenServiceServer(s *grpc.Server, srv GetTokenServiceServer) {
	s.RegisterService(&_GetTokenService_serviceDesc, srv)
}

func _GetTokenService_GetToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GetTokenServiceServer).GetToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/alpacalabs.auth.v1.GetTokenService/GetToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GetTokenServiceServer).GetToken(ctx, req.(*GetTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GetTokenService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "alpacalabs.auth.v1.GetTokenService",
	HandlerType: (*GetTokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetToken",
			Handler:    _GetTokenService_GetToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "alpacalabs/auth/v1/internal_get_token.proto",
}
