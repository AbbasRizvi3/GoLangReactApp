package main

import (
	"log"

	routes "github.com/AbbasRizvi3/GoLangReactApp/Api/Routes"
	"github.com/AbbasRizvi3/GoLangReactApp/Api/Routes/cors"
	config "github.com/AbbasRizvi3/GoLangReactApp/Config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var router *gin.Engine

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
func main() {
	config.ConnectDB()
	router = gin.Default()
	router.Use(cors.SetupCors())
	routes.SetupRoutes(router)

	router.Run(":8000")

}
