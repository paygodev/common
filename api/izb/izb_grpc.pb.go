// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: api/izb/izb.proto

package izb

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
	IZBService_Disbursement_FullMethodName       = "/izb.IZBService/Disbursement"
	IZBService_DisbursementStatus_FullMethodName = "/izb.IZBService/DisbursementStatus"
)

// IZBServiceClient is the client API for IZBService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type IZBServiceClient interface {
	Disbursement(ctx context.Context, in *IZBRequest, opts ...grpc.CallOption) (*IZBResponse, error)
	DisbursementStatus(ctx context.Context, in *IZBStatusRequest, opts ...grpc.CallOption) (*IZBStatusResponse, error)
}

type iZBServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewIZBServiceClient(cc grpc.ClientConnInterface) IZBServiceClient {
	return &iZBServiceClient{cc}
}

func (c *iZBServiceClient) Disbursement(ctx context.Context, in *IZBRequest, opts ...grpc.CallOption) (*IZBResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IZBResponse)
	err := c.cc.Invoke(ctx, IZBService_Disbursement_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iZBServiceClient) DisbursementStatus(ctx context.Context, in *IZBStatusRequest, opts ...grpc.CallOption) (*IZBStatusResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IZBStatusResponse)
	err := c.cc.Invoke(ctx, IZBService_DisbursementStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// IZBServiceServer is the server API for IZBService service.
// All implementations must embed UnimplementedIZBServiceServer
// for forward compatibility
type IZBServiceServer interface {
	Disbursement(context.Context, *IZBRequest) (*IZBResponse, error)
	DisbursementStatus(context.Context, *IZBStatusRequest) (*IZBStatusResponse, error)
	mustEmbedUnimplementedIZBServiceServer()
}

// UnimplementedIZBServiceServer must be embedded to have forward compatible implementations.
type UnimplementedIZBServiceServer struct {
}

func (UnimplementedIZBServiceServer) Disbursement(context.Context, *IZBRequest) (*IZBResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Disbursement not implemented")
}
func (UnimplementedIZBServiceServer) DisbursementStatus(context.Context, *IZBStatusRequest) (*IZBStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DisbursementStatus not implemented")
}
func (UnimplementedIZBServiceServer) mustEmbedUnimplementedIZBServiceServer() {}

// UnsafeIZBServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to IZBServiceServer will
// result in compilation errors.
type UnsafeIZBServiceServer interface {
	mustEmbedUnimplementedIZBServiceServer()
}

func RegisterIZBServiceServer(s grpc.ServiceRegistrar, srv IZBServiceServer) {
	s.RegisterService(&IZBService_ServiceDesc, srv)
}

func _IZBService_Disbursement_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IZBRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IZBServiceServer).Disbursement(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IZBService_Disbursement_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IZBServiceServer).Disbursement(ctx, req.(*IZBRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _IZBService_DisbursementStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IZBStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(IZBServiceServer).DisbursementStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: IZBService_DisbursementStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(IZBServiceServer).DisbursementStatus(ctx, req.(*IZBStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// IZBService_ServiceDesc is the grpc.ServiceDesc for IZBService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var IZBService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "izb.IZBService",
	HandlerType: (*IZBServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Disbursement",
			Handler:    _IZBService_Disbursement_Handler,
		},
		{
			MethodName: "DisbursementStatus",
			Handler:    _IZBService_DisbursementStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/izb/izb.proto",
}
