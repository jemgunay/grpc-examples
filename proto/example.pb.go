// Code generated by protoc-gen-go. DO NOT EDIT.
// source: example.proto

package example

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

type Request struct {
	Topic                string   `protobuf:"bytes,1,opt,name=topic,proto3" json:"topic,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

type Response struct {
	Data                 string   `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_15a1dc8d40dadaa6, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "example.Request")
	proto.RegisterType((*Response)(nil), "example.Response")
}

func init() { proto.RegisterFile("example.proto", fileDescriptor_15a1dc8d40dadaa6) }

var fileDescriptor_15a1dc8d40dadaa6 = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x87, 0x72, 0x95, 0xe4, 0xb9,
	0xd8, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x44, 0xb8, 0x58, 0x4b, 0xf2, 0x0b, 0x32,
	0x93, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x20, 0x1c, 0x25, 0x39, 0x2e, 0x8e, 0xa0, 0xd4,
	0xe2, 0x82, 0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x21, 0x2e, 0x96, 0x94, 0xc4, 0x92, 0x44, 0xa8, 0x02,
	0x30, 0xdb, 0xa8, 0x8d, 0x89, 0x8b, 0x33, 0xa0, 0x34, 0x29, 0x27, 0xb3, 0x38, 0x23, 0xb5, 0x48,
	0xc8, 0x98, 0x8b, 0x27, 0x34, 0x2f, 0xb1, 0xa8, 0xd2, 0x15, 0x62, 0xbc, 0x90, 0x80, 0x1e, 0xcc,
	0x5e, 0xa8, 0x2d, 0x52, 0x82, 0x48, 0x22, 0x50, 0x63, 0xed, 0xb8, 0x44, 0x61, 0xec, 0xe0, 0x92,
	0xa2, 0xd4, 0xc4, 0x5c, 0x52, 0x74, 0x1b, 0x30, 0x0a, 0xd9, 0x72, 0x89, 0x40, 0xe5, 0x49, 0xd7,
	0xae, 0xc1, 0x28, 0xe4, 0xca, 0x25, 0xe5, 0x94, 0x99, 0x92, 0x59, 0x94, 0x9a, 0x5c, 0x92, 0x99,
	0x9f, 0x97, 0x98, 0x43, 0x8e, 0x21, 0x06, 0x8c, 0x49, 0x6c, 0xe0, 0x90, 0x35, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0x04, 0xca, 0x82, 0x4b, 0x6a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// PublisherClient is the client API for Publisher service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type PublisherClient interface {
	// One request, one response.
	UnaryExample(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	// One request, multiple responses.
	ResponseStreamExample(ctx context.Context, in *Request, opts ...grpc.CallOption) (Publisher_ResponseStreamExampleClient, error)
	// Multiple requests, one response.
	RequestStreamExample(ctx context.Context, opts ...grpc.CallOption) (Publisher_RequestStreamExampleClient, error)
	// Multiple requests, multiple responses.
	BidirectionalStreamExample(ctx context.Context, opts ...grpc.CallOption) (Publisher_BidirectionalStreamExampleClient, error)
}

type publisherClient struct {
	cc *grpc.ClientConn
}

func NewPublisherClient(cc *grpc.ClientConn) PublisherClient {
	return &publisherClient{cc}
}

func (c *publisherClient) UnaryExample(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/example.Publisher/UnaryExample", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publisherClient) ResponseStreamExample(ctx context.Context, in *Request, opts ...grpc.CallOption) (Publisher_ResponseStreamExampleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Publisher_serviceDesc.Streams[0], "/example.Publisher/ResponseStreamExample", opts...)
	if err != nil {
		return nil, err
	}
	x := &publisherResponseStreamExampleClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Publisher_ResponseStreamExampleClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type publisherResponseStreamExampleClient struct {
	grpc.ClientStream
}

func (x *publisherResponseStreamExampleClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *publisherClient) RequestStreamExample(ctx context.Context, opts ...grpc.CallOption) (Publisher_RequestStreamExampleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Publisher_serviceDesc.Streams[1], "/example.Publisher/RequestStreamExample", opts...)
	if err != nil {
		return nil, err
	}
	x := &publisherRequestStreamExampleClient{stream}
	return x, nil
}

type Publisher_RequestStreamExampleClient interface {
	Send(*Request) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type publisherRequestStreamExampleClient struct {
	grpc.ClientStream
}

func (x *publisherRequestStreamExampleClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *publisherRequestStreamExampleClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *publisherClient) BidirectionalStreamExample(ctx context.Context, opts ...grpc.CallOption) (Publisher_BidirectionalStreamExampleClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Publisher_serviceDesc.Streams[2], "/example.Publisher/BidirectionalStreamExample", opts...)
	if err != nil {
		return nil, err
	}
	x := &publisherBidirectionalStreamExampleClient{stream}
	return x, nil
}

type Publisher_BidirectionalStreamExampleClient interface {
	Send(*Request) error
	Recv() (*Response, error)
	grpc.ClientStream
}

type publisherBidirectionalStreamExampleClient struct {
	grpc.ClientStream
}

func (x *publisherBidirectionalStreamExampleClient) Send(m *Request) error {
	return x.ClientStream.SendMsg(m)
}

func (x *publisherBidirectionalStreamExampleClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// PublisherServer is the server API for Publisher service.
type PublisherServer interface {
	// One request, one response.
	UnaryExample(context.Context, *Request) (*Response, error)
	// One request, multiple responses.
	ResponseStreamExample(*Request, Publisher_ResponseStreamExampleServer) error
	// Multiple requests, one response.
	RequestStreamExample(Publisher_RequestStreamExampleServer) error
	// Multiple requests, multiple responses.
	BidirectionalStreamExample(Publisher_BidirectionalStreamExampleServer) error
}

// UnimplementedPublisherServer can be embedded to have forward compatible implementations.
type UnimplementedPublisherServer struct {
}

func (*UnimplementedPublisherServer) UnaryExample(ctx context.Context, req *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryExample not implemented")
}
func (*UnimplementedPublisherServer) ResponseStreamExample(req *Request, srv Publisher_ResponseStreamExampleServer) error {
	return status.Errorf(codes.Unimplemented, "method ResponseStreamExample not implemented")
}
func (*UnimplementedPublisherServer) RequestStreamExample(srv Publisher_RequestStreamExampleServer) error {
	return status.Errorf(codes.Unimplemented, "method RequestStreamExample not implemented")
}
func (*UnimplementedPublisherServer) BidirectionalStreamExample(srv Publisher_BidirectionalStreamExampleServer) error {
	return status.Errorf(codes.Unimplemented, "method BidirectionalStreamExample not implemented")
}

func RegisterPublisherServer(s *grpc.Server, srv PublisherServer) {
	s.RegisterService(&_Publisher_serviceDesc, srv)
}

func _Publisher_UnaryExample_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PublisherServer).UnaryExample(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/example.Publisher/UnaryExample",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PublisherServer).UnaryExample(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Publisher_ResponseStreamExample_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Request)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(PublisherServer).ResponseStreamExample(m, &publisherResponseStreamExampleServer{stream})
}

type Publisher_ResponseStreamExampleServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type publisherResponseStreamExampleServer struct {
	grpc.ServerStream
}

func (x *publisherResponseStreamExampleServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _Publisher_RequestStreamExample_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PublisherServer).RequestStreamExample(&publisherRequestStreamExampleServer{stream})
}

type Publisher_RequestStreamExampleServer interface {
	SendAndClose(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type publisherRequestStreamExampleServer struct {
	grpc.ServerStream
}

func (x *publisherRequestStreamExampleServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *publisherRequestStreamExampleServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Publisher_BidirectionalStreamExample_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PublisherServer).BidirectionalStreamExample(&publisherBidirectionalStreamExampleServer{stream})
}

type Publisher_BidirectionalStreamExampleServer interface {
	Send(*Response) error
	Recv() (*Request, error)
	grpc.ServerStream
}

type publisherBidirectionalStreamExampleServer struct {
	grpc.ServerStream
}

func (x *publisherBidirectionalStreamExampleServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *publisherBidirectionalStreamExampleServer) Recv() (*Request, error) {
	m := new(Request)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _Publisher_serviceDesc = grpc.ServiceDesc{
	ServiceName: "example.Publisher",
	HandlerType: (*PublisherServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryExample",
			Handler:    _Publisher_UnaryExample_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ResponseStreamExample",
			Handler:       _Publisher_ResponseStreamExample_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "RequestStreamExample",
			Handler:       _Publisher_RequestStreamExample_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BidirectionalStreamExample",
			Handler:       _Publisher_BidirectionalStreamExample_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "example.proto",
}
