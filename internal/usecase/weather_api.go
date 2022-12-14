package usecase

import (
	"fmt"
	"time"
	"weather_api/internal/adapters/db"
	"weather_api/internal/entities"
)

type weatherApiUseCase struct {
	storage db.WeatherStorageAPI
}

func NewWeatherApiUseCase(storage db.WeatherStorageAPI) *weatherApiUseCase {
	return &weatherApiUseCase{storage: storage}
}

func (w *weatherApiUseCase) GetCities() (names entities.AllCities, err error) {
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

func (w *weatherApiUseCase) GetDetaiWeatherInCity(name string, date time.Time) (weather entities.WeatherDetails, err error) {
	weathers, err := w.storage.GetDetaiWeatherInCity(name, date)
	if len(weathers) != 1 {
		return weather, fmt.Errorf("error: not found data")
	}
	return weathers[0], err
}
