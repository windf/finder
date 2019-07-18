package service

import (
	"finder/config"
	"finder/dao"
	"github.com/labstack/echo"
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
	Echo   *echo.Echo
	Logger echo.Logger
}

func New(c *config.Config) (s *Service) {

	var err error
	s = new(Service)
	if s.dao, err = dao.NewDao(c); err != nil {
		panic(err)
	}

	s.conf = c
	s.Echo = echo.New()
	s.Logger = s.Echo.Logger

	//log
	s.initLog()
	return
}

func (s *Service) initLog() {
	s.Logger.SetPrefix(s.conf.Prefix)
	s.Logger.SetLevel(logLevels[s.conf.LogLevel])
}

func (s *Service) Close() {
	s.dao.Close()
}
