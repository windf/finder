package controllers

import (
	"finder/util"
	"github.com/labstack/echo"
)

func GetCaptcha(c echo.Context) error {

	id, value := util.GenerateCaptcha()

	res := make(map[string]string)

	res["id"] = id
	res["image"] = value

	return JsonOk(c, res)
}
