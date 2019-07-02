package dao

import (
	"mall/app/api/web/wechat/model"

	"github.com/jinzhu/gorm"
)

func (d *Dao) QueryAccountWechatsOne(query model.AccountWechatsQuery) (*model.AccountWechats, error) {
	var ac model.AccountWechats
	db := d.parseAccountWechatsQuery(query)
	err := db.First(&ac).Error
	return &ac, err
}

func (d *Dao) parseAccountWechatsQuery(query model.AccountWechatsQuery) *gorm.DB {
	db := d.orm.Model(&model.AccountWechats{})

	if query.Acid != 0 {
		db = db.Where("acid = ?", query.Acid)
	}

	if query.Uniacid != 0 {
		db = db.Where("uniacid = ?", query.Uniacid)
	}

	return db
}
