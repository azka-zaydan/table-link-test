package main

//go:generate go run github.com/swaggo/swag/cmd/swag init

import (
	"github.com/azka-zaydan/table-link-test/configs"
	"github.com/azka-zaydan/table-link-test/shared/logger"
)

var config *configs.Config

func main() {
	// app := fiber.New()
	// Initialize logger
	logger.InitLogger()

	// Initialize config
	config = configs.Get()

	// Set desired log level
	logger.SetLogLevel(config)

	// Wire everything up
	// http := InitializeService()

	// // consumers := InitializeEvent()

	// // // Start consumers
	// // consumers.Start()

	// // Run server
	// http.SetupAndServe(app)

}
