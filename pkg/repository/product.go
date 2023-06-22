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

	return product.ID, nil
}

func (p *ProductPostgres) GetProduct(id int) (product models.Product, err error) {
	err = config.DB.Where("id = ?  AND is_active = TRUE", id).Find(&product).Error
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (p *ProductPostgres) UpdateProduct(id, userId int, product models.Product) error {
	err := config.DB.Where("id = ?  AND is_active = TRUE", id).Updates(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductPostgres) DeactivateProduct(id, userId int) error {
	var product models.Product
	err := config.DB.Where("id = ? AND is_active = TRUE", id).First(&product).Error
	if err != nil {
		return err
	}

	var category models.Category
	if err = config.DB.Where("name = ?", product.Category).First(&category).Error; err != nil {
		return err
	}

	category.Amount--
	product.IsActive = false

	if err = config.DB.Save(&product).Error; err != nil {
		return err
	}

	return nil
}
