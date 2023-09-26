package main

import (
	"github.com/beego/beego/v2/client/orm"
	// don't forget this
	_ "github.com/go-sql-driver/mysql"
)

// User -
type User struct {
	ID       int    `orm:"column(id)"`
	Name     string `orm:"column(username)"`
	Password string `orm:"column(password)"`
}

func init() {
	// need to register models in init
	orm.RegisterModel(new(User))

	// need to register default database
	orm.RegisterDataBase("default", "mysql", "root:Aa1248800211@tcp(127.0.0.1:3306)/gotest?charset=utf8")
}

func main() {
	// automatically build table
	orm.RunSyncdb("default", false, true)

	// create orm object
	o := orm.NewOrm()

	// data
	user := new(User)
	user.Name = "mike"

	// insert data
	o.Insert(user)
}
