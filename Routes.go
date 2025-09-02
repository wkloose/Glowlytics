package main

import (
	"github.com/gin-gonic/gin"
	"github.com/wkloose/Glowlytics/controllers"
)

func Routes(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.POST("/ai/recommendation", controllers.GetAIRecommendation)
	}
}
