package service

import (
	"finder/config"
	"finder/dao"
	"github.com/gomodule/redigo/redis"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"github.com/mojocn/base64Captcha"
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

type customizeRdsStore struct {
	redisClient *redis.Pool
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

	customeStore := customizeRdsStore{s.dao.Cpool}
	base64Captcha.SetCustomStore(&customeStore)
	return
}

func (s *Service) initLog() {
	s.Logger.SetPrefix(s.conf.Prefix)
	s.Logger.SetLevel(logLevels[s.conf.LogLevel])
}

func (s *Service) Close() {
	s.dao.Close()
}

func (s *customizeRdsStore) Set(id string, value string) {
	c := s.redisClient.Get()
	defer c.Close()
	c.Do("SETEX", id, 60, value)
}

func (s *customizeRdsStore) Get(id string, clear bool) (value string) {
	c := s.redisClient.Get()
	defer c.Close()

	reply, err := redis.String(c.Do("GET", id))
	if err == nil {
		value = reply
	}

	if clear {
		c.Do("DEL", id)
	}
	return
}
