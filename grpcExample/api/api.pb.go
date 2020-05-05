// Code generated by protoc-gen-go. DO NOT EDIT.
// source: api.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN         HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING         HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING     HealthCheckResponse_ServingStatus = 2
	HealthCheckResponse_SERVICE_UNKNOWN HealthCheckResponse_ServingStatus = 3
)

var HealthCheckResponse_ServingStatus_name = map[int32]string{
	0: "UNKNOWN",
	1: "SERVING",
	2: "NOT_SERVING",
	3: "SERVICE_UNKNOWN",
}

var HealthCheckResponse_ServingStatus_value = map[string]int32{
	"UNKNOWN":         0,
	"SERVING":         1,
	"NOT_SERVING":     2,
	"SERVICE_UNKNOWN": 3,
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return proto.EnumName(HealthCheckResponse_ServingStatus_name, int32(x))
}

func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3, 0}
}

// The serialized message that client sends.
// Each field in the message definition has a unique number used to identify fields in the message binary format, and should not be changed once your message type is in use.
// Numbers in the range 1 through 15 take one byte to encode.
type InputRequest struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	ClientName           string   `protobuf:"bytes,2,opt,name=clientName,proto3" json:"clientName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InputRequest) Reset()         { *m = InputRequest{} }
func (m *InputRequest) String() string { return proto.CompactTextString(m) }
func (*InputRequest) ProtoMessage()    {}
func (*InputRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{0}
}

func (m *InputRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InputRequest.Unmarshal(m, b)
}
func (m *InputRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InputRequest.Marshal(b, m, deterministic)
}
func (m *InputRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InputRequest.Merge(m, src)
}
func (m *InputRequest) XXX_Size() int {
	return xxx_messageInfo_InputRequest.Size(m)
}
func (m *InputRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_InputRequest.DiscardUnknown(m)
}

var xxx_messageInfo_InputRequest proto.InternalMessageInfo

func (m *InputRequest) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *InputRequest) GetClientName() string {
	if m != nil {
		return m.ClientName
	}
	return ""
}

// What server responds to as a result of getting InputRequest.
type OutputResponse struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	ServerName           string   `protobuf:"bytes,2,opt,name=serverName,proto3" json:"serverName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutputResponse) Reset()         { *m = OutputResponse{} }
func (m *OutputResponse) String() string { return proto.CompactTextString(m) }
func (*OutputResponse) ProtoMessage()    {}
func (*OutputResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{1}
}

func (m *OutputResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputResponse.Unmarshal(m, b)
}
func (m *OutputResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputResponse.Marshal(b, m, deterministic)
}
func (m *OutputResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputResponse.Merge(m, src)
}
func (m *OutputResponse) XXX_Size() int {
	return xxx_messageInfo_OutputResponse.Size(m)
}
func (m *OutputResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OutputResponse proto.InternalMessageInfo

func (m *OutputResponse) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *OutputResponse) GetServerName() string {
	if m != nil {
		return m.ServerName
	}
	return ""
}

// Health check request to find out the readiness/liveness.
type HealthCheckRequest struct {
	Service              string   `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheckRequest) Reset()         { *m = HealthCheckRequest{} }
func (m *HealthCheckRequest) String() string { return proto.CompactTextString(m) }
func (*HealthCheckRequest) ProtoMessage()    {}
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{2}
}

func (m *HealthCheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckRequest.Unmarshal(m, b)
}
func (m *HealthCheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckRequest.Marshal(b, m, deterministic)
}
func (m *HealthCheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckRequest.Merge(m, src)
}
func (m *HealthCheckRequest) XXX_Size() int {
	return xxx_messageInfo_HealthCheckRequest.Size(m)
}
func (m *HealthCheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckRequest proto.InternalMessageInfo

func (m *HealthCheckRequest) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

// Health check response.
type HealthCheckResponse struct {
	Status               HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=api.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *HealthCheckResponse) Reset()         { *m = HealthCheckResponse{} }
func (m *HealthCheckResponse) String() string { return proto.CompactTextString(m) }
func (*HealthCheckResponse) ProtoMessage()    {}
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00212fb1f9d3bf1c, []int{3}
}

func (m *HealthCheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheckResponse.Unmarshal(m, b)
}
func (m *HealthCheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheckResponse.Marshal(b, m, deterministic)
}
func (m *HealthCheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheckResponse.Merge(m, src)
}
func (m *HealthCheckResponse) XXX_Size() int {
	return xxx_messageInfo_HealthCheckResponse.Size(m)
}
func (m *HealthCheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheckResponse proto.InternalMessageInfo

func (m *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if m != nil {
		return m.Status
	}
	return HealthCheckResponse_UNKNOWN
}

func init() {
	proto.RegisterEnum("api.HealthCheckResponse_ServingStatus", HealthCheckResponse_ServingStatus_name, HealthCheckResponse_ServingStatus_value)
	proto.RegisterType((*InputRequest)(nil), "api.InputRequest")
	proto.RegisterType((*OutputResponse)(nil), "api.OutputResponse")
	proto.RegisterType((*HealthCheckRequest)(nil), "api.HealthCheckRequest")
	proto.RegisterType((*HealthCheckResponse)(nil), "api.HealthCheckResponse")
}

func init() {
	proto.RegisterFile("api.proto", fileDescriptor_00212fb1f9d3bf1c)
}

var fileDescriptor_00212fb1f9d3bf1c = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0x41, 0x4b, 0xfb, 0x40,
	0x10, 0xc5, 0x9b, 0xf6, 0xdf, 0x96, 0x4e, 0xff, 0xb6, 0x75, 0x73, 0x30, 0x78, 0x10, 0xd9, 0x83,
	0x78, 0x0a, 0x18, 0x6f, 0x1e, 0x44, 0xac, 0x45, 0x8b, 0xb0, 0x91, 0xa4, 0xda, 0x63, 0x59, 0xc3,
	0x60, 0x82, 0x35, 0x59, 0xb3, 0x1b, 0xf1, 0xe8, 0xa7, 0xf1, 0x73, 0xca, 0x6e, 0x13, 0x4c, 0xd0,
	0x5e, 0xbc, 0x65, 0xde, 0xcc, 0xfb, 0x65, 0xde, 0xb0, 0x30, 0xe0, 0x22, 0x71, 0x45, 0x9e, 0xa9,
	0x8c, 0x74, 0xb8, 0x48, 0xe8, 0x25, 0xfc, 0x9f, 0xa7, 0xa2, 0x50, 0x01, 0xbe, 0x16, 0x28, 0x15,
	0x21, 0xf0, 0x4f, 0xe1, 0xbb, 0x72, 0xac, 0x43, 0xeb, 0x78, 0x10, 0x98, 0x6f, 0x72, 0x00, 0x10,
	0xad, 0x13, 0x4c, 0x15, 0xe3, 0x2f, 0xe8, 0xb4, 0x4d, 0xa7, 0xa6, 0xd0, 0x2b, 0x18, 0xf9, 0x85,
	0x32, 0x10, 0x29, 0xb2, 0x54, 0xe2, 0x36, 0x8a, 0xc4, 0xfc, 0x0d, 0xf3, 0x3a, 0xe5, 0x5b, 0xa1,
	0x2e, 0x90, 0x1b, 0xe4, 0x6b, 0x15, 0x4f, 0x63, 0x8c, 0x9e, 0xab, 0x7d, 0x1c, 0xe8, 0xeb, 0x99,
	0x24, 0xc2, 0x12, 0x56, 0x95, 0xf4, 0xd3, 0x02, 0xbb, 0x61, 0x28, 0xff, 0x7d, 0x0e, 0x3d, 0xa9,
	0xb8, 0x2a, 0xa4, 0x31, 0x8c, 0xbc, 0x23, 0x57, 0x47, 0xfe, 0x65, 0xd2, 0x0d, 0x35, 0x29, 0x7d,
	0x0a, 0xcd, 0x74, 0x50, 0xba, 0xa8, 0x0f, 0x3b, 0x8d, 0x06, 0x19, 0x42, 0xff, 0x9e, 0xdd, 0x32,
	0x7f, 0xc9, 0x26, 0x2d, 0x5d, 0x84, 0xb3, 0xe0, 0x61, 0xce, 0xae, 0x27, 0x16, 0x19, 0xc3, 0x90,
	0xf9, 0x8b, 0x55, 0x25, 0xb4, 0x89, 0x0d, 0x63, 0x53, 0x4c, 0x67, 0xab, 0xca, 0xd2, 0xf1, 0x2e,
	0x60, 0x78, 0x97, 0x67, 0x11, 0x4a, 0xb9, 0xd0, 0x77, 0x38, 0x81, 0x6e, 0x21, 0x04, 0xe6, 0x64,
	0xd7, 0x2c, 0x56, 0xbf, 0xfe, 0xbe, 0x6d, 0xa4, 0xe6, 0x31, 0x69, 0xcb, 0xfb, 0xb0, 0xa0, 0xb7,
	0x09, 0x40, 0xce, 0xa0, 0x6b, 0x42, 0x90, 0xbd, 0x9f, 0xb1, 0x36, 0x0c, 0x67, 0x5b, 0x5e, 0xed,
	0x5d, 0x72, 0x15, 0xc5, 0x7f, 0xf0, 0x3e, 0xf6, 0xcc, 0x9b, 0x39, 0xfd, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x69, 0x34, 0x7a, 0xbc, 0x40, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProcessTextClient is the client API for ProcessText service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessTextClient interface {
	Upper(ctx context.Context, in *InputRequest, opts ...grpc.CallOption) (*OutputResponse, error)
}

type processTextClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessTextClient(cc grpc.ClientConnInterface) ProcessTextClient {
	return &processTextClient{cc}
}

func (c *processTextClient) Upper(ctx context.Context, in *InputRequest, opts ...grpc.CallOption) (*OutputResponse, error) {
	out := new(OutputResponse)
	err := c.cc.Invoke(ctx, "/api.ProcessText/upper", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProcessTextServer is the server API for ProcessText service.
type ProcessTextServer interface {
	Upper(context.Context, *InputRequest) (*OutputResponse, error)
}

// UnimplementedProcessTextServer can be embedded to have forward compatible implementations.
type UnimplementedProcessTextServer struct {
}

func (*UnimplementedProcessTextServer) Upper(ctx context.Context, req *InputRequest) (*OutputResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Upper not implemented")
}

func RegisterProcessTextServer(s *grpc.Server, srv ProcessTextServer) {
	s.RegisterService(&_ProcessText_serviceDesc, srv)
}

func _ProcessText_Upper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProcessTextServer).Upper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.ProcessText/Upper",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProcessTextServer).Upper(ctx, req.(*InputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProcessText_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ProcessText",
	HandlerType: (*ProcessTextServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "upper",
			Handler:    _ProcessText_Upper_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}

// HealthClient is the client API for Health service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type HealthClient interface {
	Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	Watch(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type healthClient struct {
	cc grpc.ClientConnInterface
}

func NewHealthClient(cc grpc.ClientConnInterface) HealthClient {
	return &healthClient{cc}
}

func (c *healthClient) Check(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/api.Health/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *healthClient) Watch(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/api.Health/Watch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HealthServer is the server API for Health service.
type HealthServer interface {
	Check(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	Watch(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

// UnimplementedHealthServer can be embedded to have forward compatible implementations.
type UnimplementedHealthServer struct {
}

func (*UnimplementedHealthServer) Check(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Check not implemented")
}
func (*UnimplementedHealthServer) Watch(ctx context.Context, req *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

func RegisterHealthServer(s *grpc.Server, srv HealthServer) {
	s.RegisterService(&_Health_serviceDesc, srv)
}

func _Health_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Health/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Check(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Health_Watch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HealthServer).Watch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Health/Watch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HealthServer).Watch(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Health_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Health",
	HandlerType: (*HealthServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Health_Check_Handler,
		},
		{
			MethodName: "Watch",
			Handler:    _Health_Watch_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}