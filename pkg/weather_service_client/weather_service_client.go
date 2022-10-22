package weatherserviceclient

import "weather_api/internal/entities"

type WeatherService interface {
	GetCities() (towns []entities.City)
	GetWeatherLists(cities []entities.City) (listWeather []entities.ListWeather)
}
