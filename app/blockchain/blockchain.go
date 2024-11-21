package blockchain

import (
	"blockchain/app/repository"
)

const (
	Difficulty = 12
)

type BlockChain struct {
	LastHash   []byte
	Repository repository.BlockchainRepository
}

func (chain *BlockChain) AddBlock(data string) {
	value, err := chain.Repository.Get([]byte("lastHash"))
	HandleError(err)

	newBlock := CreateBlock(data, value)

	err = chain.Repository.Post(newBlock.Hash, newBlock.Serialize())
	HandleError(err)

	err = chain.Repository.Post([]byte("lastHash"), newBlock.Hash)
	HandleError(err)
	chain.LastHash = newBlock.Hash
}

func Init(rep repository.BlockchainRepository) *BlockChain {
	var lastHash []byte
	if _, err := rep.Get([]byte("lastHash")); err != nil {
		genesis := CreateBlock("Genesis Block", []byte{})
		err = rep.Post(genesis.Hash, genesis.Serialize())
		HandleError(err)
		err = rep.Post([]byte("lastHash"), genesis.Hash)
		lastHash = genesis.Hash
	} else {
		hash, err := rep.Get([]byte("lastHash"))
		HandleError(err)
		lastHash = hash
	}
	blockchain := BlockChain{lastHash, rep}
	return &blockchain
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
