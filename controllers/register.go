package controllers

import (
	"finder/model"
	"finder/util"
	"github.com/labstack/echo"
	"net/http"
	"time"
)

func RenderRegister(c echo.Context) error {
	return c.Render(http.StatusOK, "register", nil)
}

func Register(c echo.Context) error {
	userName := c.FormValue("user_name")
	password := c.FormValue("password")
	nickName := c.FormValue("nickname")
	phone := c.FormValue("phone")
	email := c.FormValue("email")
	rePassword := c.FormValue("re_password")

	if userName == "" {
		return JsonBadRequest(c, "请输入用户名")
	}

	if password == "" {
		return JsonBadRequest(c, "请输入密码")
	}

	if rePassword != password {
		return JsonBadRequest(c, "两次输入的密码不一致")
	}

	if findSrv.CheckUser(userName) {
		return JsonBadRequest(c, "用户名已经存在")
	}

	if findSrv.CheckEmail(email) {
		return JsonBadRequest(c, "邮箱已经存在")
	}

	user := &model.User{
		UserName:   userName,
		Password:   util.Md5(password),
		NickName:   nickName,
		Phone:      phone,
		Email:      email,
		Role:       model.USER,
		Status:     model.StatusOK,
		CreateTime: time.Now().Unix(),
		UpdateTime: time.Now().Unix(),
	}

	if !findSrv.AddUser(user) {
		return JsonBadRequest(c, "添加用户失败，请稍后重试")
	}

	return JsonOk(c, struct{}{})
}
