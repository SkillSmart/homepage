package page_controllers

import "github.com/gin-gonic/gin"

func GetIndexPage(c *gin.Context) {
	c.String(200, "Welcome to the show")
	return
}

func GetProjectPage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":"pong",
	})
}
