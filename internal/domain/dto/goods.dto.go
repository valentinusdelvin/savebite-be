package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateGoods struct {
	Name       string    `json:"name" binding:"required"`
	Amount     int       `json:"amount" binding:"required"`
	ExpiryDate time.Time `json:"expiry_date" binding:"required"`
}

type UpdateGoods struct {
	GoodsId    uuid.UUID `json:"goods_id" binding:"required"`
	Name       string    `json:"name" binding:"required"`
	Amount     int       `json:"amount" binding:"required"`
	ExpiryDate time.Time `json:"expiry_date" binding:"required"`
}
