// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: asset.proto

package asset

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

// AssetClient is the client API for Asset service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AssetClient interface {
	CreateAsset(ctx context.Context, in *CreateAssetReq, opts ...grpc.CallOption) (*CreateAssetResp, error)
	// create project
	CreateProject(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*ProjectResp, error)
	// get project list
	GetProjectList(ctx context.Context, in *ProjectListReq, opts ...grpc.CallOption) (*ProjectListResp, error)
	ImportAssets(ctx context.Context, in *UploadReq, opts ...grpc.CallOption) (*UploadResp, error)
	// get common assets list
	GetAssetList(ctx context.Context, in *PageReq, opts ...grpc.CallOption) (*AssetListResp, error)
	// get filtered assets list
	GetFilterAssetList(ctx context.Context, in *ProjectFilterReq, opts ...grpc.CallOption) (*AssetListResp, error)
}

type assetClient struct {
	cc grpc.ClientConnInterface
}

func NewAssetClient(cc grpc.ClientConnInterface) AssetClient {
	return &assetClient{cc}
}

func (c *assetClient) CreateAsset(ctx context.Context, in *CreateAssetReq, opts ...grpc.CallOption) (*CreateAssetResp, error) {
	out := new(CreateAssetResp)
	err := c.cc.Invoke(ctx, "/asset.Asset/CreateAsset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) CreateProject(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*ProjectResp, error) {
	out := new(ProjectResp)
	err := c.cc.Invoke(ctx, "/asset.Asset/CreateProject", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) GetProjectList(ctx context.Context, in *ProjectListReq, opts ...grpc.CallOption) (*ProjectListResp, error) {
	out := new(ProjectListResp)
	err := c.cc.Invoke(ctx, "/asset.Asset/GetProjectList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) ImportAssets(ctx context.Context, in *UploadReq, opts ...grpc.CallOption) (*UploadResp, error) {
	out := new(UploadResp)
	err := c.cc.Invoke(ctx, "/asset.Asset/ImportAssets", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) GetAssetList(ctx context.Context, in *PageReq, opts ...grpc.CallOption) (*AssetListResp, error) {
	out := new(AssetListResp)
	err := c.cc.Invoke(ctx, "/asset.Asset/GetAssetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *assetClient) GetFilterAssetList(ctx context.Context, in *ProjectFilterReq, opts ...grpc.CallOption) (*AssetListResp, error) {
	out := new(AssetListResp)
	err := c.cc.Invoke(ctx, "/asset.Asset/GetFilterAssetList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AssetServer is the server API for Asset service.
// All implementations must embed UnimplementedAssetServer
// for forward compatibility
type AssetServer interface {
	CreateAsset(context.Context, *CreateAssetReq) (*CreateAssetResp, error)
	// create project
	CreateProject(context.Context, *ProjectReq) (*ProjectResp, error)
	// get project list
	GetProjectList(context.Context, *ProjectListReq) (*ProjectListResp, error)
	ImportAssets(context.Context, *UploadReq) (*UploadResp, error)
	// get common assets list
	GetAssetList(context.Context, *PageReq) (*AssetListResp, error)
	// get filtered assets list
	GetFilterAssetList(context.Context, *ProjectFilterReq) (*AssetListResp, error)
	mustEmbedUnimplementedAssetServer()
}

// UnimplementedAssetServer must be embedded to have forward compatible implementations.
type UnimplementedAssetServer struct {
}

func (UnimplementedAssetServer) CreateAsset(context.Context, *CreateAssetReq) (*CreateAssetResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAsset not implemented")
}
func (UnimplementedAssetServer) CreateProject(context.Context, *ProjectReq) (*ProjectResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProject not implemented")
}
func (UnimplementedAssetServer) GetProjectList(context.Context, *ProjectListReq) (*ProjectListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProjectList not implemented")
}
func (UnimplementedAssetServer) ImportAssets(context.Context, *UploadReq) (*UploadResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ImportAssets not implemented")
}
func (UnimplementedAssetServer) GetAssetList(context.Context, *PageReq) (*AssetListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAssetList not implemented")
}
func (UnimplementedAssetServer) GetFilterAssetList(context.Context, *ProjectFilterReq) (*AssetListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFilterAssetList not implemented")
}
func (UnimplementedAssetServer) mustEmbedUnimplementedAssetServer() {}

// UnsafeAssetServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AssetServer will
// result in compilation errors.
type UnsafeAssetServer interface {
	mustEmbedUnimplementedAssetServer()
}

func RegisterAssetServer(s grpc.ServiceRegistrar, srv AssetServer) {
	s.RegisterService(&Asset_ServiceDesc, srv)
}

func _Asset_CreateAsset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAssetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).CreateAsset(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/asset.Asset/CreateAsset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).CreateAsset(ctx, req.(*CreateAssetReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Asset_CreateProject_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).CreateProject(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/asset.Asset/CreateProject",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).CreateProject(ctx, req.(*ProjectReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Asset_GetProjectList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).GetProjectList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/asset.Asset/GetProjectList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).GetProjectList(ctx, req.(*ProjectListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Asset_ImportAssets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).ImportAssets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/asset.Asset/ImportAssets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).ImportAssets(ctx, req.(*UploadReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Asset_GetAssetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).GetAssetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/asset.Asset/GetAssetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).GetAssetList(ctx, req.(*PageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Asset_GetFilterAssetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProjectFilterReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AssetServer).GetFilterAssetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/asset.Asset/GetFilterAssetList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AssetServer).GetFilterAssetList(ctx, req.(*ProjectFilterReq))
	}
	return interceptor(ctx, in, info, handler)
}

// Asset_ServiceDesc is the grpc.ServiceDesc for Asset service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Asset_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "asset.Asset",
	HandlerType: (*AssetServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAsset",
			Handler:    _Asset_CreateAsset_Handler,
		},
		{
			MethodName: "CreateProject",
			Handler:    _Asset_CreateProject_Handler,
		},
		{
			MethodName: "GetProjectList",
			Handler:    _Asset_GetProjectList_Handler,
		},
		{
			MethodName: "ImportAssets",
			Handler:    _Asset_ImportAssets_Handler,
		},
		{
			MethodName: "GetAssetList",
			Handler:    _Asset_GetAssetList_Handler,
		},
		{
			MethodName: "GetFilterAssetList",
			Handler:    _Asset_GetFilterAssetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "asset.proto",
}
