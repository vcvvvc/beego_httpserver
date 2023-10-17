package routers

import (
	"github.com/beego/beego/v2/server/web"
	"httpserver/controllers"
)

func init() {
	//HOME
	web.CtrlGet("/", (*controllers.HomeController).HomePage)

	//Register
	web.CtrlGet("/register", (*controllers.RegisterController).RegisterPage)
	web.CtrlPost("/user_register", (*controllers.RegisterController).Register)

	//Login
	web.CtrlGet("/login", (*controllers.LoginController).LoginPage)
	web.CtrlPost("/user_login", (*controllers.LoginController).Login)
	web.CtrlGet("/session_input", (*controllers.LoginController).TestInputGet)
	web.CtrlGet("/exit", (*controllers.LoginController).UserExit)
}
