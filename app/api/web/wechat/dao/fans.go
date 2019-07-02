package dao

import (
	"mall/app/api/web/wechat/model"

	"github.com/jinzhu/gorm"
)

func (d *Dao) CreateMcMappingFans(fans model.McMappingFans) (*model.McMappingFans, error) {
	if fans.Fanid != 0 {
		fans.Fanid = 0
	}
	err := d.orm.Create(&fans).Error
	return &fans, err
}

func (d *Dao) QueryMcMappingFansOne(q model.FansQuery) (*model.McMappingFans, error) {
	var fans model.McMappingFans

	db := d.parseMcMappingFansQuery(q)
	err := db.First(&fans).Error

	return &fans, err

}

func (d *Dao) parseMcMappingFansQuery(q model.FansQuery) *gorm.DB {
	db := d.orm.Model(&model.McMappingFans{})
	if q.Fanid != 0 {
		db = db.Where("fanid = ?", q.Fanid)
	}

	if q.Uniacid != 0 {
		db = db.Where("uniacid = ?", q.Uniacid)
	}

	if q.Openid != "" {
		db = db.Where("openid = ?", q.Openid)
	}

	return db
}
