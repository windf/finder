package controllers

import (
	"finder/model"
	"github.com/labstack/echo"
	"math"
	"net/http"
	"strconv"
)

func RenderIndex(c echo.Context) error {

	reqPage := c.QueryParam("page")
	reqPageSize := c.QueryParam("pageSize")

	page, err := strconv.Atoi(reqPage)
	if err != nil {
		page = 1
	}

	pageSize, err := strconv.Atoi(reqPageSize)
	if err != nil {
		pageSize = 8
	}

	if page < 1 {
		page = 1
	}

	if pageSize <= 0 || pageSize > 200 {
		pageSize = 8
	}

	count, err := findSrv.GetRecordCount(model.AllFind, model.ReviewSuccess)
	if err != nil {
		findSrv.Logger.Errorf("GetRecordCount  err:%s", err.Error())
		return JsonServerError(c)
	}

	totalPage := int(math.Ceil(float64(count) / float64(pageSize))) //page总数
	if page > totalPage {
		page = totalPage
	}

	result, err := findSrv.GetRecordList(page, pageSize, model.AllFind, model.ReviewSuccess)
	if err != nil {
		findSrv.Logger.Errorf("get record err:%s", err.Error())
		return JsonServerError(c)
	}

	res := map[string]interface{}{
		"title":     "首页",
		"data":      nil,
		"totalPage": totalPage,
		"pageSize":  pageSize,
		"page":      page,
	}

	if result != nil {
		res["data"] = result
	}

	return c.Render(http.StatusOK, "index", res)
}
