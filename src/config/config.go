package config

import "os"

const (
	LogLevel = "info"
)

type config struct {
	LogLevel string
}

var (
	Config config
)

// NewDevelopment initializes a defaulft config for development purposes
func (*config) NewDevelopment() *config {
	// TODO: Define the default configuration for development
	return &config{
		LogLevel: "debug",
	}
}
// NewProduction intializes a default configuration for production deployments
func (*config) NewProduction() *config {
	// TODO: Define default configuration for production settings
	return &config{
		LogLevel: "info",
	}
}

//////// [ Status checker ]

// IsProduction checks if the system is set to production
func (*config) IsProduction () bool {
	return os.Getenv("GO_ENVIRONMENT") == "production"
}

