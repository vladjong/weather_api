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
			Password: os.Getenv("DB_PASSWORD"),
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
	logrus.Info("Initializing openWeatherApi service storage interface")
	s.connectExternalService()
	s.startHTTP()
	return nil
}

func (s *Service) connectExternalService() {
	logrus.Info("Initializing openWeatherApi")
	openWeatherApi := openweather.NewOpenWeatherApi(s.cfg.WeatherAPI.Limit, os.Getenv("API_KEY"), s.cfg.WeatherAPI.Units)
	logrus.Info("Initializing service storage interface")
	postgres := postgressql.NewWeatherServiceStorage(s.postgresClient)
	logrus.Info("Initializing openWeatherApi service use case")
	useCase := usercase.NewWeatherServiceUseCase(postgres, openWeatherApi)
	logrus.Info("Adding cities in db")
	weatherService := externalservice.NewOpenWeatherApi(useCase)
	if err := weatherService.CreateCities(); err != nil {
		logrus.Error(err)
	}
	logrus.Info("Adding weathers in db")
	if err := weatherService.CreateWeathers(); err != nil {
		logrus.Error(err)
	}
}

func (s *Service) startHTTP() {
	logrus.Info("HTTP Server initializing")
	server := new(server.Server)
	logrus.Info("Initializing service storage interface")
	postgres := postgressql.NewWeatherServiceStorage(s.postgresClient)
	logrus.Info("Initializing weather api use case")
	useCase := usercase.NewWeatherApiUseCase(postgres)
	logrus.Info("Initializing handlers")
	handlers := v1.NewHandler(useCase)
	go func() {
		if err := server.Run(s.cfg.Listen.Port, handlers.NewRouter()); err != nil {
			logrus.Fatalf("error: occured while running HTTP Server: %s", err.Error)
		}
	}()
	logrus.Info("HTTP Server start")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Info("HTTP Server Shutdown")
	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error: occured on server shutdown: %s", err.Error())
	}
	if err := s.postgresClient.Close(); err != nil {
		logrus.Errorf("error: occured on db connection close: %s", err.Error())
	}
}
