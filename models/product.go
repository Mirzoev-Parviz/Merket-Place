package models

import "time"

type Category struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Name     string `json:"name"`
	Amount   int    `json:"amount" gorm:"not null; default: 0"`
	IsActive bool   `json:"is_active" gorm:"not null; default: true"`
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity" gorm:"not null; default: 1"`
	InStock     bool    `json:"in_stock"`
	IsActive    bool    `json:"is_active" gorm:"not null; default: true"`
}

type Search struct {
	Query string `json:"query"`
}

type Filter struct {
	Category   string  `json:"category"`
	Price      float64 `json:"price"`
	MerchantID int     `json:"merchant_id"`
}

type ProductInfo struct {
	Product  Product `json:"product"`
	Merchant string  `json:"merchant"`
}

type Review struct {
	ID         int    `json:"id" gorm:"primaryKey"`
	ProductID  int    `json:"product_id" gorm:"not null"`
	MerchantID int    `json:"merchant_id" gorm:"not null"`
	UserID     int    `json:"user_id" gorm:"not null"`
	Rating     int    `json:"rating" gorm:"not null"`
	Comment    string `json:"comment" gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Rating struct {
	ID        int `json:"id" gorm:"primaryKey"`
	ProductID int `json:"product_id" gorm:"not null"`
	Value     int `json:"value" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
