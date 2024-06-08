// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: api/mmp/mmp.proto

package mmp

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
	AuthService_Authntication_FullMethodName = "/mmp.AuthService/Authntication"
)

// AuthServiceClient is the client API for AuthService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthServiceClient interface {
	Authntication(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
}

type authServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthServiceClient(cc grpc.ClientConnInterface) AuthServiceClient {
	return &authServiceClient{cc}
}

func (c *authServiceClient) Authntication(ctx context.Context, in *AuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AuthResponse)
	err := c.cc.Invoke(ctx, AuthService_Authntication_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthServiceServer is the server API for AuthService service.
// All implementations must embed UnimplementedAuthServiceServer
// for forward compatibility
type AuthServiceServer interface {
	Authntication(context.Context, *AuthRequest) (*AuthResponse, error)
	mustEmbedUnimplementedAuthServiceServer()
}

// UnimplementedAuthServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthServiceServer struct {
}

func (UnimplementedAuthServiceServer) Authntication(context.Context, *AuthRequest) (*AuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authntication not implemented")
}
func (UnimplementedAuthServiceServer) mustEmbedUnimplementedAuthServiceServer() {}

// UnsafeAuthServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthServiceServer will
// result in compilation errors.
type UnsafeAuthServiceServer interface {
	mustEmbedUnimplementedAuthServiceServer()
}

func RegisterAuthServiceServer(s grpc.ServiceRegistrar, srv AuthServiceServer) {
	s.RegisterService(&AuthService_ServiceDesc, srv)
}

func _AuthService_Authntication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthServiceServer).Authntication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthService_Authntication_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthServiceServer).Authntication(ctx, req.(*AuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthService_ServiceDesc is the grpc.ServiceDesc for AuthService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mmp.AuthService",
	HandlerType: (*AuthServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Authntication",
			Handler:    _AuthService_Authntication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/mmp/mmp.proto",
}

const (
	CollectionService_MakeCollection_FullMethodName   = "/mmp.CollectionService/MakeCollection"
	CollectionService_CollectionStatus_FullMethodName = "/mmp.CollectionService/CollectionStatus"
	CollectionService_CreateCollection_FullMethodName = "/mmp.CollectionService/CreateCollection"
)

// CollectionServiceClient is the client API for CollectionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollectionServiceClient interface {
	MakeCollection(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CollectionResponse, error)
	CollectionStatus(ctx context.Context, in *CollectionResponse, opts ...grpc.CallOption) (*CollectionResponse, error)
	CreateCollection(ctx context.Context, in *InsertTransactionRequest, opts ...grpc.CallOption) (*CollectionResponse, error)
}

type collectionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCollectionServiceClient(cc grpc.ClientConnInterface) CollectionServiceClient {
	return &collectionServiceClient{cc}
}

func (c *collectionServiceClient) MakeCollection(ctx context.Context, in *CreateCollectionRequest, opts ...grpc.CallOption) (*CollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CollectionResponse)
	err := c.cc.Invoke(ctx, CollectionService_MakeCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionServiceClient) CollectionStatus(ctx context.Context, in *CollectionResponse, opts ...grpc.CallOption) (*CollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CollectionResponse)
	err := c.cc.Invoke(ctx, CollectionService_CollectionStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collectionServiceClient) CreateCollection(ctx context.Context, in *InsertTransactionRequest, opts ...grpc.CallOption) (*CollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CollectionResponse)
	err := c.cc.Invoke(ctx, CollectionService_CreateCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollectionServiceServer is the server API for CollectionService service.
// All implementations must embed UnimplementedCollectionServiceServer
// for forward compatibility
type CollectionServiceServer interface {
	MakeCollection(context.Context, *CreateCollectionRequest) (*CollectionResponse, error)
	CollectionStatus(context.Context, *CollectionResponse) (*CollectionResponse, error)
	CreateCollection(context.Context, *InsertTransactionRequest) (*CollectionResponse, error)
	mustEmbedUnimplementedCollectionServiceServer()
}

// UnimplementedCollectionServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCollectionServiceServer struct {
}

func (UnimplementedCollectionServiceServer) MakeCollection(context.Context, *CreateCollectionRequest) (*CollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeCollection not implemented")
}
func (UnimplementedCollectionServiceServer) CollectionStatus(context.Context, *CollectionResponse) (*CollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectionStatus not implemented")
}
func (UnimplementedCollectionServiceServer) CreateCollection(context.Context, *InsertTransactionRequest) (*CollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCollection not implemented")
}
func (UnimplementedCollectionServiceServer) mustEmbedUnimplementedCollectionServiceServer() {}

// UnsafeCollectionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollectionServiceServer will
// result in compilation errors.
type UnsafeCollectionServiceServer interface {
	mustEmbedUnimplementedCollectionServiceServer()
}

func RegisterCollectionServiceServer(s grpc.ServiceRegistrar, srv CollectionServiceServer) {
	s.RegisterService(&CollectionService_ServiceDesc, srv)
}

func _CollectionService_MakeCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).MakeCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_MakeCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).MakeCollection(ctx, req.(*CreateCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionService_CollectionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).CollectionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_CollectionStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).CollectionStatus(ctx, req.(*CollectionResponse))
	}
	return interceptor(ctx, in, info, handler)
}

func _CollectionService_CreateCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollectionServiceServer).CreateCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CollectionService_CreateCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollectionServiceServer).CreateCollection(ctx, req.(*InsertTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CollectionService_ServiceDesc is the grpc.ServiceDesc for CollectionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CollectionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "mmp.CollectionService",
	HandlerType: (*CollectionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "MakeCollection",
			Handler:    _CollectionService_MakeCollection_Handler,
		},
		{
			MethodName: "CollectionStatus",
			Handler:    _CollectionService_CollectionStatus_Handler,
		},
		{
			MethodName: "CreateCollection",
			Handler:    _CollectionService_CreateCollection_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/mmp/mmp.proto",
}