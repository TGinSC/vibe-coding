package ai

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AIRequest struct {
	Prompt string `json:"prompt"`
}

func AIHandler(c *gin.Context) {
	var req AIRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := CallHuggingFaceAPI(req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"response": response})
}