package config

import (
	"fiber-g/pkg/gorm"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf gorm.Conf
	CacheRedis   cache.CacheConf
}
