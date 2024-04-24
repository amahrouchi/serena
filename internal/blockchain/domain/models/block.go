package models

import "time"

// BlockStatus is the block status type
type BlockStatus string

const (
	BlockStatusPending BlockStatus = "pending"
	BlockStatusActive  BlockStatus = "active"
	BlockStatusClosed  BlockStatus = "closed"
)

// Block is the block DB model
type Block struct {
	ID           uint        `gorm:"primarykey"`
	Status       BlockStatus `gorm:"type:varchar(30)"`
	Payload      string      `gorm:"type:json"`
	Hash         *string     `gorm:"unique"`
	PreviousHash *string     `gorm:"unique"`
	CreatedAt    time.Time
}
