package service

import "finder/model"

func (s *Service) GetRecord(id int64) (result *model.Record, err error) {
	return s.dao.FindOneById(id)
}
