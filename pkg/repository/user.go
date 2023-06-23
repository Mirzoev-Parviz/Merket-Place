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

func (u *UserPostgres) GetUser(userID int) (user models.User, err error) {
	err = config.DB.Where("id = ? AND is_active = TRUE", userID).First(&user).Error
	if err != nil {
		return models.User{}, err
	}

	return user, nil
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

func (u *UserPostgres) DeactivateUser(id int) error {
	var user models.User

	if err := config.DB.Where("id = ? AND is_active = TRUE").First(&user).Error; err != nil {
		return err
	}

	user.IsActive = false

	if err := config.DB.Save(&user).Error; err != nil {
		return err
	}

	return nil
}

/*
func (u *UserPostgres) DeactivateUser(id int) error {
	tx := config.DB.Begin()

	var user models.User
	if err := tx.Where("id = ? AND is_active = TRUE", id).First(&user).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to find user with id %d: %w", id, err)
	}

	user.IsActive = false
	if err := tx.Save(&user).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to update user with id %d: %w", id, err)
	}

	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
*/
