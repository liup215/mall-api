package dao

import (
	"mall/app/service/main/member/conf"
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
