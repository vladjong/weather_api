package usercase

import (
	"fmt"
	"time"
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

func (w *weatherApiUseCase) GetWeatherInCity(name string) (weather entities.WeatherPredict, err error) {
	weathers, err := w.storage.GetWeatherInCity(name)
	if err != nil {
		return weather, err
	}
	if len(weathers) != 1 {
		return weather, fmt.Errorf("error: not found data")
	}
	dates, err := w.storage.GetDatesInCity(name)
	weathers[0].Dates = dates
	return weathers[0], err
}

func (w *weatherApiUseCase) GetDetaiWeatherInCity(name string, date string) (weather entities.WeatherDetails, err error) {
	dateTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return weather, err
	}
	weathers, err := w.storage.GetDetaiWeatherInCity(name, dateTime)
	if len(weathers) != 1 {
		return weather, fmt.Errorf("error: not found data")
	}
	return weathers[0], err
}
