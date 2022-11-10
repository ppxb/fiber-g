package main

import (
	"errors"
	"fiber-g/apps/user/internal/config"
	"fiber-g/apps/user/internal/server"
	"fiber-g/apps/user/internal/svc"
	"fiber-g/apps/user/user"
	"fiber-g/models"
	"fiber-g/pkg/utils"
	"flag"
	"fmt"
	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

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

	// insert seed data to database when first initialized
	if err := ctx.Db.AutoMigrate(
		&models.User{},
		&models.Role{},
	); err == nil && ctx.Db.Migrator().HasTable(&models.User{}) {
		if err := ctx.Db.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			roleSort := new(uint)
			*roleSort = 0
			initSuperRole := &models.Role{
				Name: "超级管理员",
				Sort: roleSort,
			}
			initSuperRole.UUID = uuid.NewString()
			if err := ctx.Db.Table("tb_role").Create(&initSuperRole).Error; err != nil {
				logx.Errorw("init role data error", logx.Field("detail", err.Error()))
			}
			initSuperUser := &models.User{
				Mobile:   "18121954650",
				Password: utils.GenPwd("111"),
				Name:     "连辰",
				Avatar:   "https://gss0.baidu.com/-vo3dSag_xI4khGko9WTAnF6hhy/zhidao/pic/item/77c6a7efce1b9d1618c3302df9deb48f8c546472.jpg",
				RoleId:   initSuperRole.Id,
			}
			initSuperUser.UUID = uuid.NewString()
			if err := ctx.Db.Table("tb_user").Create(&initSuperUser).Error; err != nil {
				logx.Errorw("init user data error", logx.Field("detail", err.Error()))
			}
		} else {
			logx.Errorw("user database error", logx.Field("detail", err.Error()))
		}

	}

	fmt.Printf("Starting user rpc server at %s...\n", c.ListenOn)
	s.Start()
}
