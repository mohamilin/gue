package models

type Session struct {
	ID           uint   `json:"id" gorm:"primaryKey"`
	UserID       uint   `json:"user_id"`
	AccessToken  string `json:"access_token" gorm:"index"`
	RefreshToken string `json:"refresh_token" gorm:"index"`
}
