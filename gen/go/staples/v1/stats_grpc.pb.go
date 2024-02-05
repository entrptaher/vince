// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: staples/v1/stats.proto

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
	Stats_RealtimeVisitors_FullMethodName = "/v1.Stats/RealtimeVisitors"
	Stats_Aggregate_FullMethodName        = "/v1.Stats/Aggregate"
	Stats_Timeseries_FullMethodName       = "/v1.Stats/Timeseries"
	Stats_BreakDown_FullMethodName        = "/v1.Stats/BreakDown"
)

// StatsClient is the client API for Stats service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StatsClient interface {
	RealtimeVisitors(ctx context.Context, in *Realtime_Request, opts ...grpc.CallOption) (*Realtime_Response, error)
	Aggregate(ctx context.Context, in *Aggregate_Request, opts ...grpc.CallOption) (*Aggregate_Response, error)
	Timeseries(ctx context.Context, in *Timeseries_Request, opts ...grpc.CallOption) (*Timeseries_Response, error)
	BreakDown(ctx context.Context, in *BreakDown_Request, opts ...grpc.CallOption) (*BreakDown_Response, error)
}

type statsClient struct {
	cc grpc.ClientConnInterface
}

func NewStatsClient(cc grpc.ClientConnInterface) StatsClient {
	return &statsClient{cc}
}

func (c *statsClient) RealtimeVisitors(ctx context.Context, in *Realtime_Request, opts ...grpc.CallOption) (*Realtime_Response, error) {
	out := new(Realtime_Response)
	err := c.cc.Invoke(ctx, Stats_RealtimeVisitors_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) Aggregate(ctx context.Context, in *Aggregate_Request, opts ...grpc.CallOption) (*Aggregate_Response, error) {
	out := new(Aggregate_Response)
	err := c.cc.Invoke(ctx, Stats_Aggregate_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) Timeseries(ctx context.Context, in *Timeseries_Request, opts ...grpc.CallOption) (*Timeseries_Response, error) {
	out := new(Timeseries_Response)
	err := c.cc.Invoke(ctx, Stats_Timeseries_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *statsClient) BreakDown(ctx context.Context, in *BreakDown_Request, opts ...grpc.CallOption) (*BreakDown_Response, error) {
	out := new(BreakDown_Response)
	err := c.cc.Invoke(ctx, Stats_BreakDown_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StatsServer is the server API for Stats service.
// All implementations must embed UnimplementedStatsServer
// for forward compatibility
type StatsServer interface {
	RealtimeVisitors(context.Context, *Realtime_Request) (*Realtime_Response, error)
	Aggregate(context.Context, *Aggregate_Request) (*Aggregate_Response, error)
	Timeseries(context.Context, *Timeseries_Request) (*Timeseries_Response, error)
	BreakDown(context.Context, *BreakDown_Request) (*BreakDown_Response, error)
	mustEmbedUnimplementedStatsServer()
}

// UnimplementedStatsServer must be embedded to have forward compatible implementations.
type UnimplementedStatsServer struct {
}

func (UnimplementedStatsServer) RealtimeVisitors(context.Context, *Realtime_Request) (*Realtime_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RealtimeVisitors not implemented")
}
func (UnimplementedStatsServer) Aggregate(context.Context, *Aggregate_Request) (*Aggregate_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Aggregate not implemented")
}
func (UnimplementedStatsServer) Timeseries(context.Context, *Timeseries_Request) (*Timeseries_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Timeseries not implemented")
}
func (UnimplementedStatsServer) BreakDown(context.Context, *BreakDown_Request) (*BreakDown_Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BreakDown not implemented")
}
func (UnimplementedStatsServer) mustEmbedUnimplementedStatsServer() {}

// UnsafeStatsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StatsServer will
// result in compilation errors.
type UnsafeStatsServer interface {
	mustEmbedUnimplementedStatsServer()
}

func RegisterStatsServer(s grpc.ServiceRegistrar, srv StatsServer) {
	s.RegisterService(&Stats_ServiceDesc, srv)
}

func _Stats_RealtimeVisitors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Realtime_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).RealtimeVisitors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stats_RealtimeVisitors_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).RealtimeVisitors(ctx, req.(*Realtime_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_Aggregate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Aggregate_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).Aggregate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stats_Aggregate_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).Aggregate(ctx, req.(*Aggregate_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_Timeseries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Timeseries_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).Timeseries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stats_Timeseries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).Timeseries(ctx, req.(*Timeseries_Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stats_BreakDown_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BreakDown_Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StatsServer).BreakDown(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stats_BreakDown_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StatsServer).BreakDown(ctx, req.(*BreakDown_Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Stats_ServiceDesc is the grpc.ServiceDesc for Stats service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stats_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "v1.Stats",
	HandlerType: (*StatsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RealtimeVisitors",
			Handler:    _Stats_RealtimeVisitors_Handler,
		},
		{
			MethodName: "Aggregate",
			Handler:    _Stats_Aggregate_Handler,
		},
		{
			MethodName: "Timeseries",
			Handler:    _Stats_Timeseries_Handler,
		},
		{
			MethodName: "BreakDown",
			Handler:    _Stats_BreakDown_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "staples/v1/stats.proto",
}
