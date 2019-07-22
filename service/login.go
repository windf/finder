package service

import (
	"finder/model"
	"finder/util"
)

func (s *Service) Login(userName, password string) (result *model.User, err error) {
	password = util.Md5(password)
	result, err = s.dao.GetUserInfoByNameAndPassword(userName, password)
	if err != nil {
		s.Logger.Errorf("GetUserInfoByNameAndPassword err:%s ", err.Error())
		return
	}

	return
}
