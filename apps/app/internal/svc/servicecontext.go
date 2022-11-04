package svc

import (
	"fiber-g/apps/app/internal/config"
	"fiber-g/apps/dept/rpc/deptclient"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	DeptRpc deptclient.Dept
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DeptRpc: deptclient.NewDept(zrpc.MustNewClient(c.DeptRpc)),
	}
}
