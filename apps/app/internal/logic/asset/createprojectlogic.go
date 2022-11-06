package asset

import (
	"context"
	"fiber-g/apps/asset/asset"

	"fiber-g/apps/app/internal/svc"
	"fiber-g/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProjectLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProjectLogic {
	return &CreateProjectLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateProjectLogic) CreateProject(req *types.ProjectReq) (resp *types.ProjectResp, err error) {
	res, err := l.svcCtx.AssetRpc.CreateProject(context.Background(), &asset.ProjectReq{
		Name:            req.Name,
		ParentProjectId: req.ParentProjectId,
	})
	if err != nil {
		return nil, err
	}

	return &types.ProjectResp{UUID: res.Uuid}, nil
}
