package dao

import (
	"github.com/jinzhu/gorm"
	//financeClient "mall/app/api/web/finance/server/http/client"
	"mall/app/api/web/wechat/conf"
	"mall/lib/database/orm"
	"mall/lib/database/redis"
)

type Dao struct {
	orm    *gorm.DB
	redigo *redis.Redis
	// fc  *financeClient.Client
}

func New(c *conf.Config) *Dao {
	db := orm.New(c.Orm)
	return &Dao{
		orm:    db,
		redigo: redis.New(c.Redis),
		// fc:  financeClient.New(c.Service.Finance),
	}
}
