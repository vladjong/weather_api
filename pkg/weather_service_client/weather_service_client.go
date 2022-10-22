package weatherserviceclient

import "weather_api/internal/entities"

type WeatherService interface {
	GetCities() (cities []entities.City)
	GetWeatherLists(cities []entities.City) (listWeather []entities.ListWeather)
}
