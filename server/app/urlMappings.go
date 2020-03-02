package app

import (
	"github.com/gin-gonic/contrib/static"
	"skillsmart/homepage/src/controllers/page_controllers"
	"skillsmart/homepage/src/controllers/user_controllers"
)

func urlMaps() {

	// This serves our static file on the default route
	r.Use(static.Serve("/", static.LocalFile("./src/views", true)))

	// This groups all routes under "/api/*"
	api := r.Group("/api")
	{
		// Handle Index
		api.GET("/", page_controllers.GetIndexPage)
	}

	usersRoutes()
}

func usersRoutes() {
	users := r.Group("/users")
	{
		users.GET("/", user_controllers.GetAllUsers)
	}
}




