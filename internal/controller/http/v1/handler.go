package v1

import (
	"weather_api/internal/usercase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	weatherUseCase usercase.WeatherAPI
}

func NewHandler(weatherUseCase usercase.WeatherAPI) *Handler {
	return &Handler{
		weatherUseCase: weatherUseCase,
	}
}

func (h *Handler) NewRouter() *gin.Engine {
	router := gin.New()
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swag))

	api := router.Group("/api")
	{
		api.GET("/cities", h.GetCities)
		api.GET("/city/:name", h.GetWeatherInCity)
	}
	return router
}
