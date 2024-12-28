# Patio - Gateway to your Microservices

## Project Overview

### Purpose

A lightweight, configurable reverse proxy gateway built using Golang. Designed to act as an intermediary between clients and one or more origin servers or microservices, providing essential features such as rate limiting and request/response modification.

### Directory Structure

The directory structure is well-organized and follows standard Go project conventions:

```bash
└── blackprince001-patio/
    ├── main.go
    ├── provisioning/
    │   └── datasources/
    │       └── gateway.yaml
    ├── go.sum
    ├── internal/
    │   ├── metrics/
    │   │   └── metrics.go
    │   ├── config/
    │   │   └── config.go
    │   └── middleware/
    │       ├── metrics.go
    │       └── ratelimiter.go
    ├── core/
    │   ├── controllers/
    │   │   └── proxy.go
    │   ├── services/
    │   │   └── services.go
    │   └── models/
    │       ├── service.go
    │       └── route.go
    ├── .dockerignore
    ├── prometheus.yml
    ├── Dockerfile
    ├── go.mod
    ├── README.md
    └── docker-compose.yml
```

## Setup and Run the Project

### Prerequisites

- Go 1.22.2 or later
- Docker and Docker Compose

### Steps to Run

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/blackprince001/patio.git
   cd patio
   ```

2. **Build and Run Using Docker Compose:**

   ```bash
   docker-compose up
   ```

   This command will build the Docker image and start the containers defined in the `docker-compose.yml` file.

3. **Health Check:**
   You can verify the health of the gateway by accessing the health check endpoint:

   ```bash
   curl -f http://localhost:<PORT>/health
   ```

   Replace `<PORT>` with the port specified in your configuration.

4. **Metrics:**
   Metrics are exposed at the `/metrics` endpoint:

   ```bash
   curl http://localhost:<PORT>/metrics
   ```

5. **Rate Limiting:**
   The rate limiting is configured to allow 5 requests per second with a burst of 10. You can test this by making multiple requests to any endpoint.

### Rate Limiting

Test rate limiting by making multiple requests to any endpoint.

## Configuration

The application configuration is read from environment variables. You can set `SERVICEPATH`, `PORT`, and `ENV` variables to customize the setup.

## Contributing

Contributions are welcome. Please submit a pull request with your changes.
