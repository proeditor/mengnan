// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.17.3
// source: college.proto

package college_rpc

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

// CollegeClient is the client API for College service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CollegeClient interface {
	AddCollege(ctx context.Context, in *AddCollegeRequest, opts ...grpc.CallOption) (*AddCollegeResponse, error)
	UpdateCollege(ctx context.Context, in *UpdateCollegeRequest, opts ...grpc.CallOption) (*UpdateCollegeResponse, error)
}

type collegeClient struct {
	cc grpc.ClientConnInterface
}

func NewCollegeClient(cc grpc.ClientConnInterface) CollegeClient {
	return &collegeClient{cc}
}

func (c *collegeClient) AddCollege(ctx context.Context, in *AddCollegeRequest, opts ...grpc.CallOption) (*AddCollegeResponse, error) {
	out := new(AddCollegeResponse)
	err := c.cc.Invoke(ctx, "/add.College/AddCollege", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *collegeClient) UpdateCollege(ctx context.Context, in *UpdateCollegeRequest, opts ...grpc.CallOption) (*UpdateCollegeResponse, error) {
	out := new(UpdateCollegeResponse)
	err := c.cc.Invoke(ctx, "/add.College/UpdateCollege", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CollegeServer is the server API for College service.
// All implementations must embed UnimplementedCollegeServer
// for forward compatibility
type CollegeServer interface {
	AddCollege(context.Context, *AddCollegeRequest) (*AddCollegeResponse, error)
	UpdateCollege(context.Context, *UpdateCollegeRequest) (*UpdateCollegeResponse, error)
	mustEmbedUnimplementedCollegeServer()
}

// UnimplementedCollegeServer must be embedded to have forward compatible implementations.
type UnimplementedCollegeServer struct {
}

func (UnimplementedCollegeServer) AddCollege(context.Context, *AddCollegeRequest) (*AddCollegeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCollege not implemented")
}
func (UnimplementedCollegeServer) UpdateCollege(context.Context, *UpdateCollegeRequest) (*UpdateCollegeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCollege not implemented")
}
func (UnimplementedCollegeServer) mustEmbedUnimplementedCollegeServer() {}

// UnsafeCollegeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CollegeServer will
// result in compilation errors.
type UnsafeCollegeServer interface {
	mustEmbedUnimplementedCollegeServer()
}

func RegisterCollegeServer(s grpc.ServiceRegistrar, srv CollegeServer) {
	s.RegisterService(&College_ServiceDesc, srv)
}

func _College_AddCollege_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCollegeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollegeServer).AddCollege(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/add.College/AddCollege",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollegeServer).AddCollege(ctx, req.(*AddCollegeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _College_UpdateCollege_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCollegeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CollegeServer).UpdateCollege(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/add.College/UpdateCollege",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CollegeServer).UpdateCollege(ctx, req.(*UpdateCollegeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// College_ServiceDesc is the grpc.ServiceDesc for College service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var College_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "add.College",
	HandlerType: (*CollegeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddCollege",
			Handler:    _College_AddCollege_Handler,
		},
		{
			MethodName: "UpdateCollege",
			Handler:    _College_UpdateCollege_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "college.proto",
}