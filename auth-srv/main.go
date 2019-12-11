package main

import (
	"fmt"
	"github.com/entere/parrot/auth-srv/handler"
	"github.com/entere/parrot/auth-srv/model"

	"github.com/entere/parrot/basic"
	"github.com/entere/parrot/basic/config"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"

	auth "github.com/entere/parrot/auth-srv/proto/auth"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("com.island.code.srv.auth"),
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

	// 服务注册
	auth.RegisterAuthHandler(service.Server(), new(handler.Service))

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("com.island.code.srv.auth", service.Server(), new(subscriber.Auth))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("com.island.code.srv.auth", service.Server(), subscriber.Handler)

	// 服务启动
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}