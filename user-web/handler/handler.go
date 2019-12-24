package handler

import (
	"context"
	auth "github.com/entere/parrot/auth-srv/proto/auth"
	"github.com/entere/parrot/pkg/rest"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix"
	"strconv"

	hystrix_go "github.com/afex/hystrix-go/hystrix"
	z "github.com/entere/parrot/basic/zap"
	user "github.com/entere/parrot/user-srv/proto/user"
	"github.com/micro/go-micro/client"
	"go.uber.org/zap"
	"net/http"
)

var (
	userClient user.UserService
	authClient auth.AuthService
	log        *z.Logger
)

func Init() {

	// 熔断处理
	hystrix_go.DefaultTimeout = 3000      // 多长时间后超时
	hystrix_go.DefaultVolumeThreshold = 1 //
	hystrix_go.DefaultErrorPercentThreshold = 1
	cl := hystrix.NewClientWrapper()(client.DefaultClient)
	cl.Init(
		client.Retries(3), //重试3次
		client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (bool, error) {
			log.Info(req.Method() + "client retry  retryCount:" + strconv.Itoa(retryCount))
			return true, nil

		}),
	)

	log = z.GetLogger()
	userClient = user.NewUserService("com.island.code.srv.user", cl)
	authClient = auth.NewAuthService("com.island.code.srv.auth", cl)
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		log.Warn("[Login] 非法请求", zap.Any("ip", r.RemoteAddr))
		rest.Error(w, "非法请求", 400)
		return
	}
	r.ParseForm()

	log.Info(r.RequestURI)
	rsp, err := userClient.GetUser(context.TODO(), &user.GetUserRequest{
		UserID: r.Form.Get("userID"),
	})

	if err != nil {
		log.Error("[userClient.GetUser] 查询用户信息失败", zap.Any("err", err))
		rest.Error(w, err.Error(), 500)
	}

	rest.Response(w, rsp.Error.Msg, rsp.Error.Code, rsp.Data)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		log.Warn("[Login] 非法请求", zap.Any("ip", r.RemoteAddr))
		rest.Error(w, "非法请求", 400)
		return
	}
	r.ParseForm()

	rsp, err := userClient.UpdateUser(context.TODO(), &user.UpdateUserRequest{
		AvatarURL: r.Form.Get("avatarURL"),
		UserID:    r.Form.Get("userID"),
		Mobile:    r.Form.Get("mobile"),
	})

	if err != nil {
		log.Error("[UpdateUser] 修改用户信息失败", zap.Any("err", err))
		rest.Error(w, err.Error(), 500)
	}

	rest.Response(w, rsp.Error.Msg, rsp.Error.Code, nil)

}
