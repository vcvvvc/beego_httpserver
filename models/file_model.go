package models

import (
	"github.com/beego/beego/v2/client/orm"
	"httpserver/util"
	"time"
)

//create table if not exists file(
//id int(8) primary key auto_increment not null,
//filename varchar(255) not null,
//filepath varchar(255) not null,
//filehash varchar(50) not null,
//filetype varchar(255),
//create_time datetime not null
//) ENGINE=InnoDB, default charset=utf8;

type File struct {
	Id         int
	FileName   string
	FilePath   string
	Filehash   string
	Filetype   string
	Createtime time.Time
}

func QueryInsertFile(file_name string, file_path string, file_hash string, file_type string) bool {
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.InsertInto("User.file", "filename, filepath, filehash, filetype, create_time").
		Values("?", "?", "?", "?", "?")

	// 导出 SQL 语句
	sql := qb.String()

	checkfile := QueryCheckFile(file_hash)
	if checkfile {
		succ := util.InsertFileDB(sql, file_name, file_path, file_hash, file_type)
		return succ
	}

	return false
}

func QueryCheckFile(hash string) bool {
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("*").From("User.file").Where("filehash = ?")

	// 导出 SQL 语句
	sql := qb.String()
	return util.CheckFileDB(sql, hash)
}
