package repository

import (
	"market_place/models"

	"gorm.io/gorm"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	CheckUser(login, password string) (models.User, error)
	CheckMerch(login, password string) (models.Merchant, error)
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
	UpdateMerchant(id int, merch models.Merchant) error
	DeleteMerchant(id int) error

	AddProductToShelf(merch models.MerchantProduct) (int, error)
	GetMerchProduct(id int) (models.MerchantProduct, error)
	UpdateMerchProduct(id int, merch models.MerchantProduct) error
	DeleteMerchProduct(id int) error

	SearchMerchProduct(query string) ([]models.MerchantProduct, error)
}

type Category interface {
	CreateNewCategory(categ models.Category) (int, error)
	GetAllCategories() ([]models.Category, error)
	GetCategoryProducts(name string) ([]models.Product, error)
}

type Product interface {
	CreateProduct(product models.Product) (int, error)
	GetProduct(id int) (models.Product, error)
	UpdateProduct(id, userId int, product models.Product) error
	DeactivateProduct(id, userid int) error
}

type Cart interface {
	CreateCart(userID int) error
	AddCartItem(userID int, item models.CartItem) (int, error)
	BuyIt(cartID int) error
}

type Repository struct {
	Authorization
	User
	Merchant
	Category
	Product
	Cart
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		User:          NewUserPostgres(db),
		Merchant:      NewMerchPostgres(db),
		Category:      NewCategoryPostgres(db),
		Product:       NewProductPostgres(db),
		Cart:          NewCartPostgres(db),
	}
}
