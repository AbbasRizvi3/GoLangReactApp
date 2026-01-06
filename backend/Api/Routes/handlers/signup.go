package handlers

import (
	config "github.com/AbbasRizvi3/GoLangReactApp/Config"
	"github.com/gin-gonic/gin"
)

type SignupInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

func SignupHandler(c *gin.Context) {
	var input SignupInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	config.Client.Database("appdb").Collection("users").InsertOne(c, input)
	c.JSON(200, gin.H{"message": "User signed up successfully"})
}
