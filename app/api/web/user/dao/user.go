package dao

import (
	"mall/app/api/web/user/model"
	"mall/lib/time"

	"github.com/jinzhu/gorm"
)

func (d *Dao) CreateUser(u model.Users) (*model.Users, error) {
	u.Joindate = time.Now()
	u.Starttime = time.Now()
	err := d.orm.Create(&u).Error
	return &u, err
}

func (d *Dao) QueryUser(query model.UserQuery) (*model.Users, error) {

	var u model.Users
	err := d.parseQuery(query).First(&u).Error

	return &u, err
}

func (d *Dao) parseQuery(query model.UserQuery) *gorm.DB {

	var u model.Users
	if query.Uid != 0 {
		u.Uid = query.Uid
	}

	if query.Username != "" {
		u.Username = query.Username
	}

	return d.orm.Where(&u)
}
