// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: api/v1/profile-data-api.proto

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
	ProfileDataService_Ping_FullMethodName              = "/pb.ProfileDataService/Ping"
	ProfileDataService_CreateProfileData_FullMethodName = "/pb.ProfileDataService/CreateProfileData"
	ProfileDataService_GetProfileData_FullMethodName    = "/pb.ProfileDataService/GetProfileData"
	ProfileDataService_UpdateProfileData_FullMethodName = "/pb.ProfileDataService/UpdateProfileData"
	ProfileDataService_DeleteProfileData_FullMethodName = "/pb.ProfileDataService/DeleteProfileData"
	ProfileDataService_ListProfileData_FullMethodName   = "/pb.ProfileDataService/ListProfileData"
)

// ProfileDataServiceClient is the client API for ProfileDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProfileDataServiceClient interface {
	Ping(ctx context.Context, in *ProfileDataRequest, opts ...grpc.CallOption) (*ProfileDataResponse, error)
	CreateProfileData(ctx context.Context, in *CreateProfileDataRequest, opts ...grpc.CallOption) (*CreateProfileDataResponse, error)
	GetProfileData(ctx context.Context, in *GetProfileDataRequest, opts ...grpc.CallOption) (*GetProfileDataResponse, error)
	UpdateProfileData(ctx context.Context, in *UpdateProfileDataRequest, opts ...grpc.CallOption) (*UpdateProfileDataResponse, error)
	DeleteProfileData(ctx context.Context, in *DeleteProfileDataRequest, opts ...grpc.CallOption) (*DeleteProfileDataResponse, error)
	ListProfileData(ctx context.Context, in *ListProfileDataRequest, opts ...grpc.CallOption) (*ListProfileDataResponse, error)
}

type profileDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProfileDataServiceClient(cc grpc.ClientConnInterface) ProfileDataServiceClient {
	return &profileDataServiceClient{cc}
}

func (c *profileDataServiceClient) Ping(ctx context.Context, in *ProfileDataRequest, opts ...grpc.CallOption) (*ProfileDataResponse, error) {
	out := new(ProfileDataResponse)
	err := c.cc.Invoke(ctx, ProfileDataService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileDataServiceClient) CreateProfileData(ctx context.Context, in *CreateProfileDataRequest, opts ...grpc.CallOption) (*CreateProfileDataResponse, error) {
	out := new(CreateProfileDataResponse)
	err := c.cc.Invoke(ctx, ProfileDataService_CreateProfileData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileDataServiceClient) GetProfileData(ctx context.Context, in *GetProfileDataRequest, opts ...grpc.CallOption) (*GetProfileDataResponse, error) {
	out := new(GetProfileDataResponse)
	err := c.cc.Invoke(ctx, ProfileDataService_GetProfileData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileDataServiceClient) UpdateProfileData(ctx context.Context, in *UpdateProfileDataRequest, opts ...grpc.CallOption) (*UpdateProfileDataResponse, error) {
	out := new(UpdateProfileDataResponse)
	err := c.cc.Invoke(ctx, ProfileDataService_UpdateProfileData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileDataServiceClient) DeleteProfileData(ctx context.Context, in *DeleteProfileDataRequest, opts ...grpc.CallOption) (*DeleteProfileDataResponse, error) {
	out := new(DeleteProfileDataResponse)
	err := c.cc.Invoke(ctx, ProfileDataService_DeleteProfileData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *profileDataServiceClient) ListProfileData(ctx context.Context, in *ListProfileDataRequest, opts ...grpc.CallOption) (*ListProfileDataResponse, error) {
	out := new(ListProfileDataResponse)
	err := c.cc.Invoke(ctx, ProfileDataService_ListProfileData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProfileDataServiceServer is the server API for ProfileDataService service.
// All implementations must embed UnimplementedProfileDataServiceServer
// for forward compatibility
type ProfileDataServiceServer interface {
	Ping(context.Context, *ProfileDataRequest) (*ProfileDataResponse, error)
	CreateProfileData(context.Context, *CreateProfileDataRequest) (*CreateProfileDataResponse, error)
	GetProfileData(context.Context, *GetProfileDataRequest) (*GetProfileDataResponse, error)
	UpdateProfileData(context.Context, *UpdateProfileDataRequest) (*UpdateProfileDataResponse, error)
	DeleteProfileData(context.Context, *DeleteProfileDataRequest) (*DeleteProfileDataResponse, error)
	ListProfileData(context.Context, *ListProfileDataRequest) (*ListProfileDataResponse, error)
	mustEmbedUnimplementedProfileDataServiceServer()
}

// UnimplementedProfileDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProfileDataServiceServer struct {
}

func (UnimplementedProfileDataServiceServer) Ping(context.Context, *ProfileDataRequest) (*ProfileDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedProfileDataServiceServer) CreateProfileData(context.Context, *CreateProfileDataRequest) (*CreateProfileDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProfileData not implemented")
}
func (UnimplementedProfileDataServiceServer) GetProfileData(context.Context, *GetProfileDataRequest) (*GetProfileDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileData not implemented")
}
func (UnimplementedProfileDataServiceServer) UpdateProfileData(context.Context, *UpdateProfileDataRequest) (*UpdateProfileDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfileData not implemented")
}
func (UnimplementedProfileDataServiceServer) DeleteProfileData(context.Context, *DeleteProfileDataRequest) (*DeleteProfileDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProfileData not implemented")
}
func (UnimplementedProfileDataServiceServer) ListProfileData(context.Context, *ListProfileDataRequest) (*ListProfileDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProfileData not implemented")
}
func (UnimplementedProfileDataServiceServer) mustEmbedUnimplementedProfileDataServiceServer() {}

// UnsafeProfileDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProfileDataServiceServer will
// result in compilation errors.
type UnsafeProfileDataServiceServer interface {
	mustEmbedUnimplementedProfileDataServiceServer()
}

func RegisterProfileDataServiceServer(s grpc.ServiceRegistrar, srv ProfileDataServiceServer) {
	s.RegisterService(&ProfileDataService_ServiceDesc, srv)
}

func _ProfileDataService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileDataServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileDataService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileDataServiceServer).Ping(ctx, req.(*ProfileDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileDataService_CreateProfileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateProfileDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileDataServiceServer).CreateProfileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileDataService_CreateProfileData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileDataServiceServer).CreateProfileData(ctx, req.(*CreateProfileDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileDataService_GetProfileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileDataServiceServer).GetProfileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileDataService_GetProfileData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileDataServiceServer).GetProfileData(ctx, req.(*GetProfileDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileDataService_UpdateProfileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileDataServiceServer).UpdateProfileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileDataService_UpdateProfileData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileDataServiceServer).UpdateProfileData(ctx, req.(*UpdateProfileDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileDataService_DeleteProfileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProfileDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileDataServiceServer).DeleteProfileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileDataService_DeleteProfileData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileDataServiceServer).DeleteProfileData(ctx, req.(*DeleteProfileDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProfileDataService_ListProfileData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProfileDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProfileDataServiceServer).ListProfileData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ProfileDataService_ListProfileData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProfileDataServiceServer).ListProfileData(ctx, req.(*ListProfileDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ProfileDataService_ServiceDesc is the grpc.ServiceDesc for ProfileDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProfileDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ProfileDataService",
	HandlerType: (*ProfileDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _ProfileDataService_Ping_Handler,
		},
		{
			MethodName: "CreateProfileData",
			Handler:    _ProfileDataService_CreateProfileData_Handler,
		},
		{
			MethodName: "GetProfileData",
			Handler:    _ProfileDataService_GetProfileData_Handler,
		},
		{
			MethodName: "UpdateProfileData",
			Handler:    _ProfileDataService_UpdateProfileData_Handler,
		},
		{
			MethodName: "DeleteProfileData",
			Handler:    _ProfileDataService_DeleteProfileData_Handler,
		},
		{
			MethodName: "ListProfileData",
			Handler:    _ProfileDataService_ListProfileData_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/profile-data-api.proto",
}
