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

	//Write
	web.CtrlGet("/article/add", (*controllers.ArticleController).WritePage)
	web.CtrlPost("/article/create", (*controllers.ArticleController).AddArticle)

	//Article
	web.CtrlGet("/article/:id([0-9]+)", (*controllers.ArticleController).ArticleContent)
	web.CtrlGet("/article/editarticle", (*controllers.ArticleController).EditArticle)
	web.CtrlPost("/article/updatearticle", (*controllers.ArticleController).UpdateArticle)
	web.CtrlGet("/article/delete)", (*controllers.ArticleController).DeleteArticle)
}
