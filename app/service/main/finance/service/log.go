package service

import (
	"errors"
	"mall/app/service/main/finance/model"
)

func (s *Service) LogList(q model.LogQuery, p model.Page) ([]model.EweiShopMemberLog, int, error) {
	return s.d.LogList(q, p)
}

func (s *Service) LogDetail(q model.LogQuery) (model.EweiShopMemberLog, error) {
	var log model.EweiShopMemberLog
	if q.Uniacid == 0 {
		return log, errors.New("无效的应用ID")
	}

	if q.Id == 0 && q.Logno == "" {
		return log, errors.New("id和logno都为空，无法确定唯一记录")
	}

	if q.Id != 0 {
		return s.d.LogById(q.Id, q.Uniacid)
	}

	if q.Logno != "" {
		return s.d.LogByLogno(q.Logno, q.Uniacid)
	}

	return log, errors.New("id和logno不能同时为空")
}
