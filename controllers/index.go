package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func RenderIndex(c echo.Context) error {
	return c.Render(http.StatusOK, "index", nil)
}
