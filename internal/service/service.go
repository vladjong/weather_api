package service

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"weather_api/config"
	v1 "weather_api/internal/controller/http/v1"
	"weather_api/pkg/server"

	"github.com/sirupsen/logrus"
)

type Service struct {
	cfg *config.Config
	// postgresClient *sqlx.DB
	// handler        v1.Handler
}

func NewService(cfg *config.Config) (service Service, err error) {
	// Добавить свагер
	// Добавить роутер
	// postrgresClient, err := postrgres.NewClient(
	// 	postrgres.PostgresConfig{
	// 		Host:     cfg.PostgresSQL.Host,
	// 		Port:     cfg.PostgresSQL.Port,
	// 		Username: cfg.PostgresSQL.Username,
	// 		Password: cfg.PostgresSQL.Password,
	// 		DBName:   cfg.PostgresSQL.DBName,
	// 		SSLMode:  cfg.PostgresSQL.SSLMode,
	// 	})
	// if err != nil {
	// 	return service, err
	// }
	return Service{
		cfg: cfg,
		// postgresClient: postrgresClient,
	}, nil
}

func (s *Service) Run() error {
	// работа с внешним апи
	// запуск htttp
	s.startHTTP()
	return nil
}

func (s *Service) startHTTP() {
	logrus.Info("HTTP Server initializing")
	server := new(server.Server)
	handlers := new(v1.Handler)
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
	// if err := s.postgresClient.Close(); err != nil {
	// 	logrus.Errorf("Error: occured on db connection close: %s", err.Error())
	// }
}
