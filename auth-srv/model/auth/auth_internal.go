package auth

import (
	"fmt"
)

// saveTokenToCache 保存token到缓存
func (s *service) saveTokenToCache(subject *Subject, token string) (err error) {
	//保存
	if err = ca.Set(tokenIDKeyPrefix+subject.UserID, token, tokenExpiredDate).Err(); err != nil {
		return fmt.Errorf("[saveTokenToCache] 保存token到缓存发生错误，err:" + err.Error())
	}
	return
}

// delTokenFromCache 清空token
func (s *service) delTokenFromCache(subject *Subject) (err error) {
	//保存
	if err = ca.Del(tokenIDKeyPrefix + subject.UserID).Err(); err != nil {
		return fmt.Errorf("[delTokenFromCache] 清空token 缓存发生错误，err:" + err.Error())
	}
	return
}

// getTokenFromCache 从缓存获取token
func (s *service) getTokenFromCache(subject *Subject) (token string, err error) {
	// 获取
	tokenCached, err := ca.Get(tokenIDKeyPrefix + subject.UserID).Result()
	if err != nil {
		return token, fmt.Errorf("[getTokenFromCache] token不存在 %s", err)
	}

	return string(tokenCached), nil
}
