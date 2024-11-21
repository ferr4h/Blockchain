package repository

type BlockchainRepository interface {
	Get(key []byte) ([]byte, error)
	Post(key []byte, value []byte) error
}
