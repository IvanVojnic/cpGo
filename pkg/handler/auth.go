package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) signUp(c *gin.Context) {
	fmt.Println("hello")
	/*var input models.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(c)
	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}*/
	c.JSON(http.StatusOK, map[string]interface{}{
		"aaa": "aaa",
	})
}

func (h *Handler) test(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"hui": "hui",
	})
	fmt.Println("hui")
}

func (h *Handler) signIn(c *gin.Context) {

}
