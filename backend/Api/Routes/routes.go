package routes

import (
	"github.com/AbbasRizvi3/GoLangReactApp/Api/Routes/handlers"
	"github.com/AbbasRizvi3/GoLangReactApp/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.Use(middleware.AuthMiddleware())
	router.GET("/Dashboard", handlers.DashboardHandler)
	router.POST("/Signup", handlers.SignupHandler)
	router.POST("/Login", handlers.LoginHandler)
}
