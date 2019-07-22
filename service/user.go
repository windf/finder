package service

import (
	"finder/model"
)

func (s *Service) GetUserList(page, pageSize int) (result []*model.User, err error) {
	//todo get cache

	result, err = s.dao.GetUserList(page, pageSize)
	if err != nil {
		s.Logger.Errorf("GetUserList err:%s ", err.Error())
		return
	}
	return
}

func (s *Service) GetUser(id int64) (result *model.User, err error) {
	result, err = s.dao.GetUserInfoById(id)
	if err != nil {
		s.Logger.Errorf("GetUserInfoById err:%s ", err.Error())
		return
	}

	return
}
