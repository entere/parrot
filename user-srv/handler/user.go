package handler

import (
	"context"
	z "github.com/entere/parrot/basic/zap"
	um "github.com/entere/parrot/user-srv/model/user"
	"go.uber.org/zap"

	user "github.com/entere/parrot/user-srv/proto/user"
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

func (s *Service) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest, rsp *user.UpdatePasswordResponse) error {

	err := userService.UpdatePassword(req.Password, req.UserID)
	if err != nil {

		rsp.Error = &user.Error{
			Code: 500,
			Msg:  err.Error(),
		}
		return err
	}
	rsp.Error = &user.Error{
		Code: 200,
		Msg:  "success",
	}

	return nil
}
