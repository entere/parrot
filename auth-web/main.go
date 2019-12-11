package main

import (
	"fmt"
	"github.com/entere/parrot/basic"
	"github.com/entere/parrot/basic/config"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"

	"github.com/entere/parrot/auth-web/handler"
	"github.com/micro/go-micro/web"
)

func main() {
	basic.Init()
	micReg := etcd.NewRegistry(registryOptions)
	// create new web service
	service := web.NewService(
		web.Name("com.island.code.web.auth"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8088"),
	)

	// initialise service
	if err := service.Init(
		web.Action(func(context *cli.Context) {
			handler.Init()

		}),
	); err != nil {
		log.Fatal(err)
	}

	// register html handler
	// service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/auth/login", handler.Login)
	service.HandleFunc("/auth/logout", handler.Logout)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
