package entities

type Town struct {
	// Id         int               `json:"-" db:"id"`
	Name string `json:"name" db:"name"`
	// Local_names map[string]string `json:"local_names"`
	Lat float64 `json:"lat" db:"lat"`
	Lon float64 `json:"lon" db:"lon"`
	// Country     string            `json:"country"`
	// State       string            `json:"state"`
}
