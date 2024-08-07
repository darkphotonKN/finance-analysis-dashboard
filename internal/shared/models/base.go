package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
	Password  string `gorm:"not null"`
	Role      string `gorm:"not null"`
}
