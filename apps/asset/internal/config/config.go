package config

import (
	"fiber-g/pkg/gorm"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	DataSource   string
	CacheRedis   cache.CacheConf
	DatabaseConf gorm.Conf
}
