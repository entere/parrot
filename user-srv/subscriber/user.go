package subscriber

import (
	"context"
	auth "github.com/entere/parrot/auth-srv/proto/auth"
	z "github.com/entere/parrot/basic/zap"
	um "github.com/entere/parrot/user-srv/model/user"
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

// LogoutLog 消息订阅 收到退出登录的消息后写退出登录日志
func LogoutLog(ctx context.Context, event *auth.LogoutEvent) (err error) {
	log.Info("[LogoutLog] 收到logout通知，userID:" + event.UserID)
	//err = userService.InsertLogoutLog(event)
	//if err != nil {
	//    log.Info("[LogoutLog] 收到退出登录通知，写入退出日志异常",zap.Any("err",err))
	//	return
	//}
	return
}
