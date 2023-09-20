package main

import (
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

func (u *UserController) HelloWorld() {
	u.Ctx.WriteString("hello, world")

}

func (u *UserController) Index() {
	u.TplName = "index.html"
	u.Render()
}

func main() {
	web.SetStaticPath("/", "blog")
	//web.CtrlGet("/", (*UserController).Index)
	web.CtrlGet("/helloworld", (*UserController).HelloWorld)
	web.Run()
}
