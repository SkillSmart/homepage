package user_controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"skillsmart/homepage/src/logg"
	"skillsmart/homepage/src/services/user_services"
)

var(
	Member *member
)

type member struct {}


// USER ROUTES (SINGLE USER INSTANCE) -------------
func GetUserById(c *gin.Context) {
	userId := c.GetInt64("user_id")
	user, err := user_services.GetUserByID(userId)
	if err != nil {
		logg.Application.Info(err.Error())
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusFound, user)
}

// Retrieve a given user by email
func GetUserByEmail(c *gin.Context) {
	email := c.GetString("email")
	user, err := user_services.GetUserByEmail(email)
	if err != nil {
		c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusFound, user)
}


// USERS ROUTES (LIST OF USER) ------------------

func GetAllUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":"NOT IMPLEMENTED YET",
	})
	c.String(200, "Welcome to the show")
	return
}

func GetAllUsersByRegion(c *gin.Context) {
	c.JSON(200, gin.H{
		"message":"NOT IMPLEMENTED YET",
	})
}

