package repository

import (
	"fmt"
	"market_place/config"
	"market_place/models"

	"gorm.io/gorm"
)

type MerchPostgres struct {
	db *gorm.DB
}

func NewMerchPostgres(db *gorm.DB) *MerchPostgres {
	return &MerchPostgres{db: db}
}

func (m *MerchPostgres) CreateMerchant(merch models.Merchant) (id int, err error) {
	if err = config.DB.Create(&merch).Error; err != nil {
		return 0, err
	}

	return merch.ID, nil
}

func (m *MerchPostgres) GetMerchant(id int) (merch models.Merchant, err error) {
	err = config.DB.Where("id = ? AND is_active = TRUE", id).First(&merch).Error
	if err != nil {
		return models.Merchant{}, err
	}

	return merch, nil
}

func (m *MerchPostgres) UpdateMerchant(id int, merch models.Merchant) error {
	err := config.DB.Where("id = ? AND is_active = TRUE", id).Updates(&merch).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *MerchPostgres) DeleteMerchant(id int) error {
	return config.DB.
		Model(&models.Merchant{}).
		Where("id = ? AND is_active = TRUE").
		Update("is_active", false).Error
}

//Adding product to merchants

func (m *MerchPostgres) SearchMerchProduct(query string) (products []models.MerchantProduct, err error) {
	productIdes, err := GetProductByQuery(query)
	if err != nil {
		return []models.MerchantProduct{}, err
	}

	err = config.DB.Where("product_id IN (?) AND is_active = TRUE", productIdes).Find(&products).Error
	if err != nil {
		return []models.MerchantProduct{}, err
	}

	return products, nil
}

func ChangeProductQuantity(id, quantity int) error {
	tx := config.DB.Begin()

	var product models.Product
	if err := config.DB.Where("id = ? AND quantity >= ? AND is_active = TRUE", id, quantity).First(&product).Error; err != nil {
		tx.Rollback()
		return err
	}

	product.Quantity -= quantity
	product.InStock = BeforeSave(product.Quantity, product.InStock)

	if err := config.DB.Where("id = ? AND is_active = TRUE", product.ID).Save(&product).Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func ExistMerchantProduct(merchantID, productID int) (bool, error) {
	var merchantP models.MerchantProduct
	err := config.DB.Where("merchant_id = ? AND product_id = ? AND is_active = TRUE",
		merchantID, productID).Find(&merchantP).Error
	if err != nil {
		return true, err
	}

	if merchantP.ID == 0 {
		return false, nil
	}

	if merchantP.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetProductPrice(id int) (price float64, err error) {
	err = config.DB.Table("products").Select("price").Where("id = ? AND is_active = TRUE", id).Find(&price).Error
	if err != nil {
		return 0, err
	}

	fmt.Printf("price is: %v\n", price)

	return price, nil
}
