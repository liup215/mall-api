package dao

import (
	"errors"
	"mall/app/api/web/groups/model"

	"github.com/jinzhu/gorm"
)

func (d *Dao) PostCategory(cat model.EweiShopGroupsCategory) (*model.EweiShopGroupsCategory, error) {
	var err error
	if cat.Id != 0 {
		err = d.orm.Save(&cat).Error
	} else {
		err = d.orm.Create(&cat).Error
	}

	return &cat, err
}

func (d *Dao) DeleteCategory(id int) error {
	if id == 0 {
		return errors.New("分类ID不能为空")
	}
	err := d.orm.Delete(&model.EweiShopGroupsCategory{Id: id}).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("抱歉，分类不存在或是已经被删除！")
		} else {
			return errors.New("删除错误," + err.Error())
		}
	}

	return nil
}

func (d *Dao) DisplayorderCategory(id, order int) error {
	if id == 0 {
		return errors.New("分类ID不能为空")
	}

	return d.orm.Model(&model.EweiShopGroupsCategory{Id: id}).Update("displayorder", order).Error
}

func (d *Dao) EnableCategory(ids []int) error {
	var cats []model.EweiShopGroupsCategory
	err := d.orm.Where("id in (?)", ids).Find(&cats).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errors.New("记录不存在")
		} else {
			return err
		}
	}

	for _, cat := range cats {
		if cat.Id == 0 {
			break
		}
		enabled := 0
		if cat.Enabled == 0 {
			enabled = 1
		}
		d.orm.Where(&cat).Update("enabled", enabled)
	}
	return nil

}
