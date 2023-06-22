package repository

import (
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
	merch, err := m.GetMerchant(id)
	if err != nil {
		return err
	}

	merch.IsActive = false

	if err = config.DB.Where("id = ?", id).Save(&merch).Error; err != nil {
		return err
	}

	return nil
}

//Adding product to merchants

func (m *MerchPostgres) AddProductToShelf(merch models.MerchantProduct) (int, error) {
	tx := config.DB.Begin()

	exist, err := ExistMerchantProduct(merch.MerchantID, merch.ProductID)
	if err != nil {
		return 0, err
	}

	if exist {
		if err := ChangeProductQuantity(merch.ProductID, merch.Quantity); err != nil {
			tx.Rollback()
			return 0, err
		}

		var test models.MerchantProduct

		if err := config.DB.Where("merchant_id = ? AND product_id = ? AND is_active = TRUE",
			merch.MerchantID, merch.ProductID).First(&test).Error; err != nil {
			tx.Rollback()
			return 0, nil
		}

		merch.Quantity += test.Quantity

		err := config.DB.Where("merchant_id = ? AND product_id = ? AND is_active = TRUE",
			merch.MerchantID, merch.ProductID).Updates(&merch).Error
		// fmt.Println(merch.MerchantID)
		if err != nil {
			return 0, err
		}

	} else {
		if err := ChangeProductQuantity(merch.ProductID, merch.Quantity); err != nil {
			tx.Rollback()
			return 0, err
		}
		if err := config.DB.Create(&merch).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return int(merch.ID), nil
}

func (m *MerchPostgres) GetMerchProduct(id int) (mp models.MerchantProduct, err error) {
	err = config.DB.Where("id = ? AND is_active = TRUE", id).First(&mp).Error
	if err != nil {
		return models.MerchantProduct{}, err
	}

	return mp, nil
}

func (m *MerchPostgres) UpdateMerchProduct(id int, merch models.MerchantProduct) error {

	/*if err := ChangeProductQuantity(merch.ProductID, merch.Quantity); err != nil {
		return err
	}*/

	err := config.DB.Where("id = ? AND is_active = TRUE", id).Updates(&merch).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *MerchPostgres) DeleteMerchProduct(id int) error {
	merch, err := m.GetMerchProduct(id)
	if err != nil {
		return err
	}

	merch.IsActive = false

	err = config.DB.Where("id = ? AND  is_active = TRUE", id).Save(&merch).Error
	if err != nil {
		return err
	}

	return nil
}

func ChangeProductQuantity(id, quantity int) error {
	tx := config.DB.Begin()

	var product models.Product
	if err := config.DB.Where("id = ? AND quantity >= ? AND is_active = TRUE", id, quantity).First(&product).Error; err != nil {
		tx.Rollback()
		return err
	}

	product.Quantity -= quantity

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
