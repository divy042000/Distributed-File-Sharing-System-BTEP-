// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: heartbeat.proto

package __

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
	HeartbeatService_SendHeartbeat_FullMethodName = "/proto.HeartbeatService/SendHeartbeat"
)

// HeartbeatServiceClient is the client API for HeartbeatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HeartbeatServiceClient interface {
	SendHeartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
}

type heartbeatServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewHeartbeatServiceClient(cc grpc.ClientConnInterface) HeartbeatServiceClient {
	return &heartbeatServiceClient{cc}
}

func (c *heartbeatServiceClient) SendHeartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, HeartbeatService_SendHeartbeat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HeartbeatServiceServer is the server API for HeartbeatService service.
// All implementations must embed UnimplementedHeartbeatServiceServer
// for forward compatibility.
type HeartbeatServiceServer interface {
	SendHeartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
	mustEmbedUnimplementedHeartbeatServiceServer()
}

// UnimplementedHeartbeatServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedHeartbeatServiceServer struct{}

func (UnimplementedHeartbeatServiceServer) SendHeartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendHeartbeat not implemented")
}
func (UnimplementedHeartbeatServiceServer) mustEmbedUnimplementedHeartbeatServiceServer() {}
func (UnimplementedHeartbeatServiceServer) testEmbeddedByValue()                          {}

// UnsafeHeartbeatServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HeartbeatServiceServer will
// result in compilation errors.
type UnsafeHeartbeatServiceServer interface {
	mustEmbedUnimplementedHeartbeatServiceServer()
}

func RegisterHeartbeatServiceServer(s grpc.ServiceRegistrar, srv HeartbeatServiceServer) {
	// If the following call pancis, it indicates UnimplementedHeartbeatServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&HeartbeatService_ServiceDesc, srv)
}

func _HeartbeatService_SendHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HeartbeatServiceServer).SendHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: HeartbeatService_SendHeartbeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HeartbeatServiceServer).SendHeartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HeartbeatService_ServiceDesc is the grpc.ServiceDesc for HeartbeatService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HeartbeatService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.HeartbeatService",
	HandlerType: (*HeartbeatServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendHeartbeat",
			Handler:    _HeartbeatService_SendHeartbeat_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "heartbeat.proto",
}
