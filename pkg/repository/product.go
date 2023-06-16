package repository

import (
	"market_place/config"
	"market_place/models"

	"gorm.io/gorm"
)

type ProductPostgres struct {
	db *gorm.DB
}

func NewProductPostgres(db *gorm.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (p *ProductPostgres) CreateProduct(userId int, product models.Product) (int, error) {
	if err := config.DB.Create(&product).Error; err != nil {
		return 0, nil
	}

	SaveProduct(product.Category)

	return product.Id, nil
}

func (p *ProductPostgres) GetProduct(id int) (product models.Product, err error) {
	err = config.DB.Where("id = ?  AND is_active = TRUE", id).Find(&product).Error
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (p *ProductPostgres) UpdateProduct(id, userId int, product models.Product) error {
	err := config.DB.Where("id = ? AND user_id = ? AND is_active = TRUE", id, userId).Updates(&product).Error
	if err != nil {
		return err
	}

	return nil
}
