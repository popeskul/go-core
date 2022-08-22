// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.4
// source: messages_proto/messages.proto

package messages_proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MessengerClient is the client API for Messenger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MessengerClient interface {
	Messages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Messenger_MessagesClient, error)
	Send(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error)
}

type messengerClient struct {
	cc grpc.ClientConnInterface
}

func NewMessengerClient(cc grpc.ClientConnInterface) MessengerClient {
	return &messengerClient{cc}
}

func (c *messengerClient) Messages(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Messenger_MessagesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Messenger_ServiceDesc.Streams[0], "/protobuf.Messenger/Messages", opts...)
	if err != nil {
		return nil, err
	}
	x := &messengerMessagesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Messenger_MessagesClient interface {
	Recv() (*Message, error)
	grpc.ClientStream
}

type messengerMessagesClient struct {
	grpc.ClientStream
}

func (x *messengerMessagesClient) Recv() (*Message, error) {
	m := new(Message)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *messengerClient) Send(ctx context.Context, in *Message, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/protobuf.Messenger/Send", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MessengerServer is the server API for Messenger service.
// All implementations must embed UnimplementedMessengerServer
// for forward compatibility
type MessengerServer interface {
	Messages(*Empty, Messenger_MessagesServer) error
	Send(context.Context, *Message) (*Message, error)
	mustEmbedUnimplementedMessengerServer()
}

// UnimplementedMessengerServer must be embedded to have forward compatible implementations.
type UnimplementedMessengerServer struct {
}

func (UnimplementedMessengerServer) Messages(*Empty, Messenger_MessagesServer) error {
	return status.Errorf(codes.Unimplemented, "method Messages not implemented")
}
func (UnimplementedMessengerServer) Send(context.Context, *Message) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Send not implemented")
}
func (UnimplementedMessengerServer) mustEmbedUnimplementedMessengerServer() {}

// UnsafeMessengerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MessengerServer will
// result in compilation errors.
type UnsafeMessengerServer interface {
	mustEmbedUnimplementedMessengerServer()
}

func RegisterMessengerServer(s grpc.ServiceRegistrar, srv MessengerServer) {
	s.RegisterService(&Messenger_ServiceDesc, srv)
}

func _Messenger_Messages_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MessengerServer).Messages(m, &messengerMessagesServer{stream})
}

type Messenger_MessagesServer interface {
	Send(*Message) error
	grpc.ServerStream
}

type messengerMessagesServer struct {
	grpc.ServerStream
}

func (x *messengerMessagesServer) Send(m *Message) error {
	return x.ServerStream.SendMsg(m)
}

func _Messenger_Send_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Message)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessengerServer).Send(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protobuf.Messenger/Send",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessengerServer).Send(ctx, req.(*Message))
	}
	return interceptor(ctx, in, info, handler)
}

// Messenger_ServiceDesc is the grpc.ServiceDesc for Messenger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Messenger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protobuf.Messenger",
	HandlerType: (*MessengerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Send",
			Handler:    _Messenger_Send_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Messages",
			Handler:       _Messenger_Messages_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "messages_proto/messages.proto",
}
