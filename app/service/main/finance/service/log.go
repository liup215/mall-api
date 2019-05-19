package service

import (
	"mall/app/service/main/finance/model"
)

func (s *Service) QueryLogs(q model.LogQuery, p model.Page) ([]model.EweiShopMemberLog, int, error) {
	return s.d.QueryLog(q, p)
}
