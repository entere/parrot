package handler

import (
	"context"
	auth "github.com/entere/parrot/auth-srv/proto/auth"
	"github.com/entere/parrot/pkg/rest"

	z "github.com/entere/parrot/basic/zap"
	"go.uber.org/zap"
	"net/http"

	user "github.com/entere/parrot/user-srv/proto/user"
	"github.com/micro/go-micro/client"
)

var (
	userClient user.UserService
	authClient auth.AuthService
	log        *z.Logger
)

func Init() {
	log = z.GetLogger()
	userClient = user.NewUserService("com.island.code.srv.user", client.DefaultClient)
	authClient = auth.NewAuthService("com.island.code.srv.auth", client.DefaultClient)
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		log.Warn("[Login] 非法请求", zap.Any("ip", r.RemoteAddr))
		rest.Error(w, "非法请求", 400)
		return
	}
	r.ParseForm()

	rsp, err := userClient.UpdatePassword(context.TODO(), &user.UpdatePasswordRequest{
		Password: r.Form.Get("password"),
		UserID:   r.Form.Get("userID"),
		// LoginName: r.Form.Get("loginName"),
	})

	if err != nil {
		log.Error("[UpdatePassword] 修改密码失败", zap.Any("err", err))
		rest.Error(w, err.Error(), 500)
	}

	rest.Response(w, rsp.Error.Msg, rsp.Error.Code, nil)

}
