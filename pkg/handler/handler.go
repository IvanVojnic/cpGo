package handler

import (
	"fmt"
	"github.com/IvanVojnic/cpGo.git/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.New()
	router.Use(CORSMiddleware())
	fmt.Println("init routes")
	auth := router.Group("/auth")
	{
		auth.POST("/signup", h.signUp)
		auth.POST("/login", h.signIn)
		auth.GET("/test", h.test)
		auth.GET("/private", h.isAuth)
	}
	userCommunicate := router.Group("")
	{
		//userCommunicate.POST("/getFriends", h.sendFriends)
		userCommunicate.POST("/findFriend", h.findUser)
		/*userCommunicate.POST("/sendRequest")
		userCommunicate.POST("/getFriendsRequest")
		userCommunicate.POST("/acceptFriendsRequest")
		userCommunicate.POST("/sendInvite")*/
	}
	return router
}
