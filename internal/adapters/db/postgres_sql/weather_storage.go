package postgressql

import (
	"fmt"
	"time"
	"weather_api/config"
	"weather_api/internal/entities"

	"github.com/jmoiron/sqlx"
)

type weatherServiceStorage struct {
	db *sqlx.DB
}

func NewWeatherServiceStorage(db *sqlx.DB) *weatherServiceStorage {
	return &weatherServiceStorage{
		db: db,
	}
}

func (w *weatherServiceStorage) CreateCity(city entities.City) (id int, err error) {
	query := fmt.Sprintf(`INSERT INTO %s (name, lat, lon, country)
							VALUES ($1, $2, $3, $4) ON CONFLICT (name)
							DO UPDATE SET (name, lat, lon, country) = (EXCLUDED.name, EXCLUDED.lat, EXCLUDED.lon, EXCLUDED.country)
							RETURNING id`, config.CitiesTable)
	row := w.db.QueryRow(query, city.Name, city.Lat, city.Lon, city.Country)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (w *weatherServiceStorage) CreateWeather(weather entities.WeatherCreate) (id int, err error) {
	query := fmt.Sprintf(`INSERT INTO %s (city_id, temp, date, info)
							VALUES ($1, $2, $3, $4)
							ON CONFLICT ON CONSTRAINT %s
							DO UPDATE SET (city_id, temp, date, info) = (EXCLUDED.city_id, EXCLUDED.temp, EXCLUDED.date, EXCLUDED.info)
							RETURNING id`, config.WeathersTable, config.ConstraintWeather)
	row := w.db.QueryRow(query, weather.CityId, weather.Temp, weather.Date, weather.Info)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (w *weatherServiceStorage) GetCities() (names []string, err error) {
	query := fmt.Sprintf(`SELECT DISTINCT c.name
							FROM %s AS w
							JOIN %s AS c ON w.city_id = c.id
							ORDER BY 1`, config.WeathersTable, config.CitiesTable)
	if err := w.db.Select(&names, query); err != nil {
		return nil, err
	}
	return names, nil
}

func (w *weatherServiceStorage) GetWeatherInCity(name string) (weathers []entities.WeatherPredict, err error) {
	query := fmt.Sprintf(`SELECT c.country, c.name, AVG(w.temp) AS av_temp
							FROM %s AS w
							JOIN %s AS c ON w.city_id = c.id
							WHERE c.name = $1
							GROUP BY c.country, c.name`, config.WeathersTable, config.CitiesTable)
	if err := w.db.Select(&weathers, query, name); err != nil {
		return weathers, err
	}
	return weathers, nil
}

func (w *weatherServiceStorage) GetDatesInCity(name string) (dates []time.Time, err error) {
	query := fmt.Sprintf(`SELECT w.date
							FROM %s AS w
							JOIN %s AS c ON w.city_id = c.id
							WHERE c.name = $1
							ORDER BY 1`, config.WeathersTable, config.CitiesTable)
	if err := w.db.Select(&dates, query, name); err != nil {
		return dates, err
	}
	return dates, err
}

func (w *weatherServiceStorage) GetDetaiWeatherInCity(name string, date time.Time) (weathers []entities.WeatherDetails, err error) {
	query := fmt.Sprintf(`SELECT c.name, w.date, w.info::text
							FROM %s AS w
							JOIN %s AS c ON w.city_id = c.id
							WHERE c.name = $1 AND w.date = $2
							ORDER BY 1`, config.WeathersTable, config.CitiesTable)
	if err := w.db.Select(&weathers, query, name, date); err != nil {
		return weathers, err
	}
	return weathers, nil
}
