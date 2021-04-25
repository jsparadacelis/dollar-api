package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Gretting(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

//DollarController resolves requests about dollar price.
func DollarController(c *gin.Context) {

	response, ok := GetDollarValue()
	if ok {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.String(http.StatusOK, response)
	}
}
