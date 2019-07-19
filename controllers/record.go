package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func GetList(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func GetDetail(c echo.Context) (err error) {
	req := new(GetDetailReq)
	if err = c.Bind(req); err != nil {
		return
	}

	if req.ID <= 0 {
		return JsonFail(c, http.StatusOK, "id不正确")
	}

	result, err := findSrv.GetRecord(1)
	if err != nil {
		findSrv.Logger.Errorf("get record err:%s", err.Error())
	}

	var res interface{}
	res = struct{}{}

	if result != nil {
		res = result
	}

	return JsonSuccess(c, http.StatusOK, res)
}
