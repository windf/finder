package service

import (
	"finder/config"
	"finder/controllers"
	"finder/dao"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Finder struct {
	dao *dao.Dao
}

func Run(httpAddress string) (err error) {
	r := new(Finder)
	if r.dao, err = dao.NewDao(config.Conf); err != nil {
		panic(err)
	}

	startServer(httpAddress)

	return
}

func startServer(httpAddress string) {
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

	e.Logger.Fatal(e.Start(httpAddress))
}
