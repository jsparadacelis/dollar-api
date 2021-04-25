package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Log error if .env file does not exist
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {

	api := gin.Default()
	api.GET("/ping", Gretting)
	api.GET("/dollar", DollarController)
	api.Run()
}
