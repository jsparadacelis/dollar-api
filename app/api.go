package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Hola mundo")

	api := gin.Default()
	api.GET("/ping", Gretting)
	api.GET("/dollar", DollarController)
	api.Run()
}
