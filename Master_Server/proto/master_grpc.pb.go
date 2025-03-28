// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.2
// source: master.proto

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
	MasterService_GetChunkLocations_FullMethodName = "/proto.MasterService/GetChunkLocations"
	MasterService_ReportChunk_FullMethodName       = "/proto.MasterService/ReportChunk"
	MasterService_SendHeartbeat_FullMethodName     = "/proto.MasterService/SendHeartbeat"
	MasterService_RegisterFile_FullMethodName      = "/proto.MasterService/RegisterFile"
	MasterService_GetFileMetadata_FullMethodName   = "/proto.MasterService/GetFileMetadata"
)

// MasterServiceClient is the client API for MasterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Master Service
type MasterServiceClient interface {
	GetChunkLocations(ctx context.Context, in *GetChunkRequest, opts ...grpc.CallOption) (*GetChunkResponse, error)
	ReportChunk(ctx context.Context, in *ChunkReport, opts ...grpc.CallOption) (*ChunkResponse, error)
	// Receives heartbeat from chunkservers
	SendHeartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error)
	// Client registers a new file and gets chunk server assignments
	RegisterFile(ctx context.Context, in *RegisterFileRequest, opts ...grpc.CallOption) (*RegisterFileResponse, error)
	// Client retrieves metadata for an existing file
	GetFileMetadata(ctx context.Context, in *GetFileMetadataRequest, opts ...grpc.CallOption) (*GetFileMetadataResponse, error)
}

type masterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterServiceClient(cc grpc.ClientConnInterface) MasterServiceClient {
	return &masterServiceClient{cc}
}

func (c *masterServiceClient) GetChunkLocations(ctx context.Context, in *GetChunkRequest, opts ...grpc.CallOption) (*GetChunkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetChunkResponse)
	err := c.cc.Invoke(ctx, MasterService_GetChunkLocations_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) ReportChunk(ctx context.Context, in *ChunkReport, opts ...grpc.CallOption) (*ChunkResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ChunkResponse)
	err := c.cc.Invoke(ctx, MasterService_ReportChunk_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) SendHeartbeat(ctx context.Context, in *HeartbeatRequest, opts ...grpc.CallOption) (*HeartbeatResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HeartbeatResponse)
	err := c.cc.Invoke(ctx, MasterService_SendHeartbeat_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) RegisterFile(ctx context.Context, in *RegisterFileRequest, opts ...grpc.CallOption) (*RegisterFileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterFileResponse)
	err := c.cc.Invoke(ctx, MasterService_RegisterFile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterServiceClient) GetFileMetadata(ctx context.Context, in *GetFileMetadataRequest, opts ...grpc.CallOption) (*GetFileMetadataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetFileMetadataResponse)
	err := c.cc.Invoke(ctx, MasterService_GetFileMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServiceServer is the server API for MasterService service.
// All implementations must embed UnimplementedMasterServiceServer
// for forward compatibility.
//
// Master Service
type MasterServiceServer interface {
	GetChunkLocations(context.Context, *GetChunkRequest) (*GetChunkResponse, error)
	ReportChunk(context.Context, *ChunkReport) (*ChunkResponse, error)
	// Receives heartbeat from chunkservers
	SendHeartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error)
	// Client registers a new file and gets chunk server assignments
	RegisterFile(context.Context, *RegisterFileRequest) (*RegisterFileResponse, error)
	// Client retrieves metadata for an existing file
	GetFileMetadata(context.Context, *GetFileMetadataRequest) (*GetFileMetadataResponse, error)
	mustEmbedUnimplementedMasterServiceServer()
}

// UnimplementedMasterServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMasterServiceServer struct{}

func (UnimplementedMasterServiceServer) GetChunkLocations(context.Context, *GetChunkRequest) (*GetChunkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChunkLocations not implemented")
}
func (UnimplementedMasterServiceServer) ReportChunk(context.Context, *ChunkReport) (*ChunkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportChunk not implemented")
}
func (UnimplementedMasterServiceServer) SendHeartbeat(context.Context, *HeartbeatRequest) (*HeartbeatResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendHeartbeat not implemented")
}
func (UnimplementedMasterServiceServer) RegisterFile(context.Context, *RegisterFileRequest) (*RegisterFileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterFile not implemented")
}
func (UnimplementedMasterServiceServer) GetFileMetadata(context.Context, *GetFileMetadataRequest) (*GetFileMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFileMetadata not implemented")
}
func (UnimplementedMasterServiceServer) mustEmbedUnimplementedMasterServiceServer() {}
func (UnimplementedMasterServiceServer) testEmbeddedByValue()                       {}

// UnsafeMasterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterServiceServer will
// result in compilation errors.
type UnsafeMasterServiceServer interface {
	mustEmbedUnimplementedMasterServiceServer()
}

func RegisterMasterServiceServer(s grpc.ServiceRegistrar, srv MasterServiceServer) {
	// If the following call pancis, it indicates UnimplementedMasterServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MasterService_ServiceDesc, srv)
}

func _MasterService_GetChunkLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetChunkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).GetChunkLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MasterService_GetChunkLocations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).GetChunkLocations(ctx, req.(*GetChunkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_ReportChunk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChunkReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).ReportChunk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MasterService_ReportChunk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).ReportChunk(ctx, req.(*ChunkReport))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_SendHeartbeat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeartbeatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).SendHeartbeat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MasterService_SendHeartbeat_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).SendHeartbeat(ctx, req.(*HeartbeatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_RegisterFile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterFileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).RegisterFile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MasterService_RegisterFile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).RegisterFile(ctx, req.(*RegisterFileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MasterService_GetFileMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFileMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServiceServer).GetFileMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MasterService_GetFileMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServiceServer).GetFileMetadata(ctx, req.(*GetFileMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MasterService_ServiceDesc is the grpc.ServiceDesc for MasterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MasterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.MasterService",
	HandlerType: (*MasterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetChunkLocations",
			Handler:    _MasterService_GetChunkLocations_Handler,
		},
		{
			MethodName: "ReportChunk",
			Handler:    _MasterService_ReportChunk_Handler,
		},
		{
			MethodName: "SendHeartbeat",
			Handler:    _MasterService_SendHeartbeat_Handler,
		},
		{
			MethodName: "RegisterFile",
			Handler:    _MasterService_RegisterFile_Handler,
		},
		{
			MethodName: "GetFileMetadata",
			Handler:    _MasterService_GetFileMetadata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master.proto",
}
