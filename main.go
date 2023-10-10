package main

import (
	"httpserver/models"
)

func init() {
	// need to register models in init
	//models.InsertDB()
	models.DeleteDB(1)
}

func main() {
	println("ssss")

}
