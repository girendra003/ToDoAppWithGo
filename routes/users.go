package routes

import (
	"github.com/girendra003/gin-mongodb-api/controllers"
	"github.com/girendra003/gin-mongodb-api/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	// router.Use(middleware.AuthMiddleware()) // Apply auth middleware to all user routes
	router.GET("/login", controllers.Login())
	router.POST("/user", controllers.CreateUser())
	router.GET("/user/:userId", controllers.GetUser())
	router.PUT("/user/:userId", controllers.UpdateUser())
	router.DELETE("/user/:userId", controllers.DeleteUser())
	router.GET("/users", middleware.AuthMiddleware(),controllers.GetAllUsers())
}

