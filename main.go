package main

import (
	"blockchain/app/blockchain"
	"blockchain/app/repository/boltDB"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"strconv"
)

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	blockchainRepository := boltDB.NewBlockchainRepository(db)

	chain := blockchain.Init(blockchainRepository)

	chain.AddBlock("First Block")
	chain.AddBlock("Second Block")
	chain.AddBlock("Third Block")

	reader := chain.InitReader()
	for {
		block := reader.Iterate()

		fmt.Printf("Prev. hash: %x\n", block.PrevHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		pow := block.ValidateHash()
		fmt.Printf("PoW: %s\n", strconv.FormatBool(pow))
		fmt.Println()

		if len(block.PrevHash) == 0 {
			break
		}
	}
}

func initDB() (*bolt.DB, error) {
	db, err := bolt.Open("my.db", 0600, nil)
	if err != nil {
		return nil, err
	}

	if err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Blockchain"))
		return err
	}); err != nil {
		return nil, err
	}
	return db, nil
}
