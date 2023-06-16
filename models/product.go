package models

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	// ProductId []int  `json:"-" gorm:"not null; default: 0"`
	Amount int `json:"amount" gorm:"not null; default: 0"`
}

type Product struct {
	Id          int     `json:"id"`
	UserId      int     `json:"user_id"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float64 `json:"price"`
	IsActive    bool    `json:"is_active"`
}
