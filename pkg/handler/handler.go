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
		category.GET("/:name", h.ShowCategoryProducts)
		category.GET("/", h.ShowAllCategories)
		// category.PUT("/:id")
		// category.DELETE("/:id")
	}

	api := router.Group("api")
	{

		user := api.Group("user")
		{
			// user.GET("/:id")
			user.PUT("/:id", h.UpdateUser)
			user.DELETE("/:id", h.DeleteUser)

			product := api.Group("product")
			{
				product.POST("/", h.UserIdentity, h.AddProduct)
				product.GET("/:id", h.ShowProduct)
				product.PUT("/:id", h.UserIdentity, h.UpdateProduct)
				product.DELETE("/:id", h.UserIdentity, h.DeleteProduct)
			}

		}

		merchant := router.Group("merchant")
		{
			merchant.POST("/sign-up", h.MerchSignUp)
			merchant.POST("/sign-in", h.MerchSignIn)
			// merchant.GET("/id", h.MerchIdentity, h.ShowID)
			merchant.GET("/:id", h.GetMerchant)
			merchant.PUT("/:id", h.MerchIdentity, h.UpdateMerchant)
			merchant.DELETE("/:id", h.MerchIdentity, h.DeleteMerchant)

			product := merchant.Group("product", h.MerchIdentity)
			{
				product.POST("/", h.AddProductToShelf)
				product.GET("/:id", h.GetMerchantProduct)
				product.PUT("/:id", h.UpdateMerchProduct)
				product.DELETE("/:id", h.DeleteMerchProduct)
			}

		}

	}
	return router
}
