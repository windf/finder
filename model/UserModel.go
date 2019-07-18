package model

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//用户表
type User struct {
	Id         int32
	Username   string
	Password   string
	Nickname   string
	Phone      string
	Email      string
	Remark     string
	Status     byte
	Role       byte
	UpdateTime time.Time
	CreateTime time.Time
}

func init() {
	orm.RegisterModel(new(User))
}
func GetUserList(page int64, pageSize int64, sort string) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * pageSize
	}
	qs.Limit(pageSize, offset).OrderBy(sort).Values(&users)
	count, _ = qs.Count()
	return users, count
}
