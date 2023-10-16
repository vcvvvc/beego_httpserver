package routers

import (
	"github.com/beego/beego/v2/server/web"
	"httpserver/controllers"
)

func init() {
	web.CtrlGet("/", (*controllers.RegisterController).RegisterPage)
	web.CtrlPost("/register", (*controllers.RegisterController).Register)
	web.CtrlGet("/login", (*controllers.LoginController).LoginPage)
	web.CtrlPost("/loginreq", (*controllers.LoginController).Login)

	//beego.Router("/home", &controllers.BlogController{}, "*:Home")
	//
	//beego.AutoRouter(&controllers.AdminController{})
}
