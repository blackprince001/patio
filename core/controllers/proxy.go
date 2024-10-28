package controllers

import (
	"context"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/blackprince001/patio/core/models"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, services []models.Service) {
	router.GET("/health", HealthCheck)

	api := router.Group("/api")

	for _, service := range services {
		for _, route := range service.Routes {
			fullPath := "/" + service.Name + route.TargetPath
			api.Handle(route.Method, fullPath, createProxyHandler(service, route))
		}
	}
}

func createProxyHandler(service models.Service, route models.Route) gin.HandlerFunc {
	return func(c *gin.Context) {
		path := route.ServicePath
		for _, param := range c.Params {
			path = strings.Replace(path, ":"+param.Key, param.Value, -1)
		}
		proxyRequest(c, service.BaseURL, path)
	}
}

func proxyRequest(c *gin.Context, base string, newPath string) {
	target, err := url.Parse(base + newPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid target URL: " + err.Error()})
		return
	}

	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.Transport = &http.Transport{
		ResponseHeaderTimeout: 10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	originalDirector := proxy.Director
	proxy.Director = func(req *http.Request) {
		originalDirector(req)
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Host = target.Host
		req.URL.Path = base + newPath
	}

	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		statusCode := http.StatusBadGateway
		errorMessage := "Proxy error"

		if err == context.DeadlineExceeded {
			statusCode = http.StatusGatewayTimeout
			errorMessage = "Timeout error"
		} else if strings.Contains(err.Error(), "connection refused") {
			errorMessage = "Service unavailable"
		}

		c.JSON(statusCode, gin.H{"error": errorMessage, "details": err.Error()})
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{Status: "OK"})
}
