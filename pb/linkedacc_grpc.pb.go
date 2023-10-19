// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: linkedacc.proto

package pb

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
	Linkedacc_GetVkUrl_FullMethodName          = "/Linkedacc/GetVkUrl"
	Linkedacc_AddVk_FullMethodName             = "/Linkedacc/AddVk"
	Linkedacc_AddTg_FullMethodName             = "/Linkedacc/AddTg"
	Linkedacc_GetLinkedAccounts_FullMethodName = "/Linkedacc/GetLinkedAccounts"
	Linkedacc_NewExternalPost_FullMethodName   = "/Linkedacc/NewExternalPost"
)

// LinkedaccClient is the client API for Linkedacc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LinkedaccClient interface {
	GetVkUrl(ctx context.Context, in *GetVkUrlRequest, opts ...grpc.CallOption) (*GetVkUrlResponse, error)
	AddVk(ctx context.Context, in *AddVkRequest, opts ...grpc.CallOption) (*AddVkResponse, error)
	AddTg(ctx context.Context, in *AddTgRequest, opts ...grpc.CallOption) (*AddTgResponse, error)
	GetLinkedAccounts(ctx context.Context, in *GetLinkedAccountsRequest, opts ...grpc.CallOption) (*GetLinkedAccountsResponse, error)
	NewExternalPost(ctx context.Context, in *NewExternalPostRequest, opts ...grpc.CallOption) (*NewExternalPostResponse, error)
}

type linkedaccClient struct {
	cc grpc.ClientConnInterface
}

func NewLinkedaccClient(cc grpc.ClientConnInterface) LinkedaccClient {
	return &linkedaccClient{cc}
}

func (c *linkedaccClient) GetVkUrl(ctx context.Context, in *GetVkUrlRequest, opts ...grpc.CallOption) (*GetVkUrlResponse, error) {
	out := new(GetVkUrlResponse)
	err := c.cc.Invoke(ctx, Linkedacc_GetVkUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkedaccClient) AddVk(ctx context.Context, in *AddVkRequest, opts ...grpc.CallOption) (*AddVkResponse, error) {
	out := new(AddVkResponse)
	err := c.cc.Invoke(ctx, Linkedacc_AddVk_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkedaccClient) AddTg(ctx context.Context, in *AddTgRequest, opts ...grpc.CallOption) (*AddTgResponse, error) {
	out := new(AddTgResponse)
	err := c.cc.Invoke(ctx, Linkedacc_AddTg_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkedaccClient) GetLinkedAccounts(ctx context.Context, in *GetLinkedAccountsRequest, opts ...grpc.CallOption) (*GetLinkedAccountsResponse, error) {
	out := new(GetLinkedAccountsResponse)
	err := c.cc.Invoke(ctx, Linkedacc_GetLinkedAccounts_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *linkedaccClient) NewExternalPost(ctx context.Context, in *NewExternalPostRequest, opts ...grpc.CallOption) (*NewExternalPostResponse, error) {
	out := new(NewExternalPostResponse)
	err := c.cc.Invoke(ctx, Linkedacc_NewExternalPost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LinkedaccServer is the server API for Linkedacc service.
// All implementations should embed UnimplementedLinkedaccServer
// for forward compatibility
type LinkedaccServer interface {
	GetVkUrl(context.Context, *GetVkUrlRequest) (*GetVkUrlResponse, error)
	AddVk(context.Context, *AddVkRequest) (*AddVkResponse, error)
	AddTg(context.Context, *AddTgRequest) (*AddTgResponse, error)
	GetLinkedAccounts(context.Context, *GetLinkedAccountsRequest) (*GetLinkedAccountsResponse, error)
	NewExternalPost(context.Context, *NewExternalPostRequest) (*NewExternalPostResponse, error)
}

// UnimplementedLinkedaccServer should be embedded to have forward compatible implementations.
type UnimplementedLinkedaccServer struct {
}

func (UnimplementedLinkedaccServer) GetVkUrl(context.Context, *GetVkUrlRequest) (*GetVkUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVkUrl not implemented")
}
func (UnimplementedLinkedaccServer) AddVk(context.Context, *AddVkRequest) (*AddVkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddVk not implemented")
}
func (UnimplementedLinkedaccServer) AddTg(context.Context, *AddTgRequest) (*AddTgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddTg not implemented")
}
func (UnimplementedLinkedaccServer) GetLinkedAccounts(context.Context, *GetLinkedAccountsRequest) (*GetLinkedAccountsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLinkedAccounts not implemented")
}
func (UnimplementedLinkedaccServer) NewExternalPost(context.Context, *NewExternalPostRequest) (*NewExternalPostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewExternalPost not implemented")
}

// UnsafeLinkedaccServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LinkedaccServer will
// result in compilation errors.
type UnsafeLinkedaccServer interface {
	mustEmbedUnimplementedLinkedaccServer()
}

func RegisterLinkedaccServer(s grpc.ServiceRegistrar, srv LinkedaccServer) {
	s.RegisterService(&Linkedacc_ServiceDesc, srv)
}

func _Linkedacc_GetVkUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetVkUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkedaccServer).GetVkUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Linkedacc_GetVkUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkedaccServer).GetVkUrl(ctx, req.(*GetVkUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Linkedacc_AddVk_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddVkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkedaccServer).AddVk(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Linkedacc_AddVk_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkedaccServer).AddVk(ctx, req.(*AddVkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Linkedacc_AddTg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddTgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkedaccServer).AddTg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Linkedacc_AddTg_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkedaccServer).AddTg(ctx, req.(*AddTgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Linkedacc_GetLinkedAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLinkedAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkedaccServer).GetLinkedAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Linkedacc_GetLinkedAccounts_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkedaccServer).GetLinkedAccounts(ctx, req.(*GetLinkedAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Linkedacc_NewExternalPost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewExternalPostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LinkedaccServer).NewExternalPost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Linkedacc_NewExternalPost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LinkedaccServer).NewExternalPost(ctx, req.(*NewExternalPostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Linkedacc_ServiceDesc is the grpc.ServiceDesc for Linkedacc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Linkedacc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Linkedacc",
	HandlerType: (*LinkedaccServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetVkUrl",
			Handler:    _Linkedacc_GetVkUrl_Handler,
		},
		{
			MethodName: "AddVk",
			Handler:    _Linkedacc_AddVk_Handler,
		},
		{
			MethodName: "AddTg",
			Handler:    _Linkedacc_AddTg_Handler,
		},
		{
			MethodName: "GetLinkedAccounts",
			Handler:    _Linkedacc_GetLinkedAccounts_Handler,
		},
		{
			MethodName: "NewExternalPost",
			Handler:    _Linkedacc_NewExternalPost_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "linkedacc.proto",
}