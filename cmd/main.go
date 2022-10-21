package main

import (
	"fmt"
	openweather "weather_api/pkg/weather_service/open_weather"
)

func main() {
	weather := openweather.NewOpenWeatherApi(1, "8f7a589bb238eb36173737ebbe1ec8c6", "metric")
	towns := weather.GetTownStructs()
	fmt.Println(towns)
}
