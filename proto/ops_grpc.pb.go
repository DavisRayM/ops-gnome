// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: ops.proto

package proto

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

// OpsClient is the client API for Ops service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OpsClient interface {
	ListSupportedDeployments(ctx context.Context, in *ListSupportedDeploymentsReq, opts ...grpc.CallOption) (*ListSupportedDeploymentsResp, error)
	GetDeploymentStatus(ctx context.Context, in *GetDeploymentStatusReq, opts ...grpc.CallOption) (*GetDeploymentStatusResp, error)
}

type opsClient struct {
	cc grpc.ClientConnInterface
}

func NewOpsClient(cc grpc.ClientConnInterface) OpsClient {
	return &opsClient{cc}
}

func (c *opsClient) ListSupportedDeployments(ctx context.Context, in *ListSupportedDeploymentsReq, opts ...grpc.CallOption) (*ListSupportedDeploymentsResp, error) {
	out := new(ListSupportedDeploymentsResp)
	err := c.cc.Invoke(ctx, "/ops.Ops/ListSupportedDeployments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *opsClient) GetDeploymentStatus(ctx context.Context, in *GetDeploymentStatusReq, opts ...grpc.CallOption) (*GetDeploymentStatusResp, error) {
	out := new(GetDeploymentStatusResp)
	err := c.cc.Invoke(ctx, "/ops.Ops/GetDeploymentStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OpsServer is the server API for Ops service.
// All implementations must embed UnimplementedOpsServer
// for forward compatibility
type OpsServer interface {
	ListSupportedDeployments(context.Context, *ListSupportedDeploymentsReq) (*ListSupportedDeploymentsResp, error)
	GetDeploymentStatus(context.Context, *GetDeploymentStatusReq) (*GetDeploymentStatusResp, error)
	mustEmbedUnimplementedOpsServer()
}

// UnimplementedOpsServer must be embedded to have forward compatible implementations.
type UnimplementedOpsServer struct {
}

func (UnimplementedOpsServer) ListSupportedDeployments(context.Context, *ListSupportedDeploymentsReq) (*ListSupportedDeploymentsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSupportedDeployments not implemented")
}
func (UnimplementedOpsServer) GetDeploymentStatus(context.Context, *GetDeploymentStatusReq) (*GetDeploymentStatusResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeploymentStatus not implemented")
}
func (UnimplementedOpsServer) mustEmbedUnimplementedOpsServer() {}

// UnsafeOpsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OpsServer will
// result in compilation errors.
type UnsafeOpsServer interface {
	mustEmbedUnimplementedOpsServer()
}

func RegisterOpsServer(s grpc.ServiceRegistrar, srv OpsServer) {
	s.RegisterService(&Ops_ServiceDesc, srv)
}

func _Ops_ListSupportedDeployments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSupportedDeploymentsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServer).ListSupportedDeployments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ops.Ops/ListSupportedDeployments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServer).ListSupportedDeployments(ctx, req.(*ListSupportedDeploymentsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Ops_GetDeploymentStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeploymentStatusReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OpsServer).GetDeploymentStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ops.Ops/GetDeploymentStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OpsServer).GetDeploymentStatus(ctx, req.(*GetDeploymentStatusReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Ops_ServiceDesc is the grpc.ServiceDesc for Ops service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Ops_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ops.Ops",
	HandlerType: (*OpsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSupportedDeployments",
			Handler:    _Ops_ListSupportedDeployments_Handler,
		},
		{
			MethodName: "GetDeploymentStatus",
			Handler:    _Ops_GetDeploymentStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ops.proto",
}
