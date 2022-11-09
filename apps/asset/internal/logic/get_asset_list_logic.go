package logic

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"fiber-g/apps/asset/asset"
	"fiber-g/apps/asset/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAssetListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAssetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAssetListLogic {
	return &GetAssetListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetAssetList get common assets list
func (l *GetAssetListLogic) GetAssetList(in *asset.PageReq) (*asset.AssetListResp, error) {
	var assets []*asset.AssetInfo
	var total int64

	err := l.svcCtx.Db.Table("assets").Count(&total).Offset(int(in.Page * in.PageSize)).Limit(int(in.PageSize)).Find(&assets).Error
	if err != nil {
		return nil, status.Error(codes.Internal, "服务器内部错误")
	}

	return &asset.AssetListResp{
		Total: total,
		Data:  assets,
	}, nil
}
