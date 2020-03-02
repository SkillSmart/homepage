package logg

import "fmt"

// Defines the interface for logging configuration events in the system
var (
	configLogger *cnfLogger
)

type cnfLogger struct{}

// Info logs the given message with LogLevelInfo
func (c *cnfLogger) Info(msg string) {
	// Generate the appropriate error message for the event
	// Log the event
	fmt.Println("CONFIG LOGGER (INFO):", msg)
}
func (c *cnfLogger) Debug(msg string) {
	fmt.Println("CONFIG LOGGER (DEBUG):", msg)
}
func (c *cnfLogger) Fatal(err error) {
	fmt.Println("CONFIG LOGGER (FATAL):", err)
}
