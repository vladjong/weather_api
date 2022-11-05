package entities

type UserList struct {
	Id     int    `json:"-" db:"id"`
	UserID int    `json:"user_id" db:"user_id"`
	Title  string `json:"title" db:"title"`
}

type ListItem struct {
	Id     int `json:"-" db:"id"`
	CityID int `json:"city_id" db:"city_id"`
	ListID int `json:"list_id" db:"list_id"`
}
