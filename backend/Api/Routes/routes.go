package routes

import (
	"github.com/AbbasRizvi3/GoLangReactApp/Api/Routes/handlers"
	"github.com/AbbasRizvi3/GoLangReactApp/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/Dashboard", middleware.AuthMiddleware(), handlers.DashboardHandler)
	router.POST("/Signup", handlers.SignupHandler)
	router.POST("/Login", handlers.LoginHandler)
	router.POST("/Logout", middleware.AuthMiddleware(), handlers.LogoutHandler)
}
