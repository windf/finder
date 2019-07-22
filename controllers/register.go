package controllers

import (
	"finder/model"
	"finder/util"
	"github.com/labstack/echo"
	"time"
)

func Register(c echo.Context) error {
	userName := c.FormValue("user_name")
	password := c.FormValue("password")
	nickName := c.FormValue("nickname")
	phone := c.FormValue("phone")
	email := c.FormValue("email")
	remark := c.FormValue("remark")

	if userName == "" {
		return JsonBadRequest(c, "请输入用户名")
	}

	if password == "" {
		return JsonBadRequest(c, "请输入密码")
	}

	if findSrv.CheckUser(userName) {
		return JsonBadRequest(c, "用户名已经存在")
	}

	user := &model.User{
		UserName:   userName,
		Password:   util.Md5(password),
		NickName:   nickName,
		Phone:      phone,
		Email:      email,
		Remark:     remark,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	if !findSrv.AddUser(user) {
		return JsonBadRequest(c, "添加用户失败，请稍后重试")
	}

	return JsonOk(c, struct{}{})
}
