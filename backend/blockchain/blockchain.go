package blockchain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

type Block struct {
	Index        int
	Timestamp    time.Time
	Data         []byte
	PreviousHash string
	Hash         string
	Nonce        int
}

type Blockchain struct {
	Chain  []*Block
	mu     sync.Mutex
}

func NewBlockchain() *Blockchain {
	bc := &Blockchain{
		Chain: make([]*Block, 0),
	}
	bc.addGenesisBlock()
	return bc
}

func (bc *Blockchain) addGenesisBlock() {
	genesisBlock := &Block{
		Index:        0,
		Timestamp:    time.Now(),
		Data:         []byte("Genesis Block"),
		PreviousHash: "0",
		Nonce:        0,
	}
	genesisBlock.Hash = bc.calculateHash(genesisBlock)
	bc.Chain = append(bc.Chain, genesisBlock)
}

func (bc *Blockchain) calculateHash(block *Block) string {
	record := fmt.Sprintf("%d%s%s%s%d", 
		block.Index, 
		block.Timestamp.String(), 
		block.Data, 
		block.PreviousHash, 
		block.Nonce,
	)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func (bc *Blockchain) mineBlock(block *Block, difficulty int) {
	target := ""
	for i := 0; i < difficulty; i++ {
		target += "0"
	}

	for {
		block.Hash = bc.calculateHash(block)
		if block.Hash[:difficulty] == target {
			break
		}
		block.Nonce++
	}

	fmt.Printf("Block mined: %s\n", block.Hash)
}

func (bc *Blockchain) AddBlock(data []byte) *Block {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	prevBlock := bc.Chain[len(bc.Chain)-1]
	newBlock := &Block{
		Index:        prevBlock.Index + 1,
		Timestamp:    time.Now(),
		Data:         data,
		PreviousHash: prevBlock.Hash,
		Nonce:        0,
	}

	bc.mineBlock(newBlock, 4)
	bc.Chain = append(bc.Chain, newBlock)
	return newBlock
}

func (bc *Blockchain) IsValid() bool {
	for i := 1; i < len(bc.Chain); i++ {
		currentBlock := bc.Chain[i]
		prevBlock := bc.Chain[i-1]

		if currentBlock.Hash != bc.calculateHash(currentBlock) {
			return false
		}

		if currentBlock.PreviousHash != prevBlock.Hash {
			return false
		}
	}
	return true
}

func (bc *Blockchain) GetBlock(hash string) (*Block, error) {
	for _, block := range bc.Chain {
		if block.Hash == hash {
			return block, nil
		}
	}
	return nil, fmt.Errorf("block with hash %s not found", hash)
}

func (bc *Blockchain) GetLatestBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) GetBlockData(hash string) (map[string]interface{}, error) {
	block, err := bc.GetBlock(hash)
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err := json.Unmarshal(block.Data, &data); err != nil {
		return nil, err
	}

	return data, nil
}