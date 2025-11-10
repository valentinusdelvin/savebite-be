package entity

import "github.com/google/uuid"

type Product struct {
	ProductId    uuid.UUID `gorm:"type:varchar(36);primaryKey"`
	StoreId      uuid.UUID `gorm:"type:varchar(36);not null"`
	Name         string    `gorm:"type:varchar(255);not null"`
	Description  string    `gorm:"type:text"`
	Price        float64   `gorm:"type:decimal(10,2);not null"`
	Stock        int       `gorm:"type:int;not null"`
	ProductImage string    `gorm:"type:varchar(255)"`
}
