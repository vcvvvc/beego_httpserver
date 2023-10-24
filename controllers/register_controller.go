package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"httpserver/util"
)

type RegisterController struct {
	web.Controller
}

func (u *RegisterController) RegisterPage() {
	u.TplName = "register.html"
}

func (u *RegisterController) Register() {
	username := u.GetString("username")
	password := u.GetString("password")
	println("username, password")

	unreg := util.SearchUserDB(username)
	if unreg == true {
		u.Data["json"] = map[string]interface{}{"code": 0, "message": "用户名已经存在"}
		err := u.ServeJSON()
		if err != nil {
			fmt.Println(err)
		}
		return
	}
	//注册用户名和密码
	//存储的密码是md5后的数据，那么在登录的验证的时候，也是需要将用户的密码md5之后和数据库里面的密码进行判断
	password = util.Md5(password)
	fmt.Println("md5后：", password)

	res := util.InsertUserDB(username, password)
	if res == true {
		u.Data["json"] = map[string]interface{}{"code": 1, "message": "注册成功"}
	} else {
		u.Data["json"] = map[string]interface{}{"code": -1, "message": "注册失败"}
	}
	err := u.ServeJSON()
	if err != nil {
		return
	}
}
