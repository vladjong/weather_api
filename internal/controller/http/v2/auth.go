package v2

import (
	"net/http"
	"weather_api/internal/entities"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input entities.User
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) SignIn(c *gin.Context) {

}
