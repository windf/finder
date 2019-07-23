package dao

import (
	"finder/model"
	"github.com/jinzhu/gorm"
)

const (
	_commentTable = "comment"
)

func (d *Dao) GetCommentList(recordId int64, page, pageSize int) (result []*model.Comment, err error) {
	search := d.dbr.Table(_commentTable)
	if recordId > 0 {
		search = search.Where("record_id=?", recordId)
	}

	offset := (page - 1) * pageSize
	err = search.Offset(offset).Limit(pageSize).Select("*").Find(&result).Order("id DESC").Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result = nil
			err = nil
		}
		return
	}
	return
}

func (d *Dao) GetCommentById(id int64) (result *model.Comment, err error) {
	result = new(model.Comment)
	err = d.dbr.Table(_commentTable).Where("id=?", id).Select("*").Take(result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result = nil
			err = nil
		}
		return
	}
	return
}

func (d *Dao) AddComment(comment *model.Comment) bool {
	err := d.dbw.Table(_commentTable).Create(comment).Error
	if err != nil {
		return false
	}
	return true
}

func (d *Dao) DeleteComment(comment *model.Comment) bool {
	err := d.dbw.Table(_commentTable).Delete(comment).Error
	if err != nil {
		return false
	}
	return true
}
