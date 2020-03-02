package logg

import "fmt"

// Defines the interface for logging configuration events in the system
var (
	applicationLogger *appLogger
)

type appLogger struct{}

// Info logs the given message with LogLevelInfo
func (c *appLogger) Info(msg string) {
	// Generate the appropriate error message for the event
	// Log the event

}
func (c *appLogger) Debug(msg string) {
	fmt.Println("CONFIG LOGGER (DEBUG):", msg)
}
func (c *appLogger) Fatal(err error ) {
	fmt.Println("CONFIG LOGGER (FATAL):", err)
}
