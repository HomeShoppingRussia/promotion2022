// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.0
// source: proto/loto.proto

package loto

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

// LotoClient is the client API for Loto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LotoClient interface {
	GetWinners(ctx context.Context, in *LotoRequest, opts ...grpc.CallOption) (*LotoResponse, error)
	UploadData(ctx context.Context, in *UploadDataRequest, opts ...grpc.CallOption) (*UploadDataResponse, error)
	FillPrizesTable(ctx context.Context, in *UploadDataRequest, opts ...grpc.CallOption) (*UploadDataResponse, error)
}

type lotoClient struct {
	cc grpc.ClientConnInterface
}

func NewLotoClient(cc grpc.ClientConnInterface) LotoClient {
	return &lotoClient{cc}
}

func (c *lotoClient) GetWinners(ctx context.Context, in *LotoRequest, opts ...grpc.CallOption) (*LotoResponse, error) {
	out := new(LotoResponse)
	err := c.cc.Invoke(ctx, "/loto.Loto/GetWinners", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lotoClient) UploadData(ctx context.Context, in *UploadDataRequest, opts ...grpc.CallOption) (*UploadDataResponse, error) {
	out := new(UploadDataResponse)
	err := c.cc.Invoke(ctx, "/loto.Loto/UploadData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *lotoClient) FillPrizesTable(ctx context.Context, in *UploadDataRequest, opts ...grpc.CallOption) (*UploadDataResponse, error) {
	out := new(UploadDataResponse)
	err := c.cc.Invoke(ctx, "/loto.Loto/FillPrizesTable", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LotoServer is the server API for Loto service.
// All implementations must embed UnimplementedLotoServer
// for forward compatibility
type LotoServer interface {
	GetWinners(context.Context, *LotoRequest) (*LotoResponse, error)
	UploadData(context.Context, *UploadDataRequest) (*UploadDataResponse, error)
	FillPrizesTable(context.Context, *UploadDataRequest) (*UploadDataResponse, error)
	mustEmbedUnimplementedLotoServer()
}

// UnimplementedLotoServer must be embedded to have forward compatible implementations.
type UnimplementedLotoServer struct {
}

func (UnimplementedLotoServer) GetWinners(context.Context, *LotoRequest) (*LotoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWinners not implemented")
}
func (UnimplementedLotoServer) UploadData(context.Context, *UploadDataRequest) (*UploadDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadData not implemented")
}
func (UnimplementedLotoServer) FillPrizesTable(context.Context, *UploadDataRequest) (*UploadDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FillPrizesTable not implemented")
}
func (UnimplementedLotoServer) mustEmbedUnimplementedLotoServer() {}

// UnsafeLotoServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LotoServer will
// result in compilation errors.
type UnsafeLotoServer interface {
	mustEmbedUnimplementedLotoServer()
}

func RegisterLotoServer(s grpc.ServiceRegistrar, srv LotoServer) {
	s.RegisterService(&Loto_ServiceDesc, srv)
}

func _Loto_GetWinners_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LotoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LotoServer).GetWinners(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loto.Loto/GetWinners",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LotoServer).GetWinners(ctx, req.(*LotoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loto_UploadData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LotoServer).UploadData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loto.Loto/UploadData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LotoServer).UploadData(ctx, req.(*UploadDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Loto_FillPrizesTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LotoServer).FillPrizesTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/loto.Loto/FillPrizesTable",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LotoServer).FillPrizesTable(ctx, req.(*UploadDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Loto_ServiceDesc is the grpc.ServiceDesc for Loto service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Loto_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "loto.Loto",
	HandlerType: (*LotoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetWinners",
			Handler:    _Loto_GetWinners_Handler,
		},
		{
			MethodName: "UploadData",
			Handler:    _Loto_UploadData_Handler,
		},
		{
			MethodName: "FillPrizesTable",
			Handler:    _Loto_FillPrizesTable_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/loto.proto",
}
