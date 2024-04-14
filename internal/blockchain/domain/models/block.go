package models

import "gorm.io/gorm"

// Block is the block DB model
type Block struct {
	gorm.Model
	Header  *BlockHeader `gorm:"embedded"`
	Payload string       `gorm:"type:json"`
	Hash    string       `gorm:"unique;index"`
	// TODO: remove created_at, updated_at
}

// BlockHeader represents the header of a block.
type BlockHeader struct {
	PreviousHash string `gorm:"unique" json:"previousHash"`
	CreationDate uint64 `json:"creationDate"` // TODO: use a date sql type here
}
