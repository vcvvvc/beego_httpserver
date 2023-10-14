package routers

import (
	"github.com/beego/beego/v2/server/web"
	"httpserver/controllers"
)

func init() {
	web.CtrlGet("/", (*controllers.UserController).Register)
	//web.CtrlGet("/Register", (*controllers.UserController).Register)

	//beego.Router("/home", &controllers.BlogController{}, "*:Home")
	//beego.Router("/article", &controllers.BlogController{}, "*:Article")
	//beego.Router("/detail", &controllers.BlogController{}, "*:Detail")
	//beego.Router("/about", &controllers.BlogController{}, "*:About")
	//beego.Router("/timeline", &controllers.BlogController{}, "*:Timeline")
	//beego.Router("/resource", &controllers.BlogController{}, "*:Resource")
	//beego.Router("/comment", &controllers.BlogController{}, "post:Comment")
	//
	//beego.AutoRouter(&controllers.AdminController{})
}
