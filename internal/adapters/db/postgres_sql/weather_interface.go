package postgressql

import (
	"time"
	"weather_api/internal/entities"
)

type WeatherStorage interface {
	CreateCity(city entities.City) (id int, err error)
	CreateWeather(weather entities.Weather, cityId int, info []byte) (id int, err error)
	GetCities() (names []string, err error)
	GetWeatherInCity(name string) (weather []entities.WeatherPredict, err error)
	GetDetaiWeatherInCity(name string, date time.Time) (weather []entities.WeatherDetails, err error)
	GetDatesInCity(name string) (dates []time.Time, err error)
}
