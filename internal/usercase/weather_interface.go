package usercase

import (
	"time"
	"weather_api/internal/entities"
)

//go:generate mockgen -source=weather_interface.go -destination=moks/mock.go

type WeatherService interface {
	CreateCities() error
	CreateWeathers() error
}

type WeatherAPI interface {
	GetCities() (names entities.AllCities, err error)
	GetWeatherInCity(name string) (entities.WeatherPredict, error)
	GetDetaiWeatherInCity(name string, date time.Time) (entities.WeatherDetails, error)
}
