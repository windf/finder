package models

import (
	"fmt"
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
func GetUserList() {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	fmt.Println(qs)
}
