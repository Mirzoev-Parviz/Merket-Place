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
	UpdateUser(id int, user models.User) error
	DeactivateUser(id int) error
}

type Merchant interface {
	CreateMerchant(merch models.Merchant) (int, error)
	GetMerchant(id int) (models.Merchant, error)
	UpdateMerchant(id int, merch models.Merchant) error
	DeleteMerchant(id int) error

	AddProductToShelf(merch models.MerchantProduct) (int, error)
	GetMerchProduct(id int) (models.MerchantProduct, error)
	UpdateMerchProduct(id int, merch models.MerchantProduct) error
	DeleteMerchProduct(id int) error
}

type Category interface {
	CreateNewCategory(category models.Category) (int, error)
	GetCategoryProducts(name string) ([]models.Product, error)
	GetAllCategories() ([]models.Category, error)
}

type Product interface {
	CreateProduct(userId int, product models.Product) (int, error)
	GetProduct(id int) (models.Product, error)
	UpdateProduct(id, userId int, product models.Product) error
	DeactivateProduct(id, userId int) error
}

type Basket interface {
}

type Service struct {
	Authorization
	User
	Merchant
	Category
	Product
	Basket
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		User:          NewUserService(repo),
		Merchant:      NewMerchService(repo),
		Category:      NewCategoryService(repo),
		Product:       NewProductService(repo),
		Basket:        NewBasketService(repo),
	}
}
