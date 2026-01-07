package handlers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	config "github.com/AbbasRizvi3/GoLangReactApp/Config"
	models "github.com/AbbasRizvi3/GoLangReactApp/Models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
)

var key = os.Getenv("JWT_SECRET_KEY")

var JwtKey = []byte(key)

func LoginHandler(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user bson.M
	err := config.Client.Database("appdb").Collection("users").
		FindOne(c, bson.M{"email": input.Email, "password": input.Password}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

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
	c.JSON(http.StatusOK, gin.H{
		"message": "User logged in successfully",
	})
}
