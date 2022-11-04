package svc

import (
	"fiber-g/apps/dept/rpc/internal/config"
	"fiber-g/apps/dept/rpc/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	DeptModel model.DeptsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:    c,
		DeptModel: model.NewDeptsModel(conn, c.CacheRedis),
	}
}
