package v1

import (
	"fmt"
	"net/http/httptest"
	"testing"
	"weather_api/internal/entities"
	mock_usercase "weather_api/internal/usercase/mocks"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"
)

func TestHandler_GetCities(t *testing.T) {
	type mockBehavior func(c *mock_usercase.MockWeatherAPI)

	cases := []struct {
		name                 string
		mockBehavior         mockBehavior
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "OK",
			mockBehavior: func(c *mock_usercase.MockWeatherAPI) {
				c.EXPECT().GetCities().Return(entities.AllCities{
					Cities: []string{"Moscow", "Tokyo"},
				}, nil)
			},
			expectedStatusCode:   200,
			expectedResponseBody: `{"data":["Moscow","Tokyo"]}`,
		},
		{
			name: "Wrong Input",
			mockBehavior: func(c *mock_usercase.MockWeatherAPI) {
				c.EXPECT().GetCities().Return(entities.AllCities{
					Cities: []string{""},
				}, fmt.Errorf("something went wrong"))
			},
			expectedStatusCode:   404,
			expectedResponseBody: `{"message":"something went wrong"}`,
		},
	}
	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()
			useCase := mock_usercase.NewMockWeatherAPI(c)
			test.mockBehavior(useCase)
			handler := NewHandler(useCase)
			r := gin.New()
			r.GET("/cities", handler.GetCities)
			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/cities", nil)
			r.ServeHTTP(w, req)
			assert.Equal(t, w.Code, test.expectedStatusCode)
			assert.Equal(t, w.Body.String(), test.expectedResponseBody)
		})
	}
}
