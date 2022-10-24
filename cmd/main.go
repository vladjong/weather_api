package main

import (
	"log"
	"weather_api/config"
	"weather_api/internal/service"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Config initializing")
	cfg := config.GetConfig()
	logrus.Info("Env variables initializing")
	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}
	logrus.Info("Running service")
	service, err := service.NewService(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	if err := service.Run(); err != nil {
		logrus.Fatal(err)
	}
}
