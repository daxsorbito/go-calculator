# CalcuCo Calculator API

A calculator API

## Running on local

1. ```docker-compose up mongo```
2. ```GO111MODULE=on go run ./cmd/calculator-server. --port 8080```
3. Open a brower and navigate to `http://localhost:8080/v1/swagger`

## Running docker-compose

1. ```docker-compose up```
2. Open a brower and navigate to `http://localhost:8080/v1/swagger`

### Features:

- [x] Swagger implementation - `go-swagger`
- [x] Implement middleware for swagger doc (not using ReDoc)
- [x] End point for `/calc`
- [x] End point for `/report`
- [x] Service and unit test for `calc` services
- [x] Dockerfile to build the app
- [x] Docker-compose to spin up the app and `mongodb`
- [ ] Service and unit test for `report` services
- [ ] API test
- [ ] CI/CD - maybe using `travis CI`
