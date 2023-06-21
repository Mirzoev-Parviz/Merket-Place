package repository

import (
	"market_place/models"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	CheckUser(login, password string) (models.User, error)
}

type User interface {
	GetUser(login string) (bool, error)
	UpdateUser(id int, user models.User) error
	DeactivateUser(id int) error
}

type Merchant interface {
	CreateMerchant(merch models.Merchant) (int, error)
}

type Category interface {
	CreateNewCategory(categ models.Category) (int, error)
	GetAllCategories() ([]models.Category, error)
	GetCategoryProducts(name string) ([]models.Product, error)
}

type Product interface {
	CreateProduct(userId int, product models.Product) (int, error)
	GetProduct(id int) (models.Product, error)
	UpdateProduct(id, userId int, product models.Product) error
	DeactivateProduct(id, userid int) error

	AddProductToShelf(m_id, id, quantiry int) (int, error)
}

type Basket interface {
	AddBasketItem(userId int, product models.Product) error
}

type Repository struct {
	Authorization
	User
	Merchant
	Category
	Product
	Basket
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		Merchant:      NewMerchPostgres(db),
		Category:      NewCategoryPostgres(db),
		Product:       NewProductPostgres(db),
		Basket:        NewBasketPostgres(db),
	}
}
