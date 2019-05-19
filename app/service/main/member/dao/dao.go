package dao

import (
	"github.com/jinzhu/gorm"
	"mall/app/service/main/member/conf"
	"mall/lib/database/orm"
)

type Dao struct {
	db *gorm.DB
}

func New(c *conf.Config) *Dao {
	db := orm.New(c.Orm)
	return &Dao{
		db: db,
	}
}
