package handler

import (
	"errors"
	"market_place/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SignUp(c *gin.Context) {
	var input models.User

	if err := c.BindJSON(&input); err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input body"})
	}

	exist, err := h.services.CheckLogin(input.Login)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": "error while checking login"})
		return
	}

	if exist {
		c.JSON(http.StatusBadRequest, gin.H{"message": "login is already taken"})
		return
	}

	id, err := h.services.CreateUser(input)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error while creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": id})
}

func (h *Handler) SignIn(c *gin.Context) {
	var input models.SignInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid input body"})
		return
	}

	token, err := h.services.GenerateToken(input.Login, input.Password)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": "no such user exists"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) ShowID(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		h.logger.Error(err.Error())
		// c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": userId})
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("errors while parsing token")
	}

	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("error while parsing token")
	}

	return idInt, nil
}
