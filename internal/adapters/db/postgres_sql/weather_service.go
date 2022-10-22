package postgressql

import (
	"fmt"
	"weather_api/config"
	"weather_api/internal/entities"

	"github.com/jmoiron/sqlx"
)

type weatherServiceStorage struct {
	db *sqlx.DB
}

type WeatherStorage interface {
	CreateCity(city entities.City) (id int, err error)
}

func NewWeatherServiceStorage(db *sqlx.DB) *weatherServiceStorage {
	return &weatherServiceStorage{
		db: db,
	}
}

func (w *weatherServiceStorage) CreateCity(city entities.City) (id int, err error) {
	query := fmt.Sprintf("INSERT INTO %s (name, lat, lon, country) VALUES ($1, $2, $3, $4) ON CONFLICT (name) DO UPDATE SET (name, lat, lon, country) = (EXCLUDED.name, EXCLUDED.lat, EXCLUDED.lon, EXCLUDED.country) RETURNING id", config.CitiesTable)
	row := w.db.QueryRow(query, city.Name, city.Lat, city.Lon, city.Country)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
