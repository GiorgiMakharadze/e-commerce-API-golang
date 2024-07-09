package models

import (
	"gorm.io/gorm"
)

type UserRole string

const (
	Admin  UserRole = "Admin"
	Seller UserRole = "Seller"
	Buyer  UserRole = "Buyer"
)

type User struct {
	gorm.Model
	Username  string   `json:"username" gorm:"uniqueIndex;not null"`
	Email     string   `json:"email" gorm:"uniqueIndex;not null"`
	Password  string   `json:"-" gorm:"not null"`
	FirstName string   `json:"first_name" gorm:"not null"`
	LastName  string   `json:"last_name" gorm:"not null"`
	Role      UserRole `json:"role" gorm:"type:varchar(20);not null"`
}

type LoginAuthInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RegisterAuthInput struct {
	Username  string   `json:"username" binding:"required"`
	Password  string   `json:"password" binding:"required"`
	Email     string   `json:"email" binding:"required"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	Role      UserRole `json:"role" binding:"required"`
}

// type Session struct {
// 	gorm.Model
// 	UserID       uint           `json:"user_id" gorm:"index;not null"`
// 	AccessToken  string         `json:"access_token" gorm:"not null"`
// 	RefreshToken string         `json:"refresh_token" gorm:"not null"`
// }
