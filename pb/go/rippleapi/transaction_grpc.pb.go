// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.19.3
// source: transaction.proto

package rippleapi

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// RippleTransactionAPIClient is the client API for RippleTransactionAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RippleTransactionAPIClient interface {
	// https://xrpl.org/rippleapi-reference.html#preparetransaction
	PrepareTransaction(ctx context.Context, in *RequestPrepareTransaction, opts ...grpc.CallOption) (*ResponsePrepareTransaction, error)
	SignTransaction(ctx context.Context, in *RequestSignTransaction, opts ...grpc.CallOption) (*ResponseSignTransaction, error)
	SubmitTransaction(ctx context.Context, in *RequestSubmitTransaction, opts ...grpc.CallOption) (*ResponseSubmitTransaction, error)
	WaitValidation(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RippleTransactionAPI_WaitValidationClient, error)
	GetTransaction(ctx context.Context, in *RequestGetTransaction, opts ...grpc.CallOption) (*ResponseGetTransaction, error)
	CombineTransaction(ctx context.Context, in *RequestCombineTransaction, opts ...grpc.CallOption) (*ResponseCombineTransaction, error)
}

type rippleTransactionAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewRippleTransactionAPIClient(cc grpc.ClientConnInterface) RippleTransactionAPIClient {
	return &rippleTransactionAPIClient{cc}
}

func (c *rippleTransactionAPIClient) PrepareTransaction(ctx context.Context, in *RequestPrepareTransaction, opts ...grpc.CallOption) (*ResponsePrepareTransaction, error) {
	out := new(ResponsePrepareTransaction)
	err := c.cc.Invoke(ctx, "/rippleapi.transaction.RippleTransactionAPI/PrepareTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rippleTransactionAPIClient) SignTransaction(ctx context.Context, in *RequestSignTransaction, opts ...grpc.CallOption) (*ResponseSignTransaction, error) {
	out := new(ResponseSignTransaction)
	err := c.cc.Invoke(ctx, "/rippleapi.transaction.RippleTransactionAPI/SignTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rippleTransactionAPIClient) SubmitTransaction(ctx context.Context, in *RequestSubmitTransaction, opts ...grpc.CallOption) (*ResponseSubmitTransaction, error) {
	out := new(ResponseSubmitTransaction)
	err := c.cc.Invoke(ctx, "/rippleapi.transaction.RippleTransactionAPI/SubmitTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rippleTransactionAPIClient) WaitValidation(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (RippleTransactionAPI_WaitValidationClient, error) {
	stream, err := c.cc.NewStream(ctx, &RippleTransactionAPI_ServiceDesc.Streams[0], "/rippleapi.transaction.RippleTransactionAPI/WaitValidation", opts...)
	if err != nil {
		return nil, err
	}
	x := &rippleTransactionAPIWaitValidationClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type RippleTransactionAPI_WaitValidationClient interface {
	Recv() (*ResponseWaitValidation, error)
	grpc.ClientStream
}

type rippleTransactionAPIWaitValidationClient struct {
	grpc.ClientStream
}

func (x *rippleTransactionAPIWaitValidationClient) Recv() (*ResponseWaitValidation, error) {
	m := new(ResponseWaitValidation)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *rippleTransactionAPIClient) GetTransaction(ctx context.Context, in *RequestGetTransaction, opts ...grpc.CallOption) (*ResponseGetTransaction, error) {
	out := new(ResponseGetTransaction)
	err := c.cc.Invoke(ctx, "/rippleapi.transaction.RippleTransactionAPI/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *rippleTransactionAPIClient) CombineTransaction(ctx context.Context, in *RequestCombineTransaction, opts ...grpc.CallOption) (*ResponseCombineTransaction, error) {
	out := new(ResponseCombineTransaction)
	err := c.cc.Invoke(ctx, "/rippleapi.transaction.RippleTransactionAPI/CombineTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RippleTransactionAPIServer is the server API for RippleTransactionAPI service.
// All implementations must embed UnimplementedRippleTransactionAPIServer
// for forward compatibility
type RippleTransactionAPIServer interface {
	// https://xrpl.org/rippleapi-reference.html#preparetransaction
	PrepareTransaction(context.Context, *RequestPrepareTransaction) (*ResponsePrepareTransaction, error)
	SignTransaction(context.Context, *RequestSignTransaction) (*ResponseSignTransaction, error)
	SubmitTransaction(context.Context, *RequestSubmitTransaction) (*ResponseSubmitTransaction, error)
	WaitValidation(*emptypb.Empty, RippleTransactionAPI_WaitValidationServer) error
	GetTransaction(context.Context, *RequestGetTransaction) (*ResponseGetTransaction, error)
	CombineTransaction(context.Context, *RequestCombineTransaction) (*ResponseCombineTransaction, error)
	mustEmbedUnimplementedRippleTransactionAPIServer()
}

// UnimplementedRippleTransactionAPIServer must be embedded to have forward compatible implementations.
type UnimplementedRippleTransactionAPIServer struct {
}

func (UnimplementedRippleTransactionAPIServer) PrepareTransaction(context.Context, *RequestPrepareTransaction) (*ResponsePrepareTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrepareTransaction not implemented")
}
func (UnimplementedRippleTransactionAPIServer) SignTransaction(context.Context, *RequestSignTransaction) (*ResponseSignTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignTransaction not implemented")
}
func (UnimplementedRippleTransactionAPIServer) SubmitTransaction(context.Context, *RequestSubmitTransaction) (*ResponseSubmitTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitTransaction not implemented")
}
func (UnimplementedRippleTransactionAPIServer) WaitValidation(*emptypb.Empty, RippleTransactionAPI_WaitValidationServer) error {
	return status.Errorf(codes.Unimplemented, "method WaitValidation not implemented")
}
func (UnimplementedRippleTransactionAPIServer) GetTransaction(context.Context, *RequestGetTransaction) (*ResponseGetTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedRippleTransactionAPIServer) CombineTransaction(context.Context, *RequestCombineTransaction) (*ResponseCombineTransaction, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CombineTransaction not implemented")
}
func (UnimplementedRippleTransactionAPIServer) mustEmbedUnimplementedRippleTransactionAPIServer() {}

// UnsafeRippleTransactionAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RippleTransactionAPIServer will
// result in compilation errors.
type UnsafeRippleTransactionAPIServer interface {
	mustEmbedUnimplementedRippleTransactionAPIServer()
}

func RegisterRippleTransactionAPIServer(s grpc.ServiceRegistrar, srv RippleTransactionAPIServer) {
	s.RegisterService(&RippleTransactionAPI_ServiceDesc, srv)
}

func _RippleTransactionAPI_PrepareTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPrepareTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RippleTransactionAPIServer).PrepareTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rippleapi.transaction.RippleTransactionAPI/PrepareTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RippleTransactionAPIServer).PrepareTransaction(ctx, req.(*RequestPrepareTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _RippleTransactionAPI_SignTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSignTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RippleTransactionAPIServer).SignTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rippleapi.transaction.RippleTransactionAPI/SignTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RippleTransactionAPIServer).SignTransaction(ctx, req.(*RequestSignTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _RippleTransactionAPI_SubmitTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestSubmitTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RippleTransactionAPIServer).SubmitTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rippleapi.transaction.RippleTransactionAPI/SubmitTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RippleTransactionAPIServer).SubmitTransaction(ctx, req.(*RequestSubmitTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _RippleTransactionAPI_WaitValidation_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(RippleTransactionAPIServer).WaitValidation(m, &rippleTransactionAPIWaitValidationServer{stream})
}

type RippleTransactionAPI_WaitValidationServer interface {
	Send(*ResponseWaitValidation) error
	grpc.ServerStream
}

type rippleTransactionAPIWaitValidationServer struct {
	grpc.ServerStream
}

func (x *rippleTransactionAPIWaitValidationServer) Send(m *ResponseWaitValidation) error {
	return x.ServerStream.SendMsg(m)
}

func _RippleTransactionAPI_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestGetTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RippleTransactionAPIServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rippleapi.transaction.RippleTransactionAPI/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RippleTransactionAPIServer).GetTransaction(ctx, req.(*RequestGetTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _RippleTransactionAPI_CombineTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestCombineTransaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RippleTransactionAPIServer).CombineTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rippleapi.transaction.RippleTransactionAPI/CombineTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RippleTransactionAPIServer).CombineTransaction(ctx, req.(*RequestCombineTransaction))
	}
	return interceptor(ctx, in, info, handler)
}

// RippleTransactionAPI_ServiceDesc is the grpc.ServiceDesc for RippleTransactionAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RippleTransactionAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rippleapi.transaction.RippleTransactionAPI",
	HandlerType: (*RippleTransactionAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrepareTransaction",
			Handler:    _RippleTransactionAPI_PrepareTransaction_Handler,
		},
		{
			MethodName: "SignTransaction",
			Handler:    _RippleTransactionAPI_SignTransaction_Handler,
		},
		{
			MethodName: "SubmitTransaction",
			Handler:    _RippleTransactionAPI_SubmitTransaction_Handler,
		},
		{
			MethodName: "GetTransaction",
			Handler:    _RippleTransactionAPI_GetTransaction_Handler,
		},
		{
			MethodName: "CombineTransaction",
			Handler:    _RippleTransactionAPI_CombineTransaction_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WaitValidation",
			Handler:       _RippleTransactionAPI_WaitValidation_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "transaction.proto",
}
