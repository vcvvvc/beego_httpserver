package models

import (
	"github.com/beego/beego/v2/client/orm"
	"httpserver/util"
	"strconv"
	"time"
)

//create table if not exists article(
//id int(8) primary key auto_increment not null,
//title varchar(30) not null,
//author varchar(20) not null,
//tags varchar(50) not null,
//short varchar(255),
//content longtext,
//create_time datetime not null
//) ENGINE=InnoDB, default charset=utf8;

type Article struct {
	Id         int
	Title      string
	Tags       string
	Short      string
	Content    string
	Author     string
	Createtime time.Time
	//Status int //Status=0为正常，1为删除，2为冻结
}

// ---------数据处理-----------
func AddArticle(article Article) bool {
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.InsertInto("User.article", "title, author, tags, short, content, create_time").
		Values("?", "?", "?", "?", "?", "?")
	//limit {number | all}：表示最多返回多少行数据，如果是all，表示返回所有数据。
	//offset number：表示跳过多少行数据，从第number+1行开始返回。

	// 导出 SQL 语句
	sql := qb.String()
	succ_add := util.InsertArticleDB(sql, article.Title, article.Author, article.Tags, article.Short, article.Content)
	return succ_add
}

func GetArticleNum() int {
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("COUNT(*) AS total").From("User.article")
	sql := qb.String()
	str_artnums := util.CountArticles(sql)
	artnums, _ := strconv.Atoi(str_artnums)
	return artnums
}

func QueryArticleId(id int) ([]Article, error) {
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("*").From("User.article").Where("id = ?")
	sql := qb.String()
	var artList []Article
	utilArtList, err := util.WhereIdArticleDB(sql, id)

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
	//reflect反射
	return artList, nil
}

func QueryUpdateArticle(article Article) bool {
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Update("User.article").Set("title = ?", "author = ?", "tags = ?", "short = ?", "content = ?").Where("id = ?")
	// 导出 SQL 语句
	sql := qb.String()

	succ_update := util.UpdateArticleDB(sql, article.Title, article.Author, article.Tags, article.Short, article.Content, article.Id)
	return succ_update
}

func QueryDeleteArticle(id int, author string) bool {
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Delete().From("User.article").Where("id = ?").And("author = ?")
	//limit {number | all}：表示最多返回多少行数据，如果是all，表示返回所有数据。
	//offset number：表示跳过多少行数据，从第number+1行开始返回。

	// 导出 SQL 语句
	sql := qb.String()
	succ_update := util.DeleteArticleDB(sql, id, author)
	return succ_update
}
