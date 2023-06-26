package repository

import (
	"errors"
	"market_place/config"
	"market_place/models"

	"gorm.io/gorm"
)

type CategoryPostgres struct {
	db *gorm.DB
}

func NewCategoryPostgres(db *gorm.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (c *CategoryPostgres) CreateNewCategory(category models.Category) (int, error) {
	if err := config.DB.Create(&category).Error; err != nil {
		return 0, err
	}

	return category.ID, nil
}

func (c *CategoryPostgres) GetCategory(name string) (category models.Category, err error) {
	err = config.DB.Where("name = ?", name).Find(&category).Error
	if err != nil {
		return models.Category{}, err
	}

	if category.ID == 0 {
		return models.Category{}, errors.New("category not found")
	}

	return category, nil
}

func (c *CategoryPostgres) GetCategoryByID(id int) (category models.Category, err error) {
	err = config.DB.Where("id = ? AND is_active = TRUE").First(&category).Error
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (c *CategoryPostgres) GetAllCategories() (categories []models.Category, err error) {
	err = config.DB.Find(&categories).Error
	if err != nil {
		return []models.Category{}, err
	}

	return categories, nil
}

func SaveProduct(name string) error {
	var catP CategoryPostgres
	c, err := catP.GetCategory(name)
	if err != nil {
		return err
	}

	c.Amount++

	if err = config.DB.Save(&c).Error; err != nil {
		return err
	}

	if err = SaveTotalAmount(c.ParentID); err != nil {
		return err
	}

	return nil
}

func SaveTotalAmount(id int) error {
	return config.DB.Model(&models.Category{}).Where("id = ?", id).
		Update("amount", gorm.Expr("amount + 1")).Error
}

func (c *CategoryPostgres) CheckCategoryName(name string) (bool, error) {
	var category models.Category
	err := config.DB.Where("name = ? AND is_active = TRUE", name).Find(&category).Error
	if err != nil {
		return true, err
	}

	if category.Name != "" {
		return true, nil
	}

	return false, nil
}
