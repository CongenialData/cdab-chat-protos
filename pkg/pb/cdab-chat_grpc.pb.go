// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.4
// source: cdab-chat.proto

package pb

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

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChatServiceClient interface {
	ClientConnect(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (ChatService_ClientConnectClient, error)
	PostMessage(ctx context.Context, in *PostMessageRequest, opts ...grpc.CallOption) (*PostMessageResponse, error)
	Quiz(ctx context.Context, opts ...grpc.CallOption) (ChatService_QuizClient, error)
}

type chatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChatServiceClient(cc grpc.ClientConnInterface) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) ClientConnect(ctx context.Context, in *ConnectionRequest, opts ...grpc.CallOption) (ChatService_ClientConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[0], "/cdab.chat.ChatService/ClientConnect", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceClientConnectClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ChatService_ClientConnectClient interface {
	Recv() (*ConnectionResponse, error)
	grpc.ClientStream
}

type chatServiceClientConnectClient struct {
	grpc.ClientStream
}

func (x *chatServiceClientConnectClient) Recv() (*ConnectionResponse, error) {
	m := new(ConnectionResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chatServiceClient) PostMessage(ctx context.Context, in *PostMessageRequest, opts ...grpc.CallOption) (*PostMessageResponse, error) {
	out := new(PostMessageResponse)
	err := c.cc.Invoke(ctx, "/cdab.chat.ChatService/PostMessage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chatServiceClient) Quiz(ctx context.Context, opts ...grpc.CallOption) (ChatService_QuizClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChatService_ServiceDesc.Streams[1], "/cdab.chat.ChatService/Quiz", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceQuizClient{stream}
	return x, nil
}

type ChatService_QuizClient interface {
	Send(*QuizRequest) error
	Recv() (*QuizResponse, error)
	grpc.ClientStream
}

type chatServiceQuizClient struct {
	grpc.ClientStream
}

func (x *chatServiceQuizClient) Send(m *QuizRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceQuizClient) Recv() (*QuizResponse, error) {
	m := new(QuizResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
// All implementations must embed UnimplementedChatServiceServer
// for forward compatibility
type ChatServiceServer interface {
	ClientConnect(*ConnectionRequest, ChatService_ClientConnectServer) error
	PostMessage(context.Context, *PostMessageRequest) (*PostMessageResponse, error)
	Quiz(ChatService_QuizServer) error
	mustEmbedUnimplementedChatServiceServer()
}

// UnimplementedChatServiceServer must be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (UnimplementedChatServiceServer) ClientConnect(*ConnectionRequest, ChatService_ClientConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientConnect not implemented")
}
func (UnimplementedChatServiceServer) PostMessage(context.Context, *PostMessageRequest) (*PostMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PostMessage not implemented")
}
func (UnimplementedChatServiceServer) Quiz(ChatService_QuizServer) error {
	return status.Errorf(codes.Unimplemented, "method Quiz not implemented")
}
func (UnimplementedChatServiceServer) mustEmbedUnimplementedChatServiceServer() {}

// UnsafeChatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChatServiceServer will
// result in compilation errors.
type UnsafeChatServiceServer interface {
	mustEmbedUnimplementedChatServiceServer()
}

func RegisterChatServiceServer(s grpc.ServiceRegistrar, srv ChatServiceServer) {
	s.RegisterService(&ChatService_ServiceDesc, srv)
}

func _ChatService_ClientConnect_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ConnectionRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ChatServiceServer).ClientConnect(m, &chatServiceClientConnectServer{stream})
}

type ChatService_ClientConnectServer interface {
	Send(*ConnectionResponse) error
	grpc.ServerStream
}

type chatServiceClientConnectServer struct {
	grpc.ServerStream
}

func (x *chatServiceClientConnectServer) Send(m *ConnectionResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ChatService_PostMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChatServiceServer).PostMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cdab.chat.ChatService/PostMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChatServiceServer).PostMessage(ctx, req.(*PostMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChatService_Quiz_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).Quiz(&chatServiceQuizServer{stream})
}

type ChatService_QuizServer interface {
	Send(*QuizResponse) error
	Recv() (*QuizRequest, error)
	grpc.ServerStream
}

type chatServiceQuizServer struct {
	grpc.ServerStream
}

func (x *chatServiceQuizServer) Send(m *QuizResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceQuizServer) Recv() (*QuizRequest, error) {
	m := new(QuizRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatService_ServiceDesc is the grpc.ServiceDesc for ChatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cdab.chat.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PostMessage",
			Handler:    _ChatService_PostMessage_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientConnect",
			Handler:       _ChatService_ClientConnect_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Quiz",
			Handler:       _ChatService_Quiz_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "cdab-chat.proto",
}