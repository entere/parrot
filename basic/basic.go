package basic

import (
	"github.com/entere/parrot/basic/config"
	"github.com/entere/parrot/basic/db"
	"github.com/entere/parrot/basic/redis"
	"github.com/entere/parrot/basic/zap"
)

func Init() {

	config.Init()
	db.Init()
	redis.Init()
	zap.Init()

}
