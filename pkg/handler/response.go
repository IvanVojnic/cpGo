package handler

import (
	"github.com/gin-gonic/gin"
	"log"
)

type error struct {
	Message string `json:"message"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	log.Fatalf(message)
	c.AbortWithStatusJSON(statusCode, error{message})
}
