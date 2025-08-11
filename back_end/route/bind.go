package route

import "github.com/gin-gonic/gin"

func BindRoutes(router *gin.Engine) {
	userGroup := router.Group("/user")
	{
		userGroup.POST("/signup", Signup())
		userGroup.POST("/signin", Signin())
		userGroup.GET("/get", GetUser())
		userGroup.GET("/getlist", GetUsers())
		userGroup.POST("/update", UpdateUser())
		userGroup.POST("/delete", DeleteUser())
	}
	teamGroup := router.Group("/team")
	{
		teamGroup.GET("/get", GetTeam())
		teamGroup.GET("/getlist", GetTeams())
		teamGroup.POST("/create", CreateTeam())
		teamGroup.POST("/update", UpdateTeam())
		teamGroup.POST("/delete", DeleteTeam())
	}
	itemGroup := router.Group("/item")
	{
		itemGroup.GET("/get", GetItem())
		itemGroup.GET("/getlist", GetItems())
		itemGroup.POST("/create", CreateItem())
		itemGroup.POST("/update", UpdateItem())
		itemGroup.POST("/delete", DeleteItem())
	}
}
