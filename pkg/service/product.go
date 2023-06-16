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

func (p *ProductService) CreateProduct(userId int, product models.Product) (int, error) {
	product.UserId = userId
	product.IsActive = true
	return p.repo.CreateProduct(userId, product)
}

func (p *ProductService) GetProduct(id int) (models.Product, error) {
	return p.repo.GetProduct(id)
}

func (p *ProductService) UpdateProduct(id, userId int, product models.Product) error {
	return p.repo.UpdateProduct(id, userId, product)
}
