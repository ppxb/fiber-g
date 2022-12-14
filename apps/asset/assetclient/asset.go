// Code generated by goctl. DO NOT EDIT!
// Source: asset.proto

package assetclient

import (
	"context"

	"fiber-g/apps/asset/asset"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AssetInfo        = asset.AssetInfo
	AssetListResp    = asset.AssetListResp
	CreateAssetReq   = asset.CreateAssetReq
	CreateAssetResp  = asset.CreateAssetResp
	PageReq          = asset.PageReq
	ProjectFilterReq = asset.ProjectFilterReq
	ProjectInfo      = asset.ProjectInfo
	ProjectListReq   = asset.ProjectListReq
	ProjectListResp  = asset.ProjectListResp
	ProjectReq       = asset.ProjectReq
	ProjectResp      = asset.ProjectResp
	UploadReq        = asset.UploadReq
	UploadResp       = asset.UploadResp

	Asset interface {
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

	defaultAsset struct {
		cli zrpc.Client
	}
)

func NewAsset(cli zrpc.Client) Asset {
	return &defaultAsset{
		cli: cli,
	}
}

func (m *defaultAsset) CreateAsset(ctx context.Context, in *CreateAssetReq, opts ...grpc.CallOption) (*CreateAssetResp, error) {
	client := asset.NewAssetClient(m.cli.Conn())
	return client.CreateAsset(ctx, in, opts...)
}

// create project
func (m *defaultAsset) CreateProject(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*ProjectResp, error) {
	client := asset.NewAssetClient(m.cli.Conn())
	return client.CreateProject(ctx, in, opts...)
}

// get project list
func (m *defaultAsset) GetProjectList(ctx context.Context, in *ProjectListReq, opts ...grpc.CallOption) (*ProjectListResp, error) {
	client := asset.NewAssetClient(m.cli.Conn())
	return client.GetProjectList(ctx, in, opts...)
}

func (m *defaultAsset) ImportAssets(ctx context.Context, in *UploadReq, opts ...grpc.CallOption) (*UploadResp, error) {
	client := asset.NewAssetClient(m.cli.Conn())
	return client.ImportAssets(ctx, in, opts...)
}

// get common assets list
func (m *defaultAsset) GetAssetList(ctx context.Context, in *PageReq, opts ...grpc.CallOption) (*AssetListResp, error) {
	client := asset.NewAssetClient(m.cli.Conn())
	return client.GetAssetList(ctx, in, opts...)
}

// get filtered assets list
func (m *defaultAsset) GetFilterAssetList(ctx context.Context, in *ProjectFilterReq, opts ...grpc.CallOption) (*AssetListResp, error) {
	client := asset.NewAssetClient(m.cli.Conn())
	return client.GetFilterAssetList(ctx, in, opts...)
}
