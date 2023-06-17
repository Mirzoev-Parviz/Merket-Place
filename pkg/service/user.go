package service

import (
	"market_place/models"
	"market_place/pkg/repository"
)

type UserService struct {
	repo repository.User
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{repo: repo}
}

func (u *UserService) CheckLogin(login string) (bool, error) {
	return u.repo.CheckLogin(login)
}
func (u *UserService) UpdateUser(id int, user models.User) error {
	user.Password = generatePasswordHash(user.Password)
	return u.repo.UpdateUser(id, user)
}
func (u *UserService) DeactivateUser(id int) error {
	return u.repo.DeactivateUser(id)
}
