package postgressql

import (
	"fmt"
	"time"
	"weather_api/config"
	"weather_api/internal/entities"

	"github.com/jmoiron/sqlx"
)

type weatherServiceStorageAPI struct {
	db *sqlx.DB
}

func NewWeatherServiceStorageAPI(db *sqlx.DB) *weatherServiceStorageAPI {
	return &weatherServiceStorageAPI{
		db: db,
	}
}

func (w *weatherServiceStorageAPI) GetCities() (names entities.AllCities, err error) {
	query := fmt.Sprintf(`SELECT DISTINCT c.name
							FROM %s AS w
							JOIN %s AS c ON w.city_id = c.id
							ORDER BY 1`, config.WeathersTable, config.CitiesTable)
	var cities []string
	if err := w.db.Select(&cities, query); err != nil {
		return names, err
	}
	return entities.AllCities{
		Cities: cities,
	}, nil
}

func (w *weatherServiceStorageAPI) GetWeatherInCity(name string) (weathers []entities.WeatherPredict, err error) {
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

func (w *weatherServiceStorageAPI) GetDatesInCity(name string) (dates []time.Time, err error) {
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

func (w *weatherServiceStorageAPI) GetDetaiWeatherInCity(name string, date time.Time) (weathers []entities.WeatherDetails, err error) {
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
