package service

import (
	apiUserModel "mall/app/api/admin/user/model"
	serviceUserModel "mall/app/service/main/user/model"
)

func (s *Service) UserList(query apiUserModel.UserListQuery) ([]apiUserModel.UserView, int, error) {
	ulist, total, err := s.userService.UserList(serviceUserModel.UserQuery{
		Page:     query.Page,
		PageSize: query.PageSize,
		Username: query.Username,
	})

	list := []apiUserModel.UserView{}

	for _, u := range ulist {
		list = append(list, userToUserView(u))
	}

	return list, total, err

}

func userToUserView(u serviceUserModel.Users) apiUserModel.UserView {
	return apiUserModel.UserView{
		Uid:       u.Uid,
		Uniacid:   u.Uniacid,
		Username:  u.Username,
		Status:    u.Status,
		Joindate:  u.Joindate,
		Lastvisit: u.Lastvisit,
	}
}

func (s *Service) ChangeStatus(uid int) (int, error) {
	return s.userService.ChangeStatus(uid)
}

func (s *Service) QueryUser(query apiUserModel.UserQuery) (*apiUserModel.UserView, error) {
	user, err := s.userService.QueryUser(serviceUserModel.UserQuery{
		Uid:      query.Uid,
		Username: query.Username,
	})
	if err != nil {
		return nil, err
	}

	view := userToUserView(*user)

	return &view, nil
}

func (s *Service) UpdateUser(u apiUserModel.UserCreateOrUpdateParam) (*apiUserModel.UserView, error) {
	user, err := s.userService.Update(serviceUserModel.Users{
		Uid:      u.Uid,
		Username: u.Username,
		Password: u.Password,
		Status:   u.Status,
	})

	if err != nil {
		return nil, err
	}

	newUser := userToUserView(*user)

	return &newUser, nil
}
