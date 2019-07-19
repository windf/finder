package controllers

import (
	"github.com/labstack/echo"
	"strconv"
)

func GetCommentList(c echo.Context) error {
	id := c.Param("id")
	reqPage := c.QueryParam("page")
	reqPageSize := c.QueryParam("pageSize")

	recordId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return JsonBadRequest(c, "参数错误")
	}

	if recordId <= 0 {
		return JsonBadRequest(c, "id不正确")
	}

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

	result, err := findSrv.GetCommentList(recordId, page, pageSize)
	if err != nil {
		findSrv.Logger.Errorf("GetCommentList err:%s", err.Error())
	}

	var res interface{}
	res = struct{}{}

	if result != nil {
		res = result
	}

	return JsonOk(c, res)
}
