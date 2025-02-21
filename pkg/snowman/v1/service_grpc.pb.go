// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: snowman/v1/service.proto

package snowman

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
	SnowmanServiceV1_List_FullMethodName  = "/snowman.v1.SnowmanServiceV1/List"
	SnowmanServiceV1_Build_FullMethodName = "/snowman.v1.SnowmanServiceV1/Build"
)

// SnowmanServiceV1Client is the client API for SnowmanServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SnowmanServiceV1Client interface {
	// Returns list of snowmen
	List(ctx context.Context, in *ListSnowmenRequest, opts ...grpc.CallOption) (*ListSnowmenResponse, error)
	// Builds snowman based on the specified parameters
	Build(ctx context.Context, in *BuildSnowmanRequest, opts ...grpc.CallOption) (*BuildSnowmanResponse, error)
}

type snowmanServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewSnowmanServiceV1Client(cc grpc.ClientConnInterface) SnowmanServiceV1Client {
	return &snowmanServiceV1Client{cc}
}

func (c *snowmanServiceV1Client) List(ctx context.Context, in *ListSnowmenRequest, opts ...grpc.CallOption) (*ListSnowmenResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListSnowmenResponse)
	err := c.cc.Invoke(ctx, SnowmanServiceV1_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snowmanServiceV1Client) Build(ctx context.Context, in *BuildSnowmanRequest, opts ...grpc.CallOption) (*BuildSnowmanResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(BuildSnowmanResponse)
	err := c.cc.Invoke(ctx, SnowmanServiceV1_Build_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnowmanServiceV1Server is the server API for SnowmanServiceV1 service.
// All implementations must embed UnimplementedSnowmanServiceV1Server
// for forward compatibility.
type SnowmanServiceV1Server interface {
	// Returns list of snowmen
	List(context.Context, *ListSnowmenRequest) (*ListSnowmenResponse, error)
	// Builds snowman based on the specified parameters
	Build(context.Context, *BuildSnowmanRequest) (*BuildSnowmanResponse, error)
	mustEmbedUnimplementedSnowmanServiceV1Server()
}

// UnimplementedSnowmanServiceV1Server must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedSnowmanServiceV1Server struct{}

func (UnimplementedSnowmanServiceV1Server) List(context.Context, *ListSnowmenRequest) (*ListSnowmenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedSnowmanServiceV1Server) Build(context.Context, *BuildSnowmanRequest) (*BuildSnowmanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Build not implemented")
}
func (UnimplementedSnowmanServiceV1Server) mustEmbedUnimplementedSnowmanServiceV1Server() {}
func (UnimplementedSnowmanServiceV1Server) testEmbeddedByValue()                          {}

// UnsafeSnowmanServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SnowmanServiceV1Server will
// result in compilation errors.
type UnsafeSnowmanServiceV1Server interface {
	mustEmbedUnimplementedSnowmanServiceV1Server()
}

func RegisterSnowmanServiceV1Server(s grpc.ServiceRegistrar, srv SnowmanServiceV1Server) {
	// If the following call pancis, it indicates UnimplementedSnowmanServiceV1Server was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&SnowmanServiceV1_ServiceDesc, srv)
}

func _SnowmanServiceV1_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSnowmenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnowmanServiceV1Server).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnowmanServiceV1_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnowmanServiceV1Server).List(ctx, req.(*ListSnowmenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnowmanServiceV1_Build_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildSnowmanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnowmanServiceV1Server).Build(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SnowmanServiceV1_Build_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnowmanServiceV1Server).Build(ctx, req.(*BuildSnowmanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SnowmanServiceV1_ServiceDesc is the grpc.ServiceDesc for SnowmanServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SnowmanServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "snowman.v1.SnowmanServiceV1",
	HandlerType: (*SnowmanServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _SnowmanServiceV1_List_Handler,
		},
		{
			MethodName: "Build",
			Handler:    _SnowmanServiceV1_Build_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "snowman/v1/service.proto",
}
