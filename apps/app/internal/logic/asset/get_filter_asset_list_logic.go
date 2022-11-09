package asset

import (
	"context"
	"fiber-g/apps/asset/asset"
	"fiber-g/pkg/errorx"
	"time"

	"fiber-g/apps/app/internal/svc"
	"fiber-g/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFilterAssetListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetFilterAssetListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFilterAssetListLogic {
	return &GetFilterAssetListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetFilterAssetListLogic) GetFilterAssetList(req *types.FilterAssetListReq) (resp *types.ResultWithData, err error) {
	if req.Page < 0 || req.PageSize < 0 || req.PageSize > 50 {
		return nil, errorx.NewDefaultError("page 或 pageSize 超过了允许的值")
	}

	res, err := l.svcCtx.AssetRpc.GetAssetList(context.Background(), &asset.PageReq{
		Page:     int64(req.Page),
		PageSize: int64(req.PageSize),
	})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.ResultWithData{
		Code: errorx.OK,
		Msg:  "请求成功",
		Data: map[string]interface{}{
			"assets": res.Data,
			"total":  res.Total,
		},
		Timestamp: time.Now().Unix(),
	}, nil
}
