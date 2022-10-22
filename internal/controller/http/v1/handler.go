package v1

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) NewRouter() *gin.Engine {
	router := gin.New()
	// router.GET("/swagger/*any", ginSwagger.WrapHandler(swag))

	api := router.Group("/api")
	{
		api.GET("/", h.get)
	}
	return router
}
