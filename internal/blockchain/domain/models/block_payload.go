package models

import "time"

// BlockPayloadItem represents the payload of a block.
type BlockPayloadItem struct {
	Author    string         `json:"author" validate:"required"`
	Data      map[string]any `json:"data" validate:"required"`
	CreatedAt time.Time      `json:"created_at" validate:"required"`
}
