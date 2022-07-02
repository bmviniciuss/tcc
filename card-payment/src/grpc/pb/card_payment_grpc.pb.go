// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: card_payment.proto

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

// CardPaymentClient is the client API for CardPayment service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardPaymentClient interface {
	ProccessCardPayment(ctx context.Context, in *ProcessCardPaymentInput, opts ...grpc.CallOption) (*Payment, error)
	GetPaymentsByClientId(ctx context.Context, in *GetPaymentsByClientIdRequest, opts ...grpc.CallOption) (*PaymentsResults, error)
}

type cardPaymentClient struct {
	cc grpc.ClientConnInterface
}

func NewCardPaymentClient(cc grpc.ClientConnInterface) CardPaymentClient {
	return &cardPaymentClient{cc}
}

func (c *cardPaymentClient) ProccessCardPayment(ctx context.Context, in *ProcessCardPaymentInput, opts ...grpc.CallOption) (*Payment, error) {
	out := new(Payment)
	err := c.cc.Invoke(ctx, "/cardpayment.CardPayment/ProccessCardPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardPaymentClient) GetPaymentsByClientId(ctx context.Context, in *GetPaymentsByClientIdRequest, opts ...grpc.CallOption) (*PaymentsResults, error) {
	out := new(PaymentsResults)
	err := c.cc.Invoke(ctx, "/cardpayment.CardPayment/GetPaymentsByClientId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardPaymentServer is the server API for CardPayment service.
// All implementations must embed UnimplementedCardPaymentServer
// for forward compatibility
type CardPaymentServer interface {
	ProccessCardPayment(context.Context, *ProcessCardPaymentInput) (*Payment, error)
	GetPaymentsByClientId(context.Context, *GetPaymentsByClientIdRequest) (*PaymentsResults, error)
	mustEmbedUnimplementedCardPaymentServer()
}

// UnimplementedCardPaymentServer must be embedded to have forward compatible implementations.
type UnimplementedCardPaymentServer struct {
}

func (UnimplementedCardPaymentServer) ProccessCardPayment(context.Context, *ProcessCardPaymentInput) (*Payment, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProccessCardPayment not implemented")
}
func (UnimplementedCardPaymentServer) GetPaymentsByClientId(context.Context, *GetPaymentsByClientIdRequest) (*PaymentsResults, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPaymentsByClientId not implemented")
}
func (UnimplementedCardPaymentServer) mustEmbedUnimplementedCardPaymentServer() {}

// UnsafeCardPaymentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardPaymentServer will
// result in compilation errors.
type UnsafeCardPaymentServer interface {
	mustEmbedUnimplementedCardPaymentServer()
}

func RegisterCardPaymentServer(s grpc.ServiceRegistrar, srv CardPaymentServer) {
	s.RegisterService(&CardPayment_ServiceDesc, srv)
}

func _CardPayment_ProccessCardPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProcessCardPaymentInput)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardPaymentServer).ProccessCardPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cardpayment.CardPayment/ProccessCardPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardPaymentServer).ProccessCardPayment(ctx, req.(*ProcessCardPaymentInput))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardPayment_GetPaymentsByClientId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPaymentsByClientIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardPaymentServer).GetPaymentsByClientId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cardpayment.CardPayment/GetPaymentsByClientId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardPaymentServer).GetPaymentsByClientId(ctx, req.(*GetPaymentsByClientIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CardPayment_ServiceDesc is the grpc.ServiceDesc for CardPayment service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CardPayment_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cardpayment.CardPayment",
	HandlerType: (*CardPaymentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProccessCardPayment",
			Handler:    _CardPayment_ProccessCardPayment_Handler,
		},
		{
			MethodName: "GetPaymentsByClientId",
			Handler:    _CardPayment_GetPaymentsByClientId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "card_payment.proto",
}
