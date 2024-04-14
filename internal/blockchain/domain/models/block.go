package models

import "gorm.io/gorm"

type Block struct {
	gorm.Model
	Header  *blockHeader   `gorm:"embedded"`
	Payload map[string]any `gorm:"type:jsonb"`
	Hash    string         `gorm:"unique;index"`
}

// blockHeader represents the header of a block.
type blockHeader struct {
	PreviousHash string `json:"previousHash"`
	CreationDate uint64 `json:"creationDate"`
}
