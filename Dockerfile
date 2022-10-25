FROM golang:1.18-alpine AS builder
WORKDIR /app

COPY ./ ./
RUN go mod download
RUN go build -o weather_api ./cmd/main.go
EXPOSE 8080
CMD [ "/app/weather_api" ]