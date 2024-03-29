package dao

import (
	"finder/model"
	"github.com/jinzhu/gorm"
)

const (
	_userTable = "user"
)

func (d *Dao) GetUserList(page, pageSize int) (result []*model.User, err error) {
	offset := (page - 1) * pageSize
	err = d.dbr.Table(_userTable).Offset(offset).Limit(pageSize).Select("*").Order("id DESC").Find(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result = nil
			err = nil
		}
		return
	}
	return
}

func (d *Dao) GetUserCount() (result int64, err error) {
	err = d.dbr.Table(_userTable).Count(&result).Error
	return
}

func (d *Dao) GetUserInfoById(id int64) (result *model.User, err error) {
	result = new(model.User)
	err = d.dbr.Table(_userTable).Where("id=?", id).Select("*").Take(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result = nil
			err = nil
		}
		return
	}
	return
}

func (d *Dao) GetUserInfoByName(userName string) (result *model.User, err error) {
	result = new(model.User)
	err = d.dbr.Table(_userTable).Where("username=?", userName).Select("*").Take(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result = nil
			err = nil
		}
		return
	}
	return
}

func (d *Dao) GetUserInfoByEmail(email string) (result *model.User, err error) {
	result = new(model.User)
	err = d.dbr.Table(_userTable).Where("email=?", email).Select("*").Take(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result = nil
			err = nil
		}
		return
	}
	return
}

func (d *Dao) GetUserInfoByNameAndPassword(userName, password string) (result *model.User, err error) {
	result = new(model.User)
	err = d.dbr.Table(_userTable).Where("username=? AND password=?", userName, password).Select("*").Take(&result).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			result = nil
			err = nil
		}
		return
	}
	return
}

func (d *Dao) AddUser(user *model.User) bool {
	err := d.dbw.Table(_userTable).Create(user).Error
	if err != nil {
		return false
	}
	return true
}

func (d *Dao) UpdateUser(userId int64, user *model.User) bool {
	err := d.dbw.Table(_userTable).Where("id=?", userId).Updates(user).Error
	if err != nil {
		return false
	}
	return true
}

func (d *Dao) DeleteUser(user *model.User) bool {
	err := d.dbw.Table(_userTable).Delete(user).Error
	if err != nil {
		return false
	}
	return true
}
