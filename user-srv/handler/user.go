package handler

import (
	"context"
	"database/sql"
	z "github.com/entere/parrot/basic/zap"
	um "github.com/entere/parrot/user-srv/model/user"
	user "github.com/entere/parrot/user-srv/proto/user"
	"go.uber.org/zap"
)

var (
	userService um.Service
	log         *z.Logger
)

// Init 初始化handler
func Init() {
	var err error
	log = z.GetLogger()
	userService, err = um.GetService()
	if err != nil {
		log.Fatal("初始化Handler错误，%s", zap.Any("err", err))
		return
	}
}

type Service struct{}

func (s *Service) GetUser(ctx context.Context, req *user.GetUserRequest, rsp *user.GetUserResponse) error {
	// time.Sleep(time.Second * 2)
	var code int32
	userData, err := userService.GetUser(req.UserID)
	if err != nil {
		if err == sql.ErrNoRows {
			code = 404
			log.Info("[GetUser] 用户不存在 userID:"+req.UserID, zap.Any("err", err))
		} else {
			code = 500
			log.Warn("[GetUser] 查询失败 userID:"+req.UserID, zap.Any("err", err))
		}
		rsp.Error = &user.Error{
			Code: code,
			Msg:  err.Error(),
		}
		return err
	}
	rsp.Error = &user.Error{
		Code: 200,
		Msg:  "Success",
	}
	rsp.Data = userData
	return nil
}

func (s *Service) UpdateUser(ctx context.Context, req *user.UpdateUserRequest, rsp *user.UpdateUserResponse) error {

	err := userService.UpdateUser(req)
	if err != nil {

		rsp.Error = &user.Error{
			Code: 500,
			Msg:  err.Error(),
		}
		return err
	}
	rsp.Error = &user.Error{
		Code: 200,
		Msg:  "Success",
	}

	return nil
}
