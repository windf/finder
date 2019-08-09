package dao

import (
	"finder/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/json-iterator/go"
	"strings"
	"time"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Dao struct {
	Cpool *redis.Pool
	dbr   *gorm.DB
	dbw   *gorm.DB
}

type RedisCfg struct {
	Addr string `toml:"addr" json:"Addr"`
	Psw  string `toml:"psw" json:"Psw"`
}

func NewDao(c *config.Config) (d *Dao, err error) {
	dao := new(Dao)
	dbr, err := gorm.Open("mysql", c.Database.Read.Addr)
	if err != nil {
		panic(err)
	}
	dbr.DB().SetMaxIdleConns(c.Database.Read.Idle)
	dbr.DB().SetMaxOpenConns(c.Database.Read.Active)
	dbr.DB().SetConnMaxLifetime(time.Duration(c.Database.Read.IdleTimeout) * time.Minute)
	dao.dbr = dbr

	dbw, err := gorm.Open("mysql", c.Database.Write.Addr)
	if err != nil {
		panic(err)
	}
	dbw.DB().SetMaxIdleConns(c.Database.Write.Idle)
	dbw.DB().SetMaxOpenConns(c.Database.Write.Active)
	dbw.DB().SetConnMaxLifetime(time.Duration(c.Database.Write.IdleTimeout) * time.Minute)
	dao.dbw = dbw

	redisCfg := RedisCfg{
		Addr: c.Redis.Addr,
		Psw:  c.Redis.Psw,
	}
	redis := NewRedisPool(&redisCfg)
	dao.Cpool = redis
	if strings.ToLower(c.Env) != "prod" {
		dbw.LogMode(true)
		dbr.LogMode(true)
	}
	return dao, nil
}

func NewRedisPool(cfg *RedisCfg) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 300 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.Addr)
			if err != nil {
				return nil, err
			}
			if len(cfg.Psw) > 0 {
				if _, err := c.Do("AUTH", cfg.Psw); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func (d *Dao) Close() {
	d.Cpool.Close()
	d.dbr.Close()
	d.dbw.Close()
}
