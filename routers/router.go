package routers

import (
	"finder/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/record", &controllers.RecordController{})
	beego.Router("/user", &controllers.UserController{}, "*:Index")
	beego.Router("/api/user", &controllers.UserController{}, "*:ApiUser")

	//admin todo restful
	beego.Router("/admin/user/add", &controllers.UserController{}, "POST:Add")
	beego.Router("/admin/user/update", &controllers.UserController{}, "POST:Update")
	beego.Router("/admin/user/delete", &controllers.UserController{}, "POST:Delete")
}
