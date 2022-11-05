package asset

import (
	"context"
	"fiber-g/apps/app/internal/svc"
	"fiber-g/apps/app/internal/types"
	"fmt"
	"github.com/google/uuid"
	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"net/http"
	"reflect"
)

const maxFileSize = 10 << 20 // 10MB

type ImportAssetLogic struct {
	logx.Logger
	r      *http.Request
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewImportAssetLogic(r *http.Request, svcCtx *svc.ServiceContext) *ImportAssetLogic {
	return &ImportAssetLogic{
		Logger: logx.WithContext(r.Context()),
		ctx:    r.Context(),
		r:      r,
		svcCtx: svcCtx,
	}
}

func (l *ImportAssetLogic) ImportAsset() (resp *types.ImportAssetResp, err error) {
	l.r.ParseMultipartForm(maxFileSize)
	file, handler, err := l.r.FormFile("file")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	logx.Infof("upload file: %+v, file size: %d, MIME header: %+v",
		handler.Filename, handler.Size, handler.Header)

	f, err := excelize.OpenReader(file)
	if err != nil {
		logx.Error("open file error")
	}

	var list []types.CreateAssetReq
	var asset types.CreateAssetReq

	rows, err := f.GetRows("Sheet1")

	values := reflect.ValueOf(&asset).Elem()

	for i, r := range rows {
		if i < 1 {
			continue
		}
		for j, v := range r {
			values.Field(j).SetString(v)
		}
		list = append(list, asset)
		fmt.Printf("%+v", list)
	}
	//入库将value改为float64

	return &types.ImportAssetResp{UUID: uuid.NewString()}, nil
}
