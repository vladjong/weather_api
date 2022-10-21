package weatherservice

import "weather_api/internal/entities"

type WeatherService interface {
	GetTownStructs() (towns []entities.Town)
}
