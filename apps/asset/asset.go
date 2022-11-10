package main

import (
	"fiber-g/apps/asset/asset"
	"fiber-g/apps/asset/internal/config"
	"fiber-g/apps/asset/internal/model"
	"fiber-g/apps/asset/internal/server"
	"fiber-g/apps/asset/internal/svc"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/asset.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		asset.RegisterAssetServer(grpcServer, server.NewAssetServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	// gorm auto create migrate and create table
	err := ctx.Db.AutoMigrate(
		&model.Asset{},
		&model.Project{},
	)
	if err != nil {
		logx.Errorw("database error", logx.Field("detail", err.Error()))
	}

	fmt.Printf("Starting asset rpc server at %s...\n", c.ListenOn)
	s.Start()
}
