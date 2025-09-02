package main

import (
	"os"
	"github.com/wkloose/Glowlytics/initializers"
	"github.com/wkloose/Glowlytics/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.ConnectToDB()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	router.Use(middleware.RequireAuth)
	Routes(router)

	
	router.Run(":" + os.Getenv("PORT"))
}
