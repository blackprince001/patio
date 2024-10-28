package services

import (
	"github.com/blackprince001/patio/core/models"
	"github.com/blackprince001/patio/internal/config"
)

// TODO write a parser that could read from yaml or text files to generate the service structs

func Setup() []models.Service {
	// Example service. This was used to reroute requests related to a fileserving to an actual fileserver
	services := []models.Service{
		{
			Name:    "File Service",
			BaseURL: config.NewConfig().ServicePath,
			Routes: []models.Route{
				{Method: "POST", ServicePath: "/upload", TargetPath: "/upload"},
				{Method: "GET", ServicePath: "/download/:shortLink", TargetPath: "/download/:shortLink"},
			},
		},
	}
	return services
}
