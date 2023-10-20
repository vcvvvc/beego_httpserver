package main

import (
	"github.com/beego/beego/v2/server/web"
	_ "httpserver/routers"
	"httpserver/util"
)

func init() {
	util.Init()
	web.BConfig.WebConfig.Session.SessionOn = true

}

func main() {
	web.Run(":8081")
}
