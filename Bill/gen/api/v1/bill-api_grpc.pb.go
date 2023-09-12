// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api/v1/bill-api.proto

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

const (
	BillService_Ping_FullMethodName       = "/pb.BillService/Ping"
	BillService_CreateBill_FullMethodName = "/pb.BillService/CreateBill"
	BillService_GetBill_FullMethodName    = "/pb.BillService/GetBill"
	BillService_UpdateBill_FullMethodName = "/pb.BillService/UpdateBill"
	BillService_DeleteBill_FullMethodName = "/pb.BillService/DeleteBill"
	BillService_ListBills_FullMethodName  = "/pb.BillService/ListBills"
)

// BillServiceClient is the client API for BillService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BillServiceClient interface {
	Ping(ctx context.Context, in *BillRequest, opts ...grpc.CallOption) (*BillResponse, error)
	CreateBill(ctx context.Context, in *CreateBillRequest, opts ...grpc.CallOption) (*CreateBillResponse, error)
	GetBill(ctx context.Context, in *GetBillRequest, opts ...grpc.CallOption) (*GetBillResponse, error)
	UpdateBill(ctx context.Context, in *UpdateBillRequest, opts ...grpc.CallOption) (*UpdateBillResponse, error)
	DeleteBill(ctx context.Context, in *DeleteBillRequest, opts ...grpc.CallOption) (*DeleteBillResponse, error)
	ListBills(ctx context.Context, in *ListBillsRequest, opts ...grpc.CallOption) (*ListBillsResponse, error)
}

type billServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBillServiceClient(cc grpc.ClientConnInterface) BillServiceClient {
	return &billServiceClient{cc}
}

func (c *billServiceClient) Ping(ctx context.Context, in *BillRequest, opts ...grpc.CallOption) (*BillResponse, error) {
	out := new(BillResponse)
	err := c.cc.Invoke(ctx, BillService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billServiceClient) CreateBill(ctx context.Context, in *CreateBillRequest, opts ...grpc.CallOption) (*CreateBillResponse, error) {
	out := new(CreateBillResponse)
	err := c.cc.Invoke(ctx, BillService_CreateBill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billServiceClient) GetBill(ctx context.Context, in *GetBillRequest, opts ...grpc.CallOption) (*GetBillResponse, error) {
	out := new(GetBillResponse)
	err := c.cc.Invoke(ctx, BillService_GetBill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billServiceClient) UpdateBill(ctx context.Context, in *UpdateBillRequest, opts ...grpc.CallOption) (*UpdateBillResponse, error) {
	out := new(UpdateBillResponse)
	err := c.cc.Invoke(ctx, BillService_UpdateBill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billServiceClient) DeleteBill(ctx context.Context, in *DeleteBillRequest, opts ...grpc.CallOption) (*DeleteBillResponse, error) {
	out := new(DeleteBillResponse)
	err := c.cc.Invoke(ctx, BillService_DeleteBill_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billServiceClient) ListBills(ctx context.Context, in *ListBillsRequest, opts ...grpc.CallOption) (*ListBillsResponse, error) {
	out := new(ListBillsResponse)
	err := c.cc.Invoke(ctx, BillService_ListBills_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillServiceServer is the server API for BillService service.
// All implementations must embed UnimplementedBillServiceServer
// for forward compatibility
type BillServiceServer interface {
	Ping(context.Context, *BillRequest) (*BillResponse, error)
	CreateBill(context.Context, *CreateBillRequest) (*CreateBillResponse, error)
	GetBill(context.Context, *GetBillRequest) (*GetBillResponse, error)
	UpdateBill(context.Context, *UpdateBillRequest) (*UpdateBillResponse, error)
	DeleteBill(context.Context, *DeleteBillRequest) (*DeleteBillResponse, error)
	ListBills(context.Context, *ListBillsRequest) (*ListBillsResponse, error)
	mustEmbedUnimplementedBillServiceServer()
}

// UnimplementedBillServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBillServiceServer struct {
}

func (UnimplementedBillServiceServer) Ping(context.Context, *BillRequest) (*BillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedBillServiceServer) CreateBill(context.Context, *CreateBillRequest) (*CreateBillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBill not implemented")
}
func (UnimplementedBillServiceServer) GetBill(context.Context, *GetBillRequest) (*GetBillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBill not implemented")
}
func (UnimplementedBillServiceServer) UpdateBill(context.Context, *UpdateBillRequest) (*UpdateBillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBill not implemented")
}
func (UnimplementedBillServiceServer) DeleteBill(context.Context, *DeleteBillRequest) (*DeleteBillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBill not implemented")
}
func (UnimplementedBillServiceServer) ListBills(context.Context, *ListBillsRequest) (*ListBillsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBills not implemented")
}
func (UnimplementedBillServiceServer) mustEmbedUnimplementedBillServiceServer() {}

// UnsafeBillServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillServiceServer will
// result in compilation errors.
type UnsafeBillServiceServer interface {
	mustEmbedUnimplementedBillServiceServer()
}

func RegisterBillServiceServer(s grpc.ServiceRegistrar, srv BillServiceServer) {
	s.RegisterService(&BillService_ServiceDesc, srv)
}

func _BillService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillServiceServer).Ping(ctx, req.(*BillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillService_CreateBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillServiceServer).CreateBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillService_CreateBill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillServiceServer).CreateBill(ctx, req.(*CreateBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillService_GetBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillServiceServer).GetBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillService_GetBill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillServiceServer).GetBill(ctx, req.(*GetBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillService_UpdateBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillServiceServer).UpdateBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillService_UpdateBill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillServiceServer).UpdateBill(ctx, req.(*UpdateBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillService_DeleteBill_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillServiceServer).DeleteBill(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillService_DeleteBill_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillServiceServer).DeleteBill(ctx, req.(*DeleteBillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BillService_ListBills_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBillsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillServiceServer).ListBills(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BillService_ListBills_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillServiceServer).ListBills(ctx, req.(*ListBillsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BillService_ServiceDesc is the grpc.ServiceDesc for BillService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BillService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.BillService",
	HandlerType: (*BillServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _BillService_Ping_Handler,
		},
		{
			MethodName: "CreateBill",
			Handler:    _BillService_CreateBill_Handler,
		},
		{
			MethodName: "GetBill",
			Handler:    _BillService_GetBill_Handler,
		},
		{
			MethodName: "UpdateBill",
			Handler:    _BillService_UpdateBill_Handler,
		},
		{
			MethodName: "DeleteBill",
			Handler:    _BillService_DeleteBill_Handler,
		},
		{
			MethodName: "ListBills",
			Handler:    _BillService_ListBills_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/bill-api.proto",
}
