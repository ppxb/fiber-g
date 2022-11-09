package logic

import (
	"context"
	"fiber-g/apps/asset/asset"
	"fiber-g/apps/asset/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFilterAssetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFilterAssetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFilterAssetListLogic {
	return &GetFilterAssetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetFilterAssetList get filtered assets list
func (l *GetFilterAssetListLogic) GetFilterAssetList(in *asset.ProjectFilterReq) (*asset.AssetListResp, error) {
	//var assets []*asset.AssetInfo
	var total int64

	db := l.svcCtx.Db.Table("assets")
	db.Count(&total)

	return &asset.AssetListResp{}, nil
}
