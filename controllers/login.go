package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type LoginController struct {
	web.Controller
}

func (l *LoginController) loginPage() {
	l.TplName = "login.html"
}

func (l *LoginController) login() {
}
