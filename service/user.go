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

func (s *Service) GetUserCount() (result int64, err error) {
	result, err = s.dao.GetUserCount()
	if err != nil {
		s.Logger.Errorf("GetUserCount err:%s ", err.Error())
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

func (s *Service) CheckUser(userName string) bool {
	result, err := s.dao.GetUserInfoByName(userName)
	if err != nil {
		return false
	}

	if result != nil {
		return true
	}

	return false
}

func (s *Service) AddUser(user *model.User) bool {
	if s.dao.AddUser(user) {
		return true
	}
	return false
}

func (s *Service) UpdateUser(userId int64, user *model.User) bool {
	if s.dao.UpdateUser(userId, user) {
		return true
	}
	return false
}

func (s *Service) DeleteUser(user *model.User) bool {
	if s.dao.DeleteUser(user) {
		return true
	}
	return false
}
