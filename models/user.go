package models

import (
	"time"

	"gorm.io/gorm"
)

type Contact struct {
	Phone   uint   `json:"phone" binding:"required"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type SuperAdmin struct {
	ID       int    `json:"id"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type User struct {
	gorm.Model
	ID       int     `json:"id" gorm:"primarykey"`
	FullName string  `json:"full_name"`
	Login    string  `json:"login" `
	Password string  `json:"password"`
	Contacts Contact `json:"contacts" gorm:"embedded;embeddedPrefix:contacts"`
	IsActive bool    `json:"is_active" gorm:"not null; default: true"`
}

type SignInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Cart struct {
	ID        int        `json:"id" gorm:"primarykey"`
	UserID    int        `json:"user_id" gorm:"not null; references: users(id)"`
	Items     []CartItem `json:"items"`
	Quantity  int        `json:"quantity"`
	TotalSum  float64    `json:"total_sum"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	IsActive  bool       `json:"is_active" gorm:"not null; default: true"`
}

type CartItem struct {
	ID         int       `json:"id" gorm:"primarykey"`
	CartID     int       `json:"cart_id"`
	ProductID  int       `json:"product_id"`
	MerchantID int       `json:"merchant_id"`
	Quantity   int       `json:"quantity"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	IsActive   bool      `json:"is_active" gorm:"not null; default: true"`
}
