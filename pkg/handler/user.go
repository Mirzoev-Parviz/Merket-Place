package handler

import (
	"market_place/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UpdateUser(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var input models.User

	if err = c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	exist, err := h.services.CheckLogin(input.Login)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if exist && input.Login != "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "use is already taken"})
		return
	}

	if err = h.services.UpdateUser(id, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updaed successfully"})
}

func getId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, err
	}

	return id, nil
}
