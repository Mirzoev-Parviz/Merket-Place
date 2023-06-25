package repository

import (
	"errors"
	"market_place/config"
	"market_place/models"

	"gorm.io/gorm"
)

type CartPostgres struct {
	db *gorm.DB
}

func NewCartPostgres(db *gorm.DB) *CartPostgres {
	return &CartPostgres{db: db}
}

func (c *CartPostgres) CreateCart(userID int) error {
	var cart models.Cart
	cart.UserID = userID

	if err := config.DB.Create(&cart).Error; err != nil {
		return err
	}

	return nil
}

func (c *CartPostgres) GetCart(userID int) (cart models.Cart, err error) {
	err = config.DB.Where("user_id = ? AND is_active = TRUE", userID).First(&cart).Error
	if err != nil {
		return models.Cart{}, err
	}

	return cart, nil
}

func DeactiveCart(userID int) error {
	var c CartPostgres
	cart, err := c.GetCart(userID)
	if err != nil {
		return err
	}

	cart.IsActive = false

	if err = config.DB.Save(&cart).Error; err != nil {
		return err
	}

	return nil
}

func (c *CartPostgres) AddCartItem(userID int, item models.CartItem) (id int, err error) {
	tx := config.DB.Begin()

	item.CartID, err = GetCartID(userID)
	if err != nil {
		return 0, err
	}

	if err = config.DB.Create(&item).Error; err != nil {
		tx.Rollback()
		return 0, err
	}
	tx.Commit()

	return item.ID, nil
}

func (c *CartPostgres) BuyIt(userID int) error {
	cartID, err := GetCartID(userID)
	if err != nil {
		return err
	}
	var items []models.CartItem
	err = config.DB.Where("cart_id = ? AND is_active = TRUE", cartID).Find(&items).Error
	if err != nil {
		return err
	}

	if len(items) == 0 {
		return errors.New("cart is empty")
	}

	for _, item := range items {
		err = ChangeMerchantQuantity(item.MerchantID, item.ProductID, item.Quantity)
		if err != nil {
			return err
		}

		err = config.DB.Model(&models.CartItem{}).Where("id = ? AND cart_id = ? AND is_active = TRUE",
			item.ID, cartID).UpdateColumn("is_active", false).Error
		if err != nil {
			return err
		}

		err = IncreaseTotalOrders(item.ProductID)
		if err != nil {
			return err
		}
	}
	return nil
}

func GetCartID(userID int) (int, error) {
	var cart models.Cart
	err := config.DB.Where("user_id = ?", userID).First(&cart).Error
	if err != nil {
		return 0, err
	}

	return cart.ID, nil
}

func ChangeMerchantQuantity(merchantID, productID, quantity int) error {
	var merchProd models.MerchantProduct
	err := config.DB.Where("product_id = ? AND merchant_id = ? AND is_active = TRUE",
		productID, merchantID).First(&merchProd).Error
	if err != nil {
		return err
	}

	tx := config.DB.Begin()
	merchProd.Quantity -= quantity
	if merchProd.Quantity < 0 {
		merchProd.Quantity = 0
	}
	merchProd.InStock = BeforeSave(merchProd.Quantity, merchProd.InStock)

	if err = config.DB.Save(&merchProd).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil

}

func (c *CartPostgres) History(userID int) (cartItems []models.CartItem, err error) {
	cartID, err := GetCartID(userID)
	if err != nil {
		return []models.CartItem{}, err
	}

	err = config.DB.Where("cart_id = ? AND is_active = FALSE", cartID).Find(&cartItems).Error
	if err != nil {
		return []models.CartItem{}, err
	}

	return cartItems, nil
}

func IncreaseTotalOrders(productID int) error {
	return config.DB.Model(&models.MerchantProduct{}).
		Where("id = ?", productID).
		UpdateColumn("total_orders", gorm.
			Expr("total_orders + ?", 1)).Error
}
