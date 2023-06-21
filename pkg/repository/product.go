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
	err := config.DB.Where("id = ? AND user_id = ? AND is_active = TRUE", id, userId).Updates(&product).Error
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductPostgres) DeactivateProduct(id, userId int) error {
	var product models.Product
	err := config.DB.Where("id = ? AND user_id = ? AND is_active = TRUE", id, userId).First(&product).Error
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

//Adding product to merchants

func (p *ProductPostgres) AddProductToShelf(m_id, id, quantity int) (int, error) {
	tx := config.DB.Begin()

	var product models.Product

	err := config.DB.Where("id = ? AND quantity >= ? AND  is_active = TRUE", id, quantity).Find(&product).Error
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	product.Quantity -= quantity

	var merch models.MerchantProduct
	merch.MerchantID = m_id
	merch.ProductID = product.ID

	if err = config.DB.Create(&merch).Error; err != nil {
		return 0, err
	}

	return int(merch.ID), nil
}
