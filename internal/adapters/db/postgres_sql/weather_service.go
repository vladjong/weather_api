package postgressql

import (
	"fmt"
	"weather_api/config"
	"weather_api/internal/entities"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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

func (w *weatherServiceStorage) CreateWeather(weather entities.Weather, cityId int, info []byte) (id int, err error) {
	query := fmt.Sprintf(`INSERT INTO %s (city_id, temp, date, info)
							VALUES ($1, $2, $3, $4)
							ON CONFLICT ON CONSTRAINT %s
							DO UPDATE SET (city_id, temp, date, info) = (EXCLUDED.city_id, EXCLUDED.temp, EXCLUDED.date, EXCLUDED.info)
							RETURNING id`, config.WeathersTable, config.ConstraintWeather)
	row := w.db.QueryRow(query, cityId, weather.Main.Temp, weather.Date, info)
	if err := row.Scan(&id); err != nil {
		logrus.Fatal(err)
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

func (w *weatherServiceStorage) GetWeatherInCity(name string) (weather entities.WeatherPredict, err error) {
	query := fmt.Sprintf(`SELECT DISTINCT c.country, c.name, array_agg(w.date) AS dates, AVG(w.temp) AS av_temp
							FROM %s AS w
							JOIN %s AS c ON w.city_id = c.id
							WHERE c.name = $1
							GROUP BY c.country, c.name`, config.WeathersTable, config.CitiesTable)
	var weathers []entities.WeatherPredict
	if err := w.db.Select(&weathers, query, name); err != nil {
		return weather, err
	}
	return weathers[0], nil
}

func (w *weatherServiceStorage) GetDetaiWeatherInCity(name string, date string) (weathers entities.WeatherDetails, err error) {
	return weathers, nil
}
