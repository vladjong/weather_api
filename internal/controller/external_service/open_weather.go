package externalservice

import (
	"weather_api/internal/usercase"
	weatherserviceclient "weather_api/pkg/weather_service_client"

	"github.com/sirupsen/logrus"
)

type openWeatherApi struct {
	weatherService weatherserviceclient.WeatherService
	weatherUseCase usercase.WeatherService
}

func NewOpenWeatherApi(weatherService weatherserviceclient.WeatherService, weatherUseCase usercase.WeatherService) *openWeatherApi {
	return &openWeatherApi{
		weatherService: weatherService,
		weatherUseCase: weatherUseCase,
	}
}

func (o *openWeatherApi) SetCities() {
	cities := o.weatherService.GetCities()
	for _, city := range cities {
		if _, err := o.weatherUseCase.CreateCity(city); err != nil {
			logrus.Fatal(err)
		}
	}
}
