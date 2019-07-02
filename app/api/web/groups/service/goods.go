package service

import (
	"mall/app/api/web/groups/model"
)

func (s *Service) QueryGoodsList(q model.GoodsQuery) ([]model.EweiShopGroupsGoods, int) {
	return s.d.QueryGoodsList(q)
}
