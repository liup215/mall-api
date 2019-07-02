package dao

import (
	"errors"
	"mall/app/api/web/system/model"
)

func (d *Dao) settingSave(key, value string) error {
	if key == "" {
		return errors.New("key不能为空")
	}

	_, err := d.settingLoad(key)

	if err == nil {
		return d.orm.Where("key = ?", key).Update("value", value).Error
	}
	return d.orm.Create(&model.CoreSettings{Key: key, Value: value}).Error

}

func (d *Dao) settingLoad(key string) (string, error) {
	if key == "" {
		return "", errors.New("key 不能为空")
	}

	var set model.CoreSettings
	err := d.orm.Where(&model.CoreSettings{Key: key}).First(&set).Error
	return set.Value, err

}
