package routes

import (
	"github.com/AbbasRizvi3/GoLangReactApp/Api/Routes/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.GET("/Dashboard", handlers.DashboardHandler)
}
