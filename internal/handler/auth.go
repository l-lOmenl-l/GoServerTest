package handler

import (
	"example/web-servise-gin/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary SignUp
// @Tags auth
// @Description create user
// @ID create-user
// @Accept json
// @Produce json
// @Param input body domain.User true "user info"
// @Router /auth/sign-up [post]

func (h *Handler) signUp(c *gin.Context) {
	var input domain.User

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

// @Summary SignIn
// @Tags auth
// @Description login
// @ID user-login
// @Accept json
// @Produce json
// @Param input body domain.SignIn true "credentials"
// @Router /auth/sign-in [post]
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
