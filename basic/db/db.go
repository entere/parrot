package db

import (
	"database/sql"
	"github.com/entere/parrot/basic/config"
	"github.com/micro/go-micro/util/log"

	"sync"
)

var (
	inited  bool
	mysqlDB *sql.DB
	m       sync.RWMutex
)

// Init 初始化数据库
func Init() {
	m.Lock()
	defer m.Unlock()
	if inited {

		log.Warn("[Init] 已经初始化过db...")
		return
	}

	// 如果配置声明使用mysql
	if config.GetMysqlConfig().GetEnabled() {
		initMysql()
	}
	// log.Info("[Init] 初始化db...")
	inited = true
}

// GetDB 获取db
func GetDB() *sql.DB {
	return mysqlDB
}
