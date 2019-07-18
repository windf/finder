package controllers

import (
	"finder/model"
	"github.com/labstack/echo"
	"net/http"
)

func GetList(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func GetDetail(c echo.Context) error {
	result, err := findSrv.GetRecord(1)
	if err != nil {
		findSrv.Logger.Errorf("get record err:%s", err.Error())
	}

	if result == nil {
		result = new(model.Record)
	}
	return c.JSON(http.StatusOK, result)
}
