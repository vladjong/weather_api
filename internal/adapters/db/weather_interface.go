package db

import (
	"time"
	"weather_api/internal/entities"
)

type WeatherStorage interface {
	CreateCity(city entities.City) (id int, err error)
	CreateWeather(weather entities.WeatherCreate) (id int, err error)
}

type WeatherStorageAPI interface {
	GetCities() (names entities.AllCities, err error)
	GetWeatherInCity(name string) (weather []entities.WeatherPredict, err error)
	GetDetaiWeatherInCity(name string, date time.Time) (weather []entities.WeatherDetails, err error)
	GetDatesInCity(name string) (dates []time.Time, err error)
}
