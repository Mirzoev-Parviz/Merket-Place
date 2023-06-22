package repository

import (
	"market_place/config"
	"market_place/models"

	"gorm.io/gorm"
)

type AuthPostgres struct {
	db *gorm.DB
}

func NewAuthPostgres(db *gorm.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (a *AuthPostgres) CreateUser(user models.User) (int, error) {
	// user.IsActive = true
	if err := config.DB.Create(&user).Error; err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (a *AuthPostgres) CheckUser(login, password string) (user models.User, err error) {
	err = config.DB.Where("login = ? AND password = ? AND is_active = TRUE", login, password).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (a *AuthPostgres) CheckMerch(login, password string) (merch models.Merchant, err error) {
	err = config.DB.Where("login = ? AND password = ? AND is_active = TRUE", login, password).First(&merch).Error
	if err != nil {
		return models.Merchant{}, err
	}

	return merch, nil
}
