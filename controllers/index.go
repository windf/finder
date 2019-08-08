package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func RenderIndex(c echo.Context) error {
	res := map[string]interface{}{
		"title": "首页",
		"data":  nil,
		//"totalPage": totalPage,
		//"pageSize":  pageSize,
		//"page":      page,
		//"isFind":    isFind,
		//"status":    status,
	}
	/*
		if result != nil {
			res["data"] = result
		}*/

	return c.Render(http.StatusOK, "index", res)
}
