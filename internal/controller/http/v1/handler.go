package v1

import (
	_ "weather_api/docs"
	usercase "weather_api/internal/usercase/weather"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	weatherUseCase usercase.WeatherAPI
}

func NewHandler(weatherUseCase usercase.WeatherAPI) *Handler {
	return &Handler{
		weatherUseCase: weatherUseCase,
	}
}

// @title Weather API
// @version 1.0
// @description This is a service that predicts the weather

// @host      localhost:8080
// @BasePath  /api/v1
func (h *Handler) NewRouter() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	api := router.Group("/api/v1")
	{
		api.GET("/cities", h.GetCities)
		api.GET("/cities/:name", h.GetWeatherInCity)
		api.GET("/detail_weather/:name/:date", h.GetDetailWeatherInCity)
	}
	return router
}
