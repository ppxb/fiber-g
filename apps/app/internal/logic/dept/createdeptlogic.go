package dept

import (
	"context"
	"fiber-g/apps/dept/rpc/dept"

	"fiber-g/apps/app/internal/svc"
	"fiber-g/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDeptLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDeptLogic {
	return &CreateDeptLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateDeptLogic) CreateDept(req *types.CreateDeptReq) (resp *types.CreateDeptResp, err error) {
	res, err := l.svcCtx.DeptRpc.CreateDept(l.ctx, &dept.CreateDeptReq{
		Name:     req.Name,
		ParentId: req.ParentId,
		Level:    req.Level,
		HeaderId: req.HeaderId,
	})
	if err != nil {
		return nil, err
	}

	return &types.CreateDeptResp{
		UUID: res.Uuid,
	}, nil
}
