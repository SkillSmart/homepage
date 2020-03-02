package config

import (
	"skillsmart/homepage/src/logg"
	"skillsmart/homepage/src/utils/error"
)

// This defines data access methods to the configuration object
func (c *config) SetLogLevel(level string) *error.ConfigErrorResponse {
	//TODO: Set configuration level check and return with error if not met

	c.LogLevel = level
	return nil
}

// Set level with default settings applied
func (c *config) SetLevelWithDefaults(level interface{}) *error.ConfigErrorResponse{
	if logg.IsLogLevel(level) {
		switch level {
		case logg.LevelInfo:
			logg.Config.Info("loglevel set to INFO")
		case logg.LevelDebug:
			logg.Config.Debug("loglevel set to DEBUG")
		}
	}
	return error.Config
}

// HELPER FUNCTIONS ----------------------
