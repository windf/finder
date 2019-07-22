package service

import (
	"finder/model"
	"iceberg/frame/util"
)

func (s *Service) Login(userName, password string) (result *model.User, err error) {
	password = util.MD5([]byte(password))
	result, err = s.dao.GetUserInfoByNameAndPassword(userName, password)
	if err != nil {
		s.Logger.Errorf("GetUserInfoByNameAndPassword err:%s ", err.Error())
		return
	}

	return
}
