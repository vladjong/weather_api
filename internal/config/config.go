package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug       bool `env:"IS_DEBUG" env-default:"false"`
	IsDevelopment bool `env:"IS_DEV" env-default:"false"`
	Listen        struct {
		Type       string `env:"LISTEN_TYPE" env-default:"port"`
		BindIP     string `env:"IS_DEV" env-default:"0.0.0.0"`
		Port       string `env:"PORT" env-default:"8080"`
		SocketFile string `env:"SOCKET_FILE" env-default:"app.sock"`
	}
	WeatherAPI struct {
		Key string `env:"KEY" env-default:"8f7a589bb238eb36173737ebbe1ec8c6"`
	}
	PostgresDB struct {
		Host     string `env:"HOST" env-default:"localhost"`
		Port     string `env:"PORT" env-default:"5432"`
		Username string `env:"USERNAME" env-default:""`
		Password string `env:"AUTH_DB" env-default:""`
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
