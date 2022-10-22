package main

import (
	"fmt"
	openweather "weather_api/pkg/weather_service_client/open_weather"
)

func main() {
	weather := openweather.NewOpenWeatherApi(1, "8f7a589bb238eb36173737ebbe1ec8c6", "metric")
	towns := weather.GetWeatherListStruct(55.75, 37.62)
	fmt.Println(towns.List[0].Main.Temp)
}
