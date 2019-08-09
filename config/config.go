package config

import (
	"github.com/BurntSushi/toml"
	"github.com/mojocn/base64Captcha"
)

type Config struct {
	Database      *Database
	Redis         *RedisConfig
	Env           string
	Cache         *Cache
	LogLevel      string
	HttpAddress   string
	Prefix        string
	Secret        string
	CaptchaConfig base64Captcha.ConfigDigit
}

type Database struct {
	Read  *Databases
	Write *Databases
}

type RedisConfig struct {
	Addr string
	Psw  string
}

type Databases struct {
	Addr        string
	Active      int
	Idle        int
	IdleTimeout int
}

type Cache struct {
	Expire int
}

var (
	Conf = &Config{}
)

func Init(filePath string) (err error) {
	_, err = toml.DecodeFile(filePath, &Conf)
	Conf.CaptchaConfig = base64Captcha.ConfigDigit{
		Height:     80,
		Width:      240,
		MaxSkew:    0.7,
		DotCount:   80,
		CaptchaLen: 5,
	}
	return
}
