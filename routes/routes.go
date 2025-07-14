package routes

import (
	_ "genomic-api/docs"
	"genomic-api/handlers"
	"genomic-api/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Swagger UI available at /swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		// Public route for login
		api.POST("/login", middleware.Login)

		// Protected group
		protected := api.Group("/")
		protected.Use(middleware.JWTAuth())

		// User endpoints
		protected.GET("/users", handlers.ListUsers)
		protected.POST("/users", handlers.CreateUser)
		protected.GET("/users/:id", handlers.GetUser)
		protected.PUT("/users/:id", handlers.UpdateUser)
		protected.DELETE("/users/:id", handlers.DeleteUser)

		// Genome endpoints
		protected.GET("/genomes", handlers.ListGenomes)
		protected.POST("/genomes", handlers.CreateGenome)
		protected.GET("/genomes/:id", handlers.GetGenome)
		protected.PUT("/genomes/:id", handlers.UpdateGenome)
		protected.DELETE("/genomes/:id", handlers.DeleteGenome)

		// Sample endpoints
		protected.GET("/samples", handlers.ListSamples)
		protected.POST("/samples", handlers.CreateSample)
		protected.GET("/samples/:id", handlers.GetSample)
		protected.PUT("/samples/:id", handlers.UpdateSample)
		protected.DELETE("/samples/:id", handlers.DeleteSample)

		// Variant file endpoints
		protected.GET("/variants", handlers.ListVariants)
		protected.POST("/variants", handlers.CreateVariant)
		protected.GET("/samples/:id/variants", handlers.GetSampleVariants)
		protected.DELETE("/variants/:id", handlers.DeleteVariant)
	}

	return r
}
