package models

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

type Basket struct {
	Id       int       `json:"-" gorm:"primarykey"`
	UserId   int       `json:"user_id"`
	Products []Product `json:"products"`
	Quantity int       `json:"quantity"`
	TotalSum float64   `json:"total_sum"`
	PreOrder bool      `json:"preorder"`
}

type CartItem struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
