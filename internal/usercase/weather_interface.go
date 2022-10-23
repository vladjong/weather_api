package usercase

import (
	"time"
	"weather_api/internal/entities"
)

type WeatherService interface {
	CreateCity(city entities.City) (id int, err error)
	CreateWeather(weather entities.Weather, cityId int, info []byte) (id int, err error)
}

type WeatherAPI interface {
	GetCities() (names entities.AllCities, err error)
	GetWeatherInCity(name string) (entities.WeatherPredict, error)
	GetDetaiWeatherInCity(name string, date time.Time) (entities.WeatherDetails, error)
}
