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

type ProjectListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProjectListLogic {
	return &ProjectListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProjectListLogic) ProjectList(req *types.PageInfoReq) (resp *types.ResultWithData, err error) {
	if req.Page < 0 || req.PageSize < 0 {
		return nil, errorx.NewDefaultError("pageSize 或 page 不能小于0")
	}
	res, err := l.svcCtx.AssetRpc.GetProjectList(context.Background(), &asset.ProjectListReq{
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
			"projects": res.Data,
		},
		Timestamp: time.Now().Unix(),
	}, nil
}
