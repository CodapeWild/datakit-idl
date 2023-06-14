// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v3

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	v3 "skywalking.apache.org/repo/goapi/collect/common/v3"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// ContinuousProfilingServiceClient is the client API for ContinuousProfilingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ContinuousProfilingServiceClient interface {
	// Query continuous profiling policy
	QueryPolicies(ctx context.Context, in *ContinuousProfilingPolicyQuery, opts ...grpc.CallOption) (*v3.Commands, error)
	// Report the profiling task when the policy threshold is reached
	// Use the returned task ID to perform the profiling task through EBPFProfilingService#collectProfilingData.
	ReportProfilingTask(ctx context.Context, in *ContinuousProfilingReport, opts ...grpc.CallOption) (*v3.Commands, error)
}

type continuousProfilingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewContinuousProfilingServiceClient(cc grpc.ClientConnInterface) ContinuousProfilingServiceClient {
	return &continuousProfilingServiceClient{cc}
}

func (c *continuousProfilingServiceClient) QueryPolicies(ctx context.Context, in *ContinuousProfilingPolicyQuery, opts ...grpc.CallOption) (*v3.Commands, error) {
	out := new(v3.Commands)
	err := c.cc.Invoke(ctx, "/skywalking.v3.ContinuousProfilingService/queryPolicies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *continuousProfilingServiceClient) ReportProfilingTask(ctx context.Context, in *ContinuousProfilingReport, opts ...grpc.CallOption) (*v3.Commands, error) {
	out := new(v3.Commands)
	err := c.cc.Invoke(ctx, "/skywalking.v3.ContinuousProfilingService/reportProfilingTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ContinuousProfilingServiceServer is the server API for ContinuousProfilingService service.
// All implementations must embed UnimplementedContinuousProfilingServiceServer
// for forward compatibility
type ContinuousProfilingServiceServer interface {
	// Query continuous profiling policy
	QueryPolicies(context.Context, *ContinuousProfilingPolicyQuery) (*v3.Commands, error)
	// Report the profiling task when the policy threshold is reached
	// Use the returned task ID to perform the profiling task through EBPFProfilingService#collectProfilingData.
	ReportProfilingTask(context.Context, *ContinuousProfilingReport) (*v3.Commands, error)
	mustEmbedUnimplementedContinuousProfilingServiceServer()
}

// UnimplementedContinuousProfilingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedContinuousProfilingServiceServer struct {
}

func (UnimplementedContinuousProfilingServiceServer) QueryPolicies(context.Context, *ContinuousProfilingPolicyQuery) (*v3.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryPolicies not implemented")
}
func (UnimplementedContinuousProfilingServiceServer) ReportProfilingTask(context.Context, *ContinuousProfilingReport) (*v3.Commands, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReportProfilingTask not implemented")
}
func (UnimplementedContinuousProfilingServiceServer) mustEmbedUnimplementedContinuousProfilingServiceServer() {
}

// UnsafeContinuousProfilingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ContinuousProfilingServiceServer will
// result in compilation errors.
type UnsafeContinuousProfilingServiceServer interface {
	mustEmbedUnimplementedContinuousProfilingServiceServer()
}

func RegisterContinuousProfilingServiceServer(s grpc.ServiceRegistrar, srv ContinuousProfilingServiceServer) {
	s.RegisterService(&_ContinuousProfilingService_serviceDesc, srv)
}

func _ContinuousProfilingService_QueryPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContinuousProfilingPolicyQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContinuousProfilingServiceServer).QueryPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skywalking.v3.ContinuousProfilingService/queryPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContinuousProfilingServiceServer).QueryPolicies(ctx, req.(*ContinuousProfilingPolicyQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _ContinuousProfilingService_ReportProfilingTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContinuousProfilingReport)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ContinuousProfilingServiceServer).ReportProfilingTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/skywalking.v3.ContinuousProfilingService/reportProfilingTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ContinuousProfilingServiceServer).ReportProfilingTask(ctx, req.(*ContinuousProfilingReport))
	}
	return interceptor(ctx, in, info, handler)
}

var _ContinuousProfilingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "skywalking.v3.ContinuousProfilingService",
	HandlerType: (*ContinuousProfilingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "queryPolicies",
			Handler:    _ContinuousProfilingService_QueryPolicies_Handler,
		},
		{
			MethodName: "reportProfilingTask",
			Handler:    _ContinuousProfilingService_ReportProfilingTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ebpf/profiling/Continuous.proto",
}
