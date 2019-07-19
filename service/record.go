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

func (s *Service) GetRecordList(page, pageSize int) (result []*model.Record, err error) {
	result, err = s.dao.GetRecordList(page, pageSize)
	if err != nil {
		s.Logger.Errorf("GetRecordList err:%s ", err.Error())
		return
	}
	return
}
