package logic

import (
	"context"
	"fiber-g/apps/dept/rpc/dept"
	"fiber-g/apps/dept/rpc/internal/svc"
	"fiber-g/apps/dept/rpc/model"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDeptLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateDeptLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDeptLogic {
	return &CreateDeptLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateDeptLogic) CreateDept(in *dept.CreateDeptReq) (*dept.CreateDeptResp, error) {
	newDept := model.Depts{
		Id:       uuid.NewString(),
		Name:     in.Name,
		Level:    in.Level,
		ParentId: in.ParentId,
		HeaderId: in.HeaderId,
	}

	_, err := l.svcCtx.DeptModel.Insert(l.ctx, &newDept)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &dept.CreateDeptResp{
		Uuid: newDept.Id,
	}, nil
}
