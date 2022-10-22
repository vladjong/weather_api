package main

import (
	"weather_api/config"
	"weather_api/internal/service"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Config initializing")
	cfg := config.GetConfig()

	logrus.Info("Running service")
	service, err := service.NewService(cfg)
	if err != nil {
		logrus.Fatal(err)
	}
	service.Run()
}
