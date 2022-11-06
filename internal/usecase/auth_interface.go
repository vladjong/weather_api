package usecase

import "weather_api/internal/entities"

//go:generate mockgen -source=weather_interface.go -destination=moks/mock.go

type Authorization interface {
	CreateUser(user entities.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}
