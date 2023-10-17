package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type BaseController struct {
	web.Controller
	IsLogin   bool
	Loginuser interface{}
}

// 判断是否登录
func (b *BaseController) UserIsLogin() {
	loginuser := b.GetSession("Loginuser")
	if loginuser != nil {
		b.IsLogin = true
		b.Loginuser = loginuser
	} else {
		b.IsLogin = false
	}
	b.Data["IsLogin"] = b.IsLogin
}
