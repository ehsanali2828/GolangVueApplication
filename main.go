package main

import (
	"github.com/gin-gonic/gin"
	"github.com/robyklein/go-jwt/controllers"
	"github.com/robyklein/go-jwt/initializers"
	"github.com/robyklein/go-jwt/middleware"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/validate", middleware.RequireAuth, controllers.Validate)
	router.Run()
}
