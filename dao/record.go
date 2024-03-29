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

func (d *Dao) GetRecordList(page, pageSize, isFind, status int) (result []*model.Record, err error) {
	offset := (page - 1) * pageSize
	search := d.dbr.Table(_recordTable)
	if isFind > model.AllFind && status > model.AllReview {
		search = search.Where("isfind=? AND status=?", isFind, status)
	} else if isFind == model.AllFind && status > model.AllReview {
		search = search.Where("status=?", status)
	} else if isFind > model.AllFind && status == model.AllReview {
		search = search.Where("isfind=?", isFind)
	}

	err = search.Select("*").Offset(offset).Limit(pageSize).Order("id DESC").Find(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result = nil
			err = nil
		}
		return
	}
	return
}

func (d *Dao) GetUserRecordList(userId int64, page, pageSize, isFind, status int) (result []*model.Record, err error) {
	offset := (page - 1) * pageSize
	search := d.dbr.Table(_recordTable)
	if isFind > model.AllFind && status > model.AllReview {
		search = search.Where("isfind=? AND status=? AND publisher_id=?", isFind, status, userId)
	} else if isFind == model.AllFind && status > model.AllReview {
		search = search.Where("status=? AND publisher_id=?", status, userId)
	} else if isFind > model.AllFind && status == model.AllReview {
		search = search.Where("isfind=? AND publisher_id=?", isFind, userId)
	} else {
		search = search.Where("publisher_id=?", userId)
	}

	err = search.Select("*").Offset(offset).Limit(pageSize).Order("id DESC").Find(&result).Error
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
	err = d.dbr.Table(_recordTable).Where("status=? AND name=?", model.ReviewSuccess, name).Select("*").Find(&result).Error
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

func (d *Dao) GetRecordCount(isFind, status int) (result int64, err error) {
	search := d.dbr.Table(_recordTable)
	if isFind > model.AllFind && status > model.AllReview {
		search = search.Where("isfind=? AND status=?", isFind, status)
	} else if isFind == model.AllFind && status > model.AllReview {
		search = search.Where("status=?", status)
	} else if isFind > model.AllFind && status == model.AllReview {
		search = search.Where("isfind=?", isFind)
	}

	err = search.Count(&result).Error
	return
}

func (d *Dao) GetUserRecordCount(userId int64, isFind, status int) (result int64, err error) {
	search := d.dbr.Table(_recordTable)
	if isFind > model.AllFind && status > model.AllReview {
		search = search.Where("isfind=? AND status=? AND publisher_id=?", isFind, status, userId)
	} else if isFind == model.AllFind && status > model.AllReview {
		search = search.Where("status=? AND publisher_id=?", status, userId)
	} else if isFind > model.AllFind && status == model.AllReview {
		search = search.Where("isfind=? AND publisher_id=?", isFind, userId)
	} else {
		search = search.Where("publisher_id=?", userId)
	}

	err = search.Count(&result).Error
	return
}
