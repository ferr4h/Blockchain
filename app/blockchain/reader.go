package blockchain

import "blockchain/app/repository"

type Reader struct {
	CurrentHash []byte
	Repository  repository.BlockchainRepository
}

func (chain *BlockChain) InitReader() *Reader {
	iterator := &Reader{chain.LastHash, chain.Repository}
	return iterator
}

func (reader *Reader) Iterate() *Block {
	var block *Block

	value, err := reader.Repository.Get(reader.CurrentHash)
	HandleError(err)
	block = Deserialize(value)
	reader.CurrentHash = block.PrevHash
	return block
}
