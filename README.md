# CalcuCo Calculator API

A calculator API

## Required

1. Go version 1.12
1. docker
1. docker-compose

## Running on local

1. ```docker-compose up mongo```
1. ```GO111MODULE=on go run ./cmd/calculator-server. --port 8080```
1. Open a brower and navigate to `http://localhost:8080/v1/swagger`

## Running docker-compose

1. ```docker-compose up```
1. Open a brower and navigate to `http://localhost:8080/v1/swagger`

### Testing `/calc` endpoint on swagger UI

1. Navigate to `http://localhost:8080/v1/swagger`
1. Click on the `Authorize` button
1. Enter this token `abc`
1. Click `Authorize`
1. Click `Close`
1. Go to `/calc` endpoint]
1. Click the `Try it out` button
1. Enter an `operation` either `add, sub, div, multi, div, sqrt, cbrt, pow or fac` - validated to only using these keywords
1. Enter values to compute in the `args` fields
1. Click `Execute` button to see the result

### Testing `/report` endpoint on swagger UI

1. Navigate to `http://localhost:8080/v1/swagger`
1. Click on the `Authorize` button
1. Enter this token `abc`
1. Click `Authorize`
1. Click `Close`
1. Go to `/report` endpoint]
1. Click the `Try it out` button
1. Enter a `start` e.g. `2018-12-01T00:00:00.000Z`
1. Enter a `end` e.g. `2020-04-23T00:00:00.000Z`
1. Click `Execute` button to see the result

### Running locally

1. `docker-compose up mongo`
1. `GO111MODULE=on go run ./cmd/calculator-server. --port 8080`
1. Navigate to `http://localhost:8080/v1/swagger`

### Running using docker-compose

1. `docker-compose up --build`, the `--build` is set just to force build
1. Navigate to `http://localhost:8080/v1/swagger` - container port is forwarded to localhost's same port `8080`

### Running test

1. `GO111MODULE=on go test ./...`

## Features

- [x] Swagger implementation - `go-swagger`
- [x] Implement middleware for swagger doc (not using ReDoc)
- [x] Handle CORS for swagger
- [x] End point for `/calc`
- [x] End point for `/report`
- [x] Service and unit test for `calc` services
- [x] Service and unit test for `report` services
- [x] Dockerfile to build the app
- [x] Docker-compose to spin up the app and `mongodb`