package models

import "time"

type Category struct {
	ID       int    `json:"id"`
	ParentID int    `json:"parent_id"`
	Name     string `json:"name"`
	Amount   int    `json:"amount" gorm:"not null; default: 0"`
}

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity" gorm:"not null; default: 1"`
	IsActive    bool    `json:"is_active" gorm:"not null; default: true"`
}

type Cart struct {
	ID        int        `json:"id" gorm:"primarykey"`
	UserID    int        `json:"user_id" gorm:"not null; references: users(id)"`
	Items     []CartItem `json:"items"`
	Quantity  int        `json:"quantity"`
	TotalSum  float64    `json:"total_sum"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

type CartItem struct {
	ID         int       `json:"id" gorm:"primarykey"`
	CartID     int       `json:"cart_id"`
	ProductID  int       `json:"product_id"`
	MerchantID int       `json:"merchant_id"`
	Quantity   int       `json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Search struct {
	Query string `json:"query"`
}
