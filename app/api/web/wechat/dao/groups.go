package dao

import (
	"mall/app/api/web/wechat/model"

	"github.com/jinzhu/gorm"
)

func (d *Dao) QueryMcGroupsOne(p model.McGroupsQuery) (*model.McGroups, error) {
	db := d.parseMcGroupsQuery(p)

	var g model.McGroups
	err := db.First(&g).Error

	if err != nil {
		return nil, err
	}

	return &g, nil
}

func (d *Dao) parseMcGroupsQuery(p model.McGroupsQuery) *gorm.DB {
	db := d.orm.Model(&model.McGroups{})
	if p.Uniacid != 0 {
		db = db.Where("uniacid = ?", p.Uniacid)
	}

	if p.Isdefault != -1 {
		db = db.Where("isdefault = ?", p.Isdefault)
	}

	return db
}
