package models

import "time"

// BlockStatus is the block status type
type BlockStatus string

const (
	BlockStatusPending BlockStatus = "pending"
	BlockStatusActive  BlockStatus = "active"
	BlockStatusClosed  BlockStatus = "closed"
)

// Block is the block DB model + json payload
type Block struct {
	ID           uint        `gorm:"primarykey" json:"id"`
	Status       BlockStatus `gorm:"type:varchar(30)" json:"status"`
	Payload      string      `gorm:"type:json" json:"payload"`
	Hash         *string     `gorm:"type:varchar(255);unique" json:"-"`
	PreviousHash *string     `gorm:"type:varchar(255);unique" json:"previous_hash"`
	CreatedAt    time.Time   `json:"created_at"`
}
