package service

import "finder/model"

func (s *Service) GetRecord(id int64) (result *model.Record, err error) {
	result, err = s.dao.GetRecordCache(id)
	if err != nil {
		return
	}

	return s.dao.GetRecordById(id)
}
