package service

import (
	"errors"
	"fmt"
	"mall/app/api/web/groups/model"
)

func (s *Service) CreateActivity(act model.EweiShopGroupsActivity) (*model.EweiShopGroupsActivity, error) {
	if act.Id != "" {
		act.Id = ""
	}

	if act.Uniacid == 0 {
		return nil, errors.New("无效的应用ID")
	}

	if act.Uid == 0 {
		return nil, errors.New("无法确定活动主办方")
	}

	fmt.Println(act.Starttime)
	fmt.Println(act.Endtime)

	return s.d.InsertActivity(act)
}

func (s *Service) UpdateActivity(uid int, act model.EweiShopGroupsActivity) (*model.EweiShopGroupsActivity, error) {
	if act.Id == "" {
		return nil, errors.New("ID不能为空")
	}

	if act.Uniacid == 0 {
		return nil, errors.New("无效的应用ID")
	}

	if act.Starttime == 0 {
		return nil, errors.New("活动的起始时间不能为空")
	}

	if act.Endtime == 0 {
		return nil, errors.New("活动的结束时间不能为空")
	}

	record, err := s.d.QueryActivityById(act.Id)
	if err != nil {
		return nil, err
	}

	if record.Uid != uid {
		return nil, errors.New("您没有操作权限")
	}
	if act.Uid == 0 {
		act.Uid = record.Uid
	}

	return s.d.UpdateActivity(act)
}

func (s *Service) QueryActivityList(q model.ActivityQuery) ([]model.EweiShopGroupsActivity, int) {
	return s.d.QueryActivityList(q)
}

func (s *Service) QueryActivityById(aid string) (*model.EweiShopGroupsActivity, error) {
	return s.d.QueryActivityById(aid)
}

func (s *Service) IncrActivityVisits(aid string) error {
	return s.d.IncrActivityVisits(aid)
}
