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

	exist, err := h.services.CheckCategoryName(input.Name)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	if exist {
		h.logger.Error("category is already exist")
		c.JSON(http.StatusBadRequest, gin.H{"message": "category is already exist"})
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

	category, err := h.services.GetCategoryByID(id)
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, category)
}

func (h *Handler) ShowAllCategories(c *gin.Context) {
	categories, err := h.services.GetAllCategories()
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
