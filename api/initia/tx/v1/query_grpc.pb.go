// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: initia/tx/v1/query.proto

package txv1

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
	Query_GasPrices_FullMethodName = "/initia.tx.v1.Query/GasPrices"
	Query_GasPrice_FullMethodName  = "/initia.tx.v1.Query/GasPrice"
)

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// GasPrices returns the gas prices for the network.
	GasPrices(ctx context.Context, in *QueryGasPricesRequest, opts ...grpc.CallOption) (*QueryGasPricesResponse, error)
	// GasPrice returns the gas price for the network.
	GasPrice(ctx context.Context, in *QueryGasPriceRequest, opts ...grpc.CallOption) (*QueryGasPriceResponse, error)
}

type queryClient struct {
	cc grpc.ClientConnInterface
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GasPrices(ctx context.Context, in *QueryGasPricesRequest, opts ...grpc.CallOption) (*QueryGasPricesResponse, error) {
	out := new(QueryGasPricesResponse)
	err := c.cc.Invoke(ctx, Query_GasPrices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GasPrice(ctx context.Context, in *QueryGasPriceRequest, opts ...grpc.CallOption) (*QueryGasPriceResponse, error) {
	out := new(QueryGasPriceResponse)
	err := c.cc.Invoke(ctx, Query_GasPrice_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
// All implementations must embed UnimplementedQueryServer
// for forward compatibility
type QueryServer interface {
	// GasPrices returns the gas prices for the network.
	GasPrices(context.Context, *QueryGasPricesRequest) (*QueryGasPricesResponse, error)
	// GasPrice returns the gas price for the network.
	GasPrice(context.Context, *QueryGasPriceRequest) (*QueryGasPriceResponse, error)
	mustEmbedUnimplementedQueryServer()
}

// UnimplementedQueryServer must be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (UnimplementedQueryServer) GasPrices(context.Context, *QueryGasPricesRequest) (*QueryGasPricesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GasPrices not implemented")
}
func (UnimplementedQueryServer) GasPrice(context.Context, *QueryGasPriceRequest) (*QueryGasPriceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GasPrice not implemented")
}
func (UnimplementedQueryServer) mustEmbedUnimplementedQueryServer() {}

// UnsafeQueryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServer will
// result in compilation errors.
type UnsafeQueryServer interface {
	mustEmbedUnimplementedQueryServer()
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_GasPrices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGasPricesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GasPrices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GasPrices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GasPrices(ctx, req.(*QueryGasPricesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GasPrice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGasPriceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GasPrice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Query_GasPrice_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GasPrice(ctx, req.(*QueryGasPriceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "initia.tx.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GasPrices",
			Handler:    _Query_GasPrices_Handler,
		},
		{
			MethodName: "GasPrice",
			Handler:    _Query_GasPrice_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "initia/tx/v1/query.proto",
}