package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sauravkarn541/bahikhata/internal/config"
)

func SetupRouter(app *config.Application) *gin.Engine {
	router := gin.New()
	// Setup cors config for api calls
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{app.Env.ClientUrl}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"Authorization"}

	// Update router params
	router.Use(gin.Recovery())
	router.Use(cors.New(corsConfig))

	// Test health checkup of backend
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Backend is up and running",
		})
	})

	// Setup api group
	api := router.Group("api")

	// Register routes
	authRouter := api.Group("/auth")
	RegisterAuthRoutes(app, authRouter)

	userRouter := api.Group("/users")
	RegisterUserRoutes(app, userRouter)

	expenseRouter := api.Group("/expenses")
	RegisterExpenseRoutes(app, expenseRouter)

	return router
}
