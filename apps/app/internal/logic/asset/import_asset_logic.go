package asset

import (
	"bytes"
	"context"
	"fiber-g/apps/app/internal/svc"
	"fiber-g/apps/app/internal/types"
	"fiber-g/apps/asset/asset"
	"fiber-g/pkg/errorx"
	"fmt"
	"io"
	"net/http"
	"time"

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

func (l *ImportAssetLogic) ImportAsset() (resp *types.ResultWithData, err error) {
	// file size bigger than defined
	err = l.r.ParseMultipartForm(maxFileSize)
	if err != nil {
		return nil, errorx.NewDefaultError("文件过大，上传文件请小于20M")
	}

	file, handler, err := l.r.FormFile("file")
	if err != nil {
		return nil, errorx.NewDefaultError("文件读取失败，请稍后再试")
	}

	buff := new(bytes.Buffer)
	if _, err := io.Copy(buff, file); err != nil {
		return nil, errorx.NewDefaultError("文件解析失败，请稍后再试")
	}

	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	logx.Infof("upload file: %+v, file size: %d, MIME header: %+v",
		handler.Filename, handler.Size, handler.Header)

	res, err := l.svcCtx.AssetRpc.ImportAssets(context.Background(), &asset.UploadReq{
		File: buff.Bytes(),
	})
	if err != nil {
		return nil, errorx.NewDefaultError(err.Error())
	}

	return &types.ResultWithData{
		Code: errorx.OK,
		Msg:  "请求成功",
		Data: map[string]interface{}{
			"success": res.Success,
			"fail":    res.Fail,
		},
		Timestamp: time.Now().Unix(),
	}, nil
}
