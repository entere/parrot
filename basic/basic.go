package basic

import (
	"github.com/entere/parrot/basic/config"
	"github.com/entere/parrot/basic/db"
	"github.com/entere/parrot/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
