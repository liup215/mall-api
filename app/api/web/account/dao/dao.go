package dao

import (
	"github.com/jinzhu/gorm"
	//financeClient "mall/app/api/web/finance/server/http/client"
	"mall/app/api/web/account/conf"
	"mall/lib/database/orm"
	redigo "mall/lib/database/redis"

	"github.com/gomodule/redigo/redis"
)

type Dao struct {
	orm    *gorm.DB
	redigo *redis.Pool
	// fc  *financeClient.Client
}

func New(c *conf.Config) *Dao {
	db := orm.New(c.Orm)
	return &Dao{
		orm:    db,
		redigo: redigo.New(c.Redis),
		// fc:  financeClient.New(c.Service.Finance),
	}
}
