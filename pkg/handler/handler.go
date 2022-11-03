package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/singUp", h.singUp)
		auth.POST("/singIn", h.singIn)
	}
	return router
}
