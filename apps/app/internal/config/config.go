package config

import (
	"fiber-g/config"
	"fiber-g/pkg/gorm"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	DatabaseConf gorm.Conf
	RedisConf    redis.RedisConf
	CasbinConf   config.CasbinConf
	DeptRpc      zrpc.RpcClientConf
	UserRpc      zrpc.RpcClientConf
	AssetRpc     zrpc.RpcClientConf
}
