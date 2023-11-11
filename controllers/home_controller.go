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
		return
	}
	tag := this.GetString("tag")
	var artList []models.Article
	if len(tag) > 0 {
		//按照指定的标签搜索
		artList, _ = models.ArticlesWithTag(tag)
		this.Data["HasFooter"] = false
	} else {
		fmt.Println("IsLogin:", this.IsLogin, this.Loginuser)
		page, _ := this.GetInt("page")
		if page <= 0 {
			page = 1
		}
		artList, _ = models.FindArticle(page)
		this.Data["PageCode"] = models.ConfigHomepagenum(page)
		this.Data["HasFooter"] = true

	}
	this.Data["Content"] = models.MakeHomeBlocks(artList, this.IsLogin)
	this.TplName = "home.html"
}
