package v2

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create items by sity name
// @Security ApiKeyAuth
// @Tags items
// @Description create items by sity name
// @ID create-items-by-sity-name
// @Accept  json
// @Produce  json
// @Success 200 {integer} 1
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v2/lists/:id/items/:city [post]
func (h *Handler) CreateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
	}
	city := c.Param("city")
	id, err := h.listUseCase.CreateItem(userId, listId, city)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get items by list id
// @Security ApiKeyAuth
// @Tags items
// @Description get items by list id
// @ID get-items-by-list-id
// @Accept  json
// @Produce  json
// @Success 200 {object} []entities.Item
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v2/lists/:id/items/ [get]
func (h *Handler) GetAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
	}
	items, err := h.listUseCase.GetAllItems(userId, id)
	if err != nil {
		NewErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	c.JSON(http.StatusOK, items)
}
