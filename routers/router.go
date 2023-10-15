package routers

import (
	"github.com/beego/beego/v2/server/web"
	"httpserver/controllers"
)

func init() {
	web.CtrlGet("/", (*controllers.RegisterController).RegisterPage)
	web.CtrlPost("/register", (*controllers.RegisterController).Register)
	web.CtrlGet("/loginPage", (*controllers.LoginController).loginPage)
	web.CtrlPost("/login", (*controllers.LoginController).login)

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
