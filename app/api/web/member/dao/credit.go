package dao

import (
	"errors"
	"mall/app/api/web/member/model"
	"sync"

	"github.com/jinzhu/gorm"
)

var updateCreditLock sync.Mutex

func (d *Dao) UpdateCredit(arg model.UpdateCreditParam) error {
	if arg.Uniacid == 0 {
		return errors.New("无效的应用ID")
	}

	if arg.Openid == "" {
		return errors.New("无效的用户Openid")
	}

	updateCreditLock.Lock()
	defer updateCreditLock.Unlock()

	tx := d.orm.Begin()

	var m model.EweiShopMember
	if err := tx.Where(&model.EweiShopMember{Uniacid: arg.Uniacid, Openid: arg.Openid}).First(&m).Error; err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return errors.New("账号不存在")
		}
		return errors.New("用户查询错误," + err.Error())
	}

	u := map[string]float64{}
	switch arg.Type {
	case 1:
		if m.Credit1+arg.Money < 0 {
			tx.Rollback()
			return errors.New("余额不足")
		} else {
			u["credit1"] = m.Credit1 + arg.Money
		}
	case 0:
		if m.Credit2+arg.Money < 0 {
			tx.Rollback()
			return errors.New("余额不足")
		} else {
			u["credit2"] = m.Credit2 + arg.Money
		}
	default:
		tx.Rollback()
		return errors.New("无效的金额类型")
	}

	err := tx.Model(&model.EweiShopMember{}).Where(map[string]interface{}{"uniacid": arg.Uniacid, "openid": arg.Openid}).Updates(u).Error
	if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}

	return err

}
