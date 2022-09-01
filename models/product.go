package models

import (
	"time"
)

type Product struct {
	ID uint `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time
	Name string `json:"name"`
	Price float32 `json:"price"`
}