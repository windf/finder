package dao

import (
	"finder/model"
	"github.com/jinzhu/gorm"
)

const (
	_recordTable = "record"
)

func (d *Dao) GetRecordById(id int64) (result *model.Record, err error) {
	result = new(model.Record)
	err = d.dbr.Table(_recordTable).Where("id=?", id).Select("*").Take(result).Error
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
	err = d.dbr.Table(_recordTable).Offset(offset).Limit(pageSize).Select("*").Find(&result).Order("id DESC").Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result = nil
			err = nil
		}
		return
	}
	return
}

func (d *Dao) GetRecordByName(name string) (result []*model.Record, err error) {
	err = d.dbr.Table(_recordTable).Where("name=?", name).Select("*").Find(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result = nil
			err = nil
		}
		return
	}
	return
}

func (d *Dao) AddRecord(record *model.Record) bool {
	err := d.dbw.Table(_recordTable).Create(record).Error
	if err != nil {
		return false
	}
	return true
}

func (d *Dao) UpdateRecord(recordId int64, record *model.Record) bool {
	err := d.dbw.Table(_recordTable).Where("id=?", recordId).Updates(record).Error
	if err != nil {
		return false
	}
	return true
}

func (d *Dao) DeleteRecord(record *model.Record) bool {
	err := d.dbw.Table(_recordTable).Delete(record).Error
	if err != nil {
		return false
	}
	return true
}
