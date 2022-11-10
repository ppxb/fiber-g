package main

import (
	"fiber-g/apps/user/internal/config"
	"fiber-g/apps/user/internal/server"
	"fiber-g/apps/user/internal/svc"
	"fiber-g/apps/user/user"
	"fiber-g/models"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		user.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	err := ctx.Db.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		logx.Errorw("user database error", logx.Field("detail", err.Error()))
	}

	fmt.Printf("Starting user rpc server at %s...\n", c.ListenOn)
	s.Start()
}
