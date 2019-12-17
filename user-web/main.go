package main

import (
	"fmt"
	"github.com/entere/parrot/basic"
	"github.com/entere/parrot/basic/config"
	z "github.com/entere/parrot/basic/zap"
	"github.com/entere/parrot/user-web/handler"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
	"go.uber.org/zap"
	"net/http"
)

var (
	log *z.Logger
)

func main() {
	basic.Init()
	log = z.GetLogger()
	micReg := etcd.NewRegistry(registryOptions)
	// create new web service
	service := web.NewService(
		web.Name("com.island.code.web.user"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8089"),
	)

	// initialise service
	if err := service.Init(
		web.Action(func(context *cli.Context) {
			handler.Init()

		}),
	); err != nil {
		log.Fatal("[server.Init] err", zap.Any("err", err))
	}

	// register html handler
	// service.Handle("/", http.FileServer(http.Dir("html")))

	updatePasswordHandler := http.HandlerFunc(handler.UpdatePassword)
	service.Handle("/password/update", handler.JWTAuthWrapper(updatePasswordHandler))

	service.HandleFunc("/password/modify", handler.UpdatePassword)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal("[server.Run] err", zap.Any("err", err))
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
