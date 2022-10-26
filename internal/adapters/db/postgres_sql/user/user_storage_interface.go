package postgressql

import "weather_api/internal/entities"

type AuthorizationStorage interface {
	CreateUser(user entities.User) (int, error)
}
