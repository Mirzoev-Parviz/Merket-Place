package repository

import (
	"fmt"
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

func (c *CartPostgres) AddCartItem(userID int, item models.CartItem) (id int, err error) {
	tx := config.DB.Begin()

	item.CartID, err = GetCartID(userID)
	if err != nil {
		return 0, err
	}

	/*err = ChangeMerchantQuantity(item.MerchantID, item.ProductID, item.Quantity)
	if err != nil {
		tx.Rollback()
		return 0, err
	}*/

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
	err = config.DB.Where("cart_id = ?", cartID).Find(&items).Error
	if err != nil {
		return err
	}

	for _, x := range items {
		err = ChangeMerchantQuantity(x.MerchantID, x.ProductID, x.Quantity)
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
	fmt.Printf("merch_id:%v\nproduct_id:%v\n", merchantID, productID)
	err := config.DB.Where("product_id = ? AND merchant_id = ? AND is_active = TRUE", productID, merchantID).First(&merchProd).Error
	if err != nil {
		return err
	}

	fmt.Println(merchProd)

	tx := config.DB.Begin()

	fmt.Println(merchProd.ID)
	merchProd.Quantity -= quantity

	if err = config.DB.Save(&merchProd).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil

}
