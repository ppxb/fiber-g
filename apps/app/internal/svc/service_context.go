package svc

import (
	"fiber-g/apps/app/internal/config"
	"fiber-g/apps/asset/assetclient"
	"fiber-g/apps/dept/rpc/deptclient"
	"fiber-g/apps/user/userclient"
	"github.com/casbin/casbin/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	Redis    *redis.Redis
	Casbin   *casbin.Enforcer
	DeptRpc  deptclient.Dept
	UserRpc  userclient.User
	AssetRpc assetclient.Asset
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := c.RedisConf.NewRedis()
	if !rds.Ping() {
		logx.Error("initialize redis failed")
		return nil
	}
	cbn, err := c.CasbinConf.NewCasbin(c.DatabaseConf)
	if err != nil {
		logx.Errorw("initialize casbin failed", logx.Field("detail", err.Error()))
		return nil
	}
	return &ServiceContext{
		Config:   c,
		Redis:    rds,
		Casbin:   cbn,
		DeptRpc:  deptclient.NewDept(zrpc.MustNewClient(c.DeptRpc)),
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		AssetRpc: assetclient.NewAsset(zrpc.MustNewClient(c.AssetRpc)),
	}
}
