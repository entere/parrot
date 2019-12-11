package handler

import (
	"context"
	am "github.com/entere/parrot/auth-srv/model/auth"
	auth "github.com/entere/parrot/auth-srv/proto/auth"
	"github.com/micro/go-micro/util/log"
)

var (
	authService am.Service
)

// Init 初始化handler
func Init() {
	var err error
	authService, err = am.GetService()
	if err != nil {
		log.Fatal("[Init] 初始化Handler错误，%s", err)
		return
	}
}

type Service struct{}

// MakeAccessToken 生成token
func (s *Service) MakeAccessToken(ctx context.Context, req *auth.MakeTokenRequest, rsp *auth.MakeTokenResponse) error {
	log.Log("[MakeAccessToken] 收到创建token请求")

	token, err := authService.MakeAccessToken(&am.Subject{
		// ID:   strconv.FormatInt(req.UserID, 10),
		ID: req.UserID,
	})
	if err != nil {
		rsp.Error = &auth.Error{
			Code: 500,
			Msg:  err.Error(),
		}

		log.Logf("[MakeAccessToken] token生成失败，err：%s", err)
		return err
	}

	rsp.Error = &auth.Error{
		Code: 200,
		Msg:  "success",
	}
	rsp.Data = &auth.MakeTokenData{
		Token: token,
	}
	return nil
}

// DelUserAccessToken 清除用户token
func (s *Service) DelUserAccessToken(ctx context.Context, req *auth.DelTokenRequest, rsp *auth.DelTokenResponse) error {
	log.Log("[DelUserAccessToken] 清除用户token")
	err := authService.DelUserAccessToken(req.Token)
	if err != nil {
		rsp.Error = &auth.Error{
			Code: 500,
			Msg:  err.Error(),
		}

		log.Logf("[DelUserAccessToken] 清除用户token失败，err：%s", err)
		return err
	}
	rsp.Error = &auth.Error{
		Code: 200,
		Msg:  "success",
	}

	return nil
}

// GetCachedAccessToken 获取缓存的token
func (s *Service) GetCachedAccessToken(ctx context.Context, req *auth.MakeTokenRequest, rsp *auth.MakeTokenResponse) error {
	log.Logf("[GetCachedAccessToken] 获取缓存的token，%d", req.UserID)
	token, err := authService.GetCachedAccessToken(&am.Subject{
		ID: req.UserID,
	})
	if err != nil {
		rsp.Error = &auth.Error{
			Code: 500,
			Msg:  err.Error(),
		}

		log.Logf("[GetCachedAccessToken] 获取缓存的token失败，err：%s", err)
		return err
	}

	rsp.Error = &auth.Error{
		Code: 200,
		Msg:  "success",
	}
	rsp.Data = &auth.MakeTokenData{
		Token: token,
	}
	return nil
}

// QueryUserByName
func (s *Service) QueryUserByName(ctx context.Context, req *auth.QueryUserRequest, rsp *auth.QueryUserResponse) error {
	log.Log("[QueryUserByName] 收到查询请求")
	user, err := authService.QueryUserByName(req.LoginName)
	if err != nil {
		rsp.Error = &auth.Error{
			Code: 500,
			Msg:  err.Error(),
		}

		log.Logf("[QueryUserByName] 查询失败，err：%s", err)
		return err
	}
	rsp.Error = &auth.Error{
		Code: 200,
		Msg:  "success",
	}
	rsp.Data = user
	return nil
}
