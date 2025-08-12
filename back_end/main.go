package main

import (
	"contribution/config"
	"contribution/database"
	"contribution/route"

	"github.com/gin-gonic/gin"
)

func main() {
	__config__ := config.Config__
	database.Open(__config__.DB_FILE)
	server := gin.Default()

	route.BindRoutes(server)

	server.Run(__config__.HttpPort)
}
