package routes

import (
	_ "genomic-api/docs"
	"genomic-api/handlers"
	"genomic-api/middleware"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Prometheus metrics
var (
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total HTTP requests",
		},
		[]string{"path", "method", "status"},
	)

	httpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "HTTP request duration in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method"},
	)
)

func init() {
	// register metrics
	prometheus.MustRegister(httpRequestsTotal, httpRequestDuration)
}

// Middleware for observability
func ObservabilityMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next() // process request

		duration := time.Since(start).Seconds()
		path := c.FullPath()
		method := c.Request.Method
		status := c.Writer.Status()

		// update metrics
		httpRequestsTotal.WithLabelValues(path, method, http.StatusText(status)).Inc()
		httpRequestDuration.WithLabelValues(path, method).Observe(duration)

		// structured log
		log.Info().
			Str("path", path).
			Str("method", method).
			Int("status", status).
			Float64("duration_s", duration).
			Str("remote", c.ClientIP()).
			Msg("http request")
	}
}

// SetupRouter wires up all routes, middlewares, and observability
func SetupRouter() *gin.Engine {
	// Create Gin router with recovery and logging disabled (we’ll add observability middleware instead)
	r := gin.New()
	r.Use(gin.Recovery(), ObservabilityMiddleware())

	// ---- Swagger ----
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// ---- Prometheus metrics ----
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// ---- API Routes ----
	api := r.Group("/api")
	{
		// Public endpoints
		api.POST("/login", middleware.Login)

		// Protected group with JWT
		protected := api.Group("/")
		protected.Use(middleware.JWTAuth())
		{
			// Users
			protected.GET("/users", handlers.ListUsers)
			protected.POST("/users", handlers.CreateUser)
			protected.GET("/users/:id", handlers.GetUser)
			protected.PUT("/users/:id", handlers.UpdateUser)
			protected.DELETE("/users/:id", handlers.DeleteUser)

			// Genomes
			protected.GET("/genomes", handlers.ListGenomes)
			protected.POST("/genomes", handlers.CreateGenome)
			protected.GET("/genomes/:id", handlers.GetGenome)
			protected.PUT("/genomes/:id", handlers.UpdateGenome)
			protected.DELETE("/genomes/:id", handlers.DeleteGenome)

			// Samples
			protected.GET("/samples", handlers.ListSamples)
			protected.POST("/samples", handlers.CreateSample)
			protected.GET("/samples/:id", handlers.GetSample)
			protected.PUT("/samples/:id", handlers.UpdateSample)
			protected.DELETE("/samples/:id", handlers.DeleteSample)

			// Sequences
			protected.GET("/sequence", handlers.ListSequenceFiles)
			protected.POST("/sequence", handlers.CreateSequenceFile)
			protected.GET("/sequence/:id", handlers.GetSequenceFile)
			protected.PUT("/sequence/:id", handlers.UpdateSequenceFile)
			protected.DELETE("/sequence/:id", handlers.DeleteSequenceFile)

			// Variants
			protected.GET("/variants", handlers.ListVariants)
			protected.POST("/variants", handlers.CreateVariant)
			protected.GET("/samples/:id/variants", handlers.GetSampleVariants)
			protected.DELETE("/variants/:id", handlers.DeleteVariant)
		}
	}

	return r
}
