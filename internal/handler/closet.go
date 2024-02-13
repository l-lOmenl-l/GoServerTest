package handler

import (
	"example/web-servise-gin/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) addTypeProduct(c *gin.Context) {

	var type_product domain.TypeProduct
	if err := c.BindJSON(&type_product); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	newType, err := h.services.AddTypeProduct(type_product.Name)
	if err != nil {

	}

	c.JSON(http.StatusOK, newType)
}
