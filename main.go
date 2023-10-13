package main

import "httpserver/models"

func init() {
	// need to register models in init
	//models.InsertDB()
	//models.DeleteDB(3)
	//models.UpdatePWD(2, "tesxxxxtvvvvv")
	models.SearchUser(5, "test_transaction")
}

func main() {
	println("")

}
