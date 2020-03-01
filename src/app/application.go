package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	r *gin.Engine
	hostString string
)



func init() {
	// Initialize the route engine
	r  = gin.Default()
}

func StartApplication(host string, port string) {
	// Load the rout maps
	urlMaps()

	// Start the server to listen on given port
	r.Run(getAddrString(host, port))
	//r.Run()
}


// getAddrString generates the custom server string to initialize
// the Gin server instance
func getAddrString(host string, port string) string {
	return fmt.Sprintf("%s:%s", host, port)
}
