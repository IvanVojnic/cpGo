package handler

import (
	"github.com/IvanVojnic/cpGo.git/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.signUp)
		auth.POST("/login", h.signIn)
		auth.GET("/test", h.test)
	}
	return router
}
