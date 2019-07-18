package dao

import "github.com/jinzhu/gorm"
import "finder/model"

const (
	_table = "record"
)

func (d *Dao) FindOneById(id int64) (result *model.Record, err error) {
	result = new(model.Record)
	err = d.dbr.Table(_table).Where("id=?", id).Select("*").First(result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result = nil
			err = nil
		}
		return
	}
	return
}
