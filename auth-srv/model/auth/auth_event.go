/**
 * @Author: entere
 * @Description:
 * @File:  auth_event.go
 * @Version: 1.0.0
 * @Date: 2019/12/18 09:07
 */

package auth

import (
	"context"
	auth "github.com/entere/parrot/auth-srv/proto/auth"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

// sendLogoutEvent 发送退出登陆成功事件
func (s *service) sendLogoutEvent(userID string) {
	ev := &auth.LogoutEvent{
		Id:       uuid.New().String(),
		SentTime: time.Now().Unix(),
		UserID:   userID,
	}

	if err := logoutPublisher.Publish(context.TODO(), ev); err != nil {
		log.Error("logoutPublisher.Publish error", zap.Any("err", err))
	} else {
		log.Info("logoutPublisher.Publish success，用户：" + userID + "退出登录")
	}

}
