package main

import (
	"os"
	"skillsmart/homepage/src/app"
	"skillsmart/homepage/src/logg"
)


const (
	//TODO: Implemnt variables from os.env
	hostName = "homepage_host_name"
	hostPort = "homepage_host_port"
)

var (
	// Declare all relevant variables to be pulled from Environment
	hpHost string
	hpPort string
)

func init() {
	// Initialize Environment variables
	//TODO: Implement the passing of env variables
	logg.Config.Info("attempting to set server with env vars")
	logg.LogEnvVariables()
	hpHost = os.Getenv(hostName)
	hpPort = os.Getenv(hostPort)
	logg.Application.Info("testing the application info logging")
	//logg.Application.Fatal(utils.ConfigError.New("what a mess!!!!"))

}

func main() {

	// Start the application
	app.StartApplication(hpHost, hpPort)
}