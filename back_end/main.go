package main

import (
	"contribution/data"
	"contribution/database"
)

func main() {
	database.Open("data.db")
	data.StoreExampleData()
	data.GetExampleData()
	data.UpdataExampleData()
	data.GetExampleData()
	data.DeleteExampleData()
	// server := gin.Default()

	// server.Run()
}
