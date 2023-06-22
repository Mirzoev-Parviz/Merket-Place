package service

import (
	"market_place/models"
	"market_place/pkg/repository"
)

type MerchService struct {
	repo repository.Merchant
}

func NewMerchService(repo repository.Merchant) *MerchService {
	return &MerchService{repo: repo}
}

func (m *MerchService) CreateMerchant(merch models.Merchant) (int, error) {
	merch.Password = generatePasswordHash(merch.Password)
	return m.repo.CreateMerchant(merch)
}

func (m *MerchService) GetMerchant(id int) (models.Merchant, error) {
	return m.repo.GetMerchant(id)
}

func (m *MerchService) UpdateMerchant(id int, merch models.Merchant) error {
	return m.repo.UpdateMerchant(id, merch)
}

func (m *MerchService) DeleteMerchant(id int) error {
	return m.repo.DeleteMerchant(id)
}

func (m *MerchService) AddProductToShelf(merch models.MerchantProduct) (int, error) {
	return m.repo.AddProductToShelf(merch)
}

func (m *MerchService) GetMerchProduct(id int) (models.MerchantProduct, error) {
	return m.repo.GetMerchProduct(id)
}

func (m *MerchService) UpdateMerchProduct(id int, merch models.MerchantProduct) error {
	return m.repo.UpdateMerchProduct(id, merch)
}

func (m *MerchService) DeleteMerchProduct(id int) error {
	return m.repo.DeleteMerchProduct(id)
}
