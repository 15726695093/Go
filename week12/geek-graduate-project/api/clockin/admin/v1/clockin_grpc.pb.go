// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// ClockinAdminServiceClient is the client API for ClockinAdminService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClockinAdminServiceClient interface {
	GetWorkTime(ctx context.Context, in *GetWorkTimeRequest, opts ...grpc.CallOption) (*GetWorkTimeReply, error)
}

type clockinAdminServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClockinAdminServiceClient(cc grpc.ClientConnInterface) ClockinAdminServiceClient {
	return &clockinAdminServiceClient{cc}
}

func (c *clockinAdminServiceClient) GetWorkTime(ctx context.Context, in *GetWorkTimeRequest, opts ...grpc.CallOption) (*GetWorkTimeReply, error) {
	out := new(GetWorkTimeReply)
	err := c.cc.Invoke(ctx, "/api.clockin.admin.v1.ClockinAdminService/GetWorkTime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClockinAdminServiceServer is the server API for ClockinAdminService service.
// All implementations must embed UnimplementedClockinAdminServiceServer
// for forward compatibility
type ClockinAdminServiceServer interface {
	GetWorkTime(context.Context, *GetWorkTimeRequest) (*GetWorkTimeReply, error)
	mustEmbedUnimplementedClockinAdminServiceServer()
}

// UnimplementedClockinAdminServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClockinAdminServiceServer struct {
}

func (UnimplementedClockinAdminServiceServer) GetWorkTime(context.Context, *GetWorkTimeRequest) (*GetWorkTimeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWorkTime not implemented")
}
func (UnimplementedClockinAdminServiceServer) mustEmbedUnimplementedClockinAdminServiceServer() {}

// UnsafeClockinAdminServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClockinAdminServiceServer will
// result in compilation errors.
type UnsafeClockinAdminServiceServer interface {
	mustEmbedUnimplementedClockinAdminServiceServer()
}

func RegisterClockinAdminServiceServer(s grpc.ServiceRegistrar, srv ClockinAdminServiceServer) {
	s.RegisterService(&ClockinAdminService_ServiceDesc, srv)
}

func _ClockinAdminService_GetWorkTime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWorkTimeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClockinAdminServiceServer).GetWorkTime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.clockin.admin.v1.ClockinAdminService/GetWorkTime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClockinAdminServiceServer).GetWorkTime(ctx, req.(*GetWorkTimeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClockinAdminService_ServiceDesc is the grpc.ServiceDesc for ClockinAdminService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClockinAdminService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.clockin.admin.v1.ClockinAdminService",
	HandlerType: (*ClockinAdminServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWorkTime",
			Handler:    _ClockinAdminService_GetWorkTime_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/clockin/admin/v1/clockin.proto",
}
