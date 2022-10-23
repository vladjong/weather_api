package usercase

import (
	"time"
	"weather_api/internal/entities"
)

type WeatherService interface {
	CreateCities() error
	CreateWeathers() error
}

type WeatherAPI interface {
	GetCities() (names entities.AllCities, err error)
	GetWeatherInCity(name string) (entities.WeatherPredict, error)
	GetDetaiWeatherInCity(name string, date time.Time) (entities.WeatherDetails, error)
}
