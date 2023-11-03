package controllers

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"httpserver/models"
	"strconv"
	"time"
)

type ArticleController struct {
	web.Controller
}

func (wr *ArticleController) WritePage() {
	author := wr.GetSession("Loginuser")
	if author == nil {
		wr.Redirect("/login", 302)
		return
	}
	wr.Data["IsLogin"] = true
	wr.TplName = "write_article.html"
}

func (wr *ArticleController) AddArticle() {
	//获取浏览器传输的数据，通过表单的name属性获取值
	id, _ := wr.GetInt("id")
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

	wr.Data["IsLogin"] = true
	fmt.Printf("title:%s,tags:%s\n", title, tags)
	//实例化model，将它出入到数据库中
	art := models.Article{id, title, tags, short, content, author.(string), time.Now()}
	succ_add := models.AddArticle(art)
	//返回数据给浏览器
	if succ_add {
		wr.Data["json"] = map[string]interface{}{"code": 3, "message": "文章添加成功"}
	} else {
		wr.Data["json"] = map[string]interface{}{"code": -3, "message": "文章添加出错"}
	}

	wr.ServeJSON()
}

func (ed *ArticleController) EditArticle() {
	id, _ := ed.GetInt("id")
	fmt.Println(id)
	//获取id所对应的文章信息
	art, err := models.QueryArticleId(id)
	if err != nil {
		fmt.Println(err)
		return
	}
	ed.Data["Title"] = art[0].Title
	ed.Data["Tags"] = art[0].Tags
	ed.Data["Short"] = art[0].Short
	ed.Data["Content"] = art[0].Content
	ed.Data["Id"] = art[0].Id

	author := ed.GetSession("Loginuser")
	if author == nil {
		ed.Data["json"] = map[string]interface{}{"code": -31, "message": "用户未登陆"}
		ed.ServeJSON()

		ed.Redirect("/login", 302)
		return
	}
	ed.Data["IsLogin"] = true
	ed.TplName = "write_article.html"
}

func (ua *ArticleController) UpdateArticle() {
	id, _ := ua.GetInt("id")
	title := ua.GetString("title")
	tags := ua.GetString("tags")
	short := ua.GetString("short")
	content := ua.GetString("content")
	author := ua.GetSession("Loginuser")
	if author == nil {
		ua.Data["json"] = map[string]interface{}{"code": -31, "message": "用户未登陆"}
		ua.ServeJSON()
		ua.Redirect("/login", 302)
		return
	}

	ua.Data["IsLogin"] = true
	art := models.Article{id, title, tags, short, content, author.(string), time.Now()}
	succ_update := models.QueryUpdateArticle(art)
	if succ_update {
		ua.Data["json"] = map[string]interface{}{"code": 4, "message": "文章更新成功"}
	} else {
		ua.Data["json"] = map[string]interface{}{"code": -41, "message": "文章更新出错"}
	}

	ua.ServeJSON()
}

func (de *ArticleController) DeleteArticle() {
	id, _ := de.GetInt("id")
	author := de.GetSession("Loginuser")
	if author == nil {
		de.Data["json"] = map[string]interface{}{"code": -31, "message": "用户未登陆"}
		de.ServeJSON()
		de.Redirect("/login", 302)
		return
	}

	de.Data["IsLogin"] = true
	succ_del := models.QueryDeleteArticle(id, author.(string))
	if succ_del {
		de.Data["json"] = map[string]interface{}{"code": 5, "message": "文章删除成功"}
	} else {
		de.Data["json"] = map[string]interface{}{"code": -51, "message": "文章删除出错"}
	}

	de.ServeJSON()
	de.Redirect("/login", 302)
}

func (ac *ArticleController) ArticleContent() {
	idStr := ac.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	fmt.Println("id:", id)
	//获取id所对应的文章信息
	art, err := models.QueryArticleId(id)
	if err != nil {
		fmt.Println(err)
	}

	ac.Data["Title"] = art[0].Title
	ac.Data["Content"] = art[0].Content
	////this.Data["Content"] = utils.SwitchMarkdownToHtml(art.Content)

	author := ac.GetSession("Loginuser")
	if author != nil {
		ac.Data["IsLogin"] = true
	}

	ac.TplName = "show_article.html"
}
