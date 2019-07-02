package dao

import (
	"mall/app/api/web/user/conf"
	"mall/lib/database/orm"

	"github.com/jinzhu/gorm"
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
