package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Name     string    `gorm:"not null"`
	Messages []Message `gorm:"foreignKey:UserID;references:ID"`
}
