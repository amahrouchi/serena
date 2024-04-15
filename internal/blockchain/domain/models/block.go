package models

import "time"

// Block is the block DB model
type Block struct {
	ID           uint    `gorm:"primarykey"`
	Payload      string  `gorm:"type:json"`
	Hash         *string `gorm:"unique"`
	PreviousHash string  `gorm:"unique"`
	CreatedAt    time.Time
}
