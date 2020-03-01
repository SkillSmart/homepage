package logg

import (
	"fmt"
	"os"
	"strings"
)

// Defines the interface for logging configuration events in the system
var (
	environmentLogger *envLogger
)

type envLogger struct{}

// Info logs the given message with LogLevelInfo
func (c *envLogger) Info(msg string) {
	// Generate the appropriate error message for the event
	// Log the event
	fmt.Println("ENVIRONMENT LOGGER (INFO):", msg)
}
func (c *envLogger) Debug(msg string) {
	fmt.Println("ENVIRONMENT LOGGER (DEBUG):", msg)
}
func (c *envLogger) Fatal(err error) {
	fmt.Println("ENVIRONMENT LOGGER (FATAL):", err)
}


// Custom Functions
func LogEnvVariables() {
	for _,v := range os.Environ() {
		variable := strings.Split(v, "=")
		fmt.Println(variable)
	}
	return
}


