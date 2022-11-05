package v2

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) CreateList(c *gin.Context) {
	return
}

func (h *Handler) GetListById(c *gin.Context) {
	return
}

func (h *Handler) GetAllList(c *gin.Context) {
	id, _ := c.Get(userCtx)
	logrus.Infoln("sdsf ", id)
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) UpdateList(c *gin.Context) {
	return
}

func (h *Handler) DeleteList(c *gin.Context) {
	return
}
