package service

import "finder/model"

func (s *Service) GetCommentList(recordId int64, page, pageSize int) (result []*model.Comment, err error) {
	result, err = s.dao.GetCommentList(recordId, page, pageSize)
	if err != nil {
		s.Logger.Errorf("GetCommentList err:%s ", err.Error())
		return
	}
	return
}

func (s *Service) GetComment(id int64) (result *model.Comment, err error) {
	result, err = s.dao.GetCommentCache(id)
	if err == nil {
		return
	}

	result, err = s.dao.GetCommentById(id)
	if err != nil {
		s.Logger.Errorf("GetCommentById err:%s ", err.Error())
		return
	}

	return
}

func (s *Service) AddComment(comment *model.Comment) bool {
	if s.dao.AddComment(comment) {
		return true
	}
	return false
}

func (s *Service) DeleteComment(comment *model.Comment) bool {
	if s.dao.DeleteComment(comment) {
		return true
	}
	return false
}
