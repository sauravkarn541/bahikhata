package main

import (
	"github.com/sauravkarn541/bahikhata/internal/config"
	"github.com/sauravkarn541/bahikhata/internal/routes"
)

var log = config.GetLogger()

func main() {
	// Initialize the logger
	config.InitializeLogger()

	app := config.App()
	config.InitJWTParams(app.Env)

	env := app.Env

	router := routes.SetupRouter(&app)
	router.Run(env.AppHost + ":" + env.AppPort)
}
