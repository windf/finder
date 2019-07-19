package dao

import "github.com/jinzhu/gorm"
import "finder/model"

const (
	_table = "record"
)

func (d *Dao) GetRecordById(id int64) (result *model.Record, err error) {
	result = new(model.Record)
	err = d.dbr.Table(_table).Where("id=?", id).Select("*").Take(result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result = nil
			err = nil
		}
		return
	}
	return
}

func (d *Dao) GetRecordList(page, pageSize int) (result []*model.Record, err error) {
	offset := (page - 1) * pageSize
	err = d.dbr.Table(_table).Offset(offset).Limit(pageSize).Select("*").Find(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result = nil
			err = nil
		}
		return
	}
	return
}
