package models

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
