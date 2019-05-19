package dao

import (
	"github.com/jinzhu/gorm"
	"mall/app/service/main/adv/conf"
	"mall/lib/database/orm"
)

type Dao struct {
	db *gorm.DB
}

func New(c *conf.Config) *Dao {
	db := orm.New(c.orm)
	return &Dao{
		db: db,
	}
}
