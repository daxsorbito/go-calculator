# build stage
FROM golang:1.12 as builder

ENV GO111MODULE=on

COPY . /go/src/go-calculator
WORKDIR /go/src/go-calculator

RUN go mod download

RUN go build -o ./calc-api ./cmd/calculator-server/. 

# final stage
FROM golang:1.12

COPY --from=builder /go/src/go-calculator/calc-api /opt/go-calc/calc-api
COPY --from=builder /go/src/go-calculator/.env /opt/go-calc/.env

WORKDIR /opt/go-calc
CMD [ "/opt/go-calc/calc-api", "--port", "8080" ]
EXPOSE 8080