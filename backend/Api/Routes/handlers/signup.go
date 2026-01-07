package handlers

import (
	"fmt"
	"net/http"
	"time"

	config "github.com/AbbasRizvi3/GoLangReactApp/Config"
	models "github.com/AbbasRizvi3/GoLangReactApp/Models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
)

func SignupHandler(c *gin.Context) {
	var input models.SignupInput
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

	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &models.Claims{
		Email: input.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}
	fmt.Println("Token issued:", tokenString)

	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)

	c.JSON(200, gin.H{"message": "User signed up successfully"})
}
