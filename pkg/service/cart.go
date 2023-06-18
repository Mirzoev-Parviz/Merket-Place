package service

import "market_place/pkg/repository"

type CartService struct {
	repo repository.Cart
}

func NewCartService(repo repository.Cart) *CartService {
	return &CartService{repo: repo}
}
