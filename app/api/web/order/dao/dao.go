package dao

import (
	"mall/app/api/web/order/conf"
	"mall/lib/database/orm"

	"github.com/jinzhu/gorm"
)

type Dao struct {
	orm *gorm.DB
}

func New(c *conf.Config) *Dao {
	db := orm.New(c.Orm)
	return &Dao{
		orm: db,
	}
}
