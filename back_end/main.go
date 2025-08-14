package main

import (
	"contribution/config"
	"contribution/database"
	"contribution/route"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	// 加载环境变量（开发环境使用）
	if err := godotenv.Load(); err != nil {
		log.Println("Note: No .env file found (required for HuggingFace API in development)")
	}
	__config__ := config.Config__
	database.Open(__config__.DB_FILE)
	server := gin.Default()

	route.BindRoutes(server)
	route.RunAI(server)

	server.Run(__config__.HttpPort)
}
