package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Listen struct {
		Port string `env:"PORT" env-default:":8080"`
	}
	WeatherAPI struct {
		Limit int    `env:"LIMIT" env-default:"1"`
		Units string `env:"UNITS" env-default:"metric"`
	}
	PostgresSQL struct {
		Host     string `env:"HOST" env-default:"db"`
		Port     string `env:"PORT" env-default:"5432"`
		Username string `env:"USERNAME" env-default:"postgres"`
		DBName   string `env:"DBNAME" env-default:"postgres"`
		SSLMode  string `env:"SSLMODE" env-default:"disable"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		logrus.Print("Read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "Weather_api - micro_server"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			logrus.Print(help)
			logrus.Fatal(err)
		}
	})
	return instance
}
