package controllers

import (
	"finder/model"
	"finder/util"
	"github.com/labstack/echo"
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

	return JsonOk(c, res)
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

func SearchRecordDetail(c echo.Context) (err error) {
	keyword := c.QueryParam("keyword")

	id, ok := util.IsNumeric(keyword)

	var res interface{}
	res = struct{}{}

	if ok {
		if id <= 0 {
			return JsonBadRequest(c, "id不正确")
		}
		result, err := findSrv.GetRecord(id)
		if err != nil {
			findSrv.Logger.Errorf("get record err:%s", err.Error())
		}

		if result != nil {
			res = []*model.Record{result}
		}
	} else {
		if keyword == "" {
			return JsonBadRequest(c, "参数错误")
		}

		result, err := findSrv.GetRecordByName(keyword)
		if err != nil {
			findSrv.Logger.Errorf("GetRecordByName err:%s", err.Error())
		}

		if result != nil {
			res = result
		}
	}

	return JsonOk(c, res)
}
