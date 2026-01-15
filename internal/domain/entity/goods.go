package entity

import (
	"time"

	"github.com/google/uuid"
)

type Goods struct {
	GoodsId    uuid.UUID
	Name       string
	Amount     int
	ExpiryDate time.Time
}
