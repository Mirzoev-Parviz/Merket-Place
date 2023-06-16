package repository

import (
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

	return category.Id, nil
}

func (c *CategoryPostgres) GetCategory(name string) (category models.Category, err error) {
	err = config.DB.Where("name = ?", name).Find(&category).Error
	if err != nil {
		return models.Category{}, err
	}

	return category, nil
}

func (c *CategoryPostgres) GetCategoryProducts(id int) (productList []models.Product, err error) {
	var categ models.Category
	err = config.DB.Where("id = ?", id).Find(&categ).Error
	if err != nil {
		return []models.Product{}, err
	}

	err = config.DB.Where("category = ? AND is_active = TRUE", categ.Name).Find(&productList).Error
	if err != nil {
		return []models.Product{}, err
	}

	return productList, nil
}

func (c *CategoryPostgres) GetAllCategories() (categories []models.Category, err error) {
	err = config.DB.Where("").Find(&categories).Error
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

	return nil
}
