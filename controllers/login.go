package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"httpserver/models"
	"httpserver/util"
)

type LoginController struct {
	web.Controller
}

func (l *LoginController) LoginPage() {
	l.TplName = "login.html"
}

func (l *LoginController) Login() {
	username := l.GetString("username")
	password := l.GetString("password")
	password_md5 := util.Md5(password)
	m_success := models.UserLogin(username, password_md5)
	if !m_success {
		l.Data["json"] = map[string]interface{}{"code": -2, "message": "密码不正确"}
		err := l.ServeJSON()
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	l.SetSession("Loginuser", username)
	l.Data["json"] = map[string]interface{}{"code": 2, "message": "账号登录成功"}
	l.ServeJSON()

}

func (c *LoginController) TestInputGet() {
	//读取session
	username := c.GetSession("username")
	//password := c.GetSession("password")
	if nameString, ok := username.(string); ok && nameString != "" {
		c.Ctx.WriteString("Username:" + username.(string))
	} else {
		fmt.Println("session error")
	}
}
