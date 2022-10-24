package usercase

import (
	"encoding/json"
	postgressql "weather_api/internal/adapters/db/postgres_sql"
	"weather_api/internal/entities"
	weatherserviceclient "weather_api/pkg/weather_service_client"
)

type weatherServiceUseCase struct {
	storage        postgressql.WeatherStorage
	weatherService weatherserviceclient.WeatherService
	cities         []entities.City
}

func NewWeatherServiceUseCase(storage postgressql.WeatherStorage, weatherService weatherserviceclient.WeatherService) *weatherServiceUseCase {
	return &weatherServiceUseCase{
		storage:        storage,
		weatherService: weatherService,
	}
}

func (w *weatherServiceUseCase) CreateCities() error {
	cities := w.weatherService.GetCities()
	for i := 0; i < len(cities); i++ {
		id, err := w.storage.CreateCity(cities[i])
		if err != nil {
			return err
		}
		cities[i].ID = id
	}
	w.cities = cities
	return nil
}

func (w *weatherServiceUseCase) CreateWeathers() error {
	weathers := w.weatherService.GetWeatherLists(w.cities)
	for _, listWeather := range weathers {
		for _, weather := range listWeather.List {
			json, err := w.getJson(weather)
			if err != nil {
				return err
			}
			_, err = w.storage.CreateWeather(entities.WeatherCreate{
				CityId: listWeather.WeatherCity.ID,
				Temp:   weather.Main.Temp,
				Date:   weather.Date,
				Info:   json,
			})
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (w *weatherServiceUseCase) getJson(weather entities.Weather) ([]byte, error) {
	infoData := entities.Info{
		Main:        weather.Main,
		InfoWeather: weather.InfoWeather,
		Clouds:      weather.Clouds,
		Wind:        weather.Wind,
		Visibility:  weather.Visibility,
	}
	infoJson, err := json.Marshal(infoData)
	if err != nil {
		return nil, err
	}
	return infoJson, nil
}
