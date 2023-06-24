package service

import (
	"market_place/models"
	"market_place/pkg/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (p *ProductService) CreateProduct(product models.Product) (int, error) {
	return p.repo.CreateProduct(product)
}

func (p *ProductService) GetProduct(id int) (models.Product, error) {
	return p.repo.GetProduct(id)
}

func (p *ProductService) UpdateProduct(id, userId int, product models.Product) error {
	return p.repo.UpdateProduct(id, userId, product)
}
func (p *ProductService) DeactivateProduct(id, userId int) error {
	return p.repo.DeactivateProduct(id, userId)
}
