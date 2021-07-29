// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package looprpc

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

// SwapServerClient is the client API for SwapServer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SwapServerClient interface {
	LoopOutTerms(ctx context.Context, in *ServerLoopOutTermsRequest, opts ...grpc.CallOption) (*ServerLoopOutTerms, error)
	NewLoopOutSwap(ctx context.Context, in *ServerLoopOutRequest, opts ...grpc.CallOption) (*ServerLoopOutResponse, error)
	LoopOutPushPreimage(ctx context.Context, in *ServerLoopOutPushPreimageRequest, opts ...grpc.CallOption) (*ServerLoopOutPushPreimageResponse, error)
	LoopOutQuote(ctx context.Context, in *ServerLoopOutQuoteRequest, opts ...grpc.CallOption) (*ServerLoopOutQuote, error)
	LoopInTerms(ctx context.Context, in *ServerLoopInTermsRequest, opts ...grpc.CallOption) (*ServerLoopInTerms, error)
	NewLoopInSwap(ctx context.Context, in *ServerLoopInRequest, opts ...grpc.CallOption) (*ServerLoopInResponse, error)
	LoopInQuote(ctx context.Context, in *ServerLoopInQuoteRequest, opts ...grpc.CallOption) (*ServerLoopInQuoteResponse, error)
	SubscribeLoopOutUpdates(ctx context.Context, in *SubscribeUpdatesRequest, opts ...grpc.CallOption) (SwapServer_SubscribeLoopOutUpdatesClient, error)
	SubscribeLoopInUpdates(ctx context.Context, in *SubscribeUpdatesRequest, opts ...grpc.CallOption) (SwapServer_SubscribeLoopInUpdatesClient, error)
	CancelLoopOutSwap(ctx context.Context, in *CancelLoopOutSwapRequest, opts ...grpc.CallOption) (*CancelLoopOutSwapResponse, error)
}

type swapServerClient struct {
	cc grpc.ClientConnInterface
}

func NewSwapServerClient(cc grpc.ClientConnInterface) SwapServerClient {
	return &swapServerClient{cc}
}

func (c *swapServerClient) LoopOutTerms(ctx context.Context, in *ServerLoopOutTermsRequest, opts ...grpc.CallOption) (*ServerLoopOutTerms, error) {
	out := new(ServerLoopOutTerms)
	err := c.cc.Invoke(ctx, "/looprpc.SwapServer/LoopOutTerms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapServerClient) NewLoopOutSwap(ctx context.Context, in *ServerLoopOutRequest, opts ...grpc.CallOption) (*ServerLoopOutResponse, error) {
	out := new(ServerLoopOutResponse)
	err := c.cc.Invoke(ctx, "/looprpc.SwapServer/NewLoopOutSwap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapServerClient) LoopOutPushPreimage(ctx context.Context, in *ServerLoopOutPushPreimageRequest, opts ...grpc.CallOption) (*ServerLoopOutPushPreimageResponse, error) {
	out := new(ServerLoopOutPushPreimageResponse)
	err := c.cc.Invoke(ctx, "/looprpc.SwapServer/LoopOutPushPreimage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapServerClient) LoopOutQuote(ctx context.Context, in *ServerLoopOutQuoteRequest, opts ...grpc.CallOption) (*ServerLoopOutQuote, error) {
	out := new(ServerLoopOutQuote)
	err := c.cc.Invoke(ctx, "/looprpc.SwapServer/LoopOutQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapServerClient) LoopInTerms(ctx context.Context, in *ServerLoopInTermsRequest, opts ...grpc.CallOption) (*ServerLoopInTerms, error) {
	out := new(ServerLoopInTerms)
	err := c.cc.Invoke(ctx, "/looprpc.SwapServer/LoopInTerms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapServerClient) NewLoopInSwap(ctx context.Context, in *ServerLoopInRequest, opts ...grpc.CallOption) (*ServerLoopInResponse, error) {
	out := new(ServerLoopInResponse)
	err := c.cc.Invoke(ctx, "/looprpc.SwapServer/NewLoopInSwap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapServerClient) LoopInQuote(ctx context.Context, in *ServerLoopInQuoteRequest, opts ...grpc.CallOption) (*ServerLoopInQuoteResponse, error) {
	out := new(ServerLoopInQuoteResponse)
	err := c.cc.Invoke(ctx, "/looprpc.SwapServer/LoopInQuote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *swapServerClient) SubscribeLoopOutUpdates(ctx context.Context, in *SubscribeUpdatesRequest, opts ...grpc.CallOption) (SwapServer_SubscribeLoopOutUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &SwapServer_ServiceDesc.Streams[0], "/looprpc.SwapServer/SubscribeLoopOutUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &swapServerSubscribeLoopOutUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SwapServer_SubscribeLoopOutUpdatesClient interface {
	Recv() (*SubscribeLoopOutUpdatesResponse, error)
	grpc.ClientStream
}

type swapServerSubscribeLoopOutUpdatesClient struct {
	grpc.ClientStream
}

func (x *swapServerSubscribeLoopOutUpdatesClient) Recv() (*SubscribeLoopOutUpdatesResponse, error) {
	m := new(SubscribeLoopOutUpdatesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *swapServerClient) SubscribeLoopInUpdates(ctx context.Context, in *SubscribeUpdatesRequest, opts ...grpc.CallOption) (SwapServer_SubscribeLoopInUpdatesClient, error) {
	stream, err := c.cc.NewStream(ctx, &SwapServer_ServiceDesc.Streams[1], "/looprpc.SwapServer/SubscribeLoopInUpdates", opts...)
	if err != nil {
		return nil, err
	}
	x := &swapServerSubscribeLoopInUpdatesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SwapServer_SubscribeLoopInUpdatesClient interface {
	Recv() (*SubscribeLoopInUpdatesResponse, error)
	grpc.ClientStream
}

type swapServerSubscribeLoopInUpdatesClient struct {
	grpc.ClientStream
}

func (x *swapServerSubscribeLoopInUpdatesClient) Recv() (*SubscribeLoopInUpdatesResponse, error) {
	m := new(SubscribeLoopInUpdatesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *swapServerClient) CancelLoopOutSwap(ctx context.Context, in *CancelLoopOutSwapRequest, opts ...grpc.CallOption) (*CancelLoopOutSwapResponse, error) {
	out := new(CancelLoopOutSwapResponse)
	err := c.cc.Invoke(ctx, "/looprpc.SwapServer/CancelLoopOutSwap", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SwapServerServer is the server API for SwapServer service.
// All implementations must embed UnimplementedSwapServerServer
// for forward compatibility
type SwapServerServer interface {
	LoopOutTerms(context.Context, *ServerLoopOutTermsRequest) (*ServerLoopOutTerms, error)
	NewLoopOutSwap(context.Context, *ServerLoopOutRequest) (*ServerLoopOutResponse, error)
	LoopOutPushPreimage(context.Context, *ServerLoopOutPushPreimageRequest) (*ServerLoopOutPushPreimageResponse, error)
	LoopOutQuote(context.Context, *ServerLoopOutQuoteRequest) (*ServerLoopOutQuote, error)
	LoopInTerms(context.Context, *ServerLoopInTermsRequest) (*ServerLoopInTerms, error)
	NewLoopInSwap(context.Context, *ServerLoopInRequest) (*ServerLoopInResponse, error)
	LoopInQuote(context.Context, *ServerLoopInQuoteRequest) (*ServerLoopInQuoteResponse, error)
	SubscribeLoopOutUpdates(*SubscribeUpdatesRequest, SwapServer_SubscribeLoopOutUpdatesServer) error
	SubscribeLoopInUpdates(*SubscribeUpdatesRequest, SwapServer_SubscribeLoopInUpdatesServer) error
	CancelLoopOutSwap(context.Context, *CancelLoopOutSwapRequest) (*CancelLoopOutSwapResponse, error)
	mustEmbedUnimplementedSwapServerServer()
}

// UnimplementedSwapServerServer must be embedded to have forward compatible implementations.
type UnimplementedSwapServerServer struct {
}

func (UnimplementedSwapServerServer) LoopOutTerms(context.Context, *ServerLoopOutTermsRequest) (*ServerLoopOutTerms, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoopOutTerms not implemented")
}
func (UnimplementedSwapServerServer) NewLoopOutSwap(context.Context, *ServerLoopOutRequest) (*ServerLoopOutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewLoopOutSwap not implemented")
}
func (UnimplementedSwapServerServer) LoopOutPushPreimage(context.Context, *ServerLoopOutPushPreimageRequest) (*ServerLoopOutPushPreimageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoopOutPushPreimage not implemented")
}
func (UnimplementedSwapServerServer) LoopOutQuote(context.Context, *ServerLoopOutQuoteRequest) (*ServerLoopOutQuote, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoopOutQuote not implemented")
}
func (UnimplementedSwapServerServer) LoopInTerms(context.Context, *ServerLoopInTermsRequest) (*ServerLoopInTerms, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoopInTerms not implemented")
}
func (UnimplementedSwapServerServer) NewLoopInSwap(context.Context, *ServerLoopInRequest) (*ServerLoopInResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewLoopInSwap not implemented")
}
func (UnimplementedSwapServerServer) LoopInQuote(context.Context, *ServerLoopInQuoteRequest) (*ServerLoopInQuoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoopInQuote not implemented")
}
func (UnimplementedSwapServerServer) SubscribeLoopOutUpdates(*SubscribeUpdatesRequest, SwapServer_SubscribeLoopOutUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeLoopOutUpdates not implemented")
}
func (UnimplementedSwapServerServer) SubscribeLoopInUpdates(*SubscribeUpdatesRequest, SwapServer_SubscribeLoopInUpdatesServer) error {
	return status.Errorf(codes.Unimplemented, "method SubscribeLoopInUpdates not implemented")
}
func (UnimplementedSwapServerServer) CancelLoopOutSwap(context.Context, *CancelLoopOutSwapRequest) (*CancelLoopOutSwapResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelLoopOutSwap not implemented")
}
func (UnimplementedSwapServerServer) mustEmbedUnimplementedSwapServerServer() {}

// UnsafeSwapServerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SwapServerServer will
// result in compilation errors.
type UnsafeSwapServerServer interface {
	mustEmbedUnimplementedSwapServerServer()
}

func RegisterSwapServerServer(s grpc.ServiceRegistrar, srv SwapServerServer) {
	s.RegisterService(&SwapServer_ServiceDesc, srv)
}

func _SwapServer_LoopOutTerms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerLoopOutTermsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapServerServer).LoopOutTerms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/looprpc.SwapServer/LoopOutTerms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapServerServer).LoopOutTerms(ctx, req.(*ServerLoopOutTermsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapServer_NewLoopOutSwap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerLoopOutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapServerServer).NewLoopOutSwap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/looprpc.SwapServer/NewLoopOutSwap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapServerServer).NewLoopOutSwap(ctx, req.(*ServerLoopOutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapServer_LoopOutPushPreimage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerLoopOutPushPreimageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapServerServer).LoopOutPushPreimage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/looprpc.SwapServer/LoopOutPushPreimage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapServerServer).LoopOutPushPreimage(ctx, req.(*ServerLoopOutPushPreimageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapServer_LoopOutQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerLoopOutQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapServerServer).LoopOutQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/looprpc.SwapServer/LoopOutQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapServerServer).LoopOutQuote(ctx, req.(*ServerLoopOutQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapServer_LoopInTerms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerLoopInTermsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapServerServer).LoopInTerms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/looprpc.SwapServer/LoopInTerms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapServerServer).LoopInTerms(ctx, req.(*ServerLoopInTermsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapServer_NewLoopInSwap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerLoopInRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapServerServer).NewLoopInSwap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/looprpc.SwapServer/NewLoopInSwap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapServerServer).NewLoopInSwap(ctx, req.(*ServerLoopInRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapServer_LoopInQuote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerLoopInQuoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapServerServer).LoopInQuote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/looprpc.SwapServer/LoopInQuote",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapServerServer).LoopInQuote(ctx, req.(*ServerLoopInQuoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SwapServer_SubscribeLoopOutUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeUpdatesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SwapServerServer).SubscribeLoopOutUpdates(m, &swapServerSubscribeLoopOutUpdatesServer{stream})
}

type SwapServer_SubscribeLoopOutUpdatesServer interface {
	Send(*SubscribeLoopOutUpdatesResponse) error
	grpc.ServerStream
}

type swapServerSubscribeLoopOutUpdatesServer struct {
	grpc.ServerStream
}

func (x *swapServerSubscribeLoopOutUpdatesServer) Send(m *SubscribeLoopOutUpdatesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SwapServer_SubscribeLoopInUpdates_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(SubscribeUpdatesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SwapServerServer).SubscribeLoopInUpdates(m, &swapServerSubscribeLoopInUpdatesServer{stream})
}

type SwapServer_SubscribeLoopInUpdatesServer interface {
	Send(*SubscribeLoopInUpdatesResponse) error
	grpc.ServerStream
}

type swapServerSubscribeLoopInUpdatesServer struct {
	grpc.ServerStream
}

func (x *swapServerSubscribeLoopInUpdatesServer) Send(m *SubscribeLoopInUpdatesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _SwapServer_CancelLoopOutSwap_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelLoopOutSwapRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SwapServerServer).CancelLoopOutSwap(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/looprpc.SwapServer/CancelLoopOutSwap",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SwapServerServer).CancelLoopOutSwap(ctx, req.(*CancelLoopOutSwapRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SwapServer_ServiceDesc is the grpc.ServiceDesc for SwapServer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SwapServer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "looprpc.SwapServer",
	HandlerType: (*SwapServerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "LoopOutTerms",
			Handler:    _SwapServer_LoopOutTerms_Handler,
		},
		{
			MethodName: "NewLoopOutSwap",
			Handler:    _SwapServer_NewLoopOutSwap_Handler,
		},
		{
			MethodName: "LoopOutPushPreimage",
			Handler:    _SwapServer_LoopOutPushPreimage_Handler,
		},
		{
			MethodName: "LoopOutQuote",
			Handler:    _SwapServer_LoopOutQuote_Handler,
		},
		{
			MethodName: "LoopInTerms",
			Handler:    _SwapServer_LoopInTerms_Handler,
		},
		{
			MethodName: "NewLoopInSwap",
			Handler:    _SwapServer_NewLoopInSwap_Handler,
		},
		{
			MethodName: "LoopInQuote",
			Handler:    _SwapServer_LoopInQuote_Handler,
		},
		{
			MethodName: "CancelLoopOutSwap",
			Handler:    _SwapServer_CancelLoopOutSwap_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SubscribeLoopOutUpdates",
			Handler:       _SwapServer_SubscribeLoopOutUpdates_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SubscribeLoopInUpdates",
			Handler:       _SwapServer_SubscribeLoopInUpdates_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "server.proto",
}
