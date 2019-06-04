package dao

import (
	"github.com/jinzhu/gorm"
	"mall/app/service/main/finance/conf"
	memberClient "mall/app/service/main/member/server/http/client"
	"mall/lib/database/orm"
)

type Dao struct {
	orm     *gorm.DB
	mclient *memberClient.Client
}

func New(c *conf.Config) *Dao {
	db := orm.New(c.Orm)
	return &Dao{
		orm: db,
	}
}
