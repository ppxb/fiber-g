package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	DeptRpc  zrpc.RpcClientConf
	UserRpc  zrpc.RpcClientConf
	AssetRpc zrpc.RpcClientConf
}
