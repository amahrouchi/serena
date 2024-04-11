package models

// Block represents a block in the blockchain.
type Block struct {
	Header  *blockHeader   `json:"header"`
	Payload map[string]any `json:"payload"`
	Hash    string         `json:"hash"`
}

// blockHeader represents the header of a block.
type blockHeader struct {
	PreviousHash string `json:"previousHash"`
	CreationDate uint64 `json:"creationDate"`
}
