package config

import "time"

func GetTownList() []string {
	return []string{
		"London",
		"Moscow",
		"Madrid",
		"Bangkok",
		"Chicago",
		"Dubai",
		"Paris",
		"Saint Petersburg",
		"Tokyo",
		"Novosibirsk",
		"Sydney",
		"Toronto",
		"Geneva",
		"Hong Kong",
		"Lagos",
		"Los Angeles",
		"Sofia",
		"Rome",
		"Seoul",
		"Oslo",
	}
}

const (
	CitiesTable       = "Cities"
	WeathersTable     = "Weathers"
	UserTable         = "Users"
	ListItemsTable    = "List_Items"
	UserListTable     = "User_Lists"
	ConstraintWeather = "unique_weather"
	Salt              = "fgfertdl123wedfsdf"
	TokenTTL          = 12 * time.Hour
	SignedKey         = "vcbr@34$re$d@43r#dsf@!1"
)
