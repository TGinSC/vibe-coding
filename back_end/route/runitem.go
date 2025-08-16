package route

import (
	"contribution/data"
	"contribution/tool"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

var count = 0

// GetItem 处理获取单个项目信息的HTTP请求
// 该函数返回一个gin.HandlerFunc，用于处理获取项目信息逻辑
func GetItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从URL参数中获取项目ID
		uid := ctx.Param("itemuid")
		if uid == "" {
			ctx.JSON(400, gin.H{"error": "Item UID is required"})
			return
		}

		// 将字符串转换为uint
		var itemUID uint
		_, err := fmt.Sscanf(uid, "%d", &itemUID)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid item UID"})
			return
		}

		// 从数据库获取项目信息
		item, err := data.NewItem().Get(itemUID)
		if err != nil {
			ctx.JSON(404, gin.H{"error": "Item not found"})
			return
		}

		// 返回项目信息
		ctx.JSON(200, gin.H{
			"item": item,
		})
	}
}

// GetItems 处理获取项目列表的HTTP请求
// 该函数返回一个gin.HandlerFunc，用于处理获取项目列表逻辑
func GetItems() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: 实现获取项目列表逻辑
		// 注意：当前数据库层似乎没有提供获取所有项目的接口
		// 这里暂时返回未实现
		ctx.JSON(501, gin.H{"error": "Not implemented"})
	}
}

// CreateItem 处理创建项目的HTTP请求
// 该函数返回一个gin.HandlerFunc，用于处理创建项目逻辑
func CreateItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求中获取项目信息
		item, e := tool.GetItem(ctx)
		if e != nil {
			ctx.JSON(400, gin.H{"error": "Invalid item data"})
			return
		}

		// 创建新项目
		count++
		item.ItemUID = uint(count)
		item.IsComplete = false
		e = data.NewItem().Create(&item)
		if e != nil {
			ctx.JSON(500, gin.H{"error": "Failed to create item"})
			return
		}

		teamuid_string := ctx.Param("teamuid")
		var teamuid uint
		// 将字符串转换为uint
		_, e = fmt.Sscanf(teamuid_string, "%d", &teamuid)
		if e != nil {
			ctx.JSON(400, gin.H{"error": "Invalid team UID"})
			return
		}
		// 将项目添加到团队
		team, e := data.NewTeam().Get(teamuid)
		if e != nil {
			ctx.JSON(404, gin.H{"error": "Team not found"})
			return
		}
		team.ItemsInclude = append(team.ItemsInclude, item.ItemUID)
		startTime := time.Now().Unix()
		_ = data.NewTime().Create(&data.Time{
			ItemUID: item.ItemUID,
			Time:    uint64(startTime),
		})
		e = data.NewTeam().Updata(&team)
		if e != nil {
			ctx.JSON(500, gin.H{"error": "Failed to update team with new item"})
			return
		}

		// 返回响应
		ctx.JSON(200, gin.H{
			"message": "Item created successfully",
			"itemUID": item.ItemUID,
		})
	}
}

// UpdateItem 处理更新项目信息的HTTP请求
// 该函数返回一个gin.HandlerFunc，用于处理更新项目信息逻辑
func UpdateItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 从请求中获取项目信息
		item, e := tool.GetItem(ctx)
		if e != nil {
			ctx.JSON(400, gin.H{"error": "Invalid item data"})
			return
		}

		// 更新项目信息

		teamuid_string := ctx.Param("teamuid")
		var teamuid uint
		// 将字符串转换为uint
		_, e = fmt.Sscanf(teamuid_string, "%d", &teamuid)
		if e != nil {
			ctx.JSON(400, gin.H{"error": "Invalid team UID"})
			return
		}
		team, e := data.NewTeam().Get(teamuid)
		if e != nil {
			ctx.JSON(404, gin.H{"error": "Team not found"})
			return
		}
		// 检查项目是否在团队中
		found := false
		for _, itemUID := range team.ItemsInclude {
			if itemUID == item.ItemUID {
				found = true
				break
			}
		}
		if !found {
			ctx.JSON(400, gin.H{"error": "Item not found in the specified team"})
			return
		}

		// 如果项目标记为完成，更新团队成员的分数
		if item.IsComplete {
			useruid := uint(item.BCB)
			user, e := data.NewUser().Get(uint(useruid))
			if e != nil {
				ctx.JSON(404, gin.H{"error": "User not found"})
				return
			}
			for _, tb := range user.TeamsBelong {
				if tb.TeamUID == teamuid {
					// 更新用户分数
					tb.Score += item.Score
					// 计算完成百分比
					var ShouldBCBcount, BCBcount = 0, 0
					for _, itemUID := range team.ItemsInclude {
						item, _ := data.NewItem().Get(itemUID)
						if useruid == uint(item.ShouldBCB) && !item.IsComplete {
							ShouldBCBcount++
						}
						if useruid == uint(item.BCB) && item.IsComplete {
							BCBcount++
						}
					}
					tb.PercentComplete = uint((float32(BCBcount) / float32(ShouldBCBcount)) * 100)

					NewTeamsBelong := make([]data.TeamBelong, 0)
					for _, tbb := range user.TeamsBelong {
						if tbb.TeamUID != teamuid {
							NewTeamsBelong = append(NewTeamsBelong, tbb)
						}
					}
					NewTeamsBelong = append(NewTeamsBelong, tb)
					user.TeamsBelong = NewTeamsBelong
					break
				}
			}
			// 更新用户信息
			e = data.NewUser().Updata(&user)
			if e != nil {
				ctx.JSON(500, gin.H{"error": "Failed to update user score"})
				return
			}
		}

		e = data.NewItem().Updata(&item)
		if e != nil {
			ctx.JSON(500, gin.H{"error": "Failed to update item"})
			return
		}

		// 返回成功响应
		ctx.JSON(200, gin.H{
			"message": "Item updated successfully",
			"itemUID": item.ItemUID,
		})
	}
}

// DeleteItem 处理删除项目的HTTP请求
// 该函数返回一个gin.HandlerFunc，用于处理删除项目逻辑
func DeleteItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		item, err := tool.GetItem(ctx)
		if err != nil {
			ctx.JSON(400, gin.H{"error": "Item UID is required"})
			return
		}
		itemUID := item.ItemUID

		// 删除项目
		err = data.NewItem().Delete(itemUID)
		if err != nil {
			ctx.JSON(500, gin.H{"error": "Failed to delete item"})
			return
		}

		// 返回成功响应
		ctx.JSON(200, gin.H{
			"message": "Item deleted successfully",
		})
	}
}
