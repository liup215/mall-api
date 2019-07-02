package dao

import (
	"errors"
	"mall/app/api/web/groups/model"
	"mall/lib/strings"
	xtime "mall/lib/time"
	"sync"

	"github.com/jinzhu/gorm"
)

var orderLock sync.Mutex

func (d *Dao) CreateOrder(order model.EweiShopGroupsOrder) (*model.EweiShopGroupsOrder, error) {
	if order.Id != 0 {
		order.Id = 0
	}

	if order.Uniacid == 0 {
		return nil, errors.New("无效的应用ID")
	}

	if order.Openid == "" {
		return nil, errors.New("无效的用户OpenID")
	}

	if order.Price == 0 {
		return nil, errors.New("支付金额不能为空")
	}

	orderLock.Lock()
	defer orderLock.Unlock()

	tx := d.orm.Begin()
	// 生成订单号
	orderno := strings.CreateNo("PT")
	for {
		count := 0
		err := tx.Model(&model.EweiShopGroupsOrder{}).Where("orderno = ?", orderno).Count(&count).Error
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				tx.Rollback()
				return nil, err
			}
		}

		if count <= 0 {
			break
		}

		orderno = strings.CreateNo("PT")

	}

	order.Orderno = orderno
	order.Createtime = xtime.Now()
	order.Status = 0

	// 如果是拼团但没有团购id，就生成新的
	if order.IsTeam == 1 && order.Teamid == 0 {
		rows, err := tx.Model(&model.EweiShopGroupsOrder{}).Select("max(teamid) as max").Rows()
		if err != nil {
			if err != gorm.ErrRecordNotFound {
				tx.Rollback()
				return nil, err
			}
		}
		maxTeamid := 0
		for rows.Next() {
			err := rows.Scan(&maxTeamid)
			if err != nil {
				tx.Rollback()
				return nil, err
			}
		}

		order.Teamid = maxTeamid + 1
	}

	if err := tx.Create(&order).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return &order, nil
}

func (d *Dao) OrderCount(q model.OrderQuery) (int, error) {
	var count int
	db := d.parseOrderQuery(q)
	err := db.Count(&count).Error
	return count, err

}

func (d *Dao) QueryOrderOne(q model.OrderQuery) (*model.EweiShopGroupsOrder, error) {
	db := d.parseOrderQuery(q)
	var o model.EweiShopGroupsOrder
	err := db.First(&o).Error
	return &o, err
}

func (d *Dao) QueryOrderList(q model.OrderQuery) ([]model.EweiShopGroupsOrder, int, error) {
	if q.Page <= 0 {
		q.Page = 1
	}

	if q.PageSize <= 0 {
		q.PageSize = 20
	}
	db := d.parseOrderQuery(q)

	list := []model.EweiShopGroupsOrder{}
	if err := db.Offset((q.Page - 1) * q.PageSize).Limit(q.PageSize).Find(&list).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return list, 0, nil
		}
		return list, 0, err
	}

	total := 0
	if err := db.Count(&total).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return list, 0, nil
		}
		return list, 0, err
	}
	return list, total, nil

}

func (d *Dao) parseOrderQuery(q model.OrderQuery) *gorm.DB {
	db := d.orm.Model(&model.EweiShopGroupsOrder{})
	if q.Id != 0 {
		db = db.Where("id = ?", q.Id)
	}

	if q.Uniacid != 0 {
		db = db.Where("uniacid = ?", q.Uniacid)
	}

	if q.Openid != "" {
		db = db.Where("openid = ?", q.Openid)
	}

	if q.Aid != "" {
		db = db.Where("aid = ?", q.Aid)
	}

	if q.Orderno != "" {
		db = db.Where("orderno = ?", q.Orderno)
	}

	if q.Status != -1 {
		q.Statuses = append(q.Statuses, q.Status)
	}

	if len(q.Statuses) > 0 {
		db = db.Where("status IN (?)", q.Statuses)
	}

	return db
}
