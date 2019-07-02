package service

import (
	"errors"
	"mall/app/api/web/groups/model"
)

func (s *Service) CreateCategory(cat model.EweiShopGroupsCategory) (*model.EweiShopGroupsCategory, error) {
	if cat.Id != 0 {
		cat.Id = 0
	}

	if cat.Uniacid == 0 {
		return nil, errors.New("无效的应用ID")
	}

	return s.d.PostCategory(cat)
}

func (s *Service) UpdateCategory(cat model.EweiShopGroupsCategory) (*model.EweiShopGroupsCategory, error) {
	if cat.Id == 0 {
		return nil, errors.New("分类ID不能为空")
	}

	if cat.Uniacid == 0 {
		return nil, errors.New("无效的应用ID")
	}

	return s.d.PostCategory(cat)
}

func (s *Service) DeleteCategory(id int) error {
	if id == 0 {
		return errors.New("分类ID不能为空")
	}

	return s.d.DeleteCategory(id)
}

func (s *Service) DisplayorderCategory(id, order int) error {
	if id == 0 {
		return errors.New("分类ID不能为空")
	}
	return s.d.DisplayorderCategory(id, order)
}

func (s *Service) EnableCategory(ids []int) error {
	return s.d.EnableCategory(ids)
}
