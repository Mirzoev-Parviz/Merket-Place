package models

import "gorm.io/gorm"

type Contact struct {
	Phone   uint   `json:"phone" binding:"required"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type User struct {
	gorm.Model
	ID       int     `json:"id" gorm:"primarykey"`
	FullName string  `json:"full_name"`
	Login    string  `json:"login" `
	Password string  `json:"password"`
	Contacts Contact `json:"contacts" gorm:"embedded;embeddedPrefix:contacts"`
	Role     string  `json:"role" binding:"required"`
	IsActive bool    `json:"is_active" gorm:"not null; default: true"`
}

type SignInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
