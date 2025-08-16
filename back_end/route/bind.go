package route

import (
	"github.com/gin-gonic/gin"
)

func BindRoutes(router *gin.Engine) {
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
		itemGroup.POST("/deltatime", GetDeltaTime())
	}
	aiGroup := router.Group("/ai")
	{
		aiGroup.POST("/assist", AIHandler())
	}
}
