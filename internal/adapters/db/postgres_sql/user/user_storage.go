package postgressql

import (
	"weather_api/internal/entities"

	"github.com/jmoiron/sqlx"
)

type authServiceStorage struct {
	db *sqlx.DB
}

func (s *authServiceStorage) CreateUser(user entities.User) (id int, err error) {
	return id, nil
}
