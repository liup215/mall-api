package dao

import (
	"mall/app/api/web/order/model"

	"github.com/jinzhu/gorm"
)

func (d *Dao) OrderSumAndTotal(p model.OrderStaticsParam) (*model.OrderStaticsResponse, error) {
	var r []model.OrderStaticsResponse
	db := d.parseOrderStaticsParam(p)
	err := db.Select("count(*) as count, sum(price) as money").Scan(&r).Error

	if err != nil {
		return nil, err
	}

	if len(r) == 0 {
		return &model.OrderStaticsResponse{}, nil
	}

	return &(r[0]), nil

}

func (d *Dao) OrderCount(p model.OrderStaticsParam) (int, error) {
	var count int
	db := d.parseOrderStaticsParam(p)
	err := db.Model(&model.EweiShopOrder{}).Count(&count).Error
	return count, err

}

func (d *Dao) OrderSum(p model.OrderStaticsParam) (float64, error) {
	type Money struct {
		Money float64
	}

	var money []Money

	db := d.parseOrderStaticsParam(p)
	err := db.Select("sum(price) as money").Model(&model.EweiShopOrder{}).Scan(&money).Error
	if err != nil {
		return 0.00, err
	}

	return money[0].Money, err
}

func (d *Dao) parseOrderStaticsParam(p model.OrderStaticsParam) *gorm.DB {
	m := map[string]interface{}{}

	if p.Uniacid != 0 {
		m["uniacid"] = p.Uniacid
	}

	if p.Openid != "" {
		m["openid"] = p.Openid
	}

	if p.Status != -1 {
		m["status"] = p.Status
	}

	return d.orm.Model(&model.EweiShopOrder{}).Where(m)
}
