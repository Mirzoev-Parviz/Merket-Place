package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "user_id"
)

func (h *Handler) UserIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "empty auth header"})
		return
	}

	headerParts := strings.Split(header, " ")

	if len(headerParts) != 2 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "invalid auth header"})
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.Set(userCtx, userId)
}
