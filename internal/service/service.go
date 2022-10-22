package service

import (
	"weather_api/config"
	"weather_api/pkg/client/postrgres"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	cfg            *config.Config
	postgresClient *sqlx.DB
}

func NewService(cfg *config.Config) (service Service, err error) {
	// Добавить свагер
	// Добавить роутер
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
	// работа с внешним апи
	// запуск htttp
	return nil
}
