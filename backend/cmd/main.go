package main

import (
	"github.com/sauravkarn541/bahikhata/internal/config"
)

var log = config.GetLogger()

func main() {
	// Initialize the logger
	config.InitializeLogger()

	app := config.App()
	config.InitJWTParams(app.Env)

	env := app.Env
	log.Info("Starting server on port " + env.AppPort)

	// Initialize the router
	// r := gin.Default()

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.Run() // Default listens on :8080
}
