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

func (o *openWeatherApi) GetTowns() (towns []entities.Town) {
	names := config.GetTownList()
	for _, name := range names {
		towns = append(towns, o.getTownStruct(name))
	}
	return towns
}

func (o *openWeatherApi) GetPredictionWeathers() (listWeather []entities.ListWeather) {
	return listWeather
}

func (o *openWeatherApi) getTownStruct(name string) entities.Town {
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
	var town []entities.Town
	if err := json.Unmarshal(jsonBytes, &town); err != nil {
		logrus.Fatal(err)
	}
	return town[0]
}
