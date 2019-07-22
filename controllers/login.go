package controllers

import (
	"finder/model"
	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	userName := c.FormValue("user_name")
	password := c.FormValue("password")

	if userName == "" || password == "" {
		return JsonBadRequest(c, "参数错误")
	}

	result, err := findSrv.Login(userName, password)
	if err != nil || result == nil {
		return JsonBadRequest(c, "用户不存在或密码错误")
	}

	if result.Status == model.StatusERR {
		return JsonBadRequest(c, "该用户已被封禁，请尝试联系管理员")
	}

	err = SetSessionId(c, result.ID)
	if err != nil {
		return JsonBadRequest(c, "登录失败，请稍后重试")
	}
	return RenderAdmin(c)
}

func Logout(c echo.Context) error {
	err := DelSessionId(c)
	if err != nil {
		return JsonBadRequest(c, "退出失败，请稍后重试")
	}

	return RenderLogin(c)
}
