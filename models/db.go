package models

import (
	"context"
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
	ID            int       `orm:"column(id);index;unique"`
	UserName      string    `orm:"column(username);unique"`
	PassWord      string    `orm:"column(password)"`
	Register_time time.Time `orm:"column(register_time)"`
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
	orm.RegisterModel(new(User))

	// need to register default database
	//orm.RegisterDataBase("default", "mysql", "6923403:zxz123456@tcp(192.168.31.172:3306)/User?charset=utf8&parseTime=true&loc=Local")
	orm.Debug = true
}

// 设置引擎为 INNODB
func (u *User) TableEngine() string {
	return "INNODB"
}

// 指定Order结构体默认绑定的表名
func (o *User) TableName() string {
	return "users"
}

func InsertDB() {
	orm.RunSyncdb("default", false, true)
	o := orm.NewOrm()
	// 在闭包内执行事务处理
	err := o.DoTx(func(ctx context.Context, txOrm orm.TxOrmer) error {
		// 准备数据
		user := new(User)
		user.UserName = "test_transaction"
		user.PassWord = "jtestvvvv000"
		user.Register_time = time.Now()

		// 插入数据
		// 使用txOrm执行SQL
		_, e := txOrm.Insert(user)
		return e
	})
	if err != nil {
		println(err)
	} else {
		println("插入成功")
	}

	//us := new(User)
	//// 对order对象赋值
	//us.UserName = "vcvc"
	//us.PassWord = "zxzmima123456"
	//us.Register_time = time.Now()
	//
	//id, err := o.Insert(us)
	//if err != nil {
	//	fmt.Println("插入失败")
	//} else {
	//	// 插入成功会返回插入数据自增字段，生成的id
	//	fmt.Println("新插入数据的id为:", id)
	//}

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
		println("其他问题")
	}

}

func SearchUser(uid int, username string) {
	o := orm.NewOrm()
	us := new(User)
	us.ID = uid
	if username != "" {
		us.UserName = username
	}

	err := o.Read(us)
	if err != nil {
		println("查询出错，", err)
	} else {
		println("查询成功", us.ID, us.UserName, us.PassWord, us.Register_time.String())
	}
}
