FROM golang:1.24.0-alpine3.21 AS builder

RUN apk update && apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download
COPY . .

RUN go build -o /currency-converter cmd/main.go

FROM alpine:latest
COPY --from=builder /currency-converter /currency-converter
EXPOSE 8001
CMD ["/currency-converter"]
