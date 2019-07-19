package service

import "finder/model"

func (s *Service) GetCommentList(recordId int64, page, pageSize int) (result []*model.Comment, err error) {
	//todo get cache

	result, err = s.dao.GetCommentList(recordId, page, pageSize)
	if err != nil {
		s.Logger.Errorf("GetCommentList err:%s ", err.Error())
		return
	}
	return
}
