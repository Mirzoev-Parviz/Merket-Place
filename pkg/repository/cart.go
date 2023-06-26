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

func (c *CartPostgres) CreateCart(userID int) error {
	var cart models.Cart
	cart.UserID = userID

	if err := config.DB.Create(&cart).Error; err != nil {
		return err
	}

	return nil
}

func (c *CartPostgres) GetCart(userID int) (cart models.Cart, err error) {
	err = config.DB.Where("user_id = ? AND is_active = TRUE", userID).First(&cart).Error
	if err != nil {
		return models.Cart{}, err
	}

	return cart, nil
}

func DeactiveCart(userID int) error {
	var c CartPostgres
	cart, err := c.GetCart(userID)
	if err != nil {
		return err
	}

	cart.IsActive = false

	if err = config.DB.Save(&cart).Error; err != nil {
		return err
	}

	return nil
}
