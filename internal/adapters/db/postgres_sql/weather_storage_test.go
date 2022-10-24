package postgressql

import (
	"testing"
	"weather_api/internal/entities"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestWeather_CreateCity(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer mockDB.Close()
	sqlxDB := sqlx.NewDb(mockDB, "sqlmock")
	r := NewWeatherServiceStorage(sqlxDB)

	cases := []struct {
		name    string
		moc     func()
		input   entities.City
		want    int
		wantErr bool
	}{
		{
			name: "OK",
			input: entities.City{
				Name:    "Moscow",
				Lat:     12.23,
				Lon:     42.24,
				Country: "RU",
			},
			want: 1,
			moc: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
				mock.ExpectQuery("INSERT INTO Cities").
					WithArgs("Moscow", 12.23, 42.24, "RU").WillReturnRows(rows)
			},
		},
		{
			name: "Empty fields",
			input: entities.City{
				Name:    "",
				Country: "",
			},
			wantErr: true,
			moc: func() {
				rows := sqlmock.NewRows([]string{"id"})
				mock.ExpectQuery("INSERT INTO Cities").
					WithArgs("", "").WillReturnRows(rows)
			},
		},
	}
	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.moc()
			got, err := r.CreateCity(testCase.input)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.want, got)
			}
		})
	}
}
