package main

import "github.com/gin-gonic/gin"

func Gretting(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong1",
	})
}

func DollarController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "dollar value",
	})
}
