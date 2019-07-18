package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)


func GetPeopleList(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func GetPersonDetail(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}