package repository

import (
	"errors"
	"market_place/config"
	"market_place/models"

	"gorm.io/gorm"
)

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

		exist, err := CheckCartItemID(item.ID)
		if err != nil {
			return err
		}

		if !exist {
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

func GetCartItems(userID int) (items []models.CartItem, err error) {
	id, err := GetCartID(userID)
	if err != nil {
		return []models.CartItem{}, err
	}

	err = config.DB.Where("cart_id = ?", id).Find(items).Error
	if err != nil {
		return []models.CartItem{}, err
	}

	return items, nil
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

func (c *CartPostgres) Later(userID, cartItemID int) error {
	var later models.Later

	cartID, err := GetCartID(userID)
	if err != nil {
		return err
	}

	later.CartID = cartID
	later.CartItemID = cartItemID

	if err := config.DB.Create(&later).Error; err != nil {
		return err
	}

	return nil
}

func (c *CartPostgres) DeleteLater(userID, cartItemID int) error {

	cartID, err := GetCartID(userID)
	if err != nil {
		return err
	}

	err = config.DB.Where("cart_id = ? AND cart_item_id = ?",
		cartID, cartItemID).Delete(&models.Later{}).Error
	if err != nil {
		return err
	}

	return nil
}

func CheckCartItemID(id int) (bool, error) {
	var cartItemID int
	err := config.DB.Model(&models.Later{}).Select("cart_item_id").
		Where("cart_item_id = ?", id).Scan(&cartItemID).Error
	if err != nil {
		return true, err
	}

	if cartItemID == id {
		return true, nil
	}

	return false, nil
}
