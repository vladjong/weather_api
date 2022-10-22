package entities

type ListWeather struct {
	List        []Weather `json:"list"`
	WeatherCity City      `json:"city"`
}

type Weather struct {
	Main        Main               `json:"main"`
	InfoWeather []InfoWeather      `json:"weather"`
	Clouds      map[string]float64 `json:"clouds"`
	Wind        map[string]float64 `json:"wind"`
	Visibility  float64            `json:"visibility"`
	Date        string             `json:"dt_txt"`
}

type Main struct {
	Temp       float64 `json:"temp" db:"temp"`
	FeelsLike  float64 `json:"feels_like"`
	TempMin    float64 `json:"temp_min"`
	TempMap    float64 `json:"temp_map"`
	Pressure   float64 `json:"pressure"`
	SeaLevel   float64 `json:"sea_level"`
	GrindLevel float64 `json:"grind_level"`
	Humidity   float64 `json:"humidity"`
	TempKf     float64 `json:"temp_kf"`
}

type InfoWeather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}

type Info struct {
	Main        Main
	InfoWeather []InfoWeather
	Clouds      map[string]float64
	Wind        map[string]float64
	Visibility  float64
}

type WeatherPredict struct {
	Country     string   `json:"country" db:"country"`
	Name        string   `json:"name" db:"name"`
	AverageTemp float64  `json:"average_temp"`
	Dates       []string `json:"dates"`
}

type WeatherDetails struct {
	CityID int     `json:"city_id" db:"city_id"`
	Temp   float64 `json:"temp" db:"temp"`
	Date   string  `json:"date" db:"date"`
	Info   Info    `json:"info" db:"info"`
}
