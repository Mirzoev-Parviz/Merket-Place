package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "user_id"
	merchCtx            = "merch_id"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		h.logger.Error("failed to get header")
		c.JSON(http.StatusBadRequest, gin.H{"message": "empty auth header"})
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		h.logger.Error("invalid format header")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid auth header"})
		return
	}

	userId, err := h.services.Authorization.ParseUserToken(headerParts[1])
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.Set(userCtx, userId)
}

func (h *Handler) MerchIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		h.logger.Error("failed to get header")
		c.JSON(http.StatusBadRequest, gin.H{"message": "empty auth header"})
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		h.logger.Error("invalid format header")
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid auth header"})
		return
	}

	userId, err := h.services.Authorization.ParseMerchantToken(headerParts[1])
	if err != nil {
		h.logger.Error(err.Error())
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.Set(merchCtx, userId)
}
