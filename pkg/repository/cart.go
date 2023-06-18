package repository

import (
	"market_place/config"
	"market_place/models"

	"gorm.io/gorm"
)

type CartPostgres struct {
	db *gorm.DB
}

func NewCartPostgres(db *gorm.DB) *CartPostgres {
	return &CartPostgres{db: db}
}

func (c *CartPostgres) AddCartItem(userId, productId int, item models.CartItem) error {
	var cart models.Cart
	err := config.DB.Where("user_id = ? AND product_id = ?", userId, productId).Find(&cart).Error
	if err != nil {
		return err
	}
	return nil
}
