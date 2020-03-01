package utils

import (
	"fmt"
	"skillsmart/homepage/src/logg"
)
// Public Interface to the config section of the error package
type config struct {}

var (
	ConfigError *config
)

// New generates a new configError
func (c *config) New(message string) *configErrorResponse {
	return &configErrorResponse{
		Message: message,
		Key:     "Define the key",
		Value:   "Define the value set",
		Errors:  []*configError{{
			Message: "Unable to set specified configuration",
			Level:   logg.LevelInfo,
		}},
	}
}

//TODO: Define attributes on configError
type configError struct {
	Message string
	Level int
}

func (c *configError) New(message string) *configError {
	//TODO: return config error from instantiator
	logg.Config.Info("creating new config error value")
	return nil
}

type configErrorResponse struct {
	Message string `json:"message"`
	Key     string `json"key_accessed"`
	Value   string `json:"value_set"`
	Errors  []*configError
}

func (c *configErrorResponse) Error() string {
	return fmt.Sprintf("error tyrying to set %s to %s: %v", c.Key, c.Value, c.Errors)
}
