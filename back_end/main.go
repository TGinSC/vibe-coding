package main

import (
	"contribution/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Open("data.db")
	server := gin.Default()

	server.Run()
}
