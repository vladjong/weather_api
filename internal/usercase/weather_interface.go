package usercase

import (
	"weather_api/internal/entities"
)

type WeatherService interface {
	CreateCity(city entities.City) (id int, err error)
	CreateWeather(weather entities.Weather, cityId int, info []byte) (id int, err error)
}

type WeatherAPI interface {
	GetCities() (names []string, err error)
	GetWeatherInCity(name string) (weathers []entities.WeatherDetails, err error)
	GetDetaiWeatherInCity(name string, date string) (weathers []entities.WeatherPredict, err error)
}
