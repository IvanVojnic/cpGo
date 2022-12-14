package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type userToGet struct {
	Email string `json:"email" binding:"required"`
}

type userId struct {
	Id string `json:"id"`
}

type dataToInvite struct {
	UserSender  int   `json:"userSender"`
	FriendsList []int `json:"friendsList"`
	Id_place    int   `json:"id_place"`
}

type inputRequestFriend struct {
	UserSender   string `json:"userSender"`
	UserReceiver int    `json:"userReceiver"`
}

type inputAcceptRequest struct {
	UserSender   int    `json:"userSender"`
	UserReceiver string `json:"userReceiver"`
}

type getRoomByUserId struct {
	Id string `json:"userId"`
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
	c.IndentedJSON(http.StatusOK, gin.H{"email": user.Email, "name": user.Name, "id": user.Id})
}

func (h *Handler) sendRequest(c *gin.Context) {
	fmt.Println("sendRequest")

	var input inputRequestFriend

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var userSender int
	userSender, _ = strconv.Atoi(input.UserSender)
	err := h.services.UserCommunicate.SendRequest(userSender, input.UserReceiver)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "request sent"})

}

func (h *Handler) getFriendsRequest(c *gin.Context) {
	fmt.Println("getFriendsRequest")

	var input userId

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var user_id int
	user_id, _ = strconv.Atoi(input.Id)
	users, err := h.services.UserCommunicate.GetFriendsRequest(user_id)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}

func (h *Handler) acceptFriendsRequest(c *gin.Context) {
	fmt.Println("acceptFriendsRequest")

	var input inputAcceptRequest

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var userReceiver int
	userReceiver, _ = strconv.Atoi(input.UserReceiver)
	message, err := h.services.UserCommunicate.AcceptFriendsRequest(input.UserSender, userReceiver)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, message)
}

func (h *Handler) sendFriends(c *gin.Context) {
	fmt.Println("sendFriends")

	var input userId

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	var user_id int
	user_id, _ = strconv.Atoi(input.Id)
	friends, err := h.services.UserCommunicate.GetAllFriends(user_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, friends)
}

func (h *Handler) sendInvite(c *gin.Context) {
	fmt.Println("sendInvite")

	var input dataToInvite

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	message, err := h.services.UserCommunicate.SendInvite(input.UserSender, input.FriendsList, input.Id_place)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.IndentedJSON(http.StatusOK, message)
}

func (h *Handler) getRooms(c *gin.Context) {
	fmt.Println("getRooms")

	var input getRoomByUserId

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(input.Id)
	var user_id int
	user_id, _ = strconv.Atoi(input.Id)
	rooms, err := h.services.UserCommunicate.GetRooms(user_id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, rooms)
}
