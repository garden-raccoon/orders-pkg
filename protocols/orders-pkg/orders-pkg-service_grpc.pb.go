// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.2
// source: orders-pkg-service.proto

package orders_pkg

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	OrdersService_CreateOrUpdateOrder_FullMethodName   = "/models.OrdersService/CreateOrUpdateOrder"
	OrdersService_GetOrders_FullMethodName             = "/models.OrdersService/GetOrders"
	OrdersService_OrdersByUserUuid_FullMethodName      = "/models.OrdersService/OrdersByUserUuid"
	OrdersService_OrderByOrderUuid_FullMethodName      = "/models.OrdersService/OrderByOrderUuid"
	OrdersService_DeleteMealsFromParams_FullMethodName = "/models.OrdersService/DeleteMealsFromParams"
	OrdersService_DeleteOrder_FullMethodName           = "/models.OrdersService/DeleteOrder"
)

// OrdersServiceClient is the client API for OrdersService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// OrderService is
type OrdersServiceClient interface {
	CreateOrUpdateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*CreateOrUpdateResponse, error)
	GetOrders(ctx context.Context, in *OrdersEmpty, opts ...grpc.CallOption) (*Orders, error)
	// OrdersByUserUuid could return either one or several orders because user could have more than 1
	OrdersByUserUuid(ctx context.Context, in *ByUserUuidReq, opts ...grpc.CallOption) (*Orders, error)
	// could be only single order by its unique id
	OrderByOrderUuid(ctx context.Context, in *ByOrderUuidReq, opts ...grpc.CallOption) (*Order, error)
	DeleteMealsFromParams(ctx context.Context, in *ByMealsUuidReq, opts ...grpc.CallOption) (*OrdersEmpty, error)
	DeleteOrder(ctx context.Context, in *OrderDeleteReq, opts ...grpc.CallOption) (*OrdersEmpty, error)
}

type ordersServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrdersServiceClient(cc grpc.ClientConnInterface) OrdersServiceClient {
	return &ordersServiceClient{cc}
}

func (c *ordersServiceClient) CreateOrUpdateOrder(ctx context.Context, in *Order, opts ...grpc.CallOption) (*CreateOrUpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateOrUpdateResponse)
	err := c.cc.Invoke(ctx, OrdersService_CreateOrUpdateOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) GetOrders(ctx context.Context, in *OrdersEmpty, opts ...grpc.CallOption) (*Orders, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Orders)
	err := c.cc.Invoke(ctx, OrdersService_GetOrders_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) OrdersByUserUuid(ctx context.Context, in *ByUserUuidReq, opts ...grpc.CallOption) (*Orders, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Orders)
	err := c.cc.Invoke(ctx, OrdersService_OrdersByUserUuid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) OrderByOrderUuid(ctx context.Context, in *ByOrderUuidReq, opts ...grpc.CallOption) (*Order, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Order)
	err := c.cc.Invoke(ctx, OrdersService_OrderByOrderUuid_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) DeleteMealsFromParams(ctx context.Context, in *ByMealsUuidReq, opts ...grpc.CallOption) (*OrdersEmpty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrdersEmpty)
	err := c.cc.Invoke(ctx, OrdersService_DeleteMealsFromParams_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *ordersServiceClient) DeleteOrder(ctx context.Context, in *OrderDeleteReq, opts ...grpc.CallOption) (*OrdersEmpty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OrdersEmpty)
	err := c.cc.Invoke(ctx, OrdersService_DeleteOrder_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrdersServiceServer is the server API for OrdersService service.
// All implementations must embed UnimplementedOrdersServiceServer
// for forward compatibility.
//
// OrderService is
type OrdersServiceServer interface {
	CreateOrUpdateOrder(context.Context, *Order) (*CreateOrUpdateResponse, error)
	GetOrders(context.Context, *OrdersEmpty) (*Orders, error)
	// OrdersByUserUuid could return either one or several orders because user could have more than 1
	OrdersByUserUuid(context.Context, *ByUserUuidReq) (*Orders, error)
	// could be only single order by its unique id
	OrderByOrderUuid(context.Context, *ByOrderUuidReq) (*Order, error)
	DeleteMealsFromParams(context.Context, *ByMealsUuidReq) (*OrdersEmpty, error)
	DeleteOrder(context.Context, *OrderDeleteReq) (*OrdersEmpty, error)
	mustEmbedUnimplementedOrdersServiceServer()
}

// UnimplementedOrdersServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrdersServiceServer struct{}

func (UnimplementedOrdersServiceServer) CreateOrUpdateOrder(context.Context, *Order) (*CreateOrUpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrUpdateOrder not implemented")
}
func (UnimplementedOrdersServiceServer) GetOrders(context.Context, *OrdersEmpty) (*Orders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrders not implemented")
}
func (UnimplementedOrdersServiceServer) OrdersByUserUuid(context.Context, *ByUserUuidReq) (*Orders, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrdersByUserUuid not implemented")
}
func (UnimplementedOrdersServiceServer) OrderByOrderUuid(context.Context, *ByOrderUuidReq) (*Order, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OrderByOrderUuid not implemented")
}
func (UnimplementedOrdersServiceServer) DeleteMealsFromParams(context.Context, *ByMealsUuidReq) (*OrdersEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteMealsFromParams not implemented")
}
func (UnimplementedOrdersServiceServer) DeleteOrder(context.Context, *OrderDeleteReq) (*OrdersEmpty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedOrdersServiceServer) mustEmbedUnimplementedOrdersServiceServer() {}
func (UnimplementedOrdersServiceServer) testEmbeddedByValue()                       {}

// UnsafeOrdersServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrdersServiceServer will
// result in compilation errors.
type UnsafeOrdersServiceServer interface {
	mustEmbedUnimplementedOrdersServiceServer()
}

func RegisterOrdersServiceServer(s grpc.ServiceRegistrar, srv OrdersServiceServer) {
	// If the following call pancis, it indicates UnimplementedOrdersServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrdersService_ServiceDesc, srv)
}

func _OrdersService_CreateOrUpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Order)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).CreateOrUpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrdersService_CreateOrUpdateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).CreateOrUpdateOrder(ctx, req.(*Order))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_GetOrders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrdersEmpty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).GetOrders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrdersService_GetOrders_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).GetOrders(ctx, req.(*OrdersEmpty))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_OrdersByUserUuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByUserUuidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).OrdersByUserUuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrdersService_OrdersByUserUuid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).OrdersByUserUuid(ctx, req.(*ByUserUuidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_OrderByOrderUuid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByOrderUuidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).OrderByOrderUuid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrdersService_OrderByOrderUuid_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).OrderByOrderUuid(ctx, req.(*ByOrderUuidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_DeleteMealsFromParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByMealsUuidReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).DeleteMealsFromParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrdersService_DeleteMealsFromParams_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).DeleteMealsFromParams(ctx, req.(*ByMealsUuidReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrdersService_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderDeleteReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrdersServiceServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrdersService_DeleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrdersServiceServer).DeleteOrder(ctx, req.(*OrderDeleteReq))
	}
	return interceptor(ctx, in, info, handler)
}

// OrdersService_ServiceDesc is the grpc.ServiceDesc for OrdersService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrdersService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "models.OrdersService",
	HandlerType: (*OrdersServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrUpdateOrder",
			Handler:    _OrdersService_CreateOrUpdateOrder_Handler,
		},
		{
			MethodName: "GetOrders",
			Handler:    _OrdersService_GetOrders_Handler,
		},
		{
			MethodName: "OrdersByUserUuid",
			Handler:    _OrdersService_OrdersByUserUuid_Handler,
		},
		{
			MethodName: "OrderByOrderUuid",
			Handler:    _OrdersService_OrderByOrderUuid_Handler,
		},
		{
			MethodName: "DeleteMealsFromParams",
			Handler:    _OrdersService_DeleteMealsFromParams_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _OrdersService_DeleteOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "orders-pkg-service.proto",
}
