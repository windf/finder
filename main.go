package main

import (
	"finder/models"
	_ "finder/routers"
	"github.com/astaxie/beego"
)

func init() {
	models.MysqlConnect()
	//todo redis
}

func main() {
	beego.Run()
}
