package handler

import (
	"market_place/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) NewCategory(c *gin.Context) {
	var input models.Category

	if err := c.BindJSON(&input); err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := h.services.CreateNewCategory(input)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": id})
}

func (h *Handler) ShowCategoryProducts(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}

	productList, err := h.services.GetCategoryProducts(id)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, productList)
}
