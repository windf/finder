package service

import (
	"finder/config"
	"finder/controllers"
	"finder/dao"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

var logLevels = map[string]log.Lvl{
	"DEBUG": log.DEBUG,
	"INFO":  log.INFO,
	"WARN":  log.WARN,
	"ERROR": log.ERROR,
}

type Finder struct {
	dao *dao.Dao
}
type Service struct {
	dao    *dao.Dao
	conf   *config.Config
	echo   *echo.Echo
	logger echo.Logger
}

func New(c *config.Config) (s *Service) {

	var err error
	s = new(Service)
	if s.dao, err = dao.NewDao(c); err != nil {
		panic(err)
	}

	s.conf = c
	s.echo = echo.New()
	s.logger = s.echo.Logger

	//log
	s.initLog()
	//http
	s.startServer()
	return
}

func (s *Service) initLog() {
	s.logger.SetPrefix(s.conf.Prefix)
	s.logger.SetLevel(logLevels[s.conf.LogLevel])
}

func (s *Service) startServer() {

	//Middleware
	s.echo.Use(middleware.Recover())
	s.echo.Use(middleware.Logger())

	// Routes
	/*
		e.POST("/users", controllers.CreateUser)
		e.GET("/users/:id", controllers.GetUser)
		e.PUT("/users/:id", controllers.UpdateUser)
		e.DELETE("/users/:id", controllers.DeleteUser)
	*/
	s.echo.GET("/person/:id", controllers.GetPersonDetail)
	s.echo.GET("/people/*", controllers.GetPeopleList)

	s.echo.Logger.Fatal(s.echo.Start(s.conf.HttpAddress))
}

func (s *Service) Close() {
	s.dao.Close()
}
