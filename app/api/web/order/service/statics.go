package service

import (
	"errors"
	"mall/app/api/web/order/model"
)

func (s *Service) OrderStatics(p model.OrderStaticsParam) (*model.OrderStaticsResponse, error) {
	if p.Uniacid == 0 {
		return nil, errors.New("无效的应用ID")
	}

	count, err := s.d.OrderCount(p)
	if err != nil {
		return nil, errors.New("订单数量查询失败," + err.Error())
	}

	money, err := s.d.OrderSum(p)
	if err != nil {
		return nil, errors.New("订单总额查询失败," + err.Error())
	}

	return &model.OrderStaticsResponse{
		Count: count,
		Money: money,
	}, nil
}
