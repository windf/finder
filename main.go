package main

import (
	"finder/config"
	"finder/controllers"
	"finder/dao"
	"flag"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"os"
)

var (
	flagSet = flag.NewFlagSet("finder", flag.ExitOnError)

	filePath = flagSet.String("config-path", "config.toml", "config file")
	logLevel = flagSet.String("log-level", "DEBUG", "log level")
)

func main() {
	flagSet.Parse(os.Args[1:])

	if err := config.Init(*filePath); err != nil {
		panic(err)
	}

	if err := Run(); err != nil {
		panic(err)
	}

	startServer()
}

func Run() (err error) {
	r := new(config.Finder)
	if r.Dao, err = dao.NewDao(config.Conf); err != nil {
		panic(err)
	}
	return
}

func startServer() {
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
