package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Database *Database
	Redis    *RedisConfig
	Env      string
	Cache    *Cache
}

type Database struct {
	Read  *Databases
	Write *Databases
}

type RedisConfig struct {
	Addr string
	Psw  string
	Dbno int
}

type Databases struct {
	Addr string
}

type Cache struct {
	Expire int
	Limit  int
}

var (
	Conf = &Config{}
)

func Init(filePath string) (err error) {
	_, err = toml.DecodeFile(filePath, &Conf)
	return
}
