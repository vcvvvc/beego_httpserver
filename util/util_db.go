package util

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/config"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

// User -
//type User struct {
//	Id int `json:"id" orm:"pk;auto;description(主键id)"` // 设置主键且自动增长
//	Name string `json:"name" orm:"index;unique;size(50);description(用户名)"` // 唯一的且加了索引的
//	Age int `json:"age" orm:"default(0);description(年龄)"`
//	Salary float64 `json:"price" orm:"digits(12);decimals(2);description(薪资)"`
//	Address string `json:"address" orm:"size(100);null;column(address);description(地址)"` // 可以为空
//	// 创建时间字段
//	CreateAt *time.Time `json:"create_at" orm:"auto_now_add;type(datetime);description(创建时间)"`
//	UpdateAt *time.Time `json:"update_at" orm:"auto_now;type(datetime);description(更新时间)"`
//}

type User struct {
	ID            int       `orm:"column(id);auto;index;unique"`
	UserName      string    `orm:"column(username);unique"`
	PassWord      string    `orm:"column(password)"`
	Register_time time.Time `orm:"column(register_time)"`
}

type Article struct {
	Id         int       `orm:"column(id);auto;index;unique"`
	Title      string    `orm:"column(title)"`
	Tags       string    `orm:"column(tags)"`
	Short      string    `orm:"column(short)"`
	Content    string    `orm:"column(content)"`
	Author     string    `orm:"column(author)"`
	Createtime time.Time `orm:"column(create_time)"`
}

type File struct {
	Id         int       `orm:"column(id);auto;index;unique"`
	FileName   string    `orm:"column(filename)"`
	FilePath   string    `orm:"column(filepath)"`
	Filehash   string    `orm:"column(filehash)"`
	Filetype   string    `orm:"column(filetype)"`
	Createtime time.Time `orm:"column(create_time)"`
}

func Init() {
	dbhost, _ := config.String("dbhost")
	dbport, _ := config.String("dbport")
	dbuser, _ := config.String("dbuser")
	dbpassword, _ := config.String("dbpassword")
	dbname, _ := config.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}
	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"
	orm.RegisterDataBase("default", "mysql", dsn)
	//orm.RegisterModel(new(User), new(Category), new(Post), new(Config), new(Comment))
	// need to register models in init
	orm.RegisterModel(new(User), new(Article), new(File))

	// need to register default database
	//orm.RegisterDataBase("default", "mysql", "6923403:zxz123456@tcp(192.168.31.172:3306)/User?charset=utf8&parseTime=true&loc=Local")
	orm.Debug = true
	//orm.RunSyncdb("default", false, true)
}

// 设置引擎为 INNODB
func (ue *User) TableEngine() string {
	return "INNODB"
}

// 设置引擎为 INNODB
func (ae *Article) TableEngine() string {
	return "INNODB"
}

// 指定Order结构体默认绑定的表名
func (o *User) TableName() string {
	return "users"
}

// 指定Order结构体默认绑定的表名
func (a *Article) TableName() string {
	return "article"
}

func InsertUserDB(username string, password string) bool {
	if username == "" {
		return false
	}

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.InsertInto("User.users", "`username`, `password`, `register_time`").
		Values("?", "?", "?")
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	res, err := o.Raw(sql, username, password, time.Now()).Exec()
	// 执行SQL语句
	//res, err := o.Raw(sql, values...).Exec()
	if err != nil {
		fmt.Println(err)
		return false
	}

	// 获取影响的行数
	id, err := res.LastInsertId()
	//num, _ := res.RowsAffected()
	//num, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)

	return true
}

func SearchUserDB(username string) bool {
	var users []User
	if username == "" {
		return false
	}

	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").
		From("User.users").
		Where("username = ?")
	sql := qb.String()
	// 执行 SQL 语句
	o := orm.NewOrm()
	id, err := o.Raw(sql, username).QueryRows(&users)
	fmt.Println(id)
	if id == 0 {
		fmt.Println(err)
		return false
	}

	fmt.Println("UID: ", users[0].ID, " ,username: ", users[0].UserName)
	return true
}

type ULogin struct {
	username string
	password string
}

func UserLoginDB(username string, password string) bool {
	// 获取 QueryBuilder 对象。需要指定数据库驱动参数。
	// 第二个返回值是错误对象，在这里略过
	var users []User
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("*").
		From("User.users").
		Where("username = ? AND password = ?")
	//limit {number | all}：表示最多返回多少行数据，如果是all，表示返回所有数据。
	//offset number：表示跳过多少行数据，从第number+1行开始返回。

	// 导出 SQL 语句
	sql := qb.String()

	// 执行 SQL 语句
	o := orm.NewOrm()
	id, err := o.Raw(sql, username, password).QueryRows(&users)
	if id == 1 {
		return true
	}
	if err != nil {
		fmt.Println(err)
	}

	return false
}

func InsertArticleDB(sql string, title string, author string, tags string, short string, content string) bool {
	o := orm.NewOrm()
	res, err := o.Raw(sql, title, author, tags, short, content, time.Now()).Exec()
	if err != nil {
		fmt.Println(err)
		return false
	}

	// 获取影响的行数
	id, err := res.LastInsertId()
	//num, _ := res.RowsAffected()
	//num, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(id)

	return true
}

func QueryArticleDB(sql string) ([]Article, error) {
	o := orm.NewOrm()
	var artList []Article
	res, err := o.Raw(sql).QueryRows(&artList)
	if err != nil {
		fmt.Println(res, err)
		return nil, err
	}

	//fmt.Println(id)
	return artList, nil
}

func CountArticles(sql string) string {
	// 构建查询对象
	o := orm.NewOrm()

	var maps []orm.Params
	_, err := o.Raw(sql).Values(&maps)
	if err != nil {
		fmt.Println(err)
	}

	nums := maps[0]["total"].(string)
	return nums
}

func WhereIdArticleDB(sql string, id int) ([]Article, error) {
	o := orm.NewOrm()
	var artList []Article
	res, err := o.Raw(sql, id).QueryRows(&artList)
	if err != nil {
		fmt.Println(res, err)
		return nil, err
	}

	//fmt.Println(id)
	return artList, nil
}

func CheckArticleAuthorDB(id int, author string) bool {
	qb, _ := orm.NewQueryBuilder("mysql")
	// 构建查询对象
	qb.Select("author = ? As result").From("User.article").Where("id = ?")
	sql := qb.String()

	o := orm.NewOrm()
	var maps []orm.Params
	res, err := o.Raw(sql, author, id).Values(&maps)
	if err != nil {
		fmt.Println(res, err)
	} else if maps[0]["result"].(string) != "1" {
		fmt.Println("result = ", maps[0]["result"])
		return false
	}
	return true
}

func UpdateArticleDB(sql string, title string, author string, tags string, short string, content string, id int) bool {
	if CheckArticleAuthorDB(id, author) == false {
		fmt.Println("作者不匹配")
		return false
	}
	o := orm.NewOrm()
	_, err := o.Raw(sql, title, author, tags, short, content, id).Exec()
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func DeleteArticleDB(sql string, id int, author string) bool {
	if CheckArticleAuthorDB(id, author) == false {
		fmt.Println("作者不匹配")
		return false
	}
	o := orm.NewOrm()
	_, err := o.Raw(sql, id, author).Exec()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func CreateTagsListDB(sql string) []string {
	o := orm.NewOrm()
	var tagslist []string
	_, err := o.Raw(sql).QueryRows(&tagslist)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return tagslist
}

func QueryArticlesTagDB(sql string, tag string) ([]Article, error) {
	o := orm.NewOrm()
	var artList []Article
	_, err := o.Raw(sql, "%"+tag+"%").QueryRows(&artList)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return artList, nil
}

func InsertFileDB(sql string, file_name string, file_path string, file_hash string, file_type string) bool {
	o := orm.NewOrm()
	_, err := o.Raw(sql, file_name, file_path, file_hash, file_type, time.Now()).Exec()
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func CheckFileDB(sql string, hash string) bool {
	o := orm.NewOrm()
	var file File
	err := o.Raw(sql, hash).QueryRow(&file)
	if err == nil {
		fmt.Println(err)
		return false
	}
	return true
}

func FindAllFileDB(sql string) ([]File, error) {
	o := orm.NewOrm()
	var File_list []File
	_, err := o.Raw(sql).QueryRows(&File_list)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return File_list, err
}
