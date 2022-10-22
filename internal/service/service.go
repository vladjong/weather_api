package service

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"weather_api/config"
	postgressql "weather_api/internal/adapters/db/postgres_sql"
	externalservice "weather_api/internal/controller/external_service"
	v1 "weather_api/internal/controller/http/v1"
	"weather_api/internal/usercase"
	"weather_api/pkg/client/postrgres"
	"weather_api/pkg/server"
	openweather "weather_api/pkg/weather_service_client/open_weather"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type Service struct {
	cfg            *config.Config
	postgresClient *sqlx.DB
}

func NewService(cfg *config.Config) (service Service, err error) {
	// Добавить свагер
	postrgresClient, err := postrgres.NewClient(
		postrgres.PostgresConfig{
			Host:     cfg.PostgresSQL.Host,
			Port:     cfg.PostgresSQL.Port,
			Username: cfg.PostgresSQL.Username,
			Password: cfg.PostgresSQL.Password,
			DBName:   cfg.PostgresSQL.DBName,
			SSLMode:  cfg.PostgresSQL.SSLMode,
		})
	if err != nil {
		return service, err
	}
	return Service{
		cfg:            cfg,
		postgresClient: postrgresClient,
	}, nil
}

func (s *Service) Run() error {
	s.connectExternalService()
	// запуск htttp
	s.startHTTP()
	return nil
}

func (s *Service) connectExternalService() {
	openWeatherApi := openweather.NewOpenWeatherApi(s.cfg.WeatherAPI.Limit, s.cfg.WeatherAPI.Key, s.cfg.WeatherAPI.Units)
	postgres := postgressql.NewWeatherServiceStorage(s.postgresClient)
	useCase := usercase.NewWeatherServiceUseCase(postgres)
	weatherService := externalservice.NewOpenWeatherApi(openWeatherApi, useCase)
	if err := weatherService.CreateCities(); err != nil {
		logrus.Fatal(err)
	}
	weatherService.CreateWeathers()
}

func (s *Service) startHTTP() {
	logrus.Info("HTTP Server initializing")
	server := new(server.Server)
	postgres := postgressql.NewWeatherServiceStorage(s.postgresClient)
	useCase := usercase.NewWeatherApiUseCase(postgres)
	handlers := v1.NewHandler(useCase)
	go func() {
		if err := server.Run(s.cfg.Listen.Port, handlers.NewRouter()); err != nil {
			logrus.Fatalf("Error: occured while running HTTP Server: %s", err.Error)
		}
	}()
	logrus.Info("HTTP Server start")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Info("HTTP Server Shutdown")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error: occured on server shutdown: %s", err.Error())
	}
	if err := s.postgresClient.Close(); err != nil {
		logrus.Errorf("Error: occured on db connection close: %s", err.Error())
	}
}
