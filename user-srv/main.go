package main

import (
	"fmt"
	"github.com/entere/parrot/basic"
	"github.com/entere/parrot/basic/config"
	z "github.com/entere/parrot/basic/zap"
	"github.com/entere/parrot/user-srv/handler"
	"github.com/entere/parrot/user-srv/model"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"go.uber.org/zap"

	user "github.com/entere/parrot/user-srv/proto/user"
)

var (
	log *z.Logger
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()
	log = z.GetLogger()
	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("com.island.code.srv.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化handler
			model.Init()
			// 初始化handler
			handler.Init()
		}),
	)

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.Service))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal("[main] service run err", zap.Any("err", err))
	}
}
func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
