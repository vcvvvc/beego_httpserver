package models

import (
	"github.com/beego/beego/v2/client/orm"
	"httpserver/util"
	"strings"
)

func HandleTagsListData(tags []string) map[string]int {
	var tagsMap = make(map[string]int)
	for _, tag := range tags {
		tagList := strings.Split(tag, ", ")
		for _, value := range tagList {
			tagsMap[value]++
		}
	}
	return tagsMap
}

// 查询标签，返回一个字段的列表
func QueryArticleTags() []string {
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("tags").From("User.article")
	//limit {number | all}：表示最多返回多少行数据，如果是all，表示返回所有数据。
	//offset number：表示跳过多少行数据，从第number+1行开始返回。

	sql := qb.String()
	lists := util.CreateTagsListDB(sql)
	return lists
}

func ArticlesWithTag(tag string) ([]Article, error) {
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("*").From("User.article").Where("tags LIKE ?")
	sql := qb.String()
	var artList []Article
	utilArtList, err := util.QueryArticlesTagDB(sql, tag)

	if err != nil {
		return nil, err
	}
	for _, utilArt := range utilArtList {
		art := Article{
			Id:         utilArt.Id,
			Title:      utilArt.Title,
			Tags:       utilArt.Tags,
			Short:      utilArt.Short,
			Content:    utilArt.Content,
			Author:     utilArt.Author,
			Createtime: utilArt.Createtime,
		}
		artList = append(artList, art)
	}
	return artList, nil
}
