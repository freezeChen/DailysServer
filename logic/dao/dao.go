package dao

import (
	"DailysServer/logic/conf"

	"github.com/freezeChen/studio-library/database/mysql"
	"github.com/freezeChen/studio-library/redis"
	"github.com/go-xorm/xorm"
)

type Dao struct {
	db    xorm.EngineInterface
	redis *redis.Redis
}

func New(c *conf.Config) *Dao {
	return &Dao{
		db:    mysql.New(c.Mysql),
		redis: redis.New(c.Redis),
	}
}
