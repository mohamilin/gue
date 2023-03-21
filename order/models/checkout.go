package models

type Checkout struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"user_id"`
	ProductID   uint   `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       uint   `json:"price"`
	TotalPrice  uint   `json:"total_price"`
	Amount      uint   `json:"amount"`
}
