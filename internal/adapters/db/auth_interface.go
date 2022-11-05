package db

import "weather_api/internal/entities"

type AuthorizationStorage interface {
	CreateUser(user entities.User) (int, error)
	GetUser(login, password string) (entities.User, error)
}
