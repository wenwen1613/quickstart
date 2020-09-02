package routers

import (
	"quickstart/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//beego.Router("/user", &controllers.UserController{})
	//beego.Router("/user/info", &controllers.UserController{}, "GET:Info")
	//beego.Router("/user/list", &controllers.UserController{}, "GET:List")
	//beego.Router("/user/add", &controllers.UserController{}, "POST:Add")
	//beego.Router("/upload", &controllers.FormController{}, "POST:Upload")
	//beego.Get("/test", func(context *context.Context) {
	//	context.Output.Body([]byte("heeo"))
	//})
	beego.Include(&controllers.MainController{}, &controllers.UserController{}, &controllers.FormController{})
}
