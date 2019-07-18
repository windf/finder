package controllers

import (
	"finder/config"
	"finder/service"
	"github.com/labstack/echo/middleware"
)

var findSrv *service.Service

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

	s.Echo.Logger.Fatal(s.echo.Start(c.HttpAddress))
}
