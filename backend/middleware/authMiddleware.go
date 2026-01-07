package middleware

import (
	"fmt"
	"net/http"

	"github.com/AbbasRizvi3/GoLangReactApp/Api/Routes/handlers"
	models "github.com/AbbasRizvi3/GoLangReactApp/Models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "No token found"})
			c.Abort()
			return
		}
		claims := &models.Claims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return handlers.JwtKey, nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		fmt.Println("Token from cookie:", tokenString)

		c.Set("email", claims.Email)
		c.Next()
	}
}
