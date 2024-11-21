package blockchain

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math"
	"math/big"
)

func (block *Block) DeriveHash() (int, []byte) {
	var intHash big.Int
	var hash [32]byte

	nonce := 0

	for nonce < math.MaxInt64 {
		data := bytes.Join(
			[][]byte{
				block.PrevHash,
				block.Data,
				ToByte(int64(nonce)),
				ToByte(int64(Difficulty)),
			},
			[]byte{},
		)
		hash = sha256.Sum256(data)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(deduceTarget()) == -1 {
			break
		} else {
			nonce++
		}

	}

	return nonce, hash[:]
}

func (block *Block) ValidateHash() bool {
	var intHash big.Int

	data := bytes.Join(
		[][]byte{
			block.PrevHash,
			block.Data,
			ToByte(int64(block.Nonce)),
			ToByte(int64(Difficulty)),
		},
		[]byte{},
	)

	hash := sha256.Sum256(data)
	intHash.SetBytes(hash[:])

	return intHash.Cmp(deduceTarget()) == -1
}

func ToByte(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	HandleError(err)
	return buff.Bytes()
}

func deduceTarget() *big.Int {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-Difficulty))
	return target
}
