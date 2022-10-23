package v1

import (
	"net/http"
	errorresponse "weather_api/pkg/error_response"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCities(c *gin.Context) {
	cities, err := h.weatherUseCase.GetCities()
	if err != nil {
		errorresponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, cities)
}

func (h *Handler) GetWeatherInCity(c *gin.Context) {
	name := c.Param("name")
	weathers, err := h.weatherUseCase.GetWeatherInCity(name)
	if err != nil {
		errorresponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, weathers)
}

func (h *Handler) GetDetaiWeatherInCity(c *gin.Context) {
	name := c.Param("name")
	date := c.Param("date")
	weathers, err := h.weatherUseCase.GetDetaiWeatherInCity(name, date)
	if err != nil {
		errorresponse.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusOK, weathers)
}
