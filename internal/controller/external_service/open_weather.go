package externalservice

import (
	"encoding/json"
	"weather_api/internal/entities"
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

func (o *openWeatherApi) CreateCities() ([]entities.City, error) {
	cities := o.weatherService.GetCities()
	for i := 0; i < len(cities); i++ {
		id, err := o.weatherUseCase.CreateCity(cities[i])
		if err != nil {
			return cities, err
		}
		cities[i].ID = id
	}
	return cities, nil
}

func (o *openWeatherApi) CreateWeathers(cities []entities.City) {
	weathers := o.weatherService.GetWeatherLists(cities)
	for _, listWeather := range weathers {
		for _, weather := range listWeather.List {
			o.weatherUseCase.CreateWeather(weather, listWeather.WeatherCity.ID, o.getJson(weather))
		}
	}
}

func (o *openWeatherApi) getJson(weather entities.Weather) []byte {
	infoData := entities.Info{
		Main:        weather.Main,
		InfoWeather: weather.InfoWeather,
		Clouds:      weather.Clouds,
		Wind:        weather.Wind,
		Visibility:  weather.Visibility,
	}
	infoJson, err := json.Marshal(infoData)
	if err != nil {
		logrus.Fatal(err)
	}
	return infoJson
}
