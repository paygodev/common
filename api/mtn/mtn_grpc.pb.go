// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: api/mtn/mtn.proto

package mtn

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	MTNService_MTNAuth_FullMethodName             = "/mtn.MTNService/MTNAuth"
	MTNService_MTNCollection_FullMethodName       = "/mtn.MTNService/MTNCollection"
	MTNService_MTNCollectionStatus_FullMethodName = "/mtn.MTNService/MTNCollectionStatus"
)

// MTNServiceClient is the client API for MTNService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MTNServiceClient interface {
	MTNAuth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	MTNCollection(ctx context.Context, in *MMPCollectionRequest, opts ...grpc.CallOption) (*CollectionResponse, error)
	MTNCollectionStatus(ctx context.Context, in *CollectionStatusRequest, opts ...grpc.CallOption) (*CollectionStatusResponse, error)
}

type mTNServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMTNServiceClient(cc grpc.ClientConnInterface) MTNServiceClient {
	return &mTNServiceClient{cc}
}

func (c *mTNServiceClient) MTNAuth(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, MTNService_MTNAuth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mTNServiceClient) MTNCollection(ctx context.Context, in *MMPCollectionRequest, opts ...grpc.CallOption) (*CollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CollectionResponse)
	err := c.cc.Invoke(ctx, MTNService_MTNCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *mTNServiceClient) MTNCollectionStatus(ctx context.Context, in *CollectionStatusRequest, opts ...grpc.CallOption) (*CollectionStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CollectionStatusResponse)
	err := c.cc.Invoke(ctx, MTNService_MTNCollectionStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MTNServiceServer is the server API for MTNService service.
// All implementations must embed UnimplementedMTNServiceServer
// for forward compatibility
type MTNServiceServer interface {
	MTNAuth(context.Context, *AuthRequest) (*AuthResponse, error)
	MTNCollection(context.Context, *MMPCollectionRequest) (*CollectionResponse, error)
	MTNCollectionStatus(context.Context, *CollectionStatusRequest) (*CollectionStatusResponse, error)
	mustEmbedUnimplementedMTNServiceServer()
}

// UnimplementedMTNServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMTNServiceServer struct {
}

func (UnimplementedMTNServiceServer) MTNAuth(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MTNAuth not implemented")
}
func (UnimplementedMTNServiceServer) MTNCollection(context.Context, *MMPCollectionRequest) (*CollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MTNCollection not implemented")
}
func (UnimplementedMTNServiceServer) MTNCollectionStatus(context.Context, *CollectionStatusRequest) (*CollectionStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MTNCollectionStatus not implemented")
}
func (UnimplementedMTNServiceServer) mustEmbedUnimplementedMTNServiceServer() {}

// UnsafeMTNServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MTNServiceServer will
// result in compilation errors.
type UnsafeMTNServiceServer interface {
	mustEmbedUnimplementedMTNServiceServer()
}

func RegisterMTNServiceServer(s grpc.ServiceRegistrar, srv MTNServiceServer) {
	s.RegisterService(&MTNService_ServiceDesc, srv)
}

func _MTNService_MTNAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MTNServiceServer).MTNAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MTNService_MTNAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MTNServiceServer).MTNAuth(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MTNService_MTNCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MMPCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MTNServiceServer).MTNCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MTNService_MTNCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MTNServiceServer).MTNCollection(ctx, req.(*MMPCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _MTNService_MTNCollectionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MTNServiceServer).MTNCollectionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MTNService_MTNCollectionStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MTNServiceServer).MTNCollectionStatus(ctx, req.(*CollectionStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MTNService_ServiceDesc is the grpc.ServiceDesc for MTNService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MTNService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mtn.MTNService",
	HandlerType: (*MTNServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MTNAuth",
			Handler:    _MTNService_MTNAuth_Handler,
		},
		{
			MethodName: "MTNCollection",
			Handler:    _MTNService_MTNCollection_Handler,
		},
		{
			MethodName: "MTNCollectionStatus",
			Handler:    _MTNService_MTNCollectionStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/mtn/mtn.proto",
}
