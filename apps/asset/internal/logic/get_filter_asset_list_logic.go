package logic

import (
	"context"
	"fiber-g/apps/asset/asset"
	"fiber-g/apps/asset/internal/svc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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
	db := l.svcCtx.Db.Table("assets")
	var assets []*asset.AssetInfo
	var total int64

	if in.Name != "" {
		switch in.Level {
		case 1:
			db = db.Where("project_name LIKE ?", "%"+in.Name+"%")
		case 2:
			db = db.Where("son_project_name LIKE ?", "%"+in.Name+"%")
		case 3:
			db = db.Where("part_project_name LIKE ?", "%"+in.Name+"%")
		}
	} else {
		return nil, status.Error(codes.Internal, "缺少必须的属性level")
	}

	db.Count(&total)
	err := db.Offset(int(in.Page * in.PageSize)).Limit(int(in.PageSize)).Find(&assets).Error
	if err != nil {
		return nil, status.Error(codes.Internal, "服务器内部错误")
	}

	return &asset.AssetListResp{
		Total: total,
		Data:  assets,
	}, nil
}
