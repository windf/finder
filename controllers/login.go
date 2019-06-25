package controllers

import "github.com/astaxie/beego"

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Index() {
	this.Ctx.WriteString("admin user front")
}

func (this *LoginController) Logout() {
	this.Ctx.WriteString("admin user front")
}