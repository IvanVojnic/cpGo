package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userToGet struct {
	Email string `json:"email" binding:"required"`
}

func (h *Handler) findUser(c *gin.Context) {
	fmt.Println("findUser")

	var input userToGet

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	user, err := h.services.UserCommunicate.FindUser(input.Email)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(user)
	c.IndentedJSON(http.StatusOK, gin.H{"email": user.Email, "name": user.Name, "id": user.Id})
}

func (h *Handler) sendFriends(c *gin.Context) {
	fmt.Println("signIn")

	var input userToGet

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	/*user, err := h.services.UserCommunicate
	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"token": token})*/
}
