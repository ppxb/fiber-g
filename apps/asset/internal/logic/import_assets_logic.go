package logic

import (
	"bytes"
	"context"
	"fiber-g/apps/asset/asset"
	"fiber-g/apps/asset/internal/model"
	"fiber-g/apps/asset/internal/svc"
	"fiber-g/pkg/utils"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"reflect"
	"strconv"
)

type ImportAssetsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewImportAssetsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ImportAssetsLogic {
	return &ImportAssetsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ImportAssetsLogic) ImportAssets(in *asset.UploadReq) (*asset.UploadResp, error) {
	reader := bytes.NewReader(in.File)
	f, err := excelize.OpenReader(reader)
	if err != nil {
		return nil, status.Error(codes.Internal, "文件解析失败，请稍后再试")
	}

	var req model.AssetImportDao
	var uploadAsset model.Asset
	var uploadAssets []model.Asset

	rows, err := f.GetRows("Sheet1")
	values := reflect.ValueOf(&req).Elem()
	for i, r := range rows {
		if i < 1 {
			continue
		}

		for j, v := range r {
			switch values.Field(j).Type().String() {
			case "float64":
				cv, _ := strconv.ParseFloat(v, 64)
				values.Field(j).SetFloat(cv)
			default:
				values.Field(j).SetString(v)
			}
		}
		utils.Struct2StructByJson(req, &uploadAsset)
		uploadAsset.UUID = uuid.NewString()
		uploadAssets = append(uploadAssets, uploadAsset)
	}

	affected := l.svcCtx.Db.Table("assets").CreateInBatches(uploadAssets, 1000).RowsAffected
	fail := int64(len(rows)) - 1 - affected

	return &asset.UploadResp{
		Success: affected,
		Fail:    fail,
	}, nil
}
