package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "httpserver/routers"
)

func main() {
	beego.Run()
}
