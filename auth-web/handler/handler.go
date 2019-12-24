package handler

import (
	"context"
	"github.com/entere/parrot/basic/common"
	z "github.com/entere/parrot/basic/zap"
	"github.com/entere/parrot/pkg/rest"
	"go.uber.org/zap"
	"net/http"
	"time"

	auth "github.com/entere/parrot/auth-srv/proto/auth"
	"github.com/micro/go-micro/client"
)

var (
	authClient auth.AuthService
	log        *z.Logger
)

func Init() {
	log = z.GetLogger()
	authClient = auth.NewAuthService("com.island.code.srv.auth", client.DefaultClient)
}

type Error struct {
	Code int32  `json:"code"`
	Msg  string `json:"msg"`
}

func Login(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		log.Warn("[Login] 非法请求", zap.Any("ip", r.RemoteAddr))
		rest.Error(w, "非法请求", 400)
		return
	}
	r.ParseForm()

	rsp, err := authClient.LoginByName(context.TODO(), &auth.LoginByNameRequest{
		LoginName: r.Form.Get("loginName"),
		LoginPwd:  r.Form.Get("loginPwd"),
		DeviceID:  r.Form.Get("deviceID"),
	})

	if err != nil {
		log.Warn("[authClient.LoginByName]用户登录失败", zap.Any("err", err))
		rest.Error(w, rsp.Error.Msg, rsp.Error.Code)
		return
	}

	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}

	rsp2, err := authClient.MakeAccessToken(context.TODO(), &auth.MakeTokenRequest{
		UserID: rsp.Data.UserID,
	})
	if err != nil {
		log.Error("[MakeAccessToken] 创建token失败", zap.Any("err", err))
		rest.Error(w, err.Error(), 500)
		return
	}
	//response["token"] = rsp2.Data.Token
	//response["error"] = &Error{
	//	Code: 200,
	//	Msg:  "success",
	//}
	rsp.Data.LoginPwd = ""
	rsp.Data.Token = rsp2.Data.Token
	response["data"] = rsp.Data
	w.Header().Add("Access-Control-Allow-Origin", "http://192.168.1.68:8081")
	w.Header().Add("Access-Control-Allow-Credentials", "true")
	w.Header().Add("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Authorization, Organization-Token, Origin, X-Requested-With, Content-Type, Accept, enctype, Referer, User-Agent, Access-Control-Request-Headers, Access-Control-Request-Method")
	w.Header().Add("Access-Control-Allow-Methords", "GET,POST,OPTIONS,HEAD,PUT,DELETE")
	w.Header().Add("Access-Control-Expose-Headers", "Set-Cookie")
	w.Header().Add("set-cookie", "application/json; charset=utf-8")

	expire := time.Now().Add(30 * time.Minute)
	cookie := http.Cookie{Name: common.CookieNameRememberMe, Value: rsp2.Data.Token, Path: "/", Expires: expire, MaxAge: 90000}
	http.SetCookie(w, &cookie)

	rest.Response(w, "Success", 200, rsp.Data)

}

// Logout 退出登录
func Logout(w http.ResponseWriter, r *http.Request) {
	// 只接受POST请求
	if r.Method != "POST" {
		log.Warn("[Logout] 非法请求", zap.Any("ip", r.RemoteAddr))
		rest.Error(w, "非法请求", 400)
		return
	}

	tokenCookie, err := r.Cookie(common.CookieNameRememberMe)
	if err != nil {
		log.Warn("[r.Cookie] tokenCookie获取失败", zap.Any("err", err))
		rest.Error(w, "非法请求", 400)
		return
	}

	// 删除token
	_, err = authClient.DelUserAccessToken(context.TODO(), &auth.DelTokenRequest{
		Token: tokenCookie.Value,
	})
	if err != nil {
		rest.Error(w, err.Error(), 500)
		return
	}

	// 清除cookie
	cookie := http.Cookie{Name: common.CookieNameRememberMe, Value: "", Path: "/", Expires: time.Now().Add(0 * time.Second), MaxAge: 0}
	http.SetCookie(w, &cookie)

	rest.Success(w, nil)

	//w.Header().Add("Content-Type", "application/json; charset=utf-8")
	//// 返回结果
	//response := map[string]interface{}{
	//	"ref": time.Now().UnixNano(),
	//	"error": &Error{
	//		Code: 200,
	//		Msg:  "success",
	//	},
	//}
	//// 返回JSON结构
	//if err := json.NewEncoder(w).Encode(response); err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
}

func UpdatePassword(w http.ResponseWriter, r *http.Request) {

	if r.Method != "POST" {
		log.Warn("[Login] 非法请求", zap.Any("ip", r.RemoteAddr))
		rest.Error(w, "非法请求", 400)
		return
	}
	r.ParseForm()

	rsp, err := authClient.UpdatePassword(context.TODO(), &auth.UpdatePasswordRequest{
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
