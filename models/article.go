package models

import (
	"github.com/beego/beego/v2/client/orm"
	"httpserver/util"
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
