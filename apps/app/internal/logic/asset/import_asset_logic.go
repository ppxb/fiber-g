package asset

import (
	"context"
	"fiber-g/pkg/errorx"
	"fmt"
	"github.com/xuri/excelize/v2"
	"net/http"
	"reflect"

	"fiber-g/apps/app/internal/svc"
	"fiber-g/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 20 // 10MB

type ImportAssetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
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
	// file size bigger than defined
	err = l.r.ParseMultipartForm(maxFileSize)
	if err != nil {
		return nil, errorx.NewDefaultError("文件过大，上传文件请小于20M")
	}

	file, handler, err := l.r.FormFile("file")
	if err != nil {
		return nil, errorx.NewDefaultError("文件读取失败，请稍后再试")
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	logx.Infof("upload file: %+v, file size: %d, MIME header: %+v",
		handler.Filename, handler.Size, handler.Header)

	f, err := excelize.OpenReader(file)
	if err != nil {
		logx.Error("open excel file error")
	}

	var asset types.UploadAsset
	var assetList []types.UploadAsset

	rows, err := f.GetRows("Sheet1")
	values := reflect.ValueOf(&asset).Elem()
	for i, r := range rows {
		if i < 1 {
			continue
		}
		for j, v := range r {
			values.Field(j).SetString(v)
		}
		assetList = append(assetList, asset)
	}

	fmt.Printf("%+v", assetList)
	return &types.ImportAssetResp{
		Success: 50,
		Fail:    0,
	}, nil
}
