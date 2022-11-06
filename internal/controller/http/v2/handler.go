package v2

import (
	_ "weather_api/docs"
	"weather_api/internal/usecase"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	weatherUseCase usecase.WeatherAPI
	userUseCase    usecase.Authorization
	listUseCase    usecase.List
}

func NewHandler(weatherUseCase usecase.WeatherAPI, userUseCase usecase.Authorization, listUseCase usecase.List) *Handler {
	return &Handler{
		weatherUseCase: weatherUseCase,
		userUseCase:    userUseCase,
		listUseCase:    listUseCase,
	}
}

// @title Weather API
// @version 2.0
// @description This is a service that predicts the weather

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func (h *Handler) NewRouter() *gin.Engine {
	router := gin.New()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
	}
	api := router.Group("/api/v2")
	{
		api.GET("/cities", h.GetCities)
		api.GET("/cities/:name", h.GetWeatherInCity)
		api.GET("/detail_weather/:name/:date", h.GetDetailWeatherInCity)

		lists := api.Group("/lists", h.userIdentity)
		{
			lists.POST("/", h.CreateList)
			lists.GET("/:id", h.GetListById)
			lists.GET("/", h.GetAllList)
			lists.PUT("/:id", h.UpdateList)
			lists.DELETE("/:title", h.DeleteList)

			items := lists.Group(":id/items")
			{
				items.POST("/:city", h.CreateItem)
				items.GET("/", h.GetAllItems)
			}
		}
	}
	return router
}
