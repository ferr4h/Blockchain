package blockchain

import (
	"bytes"
	"encoding/gob"
)

type Block struct {
	Data     []byte
	Hash     []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	nonce, hash := block.DeriveHash()
	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (block *Block) Serialize() []byte {
	var result bytes.Buffer
	encoder := gob.NewEncoder(&result)
	err := encoder.Encode(block)
	HandleError(err)
	return result.Bytes()
}

func Deserialize(data []byte) *Block {
	var block Block
	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&block)
	HandleError(err)
	return &block
}
