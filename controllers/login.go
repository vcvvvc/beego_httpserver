package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"httpserver/util"
)

type LoginController struct {
	web.Controller
}

func (u *LoginController) LoginPage() {
	u.TplName = "login.html"
}

func (u *LoginController) Login() {
	username := u.GetString("username")
	password := u.GetString("password")
	password_md5 := util.Md5(password)
	m_success := util.UserLoginDB(username, password_md5)
	if m_success != true {
		u.Data["json"] = map[string]interface{}{"code": -2, "message": "密码不正确"}
		err := u.ServeJSON()
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	u.SetSession("Loginuser", username)
	u.Data["json"] = map[string]interface{}{"code": 2, "message": "账号登录成功"}
	u.ServeJSON()
}

func (u *LoginController) TestInputGet() {
	//读取session
	username := u.GetSession("Loginuser")
	if nameString, ok := username.(string); ok && nameString != "" {
		u.Ctx.WriteString("Username:" + username.(string))
	} else {
		fmt.Println("session error")
	}
}

func (u *LoginController) UserExit() {
	u.DelSession("Loginuser")
	u.Redirect("/login", 302)
}
