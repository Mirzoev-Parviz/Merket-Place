package models

import "gorm.io/gorm"

type Merchant struct {
	ID          int     `json:"id" gorm:"primarykey"`
	Login       string  `json:"login"`
	Password    string  `json:"password"`
	Name        string  `json:"name"`
	Contacts    Contact `json:"contacts" gorm:"embedded;embeddedPrefix:contacts"`
	Description string  `json:"description"`
	IsActive    bool    `json:"is_active" gorm:"not null; default: true"`
}

type MerchantProduct struct {
	gorm.Model
	ProductID   int     `json:"product_id"`
	MerchantID  int     `json:"merchant_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Quantity    int     `json:"quantity" gorm:"not null; default: 1"`
	Price       float64 `json:"price" gorm:"references: prdocts(price)"`
	InStock     bool    `json:"in_stock"`
	TotalOrders int     `json:"total_orders"`
	Rating      float64 `json:"rating"`
	IsActive    bool    `json:"is_active" gorm:"not null; default: true"`
}

type Pagination struct {
	Page int `json:"page"`
}
