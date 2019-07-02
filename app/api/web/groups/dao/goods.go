package dao

import (
	"mall/app/api/web/groups/model"

	"github.com/jinzhu/gorm"
)

func (d *Dao) QueryGoodsList(query model.GoodsQuery) ([]model.EweiShopGroupsGoods, int) {
	db := d.parseGoodsQuery(query)

	var list []model.EweiShopGroupsGoods
	if query.Page <= 0 {
		query.Page = 1
	}

	if query.PageSize <= 0 {
		query.PageSize = 20
	}
	db.Offset((query.Page - 1) * query.PageSize).Limit(query.PageSize).Find(&list)

	var total int

	db.Count(&total)

	return list, total

}

func (d *Dao) parseGoodsQuery(q model.GoodsQuery) *gorm.DB {
	db := d.orm.Model(&model.EweiShopGroupsGoods{})

	if q.Uniacid != 0 {
		db = db.Where("uniacid = ?", q.Uniacid)
	}

	switch q.Type {
	case "sale":
		db = db.Where("deleted = ?", 0).Where("stock > ?", 0).Where("status = ?", 1)
	case "sold":
		db = db.Where("deleted = ?", 0).Where("stock <= ?", 0).Where("status = ?", 1)
	case "store":
		db = db.Where("deleted = ?", 0).Where("status = ?", 0)
	case "recycle":
		db = db.Where("deleted = ?", 1)
	default:
		db = db.Where("deleted = ?", 0).Where("stock > ?", 0).Where("status = ?", 1)
	}

	if q.Keyword != "" {
		db = db.Where("title LIKE ?", "%"+q.Keyword+"%")
	}

	return db

}
