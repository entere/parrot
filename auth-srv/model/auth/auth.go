package auth

import (
	"fmt"
	auth "github.com/entere/parrot/auth-srv/proto/auth"
	"github.com/entere/parrot/basic/common"
	"github.com/entere/parrot/basic/redis"
	z "github.com/entere/parrot/basic/zap"
	r "github.com/go-redis/redis"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
	"sync"
)

var (
	s               *service
	ca              *r.Client
	m               sync.RWMutex
	log             *z.Logger
	logoutPublisher micro.Publisher
)

// service 服务
type service struct {
}

// Service 用户服务类
type Service interface {
	// MakeAccessToken 生成token
	MakeAccessToken(subject *Subject) (ret string, err error)

	// GetCachedAccessToken 获取缓存的token
	GetCachedAccessToken(subject *Subject) (ret string, err error)

	// DelUserAccessToken 清除用户token
	DelUserAccessToken(token string) (err error)

	LoginByName(req *auth.LoginByNameRequest) (ret *auth.LoginByNameData, err error)
	UpdatePassword(password string, userID string) (err error)
}

// GetService 获取服务类
func GetService() (Service, error) {
	if s == nil {
		return nil, fmt.Errorf("[GetService] GetService 未初始化")
	}
	return s, nil
}

// Init 初始化用户服务层
func Init() {
	m.Lock()
	defer m.Unlock()

	if s != nil {
		return
	}

	ca = redis.GetRedis()
	log = z.GetLogger()

	logoutPublisher = micro.NewPublisher(common.TopicLogout, client.DefaultClient)

	s = &service{}
}
