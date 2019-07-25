package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func RenderAdmin(c echo.Context) error {
	result := map[string]interface{}{
		"userId":   GetSessionId(c),
		"name":     GetSessionName(c),
		"role":     GetUserRole(c),
		"title":    "管理后台",
		"leftMenu": "",
		"menu":     "",
	}

	return c.Render(http.StatusOK, "admin", result)
}
