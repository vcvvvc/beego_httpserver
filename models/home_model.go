package models

import (
	"bytes"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	beego "github.com/beego/beego/v2/server/web"
	"html/template"
	"httpserver/util"
	"strconv"

	"strings"
)

type HomeFooterPageCode struct {
	HasPre   bool
	HasNext  bool
	ShowPage string
	PreLink  string
	NextLink string
}

type HomeBlockParam struct {
	Id         int
	Title      string
	Tags       []TagLink
	Short      string
	Content    string
	Author     string
	CreateTime string
	//查看文章的地址
	Link string

	//修改文章的地址
	UpdateLink string
	DeleteLink string

	//记录是否登录
	IsLogin bool
}

// 标签链接
type TagLink struct {
	TagName string
	TagUrl  string
}

func FindArticle(page int) ([]Article, error) {
	num, _ := config.Int("articleListPageNum")
	page--
	return FindArticleCon(page, num)
}

func FindArticleCon(page int, num int) ([]Article, error) {
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("*").
		From("User.article").Limit(num).Offset(page * num) //limit 8, offset 0 第一条开始取8条

	// 导出 SQL 语句
	sql := qb.String()
	var artList []Article
	utilArtList, err := util.QueryArticleDB(sql)

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

// 将tags字符串转化成首页模板所需要的数据结构
func createTagsLinks(tags string) []TagLink {
	var tagLink []TagLink
	tagsPamar := strings.Split(tags, ", ")
	for _, tag := range tagsPamar {
		tagLink = append(tagLink, TagLink{tag, "/?tag=" + tag})
	}
	return tagLink
}

func MakeHomeBlocks(articles []Article, isLogin bool) template.HTML {
	htmlHome := ""
	for _, art := range articles {
		//将数据库model转换为首页模板所需要的model
		homeParam := HomeBlockParam{}
		homeParam.Id = art.Id
		homeParam.Title = art.Title
		homeParam.Tags = createTagsLinks(art.Tags)
		fmt.Println("tag-->", art.Tags)
		homeParam.Short = art.Short
		homeParam.Content = art.Content
		homeParam.Author = art.Author
		homeParam.CreateTime = art.Createtime.Format("2006-01-02 15:04:05")
		homeParam.Link = "/article/" + strconv.Itoa(art.Id)
		homeParam.UpdateLink = "/article/editarticle?id=" + strconv.Itoa(art.Id)
		homeParam.DeleteLink = "/article/delete?id=" + strconv.Itoa(art.Id)
		homeParam.IsLogin = isLogin

		//处理变量
		//ParseFile解析该文件，用于插入变量
		t, _ := template.ParseFiles("views/block/home_block.html")
		buffer := bytes.Buffer{}
		//就是将html文件替换为传入的数据
		t.Execute(&buffer, homeParam)
		htmlHome += buffer.String()
	}
	fmt.Println("htmlHome-->", htmlHome)
	return template.HTML(htmlHome)
}

// -----------翻页-----------
func ConfigHomepagenum(page int) HomeFooterPageCode {
	pageCode := HomeFooterPageCode{}
	//查询出总的条数
	nums := GetArticleNum()
	//从配置文件中读取每页显示的条数
	pageRow, _ := beego.AppConfig.Int("articleListPageNum")
	//计算出总页数
	allPageNum := (nums-1)/pageRow + 1
	pageCode.ShowPage = fmt.Sprintf("%d/%d", page, allPageNum)
	//当前页数小于等于1，那么上一页的按钮不能点击
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}
	//当前页数大于等于总页数，那么下一页的按钮不能点击
	if page >= allPageNum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}
	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)
	return pageCode
}

// 设置页数
//func SetArticleRowsNum() {
//	artcileRowsNum = ConfigHomepagenum()
//}
