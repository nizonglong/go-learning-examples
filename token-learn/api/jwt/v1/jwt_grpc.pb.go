// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: jwt/v1/jwt.proto

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

const (
	Jwt_GenerateJwt_FullMethodName = "/jwt.v1.Jwt/GenerateJwt"
	Jwt_ValidateJwt_FullMethodName = "/jwt.v1.Jwt/ValidateJwt"
)

// JwtClient is the client API for Jwt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JwtClient interface {
	GenerateJwt(ctx context.Context, in *GenerateJwtRequest, opts ...grpc.CallOption) (*GenerateJwtResponse, error)
	ValidateJwt(ctx context.Context, in *ValidateJwtRequest, opts ...grpc.CallOption) (*ValidateJwtResponse, error)
}

type jwtClient struct {
	cc grpc.ClientConnInterface
}

func NewJwtClient(cc grpc.ClientConnInterface) JwtClient {
	return &jwtClient{cc}
}

func (c *jwtClient) GenerateJwt(ctx context.Context, in *GenerateJwtRequest, opts ...grpc.CallOption) (*GenerateJwtResponse, error) {
	out := new(GenerateJwtResponse)
	err := c.cc.Invoke(ctx, Jwt_GenerateJwt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jwtClient) ValidateJwt(ctx context.Context, in *ValidateJwtRequest, opts ...grpc.CallOption) (*ValidateJwtResponse, error) {
	out := new(ValidateJwtResponse)
	err := c.cc.Invoke(ctx, Jwt_ValidateJwt_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JwtServer is the server API for Jwt service.
// All implementations must embed UnimplementedJwtServer
// for forward compatibility
type JwtServer interface {
	GenerateJwt(context.Context, *GenerateJwtRequest) (*GenerateJwtResponse, error)
	ValidateJwt(context.Context, *ValidateJwtRequest) (*ValidateJwtResponse, error)
	mustEmbedUnimplementedJwtServer()
}

// UnimplementedJwtServer must be embedded to have forward compatible implementations.
type UnimplementedJwtServer struct {
}

func (UnimplementedJwtServer) GenerateJwt(context.Context, *GenerateJwtRequest) (*GenerateJwtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateJwt not implemented")
}
func (UnimplementedJwtServer) ValidateJwt(context.Context, *ValidateJwtRequest) (*ValidateJwtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateJwt not implemented")
}
func (UnimplementedJwtServer) mustEmbedUnimplementedJwtServer() {}

// UnsafeJwtServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JwtServer will
// result in compilation errors.
type UnsafeJwtServer interface {
	mustEmbedUnimplementedJwtServer()
}

func RegisterJwtServer(s grpc.ServiceRegistrar, srv JwtServer) {
	s.RegisterService(&Jwt_ServiceDesc, srv)
}

func _Jwt_GenerateJwt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateJwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JwtServer).GenerateJwt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jwt_GenerateJwt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JwtServer).GenerateJwt(ctx, req.(*GenerateJwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Jwt_ValidateJwt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateJwtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JwtServer).ValidateJwt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Jwt_ValidateJwt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JwtServer).ValidateJwt(ctx, req.(*ValidateJwtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Jwt_ServiceDesc is the grpc.ServiceDesc for Jwt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Jwt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jwt.v1.Jwt",
	HandlerType: (*JwtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateJwt",
			Handler:    _Jwt_GenerateJwt_Handler,
		},
		{
			MethodName: "ValidateJwt",
			Handler:    _Jwt_ValidateJwt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jwt/v1/jwt.proto",
}
