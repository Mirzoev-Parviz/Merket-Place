package service

import (
	"market_place/models"
	"market_place/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GenerateToken(login, password string) (string, error)
	ParseToken(accessToken string) (int, error)
}
type User interface {
	CheckLogin(login string) (bool, error)
	UpdateUser(id int, user models.User) error
	DeactivateUser(id int) error
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

type Service struct {
	Authorization
	User
	Category
	Product
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo),
		User:          NewUserService(repo),
		Category:      NewCategoryService(repo),
		Product:       NewProductService(repo),
	}
}
