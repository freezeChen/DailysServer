package dao

import (
	"DailysServer/proto"
	"context"
	"fmt"

	"DailysServer/logic/conf"
	models "DailysServer/logic/model"

	"github.com/freezeChen/studio-library/database/mysql"
	"github.com/freezeChen/studio-library/redis"
	"github.com/freezeChen/studio-library/zlog"
	redis2 "github.com/garyburd/redigo/redis"
	"github.com/go-xorm/xorm"
)

const (
	_prefixKeyServer = "key:%d" //uid => server

)

func keyServer(uid int64) string {
	return fmt.Sprintf(_prefixKeyServer, uid)
}

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

func (d *Dao) GetContactList(uid int64) (list []*proto.Conversation, err error) {
	list = make([]*proto.Conversation, 0)
	err = d.db.SQL(`select c.other_uid as uid,u.name as userName,c.mid,m.content,m.create_time as time from contact c 
			join user u on c.other_uid = u.id
			left join message m on c.mid =m.id
			where c.owner_uid=?;`, uid).Find(&list)
	if err != nil {
		return nil, err
	}

	conn := d.redis.GetConn()
	defer conn.Close()

	for i := 0; i < len(list); i++ {
		err = conn.Send("EXISTS", keyServer(list[i].Uid))
		if err != nil {
			zlog.Errorf("redis send (EXISTS) is error(%v)", err)
			break
		}
	}
	if err = conn.Flush(); err != nil {
		zlog.Errorf("redis flush is error(%s)", err)
		return
	}
	var exist bool
	for i := 0; i < len(list); i++ {

		if exist, err = redis2.Bool(conn.Receive()); err != nil {
			return
		}

		list[i].Online = exist

	}

	return list, nil
}

func (d *Dao) Online(sid string, uid int64) error {
	conn := d.redis.GetConn()
	defer conn.Close()

	key := keyServer(uid)
	err := d.redis.Set(context.Background(), key, sid, 60*30)
	return err
}

/*func (d *Dao) userOnline(uid string) bool {
	conn := d.redis.GetConn()
	defer conn.Close()

}*/
