package handler

import (
	"market_place/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddBasketItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var product models.Product

	if err = c.BindJSON(&product); err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": userId})
}
