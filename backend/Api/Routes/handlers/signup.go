package handlers

import (
	"net/http"

	config "github.com/AbbasRizvi3/GoLangReactApp/Config"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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
	var existingUser bson.M
	err := config.Client.Database("appdb").Collection("users").
		FindOne(c, bson.M{"email": input.Email}).Decode(&existingUser)

	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
		return
	}
	config.Client.Database("appdb").Collection("users").InsertOne(c, input)
	c.JSON(200, gin.H{"message": "User signed up successfully"})
}
