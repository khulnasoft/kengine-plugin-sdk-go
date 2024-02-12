// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.1
// source: backend.proto

package pluginv2

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
	Resource_CallResource_FullMethodName = "/pluginv2.Resource/CallResource"
)

// ResourceClient is the client API for Resource service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ResourceClient interface {
	CallResource(ctx context.Context, in *CallResourceRequest, opts ...grpc.CallOption) (Resource_CallResourceClient, error)
}

type resourceClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceClient(cc grpc.ClientConnInterface) ResourceClient {
	return &resourceClient{cc}
}

func (c *resourceClient) CallResource(ctx context.Context, in *CallResourceRequest, opts ...grpc.CallOption) (Resource_CallResourceClient, error) {
	stream, err := c.cc.NewStream(ctx, &Resource_ServiceDesc.Streams[0], Resource_CallResource_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &resourceCallResourceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Resource_CallResourceClient interface {
	Recv() (*CallResourceResponse, error)
	grpc.ClientStream
}

type resourceCallResourceClient struct {
	grpc.ClientStream
}

func (x *resourceCallResourceClient) Recv() (*CallResourceResponse, error) {
	m := new(CallResourceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ResourceServer is the server API for Resource service.
// All implementations should embed UnimplementedResourceServer
// for forward compatibility
type ResourceServer interface {
	CallResource(*CallResourceRequest, Resource_CallResourceServer) error
}

// UnimplementedResourceServer should be embedded to have forward compatible implementations.
type UnimplementedResourceServer struct {
}

func (UnimplementedResourceServer) CallResource(*CallResourceRequest, Resource_CallResourceServer) error {
	return status.Errorf(codes.Unimplemented, "method CallResource not implemented")
}

// UnsafeResourceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceServer will
// result in compilation errors.
type UnsafeResourceServer interface {
	mustEmbedUnimplementedResourceServer()
}

func RegisterResourceServer(s grpc.ServiceRegistrar, srv ResourceServer) {
	s.RegisterService(&Resource_ServiceDesc, srv)
}

func _Resource_CallResource_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(CallResourceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResourceServer).CallResource(m, &resourceCallResourceServer{stream})
}

type Resource_CallResourceServer interface {
	Send(*CallResourceResponse) error
	grpc.ServerStream
}

type resourceCallResourceServer struct {
	grpc.ServerStream
}

func (x *resourceCallResourceServer) Send(m *CallResourceResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Resource_ServiceDesc is the grpc.ServiceDesc for Resource service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Resource_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pluginv2.Resource",
	HandlerType: (*ResourceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "CallResource",
			Handler:       _Resource_CallResource_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "backend.proto",
}

const (
	Data_QueryData_FullMethodName = "/pluginv2.Data/QueryData"
)

// DataClient is the client API for Data service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataClient interface {
	QueryData(ctx context.Context, in *QueryDataRequest, opts ...grpc.CallOption) (*QueryDataResponse, error)
}

type dataClient struct {
	cc grpc.ClientConnInterface
}

func NewDataClient(cc grpc.ClientConnInterface) DataClient {
	return &dataClient{cc}
}

func (c *dataClient) QueryData(ctx context.Context, in *QueryDataRequest, opts ...grpc.CallOption) (*QueryDataResponse, error) {
	out := new(QueryDataResponse)
	err := c.cc.Invoke(ctx, Data_QueryData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataServer is the server API for Data service.
// All implementations should embed UnimplementedDataServer
// for forward compatibility
type DataServer interface {
	QueryData(context.Context, *QueryDataRequest) (*QueryDataResponse, error)
}

// UnimplementedDataServer should be embedded to have forward compatible implementations.
type UnimplementedDataServer struct {
}

func (UnimplementedDataServer) QueryData(context.Context, *QueryDataRequest) (*QueryDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryData not implemented")
}

// UnsafeDataServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataServer will
// result in compilation errors.
type UnsafeDataServer interface {
	mustEmbedUnimplementedDataServer()
}

func RegisterDataServer(s grpc.ServiceRegistrar, srv DataServer) {
	s.RegisterService(&Data_ServiceDesc, srv)
}

func _Data_QueryData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataServer).QueryData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Data_QueryData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataServer).QueryData(ctx, req.(*QueryDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Data_ServiceDesc is the grpc.ServiceDesc for Data service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Data_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pluginv2.Data",
	HandlerType: (*DataServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryData",
			Handler:    _Data_QueryData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend.proto",
}

const (
	Diagnostics_CheckHealth_FullMethodName    = "/pluginv2.Diagnostics/CheckHealth"
	Diagnostics_CollectMetrics_FullMethodName = "/pluginv2.Diagnostics/CollectMetrics"
)

// DiagnosticsClient is the client API for Diagnostics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DiagnosticsClient interface {
	CheckHealth(ctx context.Context, in *CheckHealthRequest, opts ...grpc.CallOption) (*CheckHealthResponse, error)
	CollectMetrics(ctx context.Context, in *CollectMetricsRequest, opts ...grpc.CallOption) (*CollectMetricsResponse, error)
}

type diagnosticsClient struct {
	cc grpc.ClientConnInterface
}

func NewDiagnosticsClient(cc grpc.ClientConnInterface) DiagnosticsClient {
	return &diagnosticsClient{cc}
}

func (c *diagnosticsClient) CheckHealth(ctx context.Context, in *CheckHealthRequest, opts ...grpc.CallOption) (*CheckHealthResponse, error) {
	out := new(CheckHealthResponse)
	err := c.cc.Invoke(ctx, Diagnostics_CheckHealth_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *diagnosticsClient) CollectMetrics(ctx context.Context, in *CollectMetricsRequest, opts ...grpc.CallOption) (*CollectMetricsResponse, error) {
	out := new(CollectMetricsResponse)
	err := c.cc.Invoke(ctx, Diagnostics_CollectMetrics_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiagnosticsServer is the server API for Diagnostics service.
// All implementations should embed UnimplementedDiagnosticsServer
// for forward compatibility
type DiagnosticsServer interface {
	CheckHealth(context.Context, *CheckHealthRequest) (*CheckHealthResponse, error)
	CollectMetrics(context.Context, *CollectMetricsRequest) (*CollectMetricsResponse, error)
}

// UnimplementedDiagnosticsServer should be embedded to have forward compatible implementations.
type UnimplementedDiagnosticsServer struct {
}

func (UnimplementedDiagnosticsServer) CheckHealth(context.Context, *CheckHealthRequest) (*CheckHealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckHealth not implemented")
}
func (UnimplementedDiagnosticsServer) CollectMetrics(context.Context, *CollectMetricsRequest) (*CollectMetricsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectMetrics not implemented")
}

// UnsafeDiagnosticsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiagnosticsServer will
// result in compilation errors.
type UnsafeDiagnosticsServer interface {
	mustEmbedUnimplementedDiagnosticsServer()
}

func RegisterDiagnosticsServer(s grpc.ServiceRegistrar, srv DiagnosticsServer) {
	s.RegisterService(&Diagnostics_ServiceDesc, srv)
}

func _Diagnostics_CheckHealth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckHealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiagnosticsServer).CheckHealth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Diagnostics_CheckHealth_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiagnosticsServer).CheckHealth(ctx, req.(*CheckHealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Diagnostics_CollectMetrics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectMetricsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiagnosticsServer).CollectMetrics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Diagnostics_CollectMetrics_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiagnosticsServer).CollectMetrics(ctx, req.(*CollectMetricsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Diagnostics_ServiceDesc is the grpc.ServiceDesc for Diagnostics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Diagnostics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pluginv2.Diagnostics",
	HandlerType: (*DiagnosticsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CheckHealth",
			Handler:    _Diagnostics_CheckHealth_Handler,
		},
		{
			MethodName: "CollectMetrics",
			Handler:    _Diagnostics_CollectMetrics_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "backend.proto",
}

const (
	Stream_SubscribeStream_FullMethodName = "/pluginv2.Stream/SubscribeStream"
	Stream_RunStream_FullMethodName       = "/pluginv2.Stream/RunStream"
	Stream_PublishStream_FullMethodName   = "/pluginv2.Stream/PublishStream"
)

// StreamClient is the client API for Stream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamClient interface {
	// SubscribeStream called when a user tries to subscribe to a plugin/datasource
	// managed channel path – thus plugin can check subscribe permissions and communicate
	// options with Khulnasoft Core. When the first subscriber joins a channel, RunStream
	// will be called.
	SubscribeStream(ctx context.Context, in *SubscribeStreamRequest, opts ...grpc.CallOption) (*SubscribeStreamResponse, error)
	// RunStream will be initiated by Khulnasoft to consume a stream. RunStream will be
	// called once for the first client successfully subscribed to a channel path.
	// When Khulnasoft detects that there are no longer any subscribers inside a channel,
	// the call will be terminated until next active subscriber appears. Call termination
	// can happen with a delay.
	RunStream(ctx context.Context, in *RunStreamRequest, opts ...grpc.CallOption) (Stream_RunStreamClient, error)
	// PublishStream called when a user tries to publish to a plugin/datasource
	// managed channel path. Here plugin can check publish permissions and
	// modify publication data if required.
	PublishStream(ctx context.Context, in *PublishStreamRequest, opts ...grpc.CallOption) (*PublishStreamResponse, error)
}

type streamClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamClient(cc grpc.ClientConnInterface) StreamClient {
	return &streamClient{cc}
}

func (c *streamClient) SubscribeStream(ctx context.Context, in *SubscribeStreamRequest, opts ...grpc.CallOption) (*SubscribeStreamResponse, error) {
	out := new(SubscribeStreamResponse)
	err := c.cc.Invoke(ctx, Stream_SubscribeStream_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamClient) RunStream(ctx context.Context, in *RunStreamRequest, opts ...grpc.CallOption) (Stream_RunStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &Stream_ServiceDesc.Streams[0], Stream_RunStream_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &streamRunStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Stream_RunStreamClient interface {
	Recv() (*StreamPacket, error)
	grpc.ClientStream
}

type streamRunStreamClient struct {
	grpc.ClientStream
}

func (x *streamRunStreamClient) Recv() (*StreamPacket, error) {
	m := new(StreamPacket)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamClient) PublishStream(ctx context.Context, in *PublishStreamRequest, opts ...grpc.CallOption) (*PublishStreamResponse, error) {
	out := new(PublishStreamResponse)
	err := c.cc.Invoke(ctx, Stream_PublishStream_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamServer is the server API for Stream service.
// All implementations should embed UnimplementedStreamServer
// for forward compatibility
type StreamServer interface {
	// SubscribeStream called when a user tries to subscribe to a plugin/datasource
	// managed channel path – thus plugin can check subscribe permissions and communicate
	// options with Khulnasoft Core. When the first subscriber joins a channel, RunStream
	// will be called.
	SubscribeStream(context.Context, *SubscribeStreamRequest) (*SubscribeStreamResponse, error)
	// RunStream will be initiated by Khulnasoft to consume a stream. RunStream will be
	// called once for the first client successfully subscribed to a channel path.
	// When Khulnasoft detects that there are no longer any subscribers inside a channel,
	// the call will be terminated until next active subscriber appears. Call termination
	// can happen with a delay.
	RunStream(*RunStreamRequest, Stream_RunStreamServer) error
	// PublishStream called when a user tries to publish to a plugin/datasource
	// managed channel path. Here plugin can check publish permissions and
	// modify publication data if required.
	PublishStream(context.Context, *PublishStreamRequest) (*PublishStreamResponse, error)
}

// UnimplementedStreamServer should be embedded to have forward compatible implementations.
type UnimplementedStreamServer struct {
}

func (UnimplementedStreamServer) SubscribeStream(context.Context, *SubscribeStreamRequest) (*SubscribeStreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubscribeStream not implemented")
}
func (UnimplementedStreamServer) RunStream(*RunStreamRequest, Stream_RunStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method RunStream not implemented")
}
func (UnimplementedStreamServer) PublishStream(context.Context, *PublishStreamRequest) (*PublishStreamResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PublishStream not implemented")
}

// UnsafeStreamServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamServer will
// result in compilation errors.
type UnsafeStreamServer interface {
	mustEmbedUnimplementedStreamServer()
}

func RegisterStreamServer(s grpc.ServiceRegistrar, srv StreamServer) {
	s.RegisterService(&Stream_ServiceDesc, srv)
}

func _Stream_SubscribeStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubscribeStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServer).SubscribeStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stream_SubscribeStream_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServer).SubscribeStream(ctx, req.(*SubscribeStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Stream_RunStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(RunStreamRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(StreamServer).RunStream(m, &streamRunStreamServer{stream})
}

type Stream_RunStreamServer interface {
	Send(*StreamPacket) error
	grpc.ServerStream
}

type streamRunStreamServer struct {
	grpc.ServerStream
}

func (x *streamRunStreamServer) Send(m *StreamPacket) error {
	return x.ServerStream.SendMsg(m)
}

func _Stream_PublishStream_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PublishStreamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamServer).PublishStream(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Stream_PublishStream_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamServer).PublishStream(ctx, req.(*PublishStreamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Stream_ServiceDesc is the grpc.ServiceDesc for Stream service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Stream_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pluginv2.Stream",
	HandlerType: (*StreamServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubscribeStream",
			Handler:    _Stream_SubscribeStream_Handler,
		},
		{
			MethodName: "PublishStream",
			Handler:    _Stream_PublishStream_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RunStream",
			Handler:       _Stream_RunStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "backend.proto",
}
