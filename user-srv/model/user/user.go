package user

import (
	"fmt"

	"github.com/entere/parrot/basic/redis"
	z "github.com/entere/parrot/basic/zap"
	user "github.com/entere/parrot/user-srv/proto/user"
	r "github.com/go-redis/redis"
	"sync"
)

var (
	s   *service
	ca  *r.Client
	m   sync.RWMutex
	log *z.Logger
)

// service 服务
type service struct {
}

// Service 用户服务类
type Service interface {
	GetUser(userID string) (ret *user.GetUserData, err error)
	UpdateUser(req *user.UpdateUserRequest) (err error)
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

	s = &service{}
}
