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
	Title      string    `orm:"column(username)"`
	Tags       string    `orm:"column(tags)"`
	Short      string    `orm:"column(short)"`
	Content    string    `orm:"column(content)"`
	Author     string    `orm:"column(author)"`
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
	orm.RegisterModel(new(User), new(Article))

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

func InsertUser(username string, password string) bool {
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

func DeleteDB(uid int) {
	o := orm.NewOrm()
	us := new(User)
	us.ID = uid

	id, err := o.Delete(us)
	if err != nil {
		fmt.Println("删除失败")
	} else {
		// 插入成功会返回插入数据自增字段，生成的id
		fmt.Println("删除的数据id为:", id)
	}
}

func UpdatePWD(uid int, pwd string) {
	o := orm.NewOrm()
	us := new(User)
	us.ID = uid
	us.PassWord = pwd

	id, err := o.Update(us, "password")
	if err != nil {
		fmt.Println("更新失败")
	} else if id == 1 {
		// 插入成功会返回插入数据自增字段，生成的id
		fmt.Println("更新成功")
	} else {
		fmt.Println("其他问题")
	}

}

func SearchUser(username string) bool {
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

func UserLogin(username string, password string) bool {
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

	// 处理查询结果（根据实际需求）
}
