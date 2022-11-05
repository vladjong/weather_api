package externalservice

import (
	"weather_api/internal/usecase"
)

type openWeatherApi struct {
	weatherUseCase usecase.WeatherService
}

func NewOpenWeatherApi(weatherUseCase usecase.WeatherService) *openWeatherApi {
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
