package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func RenderAdmin(c echo.Context) error {
	return c.Render(http.StatusOK, "admin", nil)
}
