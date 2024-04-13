package models

//type Block struct {
//	gorm.Model
//}

// BlockDTO represents a block in the blockchain.
type BlockDTO struct {
	Header  *blockHeader   `json:"header"`
	Payload map[string]any `json:"payload"`
	Hash    string         `json:"hash"`
}

// blockHeader represents the header of a block.
type blockHeader struct {
	PreviousHash string `json:"previousHash"`
	CreationDate uint64 `json:"creationDate"`
}
