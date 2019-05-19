package dao

import (
	"errors"
	"mall/app/service/main/adv/model"
)

func (dao *Dao) AddAdv(adv model.Adv) (int, error) {
	if adv.Id != 0 || adv.Uniacid == 0 {
		return errors.New("无效的参数")
	}
	err := dao.orm.Create(&adv).Error
	return int(adv.Id), err
}

func (dao *Dao) EditAdv(adv model.Adv) error {
	if adv.Id == 0 || adv.Uniacid == 0 {
		return errors.New("参数无效")
	}
	return dao.orm.Save(&adv).Error
}

func (dao *Dao) Delete(uniacid, aid int) error {
	if uniacid == 0 || aid == 0 {
		return errors.New("Aid 或Uniacid不能为空")
	}
	return dao.orm.Delete(&model.Adv{Id: aid, Uniacid: uniacid}).Error
}
