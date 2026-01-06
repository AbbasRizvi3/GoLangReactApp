package handlers

import "github.com/gin-gonic/gin"

func DashboardHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": []string{"Welcome to the Dashboard", "User Stats", "System Overview"},
	})
}
