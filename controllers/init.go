package controllers

import (
	"finder/config"
	"finder/service"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var findSrv *service.Service

type Message struct {
	Errcode int         `json:"errcode" xml:"errcode"`
	Errmsg  string      `json:"errmsg,omitempty" xml:"errmsg"`
	Data    interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func Init(c *config.Config, s *service.Service) {
	findSrv = s

	//Middleware
	s.Echo.Use(middleware.Recover())
	s.Echo.Use(middleware.Logger())

	// Routes
	/*
		e.POST("/users", controllers.CreateUser)
		e.GET("/users/:id", controllers.GetUser)
		e.PUT("/users/:id", controllers.UpdateUser)
		e.DELETE("/users/:id", controllers.DeleteUser)
	*/
	s.Echo.GET("/record/:id", GetDetail)
	s.Echo.GET("/recordList/*", GetList)

	s.Echo.Logger.Fatal(s.Echo.Start(c.HttpAddress))
}

func JSON(c echo.Context, code int, msg string, i interface{}) error {
	return c.JSON(code, &Message{
		Errcode: code,
		Errmsg:  msg,
		Data:    i,
	})
}
