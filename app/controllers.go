package main

import "github.com/gin-gonic/gin"

func Gretting(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong1",
	})
}

//DollarController resolves requests about dollar price.
func DollarController(c *gin.Context) {
	response, ok := GetDollarValue()

	if ok {
		c.JSON(200, gin.H{
			"message": response,
		})
	}
}
