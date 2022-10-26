package usercase

import postgressql "weather_api/internal/adapters/db/postgres_sql/user"

type AuthorizationUseCase struct {
	storage postgressql.AuthorizationStorage
}
