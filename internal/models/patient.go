package models

import (
	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model        // adds ID, CreatedAt, UpdatedAt, DeletedAt
	Name       string `json:"name" gorm:"not null"`
	Age        int    `json:"age" gorm:"not null"`
	Gender     string `json:"gender" gorm:"not null"`
	Address    string `json:"address"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	// Email      *string `json:"email" gorm:"uniqueIndex"`
	// Email      string `json:"email" gorm:"uniqueIndex"`
	Notes string `json:"notes"`
}
