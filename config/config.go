package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Listen struct {
		Port string `env:"PORT" env-default:":8080"`
	}
	WeatherAPI struct {
		Key string `env:"KEY" env-default:"8f7a589bb238eb36173737ebbe1ec8c6"`
	}
	PostgresSQL struct {
		Host     string `env:"HOST" env-default:"0.0.0.0"`
		Port     string `env:"PORT" env-default:"5432"`
		Username string `env:"USERNAME" env-default:"postgres"`
		Password string `env:"AUTH_DB" env-default:"postgres"`
		DBName   string `env:"DBNAME" env-default:"weather_api"`
		SSLMode  string `env:"SSLMODE" env-default:"disable"`
	}
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		log.Print("Read application configuration")
		instance = &Config{}
		if err := cleanenv.ReadEnv(instance); err != nil {
			helpText := "Weather_api - micro_server"
			help, _ := cleanenv.GetDescription(instance, &helpText)
			log.Print(help)
			log.Fatal(err)
		}
	})
	return instance
}
