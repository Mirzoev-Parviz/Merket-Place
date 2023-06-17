package repository

import (
	"market_place/models"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUser(login, password string) (models.User, error)
}

type User interface {
	CheckLogin(login string) (bool, error)
	UpdateUser(id int, user models.User) error
	DeactivateUser(id int) error
}

type Category interface {
	CreateNewCategory(categ models.Category) (int, error)
	GetAllCategories() ([]models.Category, error)
	GetCategoryProducts(id int) ([]models.Product, error)
}

type Product interface {
	CreateProduct(userId int, product models.Product) (int, error)
	GetProduct(id int) (models.Product, error)
	UpdateProduct(id, userId int, product models.Product) error
	DeactivateProduct(id, userid int) error
}

type Repository struct {
	Authorization
	User
	Category
	Product
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		Category:      NewCategoryPostgres(db),
		Product:       NewProductPostgres(db),
	}
}
