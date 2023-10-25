package controllers

import (
	"fmt"
	"httpserver/models"
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
		fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)
		page, _ := this.GetInt("page")
		if page <= 0 {
			page = 1
		}
		var artList []models.Article
		artList, err := models.FindArticle(page)
		fmt.Println(err)
		this.Data["PageCode"] = models.ConfigHomepagenum(page)
		this.Data["HasFooter"] = true
		this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)

		this.TplName = "home.html"
	}
}
