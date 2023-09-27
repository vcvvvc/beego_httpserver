package models

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"

	// don't forget this
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	// ...
}

// User -
type User struct {
	ID       int    `orm:"column(id)"`
	Name     string `orm:"column(username)"`
	Password string `orm:"column(password)"`
}

// 设置引擎为 INNODB
func (u *User) TableEngine() string {
	return "INNODB"
}

const ConfigFile = "./conf/app.conf"

func RegisterDB() {
	err := web.LoadAppConfig("conf", ConfigFile)
	if err != nil {
		logs.Critical("An error occurred:", err)
		panic(err)
	}

	val, _ := web.AppConfig.String("name")

	logs.Info("load config name is", val)

}
