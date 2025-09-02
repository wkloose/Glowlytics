package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/wkloose/Glowlytics/services"
)

type AIRequest struct {
	Prompt string `json:"prompt"`
}

func GetAIRecommendation(c *gin.Context) {
	var req AIRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := services.CallGemini(req.Prompt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ai_result": result})
}
