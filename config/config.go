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
		Limit int    `env:"LIMIT" env-default:"1"`
		Units string `env:"UNITS" env-default:"metric"`
	}
	PostgresSQL struct {
		Host     string `env:"HOST" env-default:"localhost"`
		Port     string `env:"PORT" env-default:"5436"`
		Username string `env:"USERNAME" env-default:"postgres"`
		DBName   string `env:"DBNAME" env-default:"postgres"`
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
