package main

import (
	"contribution/database"
)

func main() {
	database.Open("data.db")
	database.StoreExampleData()
	database.GetExampleData()
	database.UpdataExampleData()
	database.GetExampleData()
	database.DeleteExampleData()
	// server := gin.Default()

	// server.Run()
}
