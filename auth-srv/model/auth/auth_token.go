package auth

import (
	"fmt"
	"github.com/entere/parrot/basic/token"
	"go.uber.org/zap"
	"time"
)

var (
	// tokenExpiredDate app token过期日期 30天
	tokenExpiredDate = 3600 * 24 * 30 * time.Second

	// tokenIDKeyPrefix tokenID 前缀
	tokenIDKeyPrefix = "token:auth:id:"

	tokenExpiredTopic = "com.island.code.topic.auth.tokenExpired"
)

// Subject token 持有者
type Subject struct {
	UserID   string `json:"userID"`
	DeviceID string `json:"deviceID"`
}

// MakeAccessToken 生成token并保存到redis
func (s *service) MakeAccessToken(subject *Subject) (tk string, err error) {
	tk, err = token.GenerateToken(subject.UserID, subject.DeviceID)
	if err != nil {
		log.Error("[token.GenerateToken] error", zap.Any("err", err))
	}
	// 保存到redis
	err = s.saveTokenToCache(subject, tk)
	if err != nil {
		return "", fmt.Errorf("[MakeAccessToken] 保存token到缓存失败，err: %s", err)
	}

	return
}

// GetCachedAccessToken 获取token
func (s *service) GetCachedAccessToken(subject *Subject) (ret string, err error) {
	ret, err = s.getTokenFromCache(subject)
	if err != nil {
		return "", fmt.Errorf("[GetCachedAccessToken] 从缓存获取token失败，err: %s", err)
	}

	return
}

// DelUserAccessToken 清除用户token
func (s *service) DelUserAccessToken(tk string) (err error) {
	// 解析token字符串
	claims, err := token.ParseToken(tk)
	if err != nil {
		return fmt.Errorf("[DelUserAccessToken] 错误的token，err: %s", err)
	}

	// 通过解析到的用户id删除
	err = s.delTokenFromCache(&Subject{
		UserID: claims.Subject,
	})

	if err != nil {
		return fmt.Errorf("[DelUserAccessToken] 清除用户token，err: %s", err)
	}

	// 广播删除

	//msg := &broker.Message{
	//	Body: []byte(claims.Subject),
	//}
	//if err := broker.Publish(tokenExpiredTopic, msg); err != nil {
	//	log.Error("[pub] 发布token删除消息失败", zap.Any("err", err))
	//} else {
	//	log.Info("[pub] 发布token删除消息成功", zap.String("msg:", string(msg.Body)))
	//}

	// 广播删除
	s.sendLogoutEvent(claims.Subject)

	return
}
