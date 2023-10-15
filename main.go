package main

import (
	"github.com/beego/beego/v2/server/web"
	"httpserver/models"
	_ "httpserver/routers"
)

func init() {
	models.Init()
	web.BConfig.WebConfig.Session.SessionOn = true
}

func main() {
	web.Run(":8081")
}
