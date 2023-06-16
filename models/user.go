package models

type User struct {
	Id       int    `json:"id" gorm:"primarykey"`
	FullName string `json:"full_name"`
	Login    string `json:"login" `
	Password string `json:"password" `
	Role     string `json:"role" `
	IsActive bool   `json:"is_active" gorm:"not null; defalut: true"`
}

type SignInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
