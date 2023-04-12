package controllers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Pong(c *gin.Context) {
	param := c.Param("pingORPong")
	if param == "" {
		c.JSON(200, gin.H{
			"response": "param empty",
		})
	}
	if param == "ping" {
		c.JSON(200, gin.H{
			"response": param,
		})
	} else if param == "pong" {
		c.JSON(200, gin.H{
			"response": os.Getenv("PING"),
		})
	}
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"response": os.Getenv("PONG"),
	})
}
