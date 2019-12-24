package handler

import (
	"context"
	"github.com/entere/parrot/basic/common"
	"github.com/entere/parrot/pkg/exnet"

	// "context"
	"fmt"
	auth "github.com/entere/parrot/auth-srv/proto/auth"
	"github.com/entere/parrot/basic/token"
	"go.uber.org/zap"

	"github.com/entere/parrot/pkg/rest"
	"net/http"
)

func JWTAuthWrapper(h http.Handler) http.Handler {
	var tk string
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO 从配置中心动态获取白名单URL
		if r.URL.Path == "/auth/login" || r.URL.Path == "/auth/register" || r.URL.Path == "/oauth/test" {
			h.ServeHTTP(w, r)
			return
		}
		ck, _ := r.Cookie(common.CookieNameRememberMe)
		if ck != nil {
			tk = ck.Value
		} else {
			header := r.Header.Get("Authorization")
			fmt.Sscanf(header, "Bearer %s", &tk)
		}

		// log.Info("[JWTAuthWrapper] Token is "+tk, zap.Any("ip", exnet.GetClientIP(r)))
		// token不存在，则状态异常，无权限
		if tk == "" {
			rest.Error(w, "非法请求，没有token", 400)
			return
		}
		claims, err := token.ParseToken(tk)
		if err != nil {
			log.Error("[toke.ParseToken] error ", zap.Any("err", err))
			rest.Error(w, "token解析失败", 500)
			return

		}
		// log.Info("userID is "+claims.Subject, zap.Any("result", claims))
		res, err := authClient.GetCachedAccessToken(context.TODO(), &auth.GetCachedTokenRequest{
			UserID: claims.Subject,
		})
		if err != nil {
			log.Error("[authClient.GetCachedAccessToken] error ", zap.Any("err", err))
			rest.Error(w, "获取服务端token失败", 500)
			return
		}
		if tk != res.Data.Token {
			log.Warn("[tk!=res.Data.Token] error ", zap.Any("ip", exnet.GetClientIP(r)))
			rest.Error(w, "请使用合法token,你的ip已被记录", 500)
			return
		}

		r.Header.Set("X-UserID", claims.Subject)

		h.ServeHTTP(w, r)
	})
}
