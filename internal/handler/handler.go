package handler

import (
	"example/web-servise-gin/internal/service"
	"github.com/gin-gonic/gin"

	_ "github.com/omen/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	users := router.Group("/users", h.userIdentity)
	{
		users.GET("/getAll", h.getAll)
		users.GET("/detail", h.detailuser)
	}

	closet := router.Group("/closet", h.userIdentity)
	{
		closet.POST("/addTypeProduct", h.addTypeProduct)
	}

	return router
}
