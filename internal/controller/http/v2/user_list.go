package v2

import (
	"net/http"
	"strconv"
	"weather_api/internal/entities"

	"github.com/gin-gonic/gin"
)

// @Summary Create list
// @Security ApiKeyAuth
// @Tags lists
// @Description create list
// @ID create-list
// @Accept  json
// @Produce  json
// @Param input body entities.UserList true "list info"
// @Success 200 {integer} 1
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v2/lists [post]
func (h *Handler) CreateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	var input entities.UserList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.listUseCase.CreateList(userId, input)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

// @Summary Get list by id
// @Security ApiKeyAuth
// @Tags lists
// @Description get list by id list
// @ID get-list-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.UserList
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v2/lists/{id} [get]
func (h *Handler) GetListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
	}
	list, err := h.listUseCase.GetListById(userId, id)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, list)
}

// @Summary Get all lists
// @Security ApiKeyAuth
// @Tags lists
// @Description get all lists
// @ID get-all-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []entities.UserList
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v2/lists/ [get]
func (h *Handler) GetAllList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	lists, err := h.listUseCase.GetAllList(userId)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, lists)
}

// @Summary Update list by id
// @Security ApiKeyAuth
// @Tags lists
// @Description update list by id
// @ID update-list-by-id
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Status"
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v2/lists/{id} [put]
func (h *Handler) UpdateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		NewErrorResponse(c, http.StatusBadRequest, "invalid list id param")
	}
	var input entities.UserList
	if err := c.BindJSON(&input); err != nil {
		NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := h.listUseCase.UpdateList(userId, id, input); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}

// @Summary Delete list by title
// @Security ApiKeyAuth
// @Tags lists
// @Description delete list by title
// @ID delete-list-by-title
// @Accept  json
// @Produce  json
// @Success 200 {string} string "Status"
// @Failure 400 {object} errorResponse
// @Failure 404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/v2/lists/{title} [delete]
func (h *Handler) DeleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	title := c.Param("title")
	if err := h.listUseCase.DeleteList(userId, title); err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"Status": "ok",
	})
}
