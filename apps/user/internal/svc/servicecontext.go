package svc

import (
	"fiber-g/apps/user/internal/config"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := c.DatabaseConf.NewGormDb()
	if err != nil {
		logx.Errorw("user database error", logx.Field("detail", err.Error()))
	}

	return &ServiceContext{
		Config: c,
		Db:     db,
	}
}
