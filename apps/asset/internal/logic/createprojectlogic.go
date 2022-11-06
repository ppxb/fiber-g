package logic

import (
	"context"
	"fiber-g/apps/asset/internal/model"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"fiber-g/apps/asset/asset"
	"fiber-g/apps/asset/internal/svc"

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
			return nil, err
		}

		return &asset.ProjectResp{
			Uuid: data.UUID,
		}, nil
	}

	return nil, err
}
