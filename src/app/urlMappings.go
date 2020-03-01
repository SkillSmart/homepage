package app

import "skillsmart/homepage/src/controllers/page_controllers"

func urlMaps() {

	// Handle Index
	r.GET("", page_controllers.GetIndexPage)
	r.GET("/user", page_controllers.GetProjectPage)
}
