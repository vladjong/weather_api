package externalservice

import (
	"weather_api/internal/usercase"
)

type openWeatherApi struct {
	weatherUseCase usercase.WeatherService
}

func NewOpenWeatherApi(weatherUseCase usercase.WeatherService) *openWeatherApi {
	return &openWeatherApi{
		weatherUseCase: weatherUseCase,
	}
}

func (o *openWeatherApi) CreateCities() error {
	return o.weatherUseCase.CreateCities()
}

func (o *openWeatherApi) CreateWeathers() (err error) {
	return o.weatherUseCase.CreateWeathers()
}
