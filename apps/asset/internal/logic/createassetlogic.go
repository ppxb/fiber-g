package logic

import (
	"context"
	"database/sql"
	"fiber-g/apps/asset/internal/model"
	"github.com/google/uuid"
	"google.golang.org/grpc/status"
	"strconv"

	"fiber-g/apps/asset/asset"
	"fiber-g/apps/asset/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAssetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAssetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAssetLogic {
	return &CreateAssetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateAssetLogic) CreateAsset(in *asset.CreateAssetReq) (*asset.CreateAssetResp, error) {
	v, _ := strconv.ParseFloat(in.Value, 64)
	newAsset := model.Assets{
		Id:             uuid.NewString(),
		Name:           in.Name,
		Serial:         in.Serial,
		ProjectId:      in.ProjectId,
		SonProjectId:   in.SonProjectId,
		PartProjectId:  in.PartProjectId,
		Type:           in.Type,
		SubDistrict:    in.SubDistrict,
		Brand:          in.Brand,
		Model:          in.Model,
		Unit:           in.Unit,
		Params:         sql.NullString{String: in.Params},
		Value:          v,
		Address:        in.Address,
		Long:           in.Long,
		Lat:            in.Lat,
		ImgUrl:         in.ImgUrl,
		SystemLoginPwd: in.SystemLoginPwd,
		SystemLoginUrl: in.SystemLoginUrl,
	}

	_, err := l.svcCtx.AssetModel.Insert(l.ctx, &newAsset)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &asset.CreateAssetResp{
		Uuid: newAsset.Id,
	}, nil
}
