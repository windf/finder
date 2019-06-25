package controllers

//前台记录
import (
	"github.com/astaxie/beego"
)

type RecordController struct {
	beego.Controller
}

func (c *RecordController) Get() {
	c.Ctx.WriteString("record front")
}
