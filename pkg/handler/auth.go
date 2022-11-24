package handler

import (
	"fmt"
	"github.com/IvanVojnic/cpGo.git/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type signInInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

const authorizationHeader = "Authorization"

func (h *Handler) signUp(c *gin.Context) {
	fmt.Println("hello")

	var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) isAuth(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth Header")
		return
	}

}

func (h *Handler) signIn(c *gin.Context) {
	fmt.Println("signIn")

	var input signInInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	userId, err2 := h.services.Authorization.GetUser(input.Email, input.Password)

	if err2 != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token, "id": userId})
}

func (h *Handler) test(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"b": "b",
	})
	fmt.Println("b")
}

/*
migrate -path ./schema -database 'postgres://postgres:vojnic@localhost:5436/postgres?sslmode=disable' up
docker run --name=cpGo-db -e POSTGRES_PASSWORD='vojnic' -p 5436:5432 -d --rm postgres
*/
