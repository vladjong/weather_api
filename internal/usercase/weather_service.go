package usercase

import (
	postgressql "weather_api/internal/adapters/db/postgres_sql"
	"weather_api/internal/entities"
)

type WeatherService interface {
	CreateCity(city entities.City) (id int, err error)
	CreateWeather(weather entities.Weather, cityId int, info []byte) (id int, err error)
}

type weatherServiceUseCase struct {
	storage postgressql.WeatherStorage
}

func NewWeatherServiceUseCase(storage postgressql.WeatherStorage) *weatherServiceUseCase {
	return &weatherServiceUseCase{storage: storage}
}

func (w *weatherServiceUseCase) CreateCity(city entities.City) (id int, err error) {
	return w.storage.CreateCity(city)
}

func (w *weatherServiceUseCase) CreateWeather(weather entities.Weather, cityId int, info []byte) (id int, err error) {
	return w.storage.CreateWeather(weather, cityId, info)
}
