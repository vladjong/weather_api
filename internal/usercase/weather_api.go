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

func (w *weatherApiUseCase) GetWeatherInCity(name string) (entities.WeatherPredictDTO, error) {
	weather, err := w.storage.GetWeatherInCity(name)
	dates := weather.Dates
	var datesStr []string
	for _, date := range dates {
		str := string(date)
		// str = strings.ReplaceAll(str, "/", "")
		datesStr = append(datesStr, str)
	}
	return entities.WeatherPredictDTO{
		Country: weather.Country,
		Name:    weather.Name,
		AvTemp:  weather.AvTemp,
		Dates:   datesStr,
	}, err
}

func (w *weatherApiUseCase) GetDetaiWeatherInCity(name string, date string) (entities.WeatherDetails, error) {
	return w.storage.GetDetaiWeatherInCity(name, date)
}
