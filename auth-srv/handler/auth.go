package handler

import (
	"context"
	"database/sql"
	am "github.com/entere/parrot/auth-srv/model/auth"
	auth "github.com/entere/parrot/auth-srv/proto/auth"
	z "github.com/entere/parrot/basic/zap"
	"go.uber.org/zap"
)

var (
	authService am.Service
	log         *z.Logger
)

// Init 初始化handler
func Init() {
	var err error
	log = z.GetLogger()
	authService, err = am.GetService()
	if err != nil {
		log.Fatal("初始化Handler错误，%s", zap.Any("err", err))
		return
	}
}

type Service struct{}

// MakeAccessToken 生成token
func (s *Service) MakeAccessToken(ctx context.Context, req *auth.MakeTokenRequest, rsp *auth.MakeTokenResponse) error {

	token, err := authService.MakeAccessToken(&am.Subject{
		// ID:   strconv.FormatInt(req.UserID, 10),
		UserID:   req.UserID,
		DeviceID: req.DeviceID,
	})
	if err != nil {
		rsp.Error = &auth.Error{
			Code: 500,
			Msg:  err.Error(),
		}
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
	err := authService.DelUserAccessToken(req.Token)
	if err != nil {
		rsp.Error = &auth.Error{
			Code: 500,
			Msg:  err.Error(),
		}
		log.Error("[DelUserAccessToken] 清除用户token失败，token:"+req.Token, zap.Any("err", err))
		return err
	}
	rsp.Error = &auth.Error{
		Code: 200,
		Msg:  "success",
	}

	return nil
}

// GetCachedAccessToken 获取缓存的token
func (s *Service) GetCachedAccessToken(ctx context.Context, req *auth.GetCachedTokenRequest, rsp *auth.GetCachedTokenResponse) error {
	token, err := authService.GetCachedAccessToken(&am.Subject{
		UserID:   req.UserID,
		DeviceID: req.DeviceID,
	})
	if err != nil {
		rsp.Error = &auth.Error{
			Code: 500,
			Msg:  err.Error(),
		}
		log.Error("[GetCachedAccessToken] 获取缓存的token失败，userID:"+req.UserID, zap.Any("err", err))
		return err
	}

	rsp.Error = &auth.Error{
		Code: 200,
		Msg:  "Success",
	}
	rsp.Data = &auth.GetCachedTokenData{
		Token: token,
	}
	return nil
}

// LoginByName
func (s *Service) LoginByName(ctx context.Context, req *auth.LoginByNameRequest, rsp *auth.LoginByNameResponse) error {
	var code int32
	user, err := authService.LoginByName(req)
	if err != nil {
		if err == sql.ErrNoRows {
			code = 404
			log.Info("[LoginByName] 用户名不存在 loginName:"+req.LoginName, zap.Any("err", err))
		} else if err.Error() == "密码错误" {
			code = 403
			log.Warn("[LoginByName] 密码错误 loginName:"+req.LoginName, zap.Any("err", err))
		} else {
			code = 500
			log.Warn("[LoginByName] 查询失败 loginName:"+req.LoginName, zap.Any("err", err))
		}
		rsp.Error = &auth.Error{
			Code: code,
			Msg:  err.Error(),
		}
		return err
	}
	rsp.Error = &auth.Error{
		Code: 200,
		Msg:  "success",
	}
	rsp.Data = user
	return nil
}

func (s *Service) UpdatePassword(ctx context.Context, req *auth.UpdatePasswordRequest, rsp *auth.UpdatePasswordResponse) error {

	err := authService.UpdatePassword(req.Password, req.UserID)
	if err != nil {

		rsp.Error = &auth.Error{
			Code: 500,
			Msg:  err.Error(),
		}
		return err
	}
	rsp.Error = &auth.Error{
		Code: 200,
		Msg:  "Success",
	}

	return nil
}
