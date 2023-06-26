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

		var quantity int
		if err := config.DB.Model(&models.MerchantProduct{}).Select("quantity").
			Where("merchant_id = ? AND product_id = ? AND is_active = TRUE", merch.MerchantID, merch.ProductID).
			Scan(&quantity).Error; err != nil {
			return 0, err
		}

		merch.Quantity += quantity
		merch.InStock = BeforeSave(merch.Quantity, merch.InStock)

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

func (m *MerchPostgres) GetAllMerchProducts() (products []models.MerchantProduct, err error) {
	err = config.DB.Where("is_active = TRUE").Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (m *MerchPostgres) UpdateMerchProduct(id int, merch models.MerchantProduct) error {
	err := config.DB.Where("id = ? AND is_active = TRUE", id).Updates(&merch).Error
	if err != nil {
		return err
	}

	return nil
}

func (m *MerchPostgres) DeleteMerchProduct(id int) error {
	err := config.DB.Model(&models.MerchantProduct{}).
		Where("id = ? AND is_active = TRUE", id).Update("is_active", false).Error
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

func GetRecomenndetIds() (ids []int, err error) {
	err = config.DB.Model(&models.MerchantProduct{}).
		Order("RANDOM()").Limit(10).Pluck("id", &ids).Error
	if err != nil {
		return nil, err
	}

	return ids, nil
}

func (m *MerchPostgres) GetRecommendetProducts() (products []models.MerchantProduct, err error) {
	ids, err := GetRecomenndetIds()
	if err != nil {
		return nil, err
	}

	err = config.DB.Where("id IN (?) AND is_active = TRUE", ids).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return products, nil
}
