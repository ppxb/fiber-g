package svc

import (
	"fiber-g/apps/app/internal/config"
	"fiber-g/apps/dept/rpc/deptclient"
	"fiber-g/apps/user/rpc/userclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	DeptRpc deptclient.Dept
	UserRpc userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DeptRpc: deptclient.NewDept(zrpc.MustNewClient(c.DeptRpc)),
		UserRpc: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
