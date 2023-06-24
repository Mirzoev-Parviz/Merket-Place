package repository

import (
	"market_place/config"
	"market_place/models"
)

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
		if err != nil {
			return 0, err
		}

	} else {
		if err := ChangeProductQuantity(merch.ProductID, merch.Quantity); err != nil {
			tx.Rollback()
			return 0, err
		}

		product, err := GetProductInfo(merch.ProductID)
		if err != nil {
			return 0, err
		}

		merch.Category = product.Category
		merch.Title = product.Title
		merch.Description = product.Description
		merch.InStock = BeforeSave(merch.Quantity, merch.InStock)

		if merch.Price == 0 {
			merch.Price = product.Price
		}

		if err := config.DB.Create(&merch).Error; err != nil {
			tx.Rollback()
			return 0, err
		}
	}
	tx.Commit()
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

func (m *MerchPostgres) GetFilterdProducts(input models.Filter) (products []models.MerchantProduct, err error) {
	query := config.DB.Model(&models.MerchantProduct{})

	if input.Category != "" {
		query = query.Where("category = ?", input.Category)
	}

	if input.MerchantID != 0 {
		query = query.Where("merchant_id = ?", input.MerchantID)
	}

	if input.Price != 0 {
		query = query.Where("price >= ?", input.Price)
	}

	err = query.Find(&products).Error
	if err != nil {
		return []models.MerchantProduct{}, err
	}

	return products, nil
}
