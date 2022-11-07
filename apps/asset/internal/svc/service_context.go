package svc

import (
	"fiber-g/apps/asset/internal/config"
	"fiber-g/apps/asset/internal/model"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config     config.Config
	AssetModel model.AssetsModel
	Db         *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	db, err := c.DatabaseConf.NewGormDb()
	if err != nil {
		logx.Errorw("database error", logx.Field("detail", err.Error()))
	}

	return &ServiceContext{
		Config:     c,
		AssetModel: model.NewAssetsModel(conn, c.CacheRedis),
		Db:         db,
	}
}
