package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) detailuser(c *gin.Context) {
	id, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
	}

	detail, err := h.services.Users.GetDetail(id.(int))
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, detail)
}

func (h *Handler) getAll(c *gin.Context) {
	AllUsers, err := h.services.GetAll()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, AllUsers)
}
