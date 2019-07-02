package dao

import (
	"mall/app/api/web/finance/conf"
	memberClient "mall/app/api/web/member/server/http/client"
	"mall/lib/database/orm"

	"github.com/jinzhu/gorm"
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
