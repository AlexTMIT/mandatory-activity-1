// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: grpc/lamport.proto

package lamport_service

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ChittychatService_GetMessage_FullMethodName          = "/chittychat.chittychatService/GetMessage"
	ChittychatService_ProcessJoinRequest_FullMethodName  = "/chittychat.chittychatService/ProcessJoinRequest"
	ChittychatService_ProcessLeaveRequest_FullMethodName = "/chittychat.chittychatService/ProcessLeaveRequest"
)

// ChittychatServiceClient is the client API for ChittychatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChittychatServiceClient interface {
	GetMessage(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error)
	ProcessJoinRequest(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error)
	ProcessLeaveRequest(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error)
}

type chittychatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewChittychatServiceClient(cc grpc.ClientConnInterface) ChittychatServiceClient {
	return &chittychatServiceClient{cc}
}

func (c *chittychatServiceClient) GetMessage(ctx context.Context, in *ChatRequest, opts ...grpc.CallOption) (*ChatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChatResponse)
	err := c.cc.Invoke(ctx, ChittychatService_GetMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittychatServiceClient) ProcessJoinRequest(ctx context.Context, in *JoinRequest, opts ...grpc.CallOption) (*JoinResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinResponse)
	err := c.cc.Invoke(ctx, ChittychatService_ProcessJoinRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chittychatServiceClient) ProcessLeaveRequest(ctx context.Context, in *LeaveRequest, opts ...grpc.CallOption) (*LeaveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LeaveResponse)
	err := c.cc.Invoke(ctx, ChittychatService_ProcessLeaveRequest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChittychatServiceServer is the server API for ChittychatService service.
// All implementations must embed UnimplementedChittychatServiceServer
// for forward compatibility.
type ChittychatServiceServer interface {
	GetMessage(context.Context, *ChatRequest) (*ChatResponse, error)
	ProcessJoinRequest(context.Context, *JoinRequest) (*JoinResponse, error)
	ProcessLeaveRequest(context.Context, *LeaveRequest) (*LeaveResponse, error)
	mustEmbedUnimplementedChittychatServiceServer()
}

// UnimplementedChittychatServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedChittychatServiceServer struct{}

func (UnimplementedChittychatServiceServer) GetMessage(context.Context, *ChatRequest) (*ChatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessage not implemented")
}
func (UnimplementedChittychatServiceServer) ProcessJoinRequest(context.Context, *JoinRequest) (*JoinResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessJoinRequest not implemented")
}
func (UnimplementedChittychatServiceServer) ProcessLeaveRequest(context.Context, *LeaveRequest) (*LeaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProcessLeaveRequest not implemented")
}
func (UnimplementedChittychatServiceServer) mustEmbedUnimplementedChittychatServiceServer() {}
func (UnimplementedChittychatServiceServer) testEmbeddedByValue()                           {}

// UnsafeChittychatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChittychatServiceServer will
// result in compilation errors.
type UnsafeChittychatServiceServer interface {
	mustEmbedUnimplementedChittychatServiceServer()
}

func RegisterChittychatServiceServer(s grpc.ServiceRegistrar, srv ChittychatServiceServer) {
	// If the following call pancis, it indicates UnimplementedChittychatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ChittychatService_ServiceDesc, srv)
}

func _ChittychatService_GetMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittychatServiceServer).GetMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChittychatService_GetMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittychatServiceServer).GetMessage(ctx, req.(*ChatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittychatService_ProcessJoinRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittychatServiceServer).ProcessJoinRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChittychatService_ProcessJoinRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittychatServiceServer).ProcessJoinRequest(ctx, req.(*JoinRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChittychatService_ProcessLeaveRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LeaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChittychatServiceServer).ProcessLeaveRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ChittychatService_ProcessLeaveRequest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChittychatServiceServer).ProcessLeaveRequest(ctx, req.(*LeaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ChittychatService_ServiceDesc is the grpc.ServiceDesc for ChittychatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChittychatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "chittychat.chittychatService",
	HandlerType: (*ChittychatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessage",
			Handler:    _ChittychatService_GetMessage_Handler,
		},
		{
			MethodName: "ProcessJoinRequest",
			Handler:    _ChittychatService_ProcessJoinRequest_Handler,
		},
		{
			MethodName: "ProcessLeaveRequest",
			Handler:    _ChittychatService_ProcessLeaveRequest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/lamport.proto",
}
