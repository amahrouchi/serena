package requests

// WriteRequest represents a request to write data to the blockchain.
type WriteRequest struct {
	Author string         `json:"author" validate:"required"`
	Data   map[string]any `json:"data" validate:"required"`
}
