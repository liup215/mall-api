package dao

import (
	"github.com/jinzhu/gorm"
	//financeClient "mall/app/api/web/finance/server/http/client"
	"mall/app/api/web/member/conf"
	"mall/lib/database/orm"
)

type Dao struct {
	orm *gorm.DB
	// fc  *financeClient.Client
}

func New(c *conf.Config) *Dao {
	db := orm.New(c.Orm)
	return &Dao{
		orm: db,
		// fc:  financeClient.New(c.Service.Finance),
	}
}
