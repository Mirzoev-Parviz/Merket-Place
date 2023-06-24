package service

import (
	"market_place/models"
	"market_place/pkg/repository"
)

type CategoryService struct {
	repo repository.Category
}

func NewCategoryService(repo repository.Category) *CategoryService {
	return &CategoryService{repo: repo}
}

func (c *CategoryService) CreateNewCategory(category models.Category) (int, error) {
	return c.repo.CreateNewCategory(category)
}
func (c *CategoryService) GetAllCategories() ([]models.Category, error) {
	return c.repo.GetAllCategories()
}
func (c *CategoryService) GetCategoryProducts(name string) ([]models.Product, error) {
	return c.repo.GetCategoryProducts(name)
}
func (c *CategoryService) CheckCategoryName(name string) (bool, error) {
	return c.repo.CheckCategoryName(name)
}
