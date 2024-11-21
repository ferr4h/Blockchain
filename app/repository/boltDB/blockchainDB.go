package boltDB

import (
	"errors"
	"github.com/boltdb/bolt"
)

type BlockchainRepository struct {
	db *bolt.DB
}

func NewBlockchainRepository(db *bolt.DB) *BlockchainRepository {
	return &BlockchainRepository{db}
}

func (b BlockchainRepository) Get(key []byte) ([]byte, error) {
	var value []byte

	err := b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Blockchain"))
		value = b.Get(key)
		return nil
	})

	if value == nil {
		return nil, errors.New("not found")
	}

	return value, err
}

func (b BlockchainRepository) Post(key []byte, value []byte) error {
	return b.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Blockchain"))
		return b.Put(key, []byte(value))
	})
}
