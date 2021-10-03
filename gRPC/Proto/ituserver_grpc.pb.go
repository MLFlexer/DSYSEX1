// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package Proto

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

// CourcesClient is the client API for Cources service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CourcesClient interface {
	GetCourses(ctx context.Context, in *CourseRequest, opts ...grpc.CallOption) (*CourseReply, error)
}

type courcesClient struct {
	cc grpc.ClientConnInterface
}

func NewCourcesClient(cc grpc.ClientConnInterface) CourcesClient {
	return &courcesClient{cc}
}

func (c *courcesClient) GetCourses(ctx context.Context, in *CourseRequest, opts ...grpc.CallOption) (*CourseReply, error) {
	out := new(CourseReply)
	err := c.cc.Invoke(ctx, "/gRPC.Cources/GetCourses", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CourcesServer is the server API for Cources service.
// All implementations must embed UnimplementedCourcesServer
// for forward compatibility
type CourcesServer interface {
	GetCourses(context.Context, *CourseRequest) (*CourseReply, error)
	mustEmbedUnimplementedCourcesServer()
}

// UnimplementedCourcesServer must be embedded to have forward compatible implementations.
type UnimplementedCourcesServer struct {
}

func (UnimplementedCourcesServer) GetCourses(context.Context, *CourseRequest) (*CourseReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCourses not implemented")
}
func (UnimplementedCourcesServer) mustEmbedUnimplementedCourcesServer() {}

// UnsafeCourcesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CourcesServer will
// result in compilation errors.
type UnsafeCourcesServer interface {
	mustEmbedUnimplementedCourcesServer()
}

func RegisterCourcesServer(s grpc.ServiceRegistrar, srv CourcesServer) {
	s.RegisterService(&Cources_ServiceDesc, srv)
}

func _Cources_GetCourses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CourseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CourcesServer).GetCourses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/gRPC.Cources/GetCourses",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CourcesServer).GetCourses(ctx, req.(*CourseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cources_ServiceDesc is the grpc.ServiceDesc for Cources service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cources_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gRPC.Cources",
	HandlerType: (*CourcesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCourses",
			Handler:    _Cources_GetCourses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/ituserver.proto",
}