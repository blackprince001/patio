package main

import (
	"log"

	"github.com/blackprince001/patio/core/controllers"
	"github.com/blackprince001/patio/core/services"
	"github.com/blackprince001/patio/internal/config"
	"github.com/blackprince001/patio/internal/middleware"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/time/rate"
)

func main() {
	cfg := config.NewConfig()

	router := gin.Default()

	router.Use(middleware.MetricsMiddleware())
	router.Use(gin.Recovery()) // Gin's built-in recovery middleware

	// Prometheus metrics endpoint
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// Create a rate limiter: 5 requests per second with burst of 10
	rateLimiter := middleware.NewRateLimiter(rate.Limit(5), 10)

	// Apply rate limiting middleware
	router.Use(rateLimiter.RateLimit())

	services := services.Setup()
	controllers.SetupRoutes(router, services)

	log.Printf("Starting server on :%s", cfg.Port)
	router.Run(":" + cfg.Port)
}
