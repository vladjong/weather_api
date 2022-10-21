package entities

type ListWeather struct {
	List []Weather `json:"list"`
}

type Weather struct {
	Main        map[string]float64 `json:"main"`
	InfoWeather []InfoWeather      `json:"weather"`
	Clouds      string             `json:"clouds"`
	Wind        map[string]float64 `json:"wind"`
	Visibility  string             `json:"visibility"`
	Date        string             `json:"dt_txt"`
}

type InfoWeather struct {
	Main        string `json:"main"`
	Description string `json:"description"`
}
