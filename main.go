package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"testecho/controllers"
)

func main()  {
	e := echo.New()

	//Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())


	// Routes
	/*
	e.POST("/users", controllers.CreateUser)
	e.GET("/users/:id", controllers.GetUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
	*/
	e.GET("/person/:id", controllers.GetPersonDetail)
	e.GET("/people/*", controllers.GetPeopleList)


	e.Logger.Fatal(e.Start(":1323"))
}
