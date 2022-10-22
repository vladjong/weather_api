package postgressql

import "weather_api/internal/entities"

type WeatherStorage interface {
	CreateCity(city entities.City) (id int, err error)
	CreateWeather(weather entities.Weather, cityId int, info []byte) (id int, err error)
}
