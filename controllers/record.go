package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetList(c echo.Context) error {

	return c.JSON(http.StatusOK, nil)
}

func GetDetail(c echo.Context) (err error) {
	id := c.Param("id")
	recordId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	if recordId <= 0 {
		return JsonBadRequest(c, "id不正确")
	}

	result, err := findSrv.GetRecord(recordId)
	if err != nil {
		findSrv.Logger.Errorf("get record err:%s", err.Error())
	}

	var res interface{}
	res = struct{}{}

	if result != nil {
		res = result
	}

	return JsonOk(c, res)
}
