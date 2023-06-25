package service

import (
	"market_place/models"
	"market_place/pkg/repository"
)

type CartService struct {
	repo repository.Cart
}

func NewCartService(repo repository.Cart) *CartService {
	return &CartService{repo: repo}
}

func (c *CartService) CreateCart(userID int) error {
	return c.repo.CreateCart(userID)
}

func (c *CartService) AddCartItem(userID int, item models.CartItem) (int, error) {
	return c.repo.AddCartItem(userID, item)
}

func (c *CartService) BuyIt(userID int) error {
	return c.repo.BuyIt(userID)
}

func (c *CartService) History(userID int) (cartItems []models.CartItem, err error) {
	return c.repo.History(userID)
}

func (c *CartService) Later(userID, cartItemID int) error {
	return c.repo.Later(userID, cartItemID)
}
func (c *CartService) DeleteLater(userID, cartItemID int) error {
	return c.repo.DeleteLater(userID, cartItemID)
}
