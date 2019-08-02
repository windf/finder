package service

import "finder/model"

func (s *Service) GetRecord(id int64) (result *model.Record, err error) {
	result, err = s.dao.GetRecordCache(id)
	if err == nil {
		return
	}

	result, err = s.dao.GetRecordById(id)
	if err != nil {
		s.Logger.Errorf("GetRecordById err:%s ", err.Error())
		return
	}

	return
}

func (s *Service) GetRecordByName(name string) (result []*model.Record, err error) {
	result, err = s.dao.GetRecordByName(name)
	if err != nil {
		s.Logger.Errorf("GetRecordByName err:%s ", err.Error())
		return
	}

	return
}

func (s *Service) GetRecordList(page, pageSize, isFind, status int) (result []*model.Record, err error) {
	result, err = s.dao.GetRecordList(page, pageSize, isFind, status)
	if err != nil {
		s.Logger.Errorf("GetRecordList err:%s ", err.Error())
		return
	}
	return
}

func (s *Service) GetRecordCount(isFind, status int) (result int64, err error) {
	result, err = s.dao.GetRecordCount(isFind, status)
	if err != nil {
		s.Logger.Errorf("GetRecordCount err:%s ", err.Error())
		return
	}
	return
}

func (s *Service) GetUserRecordCount(userId int64, isFind, status int) (result int64, err error) {
	result, err = s.dao.GetUserRecordCount(userId, isFind, status)
	if err != nil {
		s.Logger.Errorf("GetUserRecordCount err:%s ", err.Error())
		return
	}
	return
}

func (s *Service) GetUserRecordList(userId int64, page, pageSize, isFind, status int) (result []*model.Record, err error) {
	result, err = s.dao.GetUserRecordList(userId, page, pageSize, isFind, status)
	if err != nil {
		s.Logger.Errorf("GetUserRecordList err:%s ", err.Error())
		return
	}
	return
}

func (s *Service) AddRecord(record *model.Record) bool {
	if s.dao.AddRecord(record) {
		return true
	}
	return false
}

func (s *Service) UpdateRecord(userId int64, record *model.Record) bool {
	if s.dao.UpdateRecord(userId, record) {
		return true
	}
	return false
}

func (s *Service) DeleteRecord(record *model.Record) bool {
	if s.dao.DeleteRecord(record) {
		return true
	}
	return false
}
