package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func GetRecordList(c echo.Context) error {
	reqPage := c.QueryParam("page")
	reqPageSize := c.QueryParam("pageSize")

	page, err := strconv.Atoi(reqPage)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	pageSize, err := strconv.Atoi(reqPageSize)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	if page < 1 {
		page = 1
	}

	if pageSize <= 0 || pageSize > 200 {
		pageSize = 10
	}

	result, err := findSrv.GetRecordList(page, pageSize)
	if err != nil {
		findSrv.Logger.Errorf("get record err:%s", err.Error())
	}

	var res interface{}
	res = struct{}{}

	if result != nil {
		res = result
	}

	return c.JSON(http.StatusOK, res)
}

func GetRecordDetail(c echo.Context) (err error) {
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
