package service

import (
	"market_place/models"
	"market_place/pkg/repository"
)

type BasketService struct {
	repo repository.Basket
}

func NewBasketService(repo repository.Basket) *BasketService {
	return &BasketService{repo: repo}
}

func (b *BasketService) AddProduct(userId int, product models.Product) error {
	return b.repo.AddBasketItem(userId, product)
}
