package service

import (
	"market_place/models"
	"market_place/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateUserToken(login, password string) (string, error)
	GenerateMerchantToken(login, password string) (string, error)
	ParseUserToken(accessToken string) (int, error)
	ParseMerchantToken(accessToken string) (int, error)
}
type User interface {
	CheckLogin(login string) (bool, error)
	GetUser(userID int) (models.User, error)
	UpdateUser(id int, user models.User) error
	DeactivateUser(id int) error
}

type Merchant interface {
	CreateMerchant(merch models.Merchant) (int, error)
	GetMerchant(id int) (models.Merchant, error)
	GetAllMerchantProducts() ([]models.MerchantProduct, error)
	UpdateMerchant(id int, merch models.Merchant) error
	DeleteMerchant(id int) error

	AddProductToShelf(merch models.MerchantProduct) (int, error)
	GetMerchProduct(id int) (models.MerchantProduct, error)
	UpdateMerchProduct(id int, merch models.MerchantProduct) error
	DeleteMerchProduct(id int) error

	SearchMerchProduct(query string) ([]models.MerchantProduct, error)
	GetFilterdProducts(input models.Filter) ([]models.MerchantProduct, error)
	GetRecommendetProducts() (products []models.MerchantProduct, err error)
}

type Category interface {
	CreateNewCategory(category models.Category) (int, error)
	GetCategoryByID(id int) (category models.Category, err error)
	GetAllCategories() ([]models.Category, error)
	CheckCategoryName(name string) (bool, error)
}

type Product interface {
	CreateProduct(product models.Product) (int, error)
	GetProduct(id int) (models.Product, error)
	UpdateProduct(id, userId int, product models.Product) error
	DeactivateProduct(id, userId int) error
}

type Cart interface {
	CreateCart(userId int) error
	AddCartItem(userID int, item models.CartItem) (int, error)
	BuyIt(userID int) error
	History(userID int) (cartItems []models.CartItem, err error)
	Later(userID, cartItemID int) error
	DeleteLater(userID, cartItemID int) error
}

type Review interface {
	CreateReview(review models.Review) error
	CalculateProductRating(productID int) error
	GetReview(id int) (review models.Review, err error)
	GetAllReviews() (reviews []models.Review, err error)
	GetMerchantProductReview(merchantID, merchantProductID int) (reviews []models.Review, err error)
	UpdateReview(input models.Review) error
	DeleteReview(id, userID int) error
}

type Service struct {
	Authorization
	User
	Merchant
	Category
	Product
	Cart
	Review
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		User:          NewUserService(repo),
		Merchant:      NewMerchService(repo),
		Category:      NewCategoryService(repo),
		Product:       NewProductService(repo),
		Cart:          NewCartService(repo),
		Review:        NewReviewService(repo),
	}
}
