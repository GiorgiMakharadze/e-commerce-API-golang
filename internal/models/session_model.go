package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	UserID       uint   `json:"user_id" gorm:"index;not null"`
	AccessToken  string `json:"access_token" gorm:"not null"`
	RefreshToken string `json:"refresh_token" gorm:"not null"`
}
