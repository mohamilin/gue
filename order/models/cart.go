package models

type Cart struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	UserID      uint   `json:"user_id"`
	ProductID   uint   `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       uint   `json:"price"`
	Amount      uint   `json:"amount"`
}
