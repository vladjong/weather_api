package openweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"weather_api/internal/config"
	"weather_api/internal/entities"

	"github.com/sirupsen/logrus"
)

type ConfigOpenWeather struct {
	Limit int
	Appid string
	Units string
}

type openWeatherApi struct {
	cfg *ConfigOpenWeather
}

func NewOpenWeatherApi(limit int, appid, units string) *openWeatherApi {
	return &openWeatherApi{
		cfg: &ConfigOpenWeather{
			Limit: limit,
			Appid: appid,
			Units: units,
		},
	}
}

func (o *openWeatherApi) GetCities() (cities []entities.City) {
	names := config.GetTownList()
	for _, name := range names {
		cities = append(cities, o.getCityStruct(name))
	}
	return cities
}

func (o *openWeatherApi) GetWeatherLists(cities []entities.City) (listWeather []entities.ListWeather) {
	for _, city := range cities {
		listWeather = append(listWeather, o.getWeatherListStruct(city.Lat, city.Lon))
	}
	return listWeather
}

func (o *openWeatherApi) getWeatherListStruct(lat, lon float64) (listWeather entities.ListWeather) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/forecast?lat=%f&lon=%f&units=%s&appid=%s",
		lat, lon, o.cfg.Units, o.cfg.Appid)
	response, err := http.Get(url)
	if err != nil {
		logrus.Fatal(err)
	}
	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			logrus.Fatal(err)
		}
	}()
	if err := json.Unmarshal(jsonBytes, &listWeather); err != nil {
		logrus.Fatal(err)
	}
	return listWeather
}

func (o *openWeatherApi) getCityStruct(name string) entities.City {
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&limit=%d&appid=%s",
		name, o.cfg.Limit, o.cfg.Appid)
	response, err := http.Get(url)
	if err != nil {
		logrus.Fatal(err)
	}
	jsonBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		logrus.Fatal(err)
	}
	defer func() {
		if err := response.Body.Close(); err != nil {
			logrus.Fatal(err)
		}
	}()
	var town []entities.City
	if err := json.Unmarshal(jsonBytes, &town); err != nil {
		logrus.Fatal(err)
	}
	return town[0]
}
