// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: cards.proto

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

// CardsClient is the client API for Cards service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardsClient interface {
	GenerateCard(ctx context.Context, in *CreateCardRequest, opts ...grpc.CallOption) (*FullCard, error)
}

type cardsClient struct {
	cc grpc.ClientConnInterface
}

func NewCardsClient(cc grpc.ClientConnInterface) CardsClient {
	return &cardsClient{cc}
}

func (c *cardsClient) GenerateCard(ctx context.Context, in *CreateCardRequest, opts ...grpc.CallOption) (*FullCard, error) {
	out := new(FullCard)
	err := c.cc.Invoke(ctx, "/cards.Cards/GenerateCard", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardsServer is the server API for Cards service.
// All implementations must embed UnimplementedCardsServer
// for forward compatibility
type CardsServer interface {
	GenerateCard(context.Context, *CreateCardRequest) (*FullCard, error)
	mustEmbedUnimplementedCardsServer()
}

// UnimplementedCardsServer must be embedded to have forward compatible implementations.
type UnimplementedCardsServer struct {
}

func (UnimplementedCardsServer) GenerateCard(context.Context, *CreateCardRequest) (*FullCard, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateCard not implemented")
}
func (UnimplementedCardsServer) mustEmbedUnimplementedCardsServer() {}

// UnsafeCardsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardsServer will
// result in compilation errors.
type UnsafeCardsServer interface {
	mustEmbedUnimplementedCardsServer()
}

func RegisterCardsServer(s grpc.ServiceRegistrar, srv CardsServer) {
	s.RegisterService(&Cards_ServiceDesc, srv)
}

func _Cards_GenerateCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardsServer).GenerateCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cards.Cards/GenerateCard",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardsServer).GenerateCard(ctx, req.(*CreateCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Cards_ServiceDesc is the grpc.ServiceDesc for Cards service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Cards_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cards.Cards",
	HandlerType: (*CardsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateCard",
			Handler:    _Cards_GenerateCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cards.proto",
}