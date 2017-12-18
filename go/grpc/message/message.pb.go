// Code generated by protoc-gen-go. DO NOT EDIT.
// source: grpc/message/message.proto

/*
Package message is a generated protocol buffer package.

It is generated from these files:
	grpc/message/message.proto

It has these top-level messages:
	Request
	Response
*/
package message

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Request struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Request) Reset()                    { *m = Request{} }
func (m *Request) String() string            { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()               {}
func (*Request) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Request) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Response struct {
	Message string `protobuf:"bytes,1,opt,name=message" json:"message,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Response) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "message.Request")
	proto.RegisterType((*Response)(nil), "message.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Echo service

type EchoClient interface {
	Echo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	ClientStream(ctx context.Context, opts ...grpc.CallOption) (Echo_ClientStreamClient, error)
	ServerStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Echo_ServerStreamClient, error)
	BilateralStream(ctx context.Context, opts ...grpc.CallOption) (Echo_BilateralStreamClient, error)
}

type echoClient struct {
	cc *grpc.ClientConn
}

func NewEchoClient(cc *grpc.ClientConn) EchoClient {
	return &echoClient{cc}
}

func (c *echoClient) Echo(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/message.Echo/Echo", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *echoClient) ClientStream(ctx context.Context, opts ...grpc.CallOption) (Echo_ClientStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Echo_serviceDesc.Streams[0], c.cc, "/message.Echo/clientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoClientStreamClient{stream}
	return x, nil
}

type Echo_ClientStreamClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type echoClientStreamClient struct {
	grpc.ClientStream
}

func (x *echoClientStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoClientStreamClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoClient) ServerStream(ctx context.Context, in *Request, opts ...grpc.CallOption) (Echo_ServerStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Echo_serviceDesc.Streams[1], c.cc, "/message.Echo/serverStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Echo_ServerStreamClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type echoServerStreamClient struct {
	grpc.ClientStream
}

func (x *echoServerStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *echoClient) BilateralStream(ctx context.Context, opts ...grpc.CallOption) (Echo_BilateralStreamClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Echo_serviceDesc.Streams[2], c.cc, "/message.Echo/bilateralStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &echoBilateralStreamClient{stream}
	return x, nil
}

type Echo_BilateralStreamClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type echoBilateralStreamClient struct {
	grpc.ClientStream
}

func (x *echoBilateralStreamClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *echoBilateralStreamClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Echo service
type EchoServer interface {
	Echo(context.Context, *Request) (*Response, error)
	ClientStream(Echo_ClientStreamServer) error
	ServerStream(*Request, Echo_ServerStreamServer) error
	BilateralStream(Echo_BilateralStreamServer) error
}

func RegisterEchoServer(s *grpc.Server, srv EchoServer) {
	s.RegisterService(&_Echo_serviceDesc, srv)
}

func _Echo_Echo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EchoServer).Echo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/message.Echo/Echo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EchoServer).Echo(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Echo_ClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServer).ClientStream(&echoClientStreamServer{stream})
}

type Echo_ClientStreamServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type echoClientStreamServer struct {
	grpc.ServerStream
}

func (x *echoClientStreamServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoClientStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Echo_ServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EchoServer).ServerStream(m, &echoServerStreamServer{stream})
}

type Echo_ServerStreamServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type echoServerStreamServer struct {
	grpc.ServerStream
}

func (x *echoServerStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _Echo_BilateralStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(EchoServer).BilateralStream(&echoBilateralStreamServer{stream})
}

type Echo_BilateralStreamServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type echoBilateralStreamServer struct {
	grpc.ServerStream
}

func (x *echoBilateralStreamServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *echoBilateralStreamServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Echo_serviceDesc = grpc.ServiceDesc{
	ServiceName: "message.Echo",
	HandlerType: (*EchoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Echo",
			Handler:    _Echo_Echo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "clientStream",
			Handler:       _Echo_ClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "serverStream",
			Handler:       _Echo_ServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "bilateralStream",
			Handler:       _Echo_BilateralStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/message/message.proto",
}

func init() { proto.RegisterFile("grpc/message/message.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 170 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4a, 0x2f, 0x2a, 0x48,
	0xd6, 0xcf, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0x85, 0xd1, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9,
	0x42, 0xec, 0x50, 0xae, 0x92, 0x32, 0x17, 0x7b, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x89, 0x90,
	0x04, 0x17, 0x4c, 0x54, 0x82, 0x51, 0x81, 0x51, 0x83, 0x33, 0x08, 0xae, 0x48, 0x85, 0x8b, 0x23,
	0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0x15, 0xb7, 0x2a, 0xa3, 0xa7, 0x8c, 0x5c, 0x2c, 0xae,
	0xc9, 0x19, 0xf9, 0x42, 0xba, 0x50, 0x5a, 0x40, 0x0f, 0x66, 0x29, 0xd4, 0x0a, 0x29, 0x41, 0x24,
	0x11, 0x88, 0x79, 0x4a, 0x0c, 0x42, 0xe6, 0x5c, 0x3c, 0xc9, 0x39, 0x99, 0xa9, 0x79, 0x25, 0xc1,
	0x25, 0x45, 0xa9, 0x89, 0xb9, 0x44, 0x6a, 0xd3, 0x60, 0x04, 0x69, 0x2c, 0x4e, 0x2d, 0x2a, 0x4b,
	0x2d, 0x22, 0x49, 0xa3, 0x01, 0xa3, 0x90, 0x0d, 0x17, 0x7f, 0x52, 0x66, 0x4e, 0x62, 0x49, 0x6a,
	0x51, 0x62, 0x0e, 0x89, 0x96, 0x1a, 0x30, 0x26, 0xb1, 0x81, 0x83, 0xd0, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xc9, 0x79, 0x1d, 0x1d, 0x60, 0x01, 0x00, 0x00,
}
