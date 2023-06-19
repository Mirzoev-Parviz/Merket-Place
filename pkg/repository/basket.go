package repository

import (
	"market_place/config"
	"market_place/models"

	"gorm.io/gorm"
)

type BasketPostgres struct {
	db *gorm.DB
}

func NewBasketPostgres(db *gorm.DB) *BasketPostgres {
	return &BasketPostgres{db: db}
}

func (b *BasketPostgres) AddBasketItem(userId int, product models.Product) error {
	var basket models.Basket
	err := config.DB.Where("userId = ?", userId).First(&basket).Error
	if err != nil {
		return err
	}

	basket.ProductId = append(basket.ProductId, byte(product.Id))
	basket.TotalSum += product.Price

	if err = config.DB.Save(&basket).Error; err != nil {
		return err
	}

	return nil
}
