package service

import (
	"errors"
	"mall/app/api/web/groups/model"
	"mall/lib/net/http/middleware/auth"

	"github.com/jinzhu/gorm"
)

func (s *Service) QueryActVisitorList(actId string) ([]model.EweiShopGroupsVisitor, int) {
	return s.d.QueryVisitorList(model.VisitorQuery{
		Aid: actId,
	})
}

func (s *Service) CheckVisitor(u *auth.CurrentUser, actId string) error {
	if u.Usertype == 0 {
		return nil
	}
	q := model.VisitorQuery{
		Uniacid: u.Uniacid,
		Mid:     u.Id,
		Openid:  u.Openid,
		Aid:     actId,
	}

	_, err := s.d.QueryVisitorOne(q)

	if err == gorm.ErrRecordNotFound {
		v := model.EweiShopGroupsVisitor{
			Uniacid:  u.Uniacid,
			Mid:      u.Id,
			Aid:      actId,
			Openid:   u.Openid,
			Nickname: u.Nickname,
			Avatar:   u.Avatar,
		}

		_, err = s.d.CreateVisitor(v)
	}

	return err

}

func (s *Service) ShareVisitorInfo(q model.VisitorQuery) (*model.EweiShopGroupsVisitor, error) {
	if q.Uniacid == 0 {
		return nil, errors.New("无效的应用ID")
	}

	if q.Aid == "" {
		return nil, errors.New("无效的活动ID")
	}

	if q.Openid == "" {
		return nil, errors.New("无效的用户Openid")
	}

	// 活动的分享数增加1
	s.d.IncrActivityShares(q.Aid)
	// 获取分享者信息
	v, err := s.d.QueryVisitorOne(q)
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("用户未参与此活动")
	}

	return v, err
}
