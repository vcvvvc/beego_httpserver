package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"httpserver/models"
	"time"
)

type ArticleController struct {
	web.Controller
}

func (wr *ArticleController) WritePage() {
	wr.TplName = "write_article.html"
}

func (wr *ArticleController) AddArticle() {
	//获取浏览器传输的数据，通过表单的name属性获取值
	title := wr.GetString("title")
	tags := wr.GetString("tags")
	short := wr.GetString("short")
	content := wr.GetString("content")
	author := wr.GetSession("Loginuser")
	if author == nil {
		wr.Data["json"] = map[string]interface{}{"code": -31, "message": "用户未登陆"}
		wr.ServeJSON()

		wr.Redirect("/login", 302)
		return
	}
	fmt.Printf("title:%s,tags:%s\n", title, tags)
	//实例化model，将它出入到数据库中
	art := models.Article{0, title, author.(string), tags, short, content, time.Now()}
	succ_add := models.AddArticle(art)
	//返回数据给浏览器
	if succ_add {
		wr.Data["json"] = map[string]interface{}{"code": 3, "message": "文章添加成功"}
	} else {
		wr.Data["json"] = map[string]interface{}{"code": -3, "message": "文章添加出错"}
	}

	wr.ServeJSON()
}

func (ua *ArticleController) UpdateArticle() {
}

//func ()
