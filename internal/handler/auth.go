package handler

import (
	"example/web-servise-gin/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegistryUser Registry User
//
//	@Summary		Registry User
//	@Description	Registry User
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			input	body		domain.SignUp	true	"credentials"
//	@Router			/auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	var input domain.SignUp

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

// AuthUser Auth User
//
//	@Summary		Auth User
//	@Description	Auth User
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			input	body		domain.SignIn	true	"credentials"
//	@Router			/auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
	var input domain.SignIn

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"token": token})
}
