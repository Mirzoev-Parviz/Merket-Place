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
		category.GET("/", h.ShowAllCategories)
		// category.PUT("/:id")
		// category.DELETE("/:id")
	}

	api := router.Group("api")
	{

		user := api.Group("user")
		{
			user.GET("/:id", h.GetUser)
			user.PUT("/", h.UserIdentity, h.UpdateUser)
			user.DELETE("/", h.UserIdentity, h.DeleteUser)
			user.POST("/", h.UserIdentity, h.BuyIt)
			user.GET("/history", h.UserIdentity, h.History)
			user.POST("/review", h.UserIdentity, h.CreateReview)
			user.POST("/later", h.UserIdentity, h.Later)
			user.DELETE("/later", h.UserIdentity, h.DeleteLater)

			cart := api.Group("cart", h.UserIdentity)
			{
				cart.POST("/item", h.AddCartItem)
			}

		}

		product := api.Group("product")
		{
			product.POST("/", h.AddProduct)
			product.GET("/:id", h.ShowProduct)
			product.PUT("/:id", h.UpdateProduct)
			product.DELETE("/:id", h.DeleteProduct)
		}

		merchant := api.Group("merchant")
		{
			merchant.POST("/sign-up", h.MerchSignUp)
			merchant.POST("/sign-in", h.MerchSignIn)
			merchant.GET("/:id", h.GetMerchant)
			merchant.PUT("/:id", h.MerchIdentity, h.UpdateMerchant)
			merchant.DELETE("/:id", h.MerchIdentity, h.DeleteMerchant)

			product := merchant.Group("product")
			{
				product.POST("/", h.MerchIdentity, h.AddProductToShelf)
				product.GET("/:id", h.MerchIdentity, h.GetMerchantProduct)
				product.GET("/", h.GetAllMerchantProducts)
				product.PUT("/:id", h.MerchIdentity, h.UpdateMerchProduct)
				product.DELETE("/:id", h.MerchIdentity, h.DeleteMerchProduct)

				product.POST("/search", h.SearchMerchProducts)
				product.POST("/filter", h.GetFilterdProducts)
			}
		}

	}
	return router
}
