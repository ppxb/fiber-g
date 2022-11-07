package logic

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"fiber-g/apps/asset/asset"
	"fiber-g/apps/asset/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProjectListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProjectListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProjectListLogic {
	return &GetProjectListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProjectListLogic) GetProjectList(in *asset.ProjectListReq) (*asset.ProjectListResp, error) {
	var projects []*asset.ProjectInfo
	var total int64

	err := l.svcCtx.Db.Table("projects").Count(&total).Offset(int(in.Page * in.PageSize)).Limit(int(in.PageSize)).Find(&projects).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, status.Error(codes.Internal, "服务器内部故障")
	}

	return &asset.ProjectListResp{
		Total: uint64(total),
		Data:  projects,
	}, nil
}
