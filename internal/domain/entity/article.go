package entity

import (
	"time"

	"github.com/google/uuid"
)

type Article struct {
	ArticleId uuid.UUID
	Title     string
	Content   string
	Thumbnail string
	CreatedAt time.Time
}
