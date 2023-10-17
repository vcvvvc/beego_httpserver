package controllers

import (
	"fmt"
)

type HomeController struct {
	//beego.Controller
	BaseController
}

func (this *HomeController) HomePage() {
	this.BaseController.UserIsLogin()
	fmt.Println("IsLogin: ", this.IsLogin, ", username: ", this.Loginuser)
	if this.IsLogin == false {
		this.Redirect("/login", 302) //若Session中无用户ID则302重定向至登陆页面
	} else {
		this.TplName = "home.html"
	}
}
