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

type CreateAssetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAssetLogic {
	return &CreateAssetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAssetLogic) CreateAsset(req *types.Asset) (resp *types.ResultWithData, err error) {
	res, err := l.svcCtx.AssetRpc.CreateAsset(l.ctx, &asset.CreateAssetReq{
		Name:           req.Name,
		Serial:         req.Serial,
		ProjectId:      req.ProjectId,
		SonProjectId:   req.SonProjectId,
		PartProjectId:  req.PartProjectId,
		Type:           req.Type,
		SubDistrict:    req.SubDistrict,
		Brand:          req.Brand,
		Unit:           req.Unit,
		Params:         req.Params,
		Value:          "0",
		Address:        req.Address,
		Long:           req.Long,
		Lat:            req.Lat,
		ImgUrl:         req.ImgUrl,
		SystemLoginPwd: req.SystemLoginPwd,
		SystemLoginUrl: req.SystemLoginUrl,
	})
	if err != nil {
		return nil, err
	}

	return &types.ResultWithData{
		Code: errorx.OK,
		Msg:  "项目创建成功",
		Data: map[string]interface{}{
			"uuid": res.Uuid,
		},
		Timestamp: time.Now().Unix(),
	}, nil
}
