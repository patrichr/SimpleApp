// Code generated by protoc-gen-go. DO NOT EDIT.
// source: contract.proto

package grpc

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

// The request message containing the user's name.
type RegRequest struct {
	TaskName             string   `protobuf:"bytes,1,opt,name=task_name,json=taskName,proto3" json:"task_name,omitempty"`
	TaskOwner            string   `protobuf:"bytes,2,opt,name=task_owner,json=taskOwner,proto3" json:"task_owner,omitempty"`
	Description          string   `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	Status               string   `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegRequest) Reset()         { *m = RegRequest{} }
func (m *RegRequest) String() string { return proto.CompactTextString(m) }
func (*RegRequest) ProtoMessage()    {}
func (*RegRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d19debeba7dea55a, []int{0}
}

func (m *RegRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegRequest.Unmarshal(m, b)
}
func (m *RegRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegRequest.Marshal(b, m, deterministic)
}
func (m *RegRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegRequest.Merge(m, src)
}
func (m *RegRequest) XXX_Size() int {
	return xxx_messageInfo_RegRequest.Size(m)
}
func (m *RegRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RegRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RegRequest proto.InternalMessageInfo

func (m *RegRequest) GetTaskName() string {
	if m != nil {
		return m.TaskName
	}
	return ""
}

func (m *RegRequest) GetTaskOwner() string {
	if m != nil {
		return m.TaskOwner
	}
	return ""
}

func (m *RegRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *RegRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

// The response message containing the greetings
type RegResponse struct {
	HttpStatus           int32    `protobuf:"varint,1,opt,name=http_status,json=httpStatus,proto3" json:"http_status,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegResponse) Reset()         { *m = RegResponse{} }
func (m *RegResponse) String() string { return proto.CompactTextString(m) }
func (*RegResponse) ProtoMessage()    {}
func (*RegResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_d19debeba7dea55a, []int{1}
}

func (m *RegResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegResponse.Unmarshal(m, b)
}
func (m *RegResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegResponse.Marshal(b, m, deterministic)
}
func (m *RegResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegResponse.Merge(m, src)
}
func (m *RegResponse) XXX_Size() int {
	return xxx_messageInfo_RegResponse.Size(m)
}
func (m *RegResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RegResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RegResponse proto.InternalMessageInfo

func (m *RegResponse) GetHttpStatus() int32 {
	if m != nil {
		return m.HttpStatus
	}
	return 0
}

func (m *RegResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*RegRequest)(nil), "grpc.RegRequest")
	proto.RegisterType((*RegResponse)(nil), "grpc.RegResponse")
}

func init() {
	proto.RegisterFile("contract.proto", fileDescriptor_d19debeba7dea55a)
}

var fileDescriptor_d19debeba7dea55a = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbd, 0x4b, 0x03, 0x41,
	0x10, 0xc5, 0x3d, 0x8d, 0xf9, 0x98, 0x80, 0x1f, 0x5b, 0xc8, 0x62, 0x10, 0xc3, 0x55, 0x56, 0x07,
	0x2a, 0x58, 0x59, 0xa5, 0x31, 0x95, 0x86, 0xb5, 0x10, 0x6c, 0xc2, 0x7a, 0x19, 0x2e, 0x87, 0x77,
	0xbb, 0xeb, 0xce, 0x84, 0x58, 0xfb, 0x97, 0xcb, 0x4e, 0x4e, 0x4c, 0xf7, 0xe6, 0xf7, 0x06, 0xe6,
	0xcd, 0x83, 0x93, 0xd2, 0x3b, 0x8e, 0xb6, 0xe4, 0x22, 0x44, 0xcf, 0x5e, 0xf5, 0xaa, 0x18, 0xca,
	0xfc, 0x27, 0x03, 0x30, 0x58, 0x19, 0xfc, 0xda, 0x20, 0xb1, 0x9a, 0xc0, 0x88, 0x2d, 0x7d, 0x2e,
	0x9d, 0x6d, 0x51, 0x67, 0xd3, 0xec, 0x66, 0x64, 0x86, 0x09, 0x3c, 0xdb, 0x16, 0xd5, 0x15, 0x80,
	0x98, 0x7e, 0xeb, 0x30, 0xea, 0x43, 0x71, 0x65, 0xfd, 0x25, 0x01, 0x35, 0x85, 0xf1, 0x0a, 0xa9,
	0x8c, 0x75, 0xe0, 0xda, 0x3b, 0x7d, 0x24, 0xfe, 0x3e, 0x52, 0x17, 0xd0, 0x27, 0xb6, 0xbc, 0x21,
	0xdd, 0x13, 0xb3, 0x9b, 0xf2, 0x39, 0x8c, 0x25, 0x03, 0x05, 0xef, 0x08, 0xd5, 0x35, 0x8c, 0xd7,
	0xcc, 0x61, 0xd9, 0xed, 0xa6, 0x18, 0xc7, 0x06, 0x12, 0x7a, 0x15, 0xa2, 0x34, 0x0c, 0x5a, 0x24,
	0xb2, 0x15, 0x76, 0x29, 0xfe, 0xc6, 0xbb, 0x47, 0x18, 0x3c, 0x45, 0x44, 0xc6, 0xa8, 0x6e, 0x61,
	0x68, 0xb0, 0xaa, 0x29, 0xe9, 0xb3, 0x22, 0x3d, 0x5b, 0xfc, 0x3f, 0x7a, 0x79, 0xbe, 0x47, 0x76,
	0x67, 0xf3, 0x83, 0xd9, 0x03, 0x4c, 0x6a, 0xbf, 0x33, 0xf0, 0xdb, 0xb6, 0xa1, 0x41, 0x2a, 0xd6,
	0xd8, 0x34, 0x7e, 0xeb, 0x63, 0xb3, 0x9a, 0x9d, 0xce, 0x93, 0x7e, 0x4b, 0x7a, 0x91, 0x2a, 0x5c,
	0x64, 0xef, 0x52, 0xe2, 0x47, 0x5f, 0x1a, 0xbd, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x03,
	0xad, 0x05, 0x63, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// GreeterClient is the client API for Greeter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreeterClient interface {
	Register(ctx context.Context, in *RegRequest, opts ...grpc.CallOption) (*RegResponse, error)
}

type greeterClient struct {
	cc grpc.ClientConnInterface
}

func NewGreeterClient(cc grpc.ClientConnInterface) GreeterClient {
	return &greeterClient{cc}
}

func (c *greeterClient) Register(ctx context.Context, in *RegRequest, opts ...grpc.CallOption) (*RegResponse, error) {
	out := new(RegResponse)
	err := c.cc.Invoke(ctx, "/grpc.Greeter/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreeterServer is the server API for Greeter service.
type GreeterServer interface {
	Register(context.Context, *RegRequest) (*RegResponse, error)
}

// UnimplementedGreeterServer can be embedded to have forward compatible implementations.
type UnimplementedGreeterServer struct {
}

func (*UnimplementedGreeterServer) Register(ctx context.Context, req *RegRequest) (*RegResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterGreeterServer(s *grpc.Server, srv GreeterServer) {
	s.RegisterService(&_Greeter_serviceDesc, srv)
}

func _Greeter_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreeterServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Greeter/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreeterServer).Register(ctx, req.(*RegRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Greeter_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Greeter",
	HandlerType: (*GreeterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _Greeter_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract.proto",
}
