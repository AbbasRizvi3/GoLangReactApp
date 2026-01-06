package handlers

import "github.com/gin-gonic/gin"

func LogoutHandler(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.JSON(200, gin.H{"msg": "Logged out successfully"})

}
