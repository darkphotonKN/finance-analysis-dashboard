package models

import (
	"github.com/darkphotonKN/finance-analysis-dashboard/internal/shared/constants"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string             `gorm:"not null"`
	FirstName string             `gorm:"not null"`
	LastName  string             `gorm:"not null"`
	Password  string             `gorm:"not null"`
	Role      constants.UserRole `gorm:"not null"`
}
