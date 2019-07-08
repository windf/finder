package controllers

import (
	"finder/models"
	"github.com/astaxie/beego"
)

//用户管理

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	//todo 搜索
	page, _ := this.GetInt64("page")
	pageSize, _ := this.GetInt64("rows")
	sort := this.GetString("sort")
	order := this.GetString("order")
	if len(order) > 0 {
		if order == "desc" {
			sort = "-" + sort
		}
	} else {
		sort = "Id"
	}
	this.Ctx.WriteString("user index")
	models.GetUserList(page, pageSize, sort)
	//todo 分页模板
}

func (this *UserController) Add() {
	this.Ctx.WriteString("admin user front")
}

func (this *UserController) Update() {
	this.Ctx.WriteString("admin user front")
}

func (this *UserController) Delete() {
	this.Ctx.WriteString("admin user front")
}

func (this *UserController) ApiUser() {
	//todo 随机数
	this.Data["json"] = &mystruct
	this.ServeJSON()
}

//todo 随机接口json
/*
SELECT *
         FROM user AS r1 JOIN
              (SELECT (RAND() *
                            (SELECT MAX(id)
                               FROM user)) AS id)
               AS r2
        WHERE r1.id >= r2.id
        ORDER BY r1.id ASC
        LIMIT 1;
*/
