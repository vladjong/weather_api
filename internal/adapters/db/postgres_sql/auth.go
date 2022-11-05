package postgressql

import (
	"fmt"
	"weather_api/config"
	"weather_api/internal/entities"

	"github.com/jmoiron/sqlx"
)

type authServiceStorage struct {
	db *sqlx.DB
}

func NewAuthServiceStorage(db *sqlx.DB) *authServiceStorage {
	return &authServiceStorage{
		db: db,
	}
}

func (s *authServiceStorage) CreateUser(user entities.User) (id int, err error) {
	query := fmt.Sprintf("INSERT INTO %s (login, password_hash) VALUES ($1, $2) RETURNING id", config.UserTable)
	row := s.db.QueryRow(query, user.Login, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (s *authServiceStorage) GetUser(login, password string) (user entities.User, err error) {
	query := fmt.Sprintf("SELECT id FROM %s WHERE login=$1 AND password_hash=$2", config.UserTable)
	err = s.db.Get(&user, query, login, password)
	return user, err
}
