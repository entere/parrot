package main

import (
	"fmt"
	"github.com/entere/parrot/basic"
	"github.com/entere/parrot/basic/config"
	z "github.com/entere/parrot/basic/zap"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"go.uber.org/zap"
	"net/http"

	"github.com/entere/parrot/auth-web/handler"
	"github.com/micro/go-micro/web"
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
		log.Fatal("[server.Init] err", zap.Any("err", err))
	}

	// register html handler
	// service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/auth/login", handler.Login)
	service.HandleFunc("/auth/logout", handler.Logout)
	service.HandleFunc("/auth/password/modify", handler.UpdatePassword)

	updatePasswordHandler := http.HandlerFunc(handler.UpdatePassword)
	service.Handle("/auth/password/update", handler.JWTAuthWrapper(updatePasswordHandler))

	// run service
	if err := service.Run(); err != nil {
		log.Fatal("[server.Run] err", zap.Any("err", err))
	}
}
func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
