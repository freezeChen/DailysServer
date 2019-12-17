package dao

import (
	"DailysServer/logic/conf"
	models "DailysServer/logic/model"

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

func (d *Dao) NewSession() *xorm.Session {
	return d.db.NewSession()
}

func (d *Dao) GetContact(ownerId, otherId int64) *models.Contact {
	var item models.Contact
	has, err := d.db.Where("owner_uid=? and other_uid=?", ownerId, otherId).Get(&item)
	if err != nil {
		return nil
	}
	if !has {
		return nil
	}

	return &item
}
