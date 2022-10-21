package weatherserviceclient

import "weather_api/internal/entities"

type WeatherService interface {
	GetTowns() (towns []entities.Town)
	GetPredictionWeathers() (listWeather []entities.ListWeather)
}
