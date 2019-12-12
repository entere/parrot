package handler

import (
	"context"
	"encoding/json"
	z "github.com/entere/parrot/basic/zap"
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
		log.Error("[Login] 非法请求", zap.Any("ip", r.RemoteAddr))
		http.Error(w, "非法请求", 400)
		return
	}
	r.ParseForm()

	rsp, err := authClient.QueryUserByName(context.TODO(), &auth.QueryUserRequest{
		LoginName: r.Form.Get("loginName"),
	})

	if err != nil {
		log.Error("[QueryUserByName] 查询用户失败", zap.Any("err", err))
		http.Error(w, err.Error(), 500)
	}

	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}
	if rsp.Data.LoginPwd == r.Form.Get("loginPwd") {

		rsp2, err := authClient.MakeAccessToken(context.TODO(), &auth.MakeTokenRequest{
			UserID: rsp.Data.UserID,
		})
		if err != nil {
			log.Error("[MakeAccessToken] 创建token失败", zap.Any("err", err))
			http.Error(w, err.Error(), 500)
			return
		}
		//response["token"] = rsp2.Data.Token
		response["error"] = &Error{
			Code: 200,
			Msg:  "success",
		}
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
		cookie := http.Cookie{Name: "remember-me-token", Value: rsp2.Data.Token, Path: "/", Expires: expire, MaxAge: 90000}
		http.SetCookie(w, &cookie)
	} else {

		response["error"] = &Error{
			Code: 403,
			Msg:  "密码错误",
		}
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// Logout 退出登录
func Logout(w http.ResponseWriter, r *http.Request) {
	// 只接受POST请求
	if r.Method != "POST" {
		log.Error("[Logout] 非法请求", zap.Any("ip", r.RemoteAddr))
		http.Error(w, "非法请求", 400)
		return
	}

	tokenCookie, err := r.Cookie("remember-me-token")
	if err != nil {
		log.Error("[r.Cookie] tokenCookie获取失败", zap.Any("err", err))
		http.Error(w, "非法请求", 400)
		return
	}

	// 删除token
	_, err = authClient.DelUserAccessToken(context.TODO(), &auth.DelTokenRequest{
		Token: tokenCookie.Value,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// 清除cookie
	cookie := http.Cookie{Name: "remember-me-token", Value: "", Path: "/", Expires: time.Now().Add(0 * time.Second), MaxAge: 0}
	http.SetCookie(w, &cookie)

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回结果
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
		"error": &Error{
			Code: 200,
			Msg:  "success",
		},
	}

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
