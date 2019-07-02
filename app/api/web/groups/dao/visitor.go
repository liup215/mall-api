package dao

import (
	"errors"
	"mall/app/api/web/groups/model"
	xtime "mall/lib/time"

	"github.com/jinzhu/gorm"
)

func (d *Dao) CreateVisitor(v model.EweiShopGroupsVisitor) (*model.EweiShopGroupsVisitor, error) {
	if v.Uniacid == 0 {
		return nil, errors.New("无效的应用ID")
	}

	if v.Mid == 0 {
		return nil, errors.New("无效的用户id")
	}

	if v.Openid == "" {
		return nil, errors.New("无效的用户OpenID")
	}
	if v.Id != 0 {
		v.Id = 0
	}

	v.Createdtime = xtime.Now()

	err := d.orm.Create(&v).Error
	return &v, err
}

func (d *Dao) QueryVisitorList(q model.VisitorQuery) ([]model.EweiShopGroupsVisitor, int) {
	db := d.parseVisitorQuery(q)

	list := []model.EweiShopGroupsVisitor{}
	if err := db.Find(&list).Error; err != nil {
		return list, 0
	}
	total := 0
	if err := db.Count(&total).Error; err != nil {
		return list, 0
	}

	return list, total

}

func (d *Dao) QueryVisitorOne(q model.VisitorQuery) (*model.EweiShopGroupsVisitor, error) {
	db := d.parseVisitorQuery(q)

	v := model.EweiShopGroupsVisitor{}
	err := db.First(&v).Error

	return &v, err

}

func (d *Dao) parseVisitorQuery(q model.VisitorQuery) *gorm.DB {
	db := d.orm.Model(&model.EweiShopGroupsVisitor{})
	if q.Id != 0 {
		db = db.Where("id = ?", q.Id)
	}

	if q.Mid != 0 {
		db = db.Where("mid = ?", q.Mid)
	}

	if q.Uniacid != 0 {
		db = db.Where("uniacid = ?", q.Uniacid)
	}

	if q.Aid != "" {
		db = db.Where("aid = ?", q.Aid)
	}

	if q.Openid != "" {
		db = db.Where("openid = ?", q.Openid)
	}
	return db
}
