package entities

type City struct {
	ID      int     `json:"-" db:"id"`
	Name    string  `json:"name" db:"name"`
	Lat     float64 `json:"lat" db:"lat"`
	Lon     float64 `json:"lon" db:"lon"`
	Country string  `json:"country" db:"country"`
}

type AllCities struct {
	Cities []string `json:"data"`
}
