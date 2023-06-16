package repository

import (
	"market_place/config"
	"market_place/models"

	"gorm.io/gorm"
)

type UserPostgres struct {
	db *gorm.DB
}

func NewUserPostgres(db *gorm.DB) *UserPostgres {
	return &UserPostgres{db: db}
}

func (u *UserPostgres) CheckLogin(login string) (bool, error) {
	var user models.User
	err := config.DB.Where("login = ? AND is_active = TRUE", login).Find(&user).Error
	if err != nil {
		return false, err
	}

	if user.Login == login {
		return true, nil
	}

	return false, nil
}

func (u *UserPostgres) UpdateUser(id int, user models.User) error {
	err := config.DB.Where("id = ? AND is_active = TRUE", id).Updates(&user).Error
	if err != nil {
		return err
	}

	return nil
}
