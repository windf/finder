package dao

import (
	"finder/model"
	"github.com/jinzhu/gorm"
)

const (
	_commentTable = "comment"
)

func (d *Dao) GetCommentList(recordId int64, page, pageSize int) (result []*model.Comment, err error) {
	offset := (page - 1) * pageSize
	err = d.dbr.Table(_commentTable).Where("record_id=?", recordId).Offset(offset).Limit(pageSize).Select("*").Find(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result = nil
			err = nil
		}
		return
	}
	return
}
