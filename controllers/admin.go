package controllers

import (
	"github.com/labstack/echo"
	"net/http"
)

func RenderAdmin(c echo.Context) error {
	result := map[string]interface{}{
		"name": GetSessionName(c),
		"role": GetUserRole(c),
	}

	return c.Render(http.StatusOK, "admin", result)
}
