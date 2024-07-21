// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: grpc/Perun_api.proto

package perun_api

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
	PerunAPI_Version_FullMethodName        = "/perun_api.PerunAPI/Version"
	PerunAPI_ConnectVelez_FullMethodName   = "/perun_api.PerunAPI/ConnectVelez"
	PerunAPI_ListNodes_FullMethodName      = "/perun_api.PerunAPI/ListNodes"
	PerunAPI_CreateService_FullMethodName  = "/perun_api.PerunAPI/CreateService"
	PerunAPI_RefreshService_FullMethodName = "/perun_api.PerunAPI/RefreshService"
	PerunAPI_DeployService_FullMethodName  = "/perun_api.PerunAPI/DeployService"
	PerunAPI_DeployResource_FullMethodName = "/perun_api.PerunAPI/DeployResource"
)

// PerunAPIClient is the client API for PerunAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PerunAPIClient interface {
	Version(ctx context.Context, in *Version_Request, opts ...grpc.CallOption) (*Version_Response, error)
	// ConnectVelez - registers new working node with Velez running
	ConnectVelez(ctx context.Context, in *ConnectVelez_Request, opts ...grpc.CallOption) (*ConnectVelez_Response, error)
	// ListNodes - returns list of working nodes (Velez) that handle service maintenance
	ListNodes(ctx context.Context, in *ListNodes_Request, opts ...grpc.CallOption) (*ListNodes_Response, error)
	// CreateService - registers new service and updates it's information
	CreateService(ctx context.Context, in *CreateService_Request, opts ...grpc.CallOption) (*CreateService_Response, error)
	// RefreshService - refreshes service info according to config.yaml
	RefreshService(ctx context.Context, in *RefreshService_Request, opts ...grpc.CallOption) (*RefreshService_Response, error)
	// Deploys (or redeploys) service
	DeployService(ctx context.Context, in *DeployService_Request, opts ...grpc.CallOption) (*DeployService_Response, error)
	DeployResource(ctx context.Context, in *DeployResource_Request, opts ...grpc.CallOption) (*DeployResource_Response, error)
}

type perunAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewPerunAPIClient(cc grpc.ClientConnInterface) PerunAPIClient {
	return &perunAPIClient{cc}
}

func (c *perunAPIClient) Version(ctx context.Context, in *Version_Request, opts ...grpc.CallOption) (*Version_Response, error) {
	out := new(Version_Response)
	err := c.cc.Invoke(ctx, PerunAPI_Version_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *perunAPIClient) ConnectVelez(ctx context.Context, in *ConnectVelez_Request, opts ...grpc.CallOption) (*ConnectVelez_Response, error) {
	out := new(ConnectVelez_Response)
	err := c.cc.Invoke(ctx, PerunAPI_ConnectVelez_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *perunAPIClient) ListNodes(ctx context.Context, in *ListNodes_Request, opts ...grpc.CallOption) (*ListNodes_Response, error) {
	out := new(ListNodes_Response)
	err := c.cc.Invoke(ctx, PerunAPI_ListNodes_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *perunAPIClient) CreateService(ctx context.Context, in *CreateService_Request, opts ...grpc.CallOption) (*CreateService_Response, error) {
	out := new(CreateService_Response)
	err := c.cc.Invoke(ctx, PerunAPI_CreateService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *perunAPIClient) RefreshService(ctx context.Context, in *RefreshService_Request, opts ...grpc.CallOption) (*RefreshService_Response, error) {
	out := new(RefreshService_Response)
	err := c.cc.Invoke(ctx, PerunAPI_RefreshService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *perunAPIClient) DeployService(ctx context.Context, in *DeployService_Request, opts ...grpc.CallOption) (*DeployService_Response, error) {
	out := new(DeployService_Response)
	err := c.cc.Invoke(ctx, PerunAPI_DeployService_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *perunAPIClient) DeployResource(ctx context.Context, in *DeployResource_Request, opts ...grpc.CallOption) (*DeployResource_Response, error) {
	out := new(DeployResource_Response)
	err := c.cc.Invoke(ctx, PerunAPI_DeployResource_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PerunAPIServer is the server API for PerunAPI service.
// All implementations must embed UnimplementedPerunAPIServer
// for forward compatibility
type PerunAPIServer interface {
	Version(context.Context, *Version_Request) (*Version_Response, error)
	// ConnectVelez - registers new working node with Velez running
	ConnectVelez(context.Context, *ConnectVelez_Request) (*ConnectVelez_Response, error)
	// ListNodes - returns list of working nodes (Velez) that handle service maintenance
	ListNodes(context.Context, *ListNodes_Request) (*ListNodes_Response, error)
	// CreateService - registers new service and updates it's information
	CreateService(context.Context, *CreateService_Request) (*CreateService_Response, error)
	// RefreshService - refreshes service info according to config.yaml
	RefreshService(context.Context, *RefreshService_Request) (*RefreshService_Response, error)
	// Deploys (or redeploys) service
	DeployService(context.Context, *DeployService_Request) (*DeployService_Response, error)
	DeployResource(context.Context, *DeployResource_Request) (*DeployResource_Response, error)
	mustEmbedUnimplementedPerunAPIServer()
}

// UnimplementedPerunAPIServer must be embedded to have forward compatible implementations.
type UnimplementedPerunAPIServer struct {
}

func (UnimplementedPerunAPIServer) Version(context.Context, *Version_Request) (*Version_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedPerunAPIServer) ConnectVelez(context.Context, *ConnectVelez_Request) (*ConnectVelez_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConnectVelez not implemented")
}
func (UnimplementedPerunAPIServer) ListNodes(context.Context, *ListNodes_Request) (*ListNodes_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNodes not implemented")
}
func (UnimplementedPerunAPIServer) CreateService(context.Context, *CreateService_Request) (*CreateService_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateService not implemented")
}
func (UnimplementedPerunAPIServer) RefreshService(context.Context, *RefreshService_Request) (*RefreshService_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshService not implemented")
}
func (UnimplementedPerunAPIServer) DeployService(context.Context, *DeployService_Request) (*DeployService_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployService not implemented")
}
func (UnimplementedPerunAPIServer) DeployResource(context.Context, *DeployResource_Request) (*DeployResource_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployResource not implemented")
}
func (UnimplementedPerunAPIServer) mustEmbedUnimplementedPerunAPIServer() {}

// UnsafePerunAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PerunAPIServer will
// result in compilation errors.
type UnsafePerunAPIServer interface {
	mustEmbedUnimplementedPerunAPIServer()
}

func RegisterPerunAPIServer(s grpc.ServiceRegistrar, srv PerunAPIServer) {
	s.RegisterService(&PerunAPI_ServiceDesc, srv)
}

func _PerunAPI_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Version_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerunAPIServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerunAPI_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerunAPIServer).Version(ctx, req.(*Version_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerunAPI_ConnectVelez_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConnectVelez_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerunAPIServer).ConnectVelez(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerunAPI_ConnectVelez_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerunAPIServer).ConnectVelez(ctx, req.(*ConnectVelez_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerunAPI_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodes_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerunAPIServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerunAPI_ListNodes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerunAPIServer).ListNodes(ctx, req.(*ListNodes_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerunAPI_CreateService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateService_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerunAPIServer).CreateService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerunAPI_CreateService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerunAPIServer).CreateService(ctx, req.(*CreateService_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerunAPI_RefreshService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshService_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerunAPIServer).RefreshService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerunAPI_RefreshService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerunAPIServer).RefreshService(ctx, req.(*RefreshService_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerunAPI_DeployService_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployService_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerunAPIServer).DeployService(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerunAPI_DeployService_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerunAPIServer).DeployService(ctx, req.(*DeployService_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _PerunAPI_DeployResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployResource_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PerunAPIServer).DeployResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PerunAPI_DeployResource_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PerunAPIServer).DeployResource(ctx, req.(*DeployResource_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// PerunAPI_ServiceDesc is the grpc.ServiceDesc for PerunAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PerunAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "perun_api.PerunAPI",
	HandlerType: (*PerunAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _PerunAPI_Version_Handler,
		},
		{
			MethodName: "ConnectVelez",
			Handler:    _PerunAPI_ConnectVelez_Handler,
		},
		{
			MethodName: "ListNodes",
			Handler:    _PerunAPI_ListNodes_Handler,
		},
		{
			MethodName: "CreateService",
			Handler:    _PerunAPI_CreateService_Handler,
		},
		{
			MethodName: "RefreshService",
			Handler:    _PerunAPI_RefreshService_Handler,
		},
		{
			MethodName: "DeployService",
			Handler:    _PerunAPI_DeployService_Handler,
		},
		{
			MethodName: "DeployResource",
			Handler:    _PerunAPI_DeployResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/Perun_api.proto",
}
