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
	CreateAssetReq  = asset.CreateAssetReq
	CreateAssetResp = asset.CreateAssetResp
	ProjectReq      = asset.ProjectReq
	ProjectResp     = asset.ProjectResp

	Asset interface {
		CreateAsset(ctx context.Context, in *CreateAssetReq, opts ...grpc.CallOption) (*CreateAssetResp, error)
		CreateProject(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*ProjectResp, error)
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

func (m *defaultAsset) CreateProject(ctx context.Context, in *ProjectReq, opts ...grpc.CallOption) (*ProjectResp, error) {
	client := asset.NewAssetClient(m.cli.Conn())
	return client.CreateProject(ctx, in, opts...)
}
