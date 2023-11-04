package controllers

import (
	"github.com/beego/beego/v2/server/web"
	"httpserver/models"
)

type TagsController struct {
	web.Controller
}

func (tg *TagsController) GetTagsList() {
	tags := models.QueryArticleTags()
	tg.Data["Tags"] = models.HandleTagsListData(tags)
	tg.TplName = "tags.html"

}
