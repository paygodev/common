// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.0
// source: api/sms/sms.proto

package sms

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
	SMSService_SendSMS_FullMethodName = "/api.SMSService/SendSMS"
)

// SMSServiceClient is the client API for SMSService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SMSServiceClient interface {
	SendSMS(ctx context.Context, in *SMSRequest, opts ...grpc.CallOption) (*SMSResponse, error)
}

type sMSServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSMSServiceClient(cc grpc.ClientConnInterface) SMSServiceClient {
	return &sMSServiceClient{cc}
}

func (c *sMSServiceClient) SendSMS(ctx context.Context, in *SMSRequest, opts ...grpc.CallOption) (*SMSResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SMSResponse)
	err := c.cc.Invoke(ctx, SMSService_SendSMS_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SMSServiceServer is the server API for SMSService service.
// All implementations must embed UnimplementedSMSServiceServer
// for forward compatibility
type SMSServiceServer interface {
	SendSMS(context.Context, *SMSRequest) (*SMSResponse, error)
	mustEmbedUnimplementedSMSServiceServer()
}

// UnimplementedSMSServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSMSServiceServer struct {
}

func (UnimplementedSMSServiceServer) SendSMS(context.Context, *SMSRequest) (*SMSResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSMS not implemented")
}
func (UnimplementedSMSServiceServer) mustEmbedUnimplementedSMSServiceServer() {}

// UnsafeSMSServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SMSServiceServer will
// result in compilation errors.
type UnsafeSMSServiceServer interface {
	mustEmbedUnimplementedSMSServiceServer()
}

func RegisterSMSServiceServer(s grpc.ServiceRegistrar, srv SMSServiceServer) {
	s.RegisterService(&SMSService_ServiceDesc, srv)
}

func _SMSService_SendSMS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SMSServiceServer).SendSMS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SMSService_SendSMS_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SMSServiceServer).SendSMS(ctx, req.(*SMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SMSService_ServiceDesc is the grpc.ServiceDesc for SMSService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SMSService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.SMSService",
	HandlerType: (*SMSServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendSMS",
			Handler:    _SMSService_SendSMS_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/sms/sms.proto",
}
