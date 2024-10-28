# syntax=docker/dockerfile:1

# Build the application from source
FROM golang:1.22.2 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o patio

# Run the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...

# Deploy the application binary into a lean image
FROM ubuntu:latest AS build-release-stage

RUN useradd -m nonroot

USER nonroot

WORKDIR /server

COPY --from=build-stage /app/patio /server/

EXPOSE 5455

CMD ["/server/patio"]
