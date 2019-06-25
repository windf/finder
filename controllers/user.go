package controllers
import (
	"github.com/astaxie/beego"
)
//用户管理

type UserController struct {
	beego.Controller
}

func (this *UserController) Index() {
	//todo 搜索
	this.Ctx.WriteString("admin user front")
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