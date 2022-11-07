package logic

import (
	"context"
	"fiber-g/apps/asset/asset"
	"fiber-g/apps/asset/internal/model"
	"fiber-g/apps/asset/internal/svc"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProjectLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProjectLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProjectLogic {
	return &CreateProjectLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProjectLogic) CreateProject(in *asset.ProjectReq) (*asset.ProjectResp, error) {
	var p model.Project
	err := l.svcCtx.Db.Where("name = ?", in.Name).First(&p).Error
	if err == gorm.ErrRecordNotFound {
		data := &model.Project{
			UUID:            uuid.NewString(),
			Name:            in.Name,
			ParentProjectId: in.ParentProjectId,
		}
		err = l.svcCtx.Db.Create(data).Error
		if err != nil {
			return nil, status.Error(codes.Internal, "服务器内部故障")
		}

		return &asset.ProjectResp{
			Uuid: data.UUID,
		}, nil
	}

	return nil, status.Error(codes.AlreadyExists, "项目名已存在")
}
