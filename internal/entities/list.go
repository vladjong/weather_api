package entities

type UserLists struct {
	Id         int `json:"-" db:"id"`
	UserID     int `json:"user_id" db:"user_id"`
	LikeListID int `json:"like_list_id" db:"like_list_id"`
}

type LikeLists struct {
	Id     int `json:"-" db:"id"`
	CityID int `json:"city_id" db:"city_id"`
}
