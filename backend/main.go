package main

import (
	"log"

	config "github.com/AbbasRizvi3/GoLangReactApp/Config"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
}
func main() {
	config.ConnectDB()
}
