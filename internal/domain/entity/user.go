package entity

import "github.com/google/uuid"

type User struct {
	UserId    uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	Email     string    `gorm:"type:varchar(255);uniqueIndex"`
	Password  string    `gorm:"type:varchar(255);not null"`
	FirstName string    `gorm:"type:varchar(100);not null"`
	LastName  string    `gorm:"type:varchar(100);not null"`
	IsAdmin   bool      `gorm:"type:boolean;not null"`
}
