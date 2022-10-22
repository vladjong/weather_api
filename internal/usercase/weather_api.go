package usercase

import (
	postgressql "weather_api/internal/adapters/db/postgres_sql"
	"weather_api/internal/entities"
)

type weatherApiUseCase struct {
	storage postgressql.WeatherStorageApi
}

func NewWeatherApiUseCase(storage postgressql.WeatherStorageApi) *weatherApiUseCase {
	return &weatherApiUseCase{storage: storage}
}

func (w *weatherApiUseCase) GetCities() (names []string, err error) {
	return w.storage.GetCities()
}

func (w *weatherApiUseCase) GetWeatherInCity(name string) (weathers []entities.WeatherDetails, err error) {
	return w.storage.GetWeatherInCity(name)
}

func (w *weatherApiUseCase) GetDetaiWeatherInCity(name string, date string) (weathers []entities.WeatherPredict, err error) {
	return w.storage.GetDetaiWeatherInCity(name, date)
}
