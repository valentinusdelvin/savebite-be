package entity

import "github.com/google/uuid"

type Store struct {
	StoreId     uuid.UUID `gorm:"type:uuid;primary_key"`
	userId      uuid.UUID `gorm:"type:uuid;not null"`
	name        string    `gorm:"type:varchar(255);not null"`
	description string    `gorm:"type:text"`

	user User `gorm:"foreignKey:userId;references:userId"`
}
