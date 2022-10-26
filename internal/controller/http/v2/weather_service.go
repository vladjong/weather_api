package v2

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Get Cities
// @Tags weather
// @Description get all cities
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.AllCities
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /cities [get]
func (h *Handler) GetCities(c *gin.Context) {
	cities, err := h.weatherUseCase.GetCities()
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, cities)
}

// @Summary Get prediction weather in City
// @Tags weather
// @Description get string by NANE
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.WeatherPredict
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /cities/{name} [get]
func (h *Handler) GetWeatherInCity(c *gin.Context) {
	name := c.Param("name")
	weathers, err := h.weatherUseCase.GetWeatherInCity(name)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, weathers)
}

// @Summary Get all information about the weather in the city on the exact day
// @Description get STRING by NAME and DATE (YYYY-MM-DDTHH:MM:SSZ)
// @Tags weather
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.WeatherDetails
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /detail_weather/{name}/{date} [get]
func (h *Handler) GetDetailWeatherInCity(c *gin.Context) {
	name := c.Param("name")
	date := c.Param("date")
	dateTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	weathers, err := h.weatherUseCase.GetDetaiWeatherInCity(name, dateTime)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, weathers)
}
