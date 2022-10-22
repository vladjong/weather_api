package config

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
	ConstraintWeather = "unique_weather"
)
