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
