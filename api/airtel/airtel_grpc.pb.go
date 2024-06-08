// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: api/airtel/airtel.proto

package airtel

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
	AirtelService_AirtelAuth_FullMethodName             = "/airtel.AirtelService/AirtelAuth"
	AirtelService_AirtelCollection_FullMethodName       = "/airtel.AirtelService/AirtelCollection"
	AirtelService_AirtelCollectionStatus_FullMethodName = "/airtel.AirtelService/AirtelCollectionStatus"
)

// AirtelServiceClient is the client API for AirtelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AirtelServiceClient interface {
	AirtelAuth(ctx context.Context, in *AirtelAuthRequest, opts ...grpc.CallOption) (*AirtelAuthResponse, error)
	AirtelCollection(ctx context.Context, in *MMPAirtelCollectionRequest, opts ...grpc.CallOption) (*AirtelCollectionResponse, error)
	AirtelCollectionStatus(ctx context.Context, in *AirtelCollectionStatusRequest, opts ...grpc.CallOption) (*AirtelCollectionStatusResponse, error)
}

type airtelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAirtelServiceClient(cc grpc.ClientConnInterface) AirtelServiceClient {
	return &airtelServiceClient{cc}
}

func (c *airtelServiceClient) AirtelAuth(ctx context.Context, in *AirtelAuthRequest, opts ...grpc.CallOption) (*AirtelAuthResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AirtelAuthResponse)
	err := c.cc.Invoke(ctx, AirtelService_AirtelAuth_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airtelServiceClient) AirtelCollection(ctx context.Context, in *MMPAirtelCollectionRequest, opts ...grpc.CallOption) (*AirtelCollectionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AirtelCollectionResponse)
	err := c.cc.Invoke(ctx, AirtelService_AirtelCollection_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *airtelServiceClient) AirtelCollectionStatus(ctx context.Context, in *AirtelCollectionStatusRequest, opts ...grpc.CallOption) (*AirtelCollectionStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AirtelCollectionStatusResponse)
	err := c.cc.Invoke(ctx, AirtelService_AirtelCollectionStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AirtelServiceServer is the server API for AirtelService service.
// All implementations must embed UnimplementedAirtelServiceServer
// for forward compatibility
type AirtelServiceServer interface {
	AirtelAuth(context.Context, *AirtelAuthRequest) (*AirtelAuthResponse, error)
	AirtelCollection(context.Context, *MMPAirtelCollectionRequest) (*AirtelCollectionResponse, error)
	AirtelCollectionStatus(context.Context, *AirtelCollectionStatusRequest) (*AirtelCollectionStatusResponse, error)
	mustEmbedUnimplementedAirtelServiceServer()
}

// UnimplementedAirtelServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAirtelServiceServer struct {
}

func (UnimplementedAirtelServiceServer) AirtelAuth(context.Context, *AirtelAuthRequest) (*AirtelAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AirtelAuth not implemented")
}
func (UnimplementedAirtelServiceServer) AirtelCollection(context.Context, *MMPAirtelCollectionRequest) (*AirtelCollectionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AirtelCollection not implemented")
}
func (UnimplementedAirtelServiceServer) AirtelCollectionStatus(context.Context, *AirtelCollectionStatusRequest) (*AirtelCollectionStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AirtelCollectionStatus not implemented")
}
func (UnimplementedAirtelServiceServer) mustEmbedUnimplementedAirtelServiceServer() {}

// UnsafeAirtelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AirtelServiceServer will
// result in compilation errors.
type UnsafeAirtelServiceServer interface {
	mustEmbedUnimplementedAirtelServiceServer()
}

func RegisterAirtelServiceServer(s grpc.ServiceRegistrar, srv AirtelServiceServer) {
	s.RegisterService(&AirtelService_ServiceDesc, srv)
}

func _AirtelService_AirtelAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AirtelAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AirtelServiceServer).AirtelAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AirtelService_AirtelAuth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AirtelServiceServer).AirtelAuth(ctx, req.(*AirtelAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AirtelService_AirtelCollection_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MMPAirtelCollectionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AirtelServiceServer).AirtelCollection(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AirtelService_AirtelCollection_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AirtelServiceServer).AirtelCollection(ctx, req.(*MMPAirtelCollectionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AirtelService_AirtelCollectionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AirtelCollectionStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AirtelServiceServer).AirtelCollectionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AirtelService_AirtelCollectionStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AirtelServiceServer).AirtelCollectionStatus(ctx, req.(*AirtelCollectionStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AirtelService_ServiceDesc is the grpc.ServiceDesc for AirtelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AirtelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "airtel.AirtelService",
	HandlerType: (*AirtelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AirtelAuth",
			Handler:    _AirtelService_AirtelAuth_Handler,
		},
		{
			MethodName: "AirtelCollection",
			Handler:    _AirtelService_AirtelCollection_Handler,
		},
		{
			MethodName: "AirtelCollectionStatus",
			Handler:    _AirtelService_AirtelCollectionStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/airtel/airtel.proto",
}