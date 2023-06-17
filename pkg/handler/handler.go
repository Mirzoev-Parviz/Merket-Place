package handler

import (
	"market_place/pkg/logging"
	"market_place/pkg/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
	logger   logging.Logger
}

func NewHandler(services *service.Service, logger logging.Logger) *Handler {
	return &Handler{
		services: services,
		logger:   logger,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "it's working"})
	})

	auth := router.Group("auth")
	{
		auth.POST("/sign-up", h.SignUp)
		auth.POST("/sign-in", h.SignIn)
		auth.GET("/id", h.UserIdentity, h.ShowID)
	}

	category := router.Group("category")
	{
		category.POST("/", h.NewCategory)
		category.GET("/:id", h.ShowCategoryProducts)
		// category.PUT("/:id")
		// category.DELETE("/:id")
	}

	api := router.Group("api")
	{

		product := api.Group("product")
		{
			product.POST("/", h.UserIdentity, h.AddProduct)
			product.GET("/:id", h.ShowProduct)
			product.PUT("/:id", h.UserIdentity, h.UpdateProduct)
			// product.DELETE("/:id")
		}

		user := router.Group("user")
		{
			// user.GET("/:id")
			user.PUT("/:id", h.UpdateUser)
			user.DELETE("/:id", h.DeleteUser)
		}

	}
	return router
}
