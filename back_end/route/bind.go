package route

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func BindRoutes(router *gin.Engine) {
	// 配置CORS中间件
	config := cors.Config{
		AllowAllOrigins:  true, // 允许所有来源
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false, // 设为false以避免与AllowAllOrigins冲突
		MaxAge:           12 * time.Hour,
	}
	router.Use(cors.New(config))

	userGroup := router.Group("/user")
	{
		userGroup.POST("/signup", Signup())
		userGroup.POST("/signin", Signin())
		userGroup.GET("/get/:uid", GetUser())
		userGroup.GET("/getlist", GetUsers())
		userGroup.POST("/update", UpdateUser())
		userGroup.POST("/delete", DeleteUser())
		userGroup.POST("/jointeam", JoinTeam())
		userGroup.POST("/leaveteam", LeaveTeam())
		userGroup.POST("/updatepassword", UpdateUserPassword())
	}
	teamGroup := router.Group("/team")
	{
		teamGroup.GET("/get/:teamuid", GetTeam())
		teamGroup.GET("/getlist", GetTeams())
		teamGroup.POST("/create", CreateTeam())
		teamGroup.POST("/update", UpdateTeam())
		teamGroup.POST("/delete", DeleteTeam())
		teamGroup.POST("/updatepassword", UpdateTeamPassword())
	}
	itemGroup := router.Group("/item")
	{
		itemGroup.GET("/get/:itemuid", GetItem())
		itemGroup.GET("/getlist", GetItems())
		itemGroup.POST("/create/:teamuid", CreateItem())
		itemGroup.POST("/update/:teamuid", UpdateItem())
		itemGroup.POST("/delete", DeleteItem())
	}
}
