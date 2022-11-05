package svc

import (
	"fiber-g/apps/asset/internal/config"
	"fiber-g/apps/asset/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config     config.Config
	AssetModel model.AssetsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:     c,
		AssetModel: model.NewAssetsModel(conn, c.CacheRedis),
	}
}
