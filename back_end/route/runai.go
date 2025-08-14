package route

import (
	"github.com/gin-gonic/gin"
	"contribution/back_end/ai"
)

func RunAI(r *gin.Engine) {
	aiGroup := r.Group("/ai")
	{
		aiGroup.POST("/assist", ai.AIHandler)
	}
}