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

func (l *CreateProjectLogic) CreateProject(req *types.Project) (resp *types.ResultWithData, err error) {
	res, err := l.svcCtx.AssetRpc.CreateProject(context.Background(), &asset.ProjectReq{
		Name:            req.Name,
		ParentProjectId: req.ParentProjectId,
	})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
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
