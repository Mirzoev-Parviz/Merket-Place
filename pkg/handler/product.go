package handler

import (
	"market_place/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) AddProduct(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var input models.Product

	if err := c.BindJSON(&input); err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, err := h.services.CreateProduct(userId, input)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

func (h *Handler) ShowProduct(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	product, err := h.services.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if product.Id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
		return
	}

	c.JSON(http.StatusOK, product)
}

func (h *Handler) UpdateProduct(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	id, err := getId(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	var input models.Product

	if err = c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err = h.services.UpdateProduct(id, userId, input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "successfully updated"})
}

func (h *Handler) DeleteProduct(c *gin.Context) {
	id, err := getId(c)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	userId, err := getUserId(c)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if err = h.services.DeactivateProduct(id, userId); err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}
