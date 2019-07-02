package dao

import (
	"errors"
	"mall/app/api/web/member/model"
	"mall/lib/strings"
	"mall/lib/time"
)

func (d *Dao) InsertLog(log model.EweiShopMemberLog) (int, error) {
	if log.Uniacid == 0 {
		return 0, errors.New("无效的应用ID")
	}
	if log.Openid == "" {
		return 0, errors.New("无效的Openid")
	}
	log.Createtime = time.Now()
	logno := ""
	for {
		logno = strings.CreateNo("RC")

		count := d.CountLog(&model.EweiShopMemberLogQuery{
			Uniacid: log.Uniacid,
			Logno:   logno,
		})

		if count == 0 {
			break
		}
	}
	log.Logno = strings.CreateNo("RC")
	err := d.orm.Create(&log).Error

	return log.Id, err

}

func (d *Dao) DeleteLog(q *model.EweiShopMemberLogQuery) error {
	return nil
}

func (d *Dao) CountLog(q *model.EweiShopMemberLogQuery) int {
	count := 0

	d.orm.Model(&model.EweiShopMemberLog{}).Where(parseLogQuery(q)).Count(&count)
	return count
}

func parseLogQuery(q *model.EweiShopMemberLogQuery) map[string]interface{} {
	m := map[string]interface{}{}

	if q.Uniacid != 0 {
		m["uniacid"] = q.Uniacid
	}

	if q.Id != 0 {
		m["id"] = q.Id
	}

	if q.Openid != "" {
		m["openid"] = q.Openid
	}

	m["type"] = q.Type
	m["status"] = q.Status
	return m
}
