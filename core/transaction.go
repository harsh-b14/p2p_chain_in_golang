package core

import (
	"crypto/sha256"
	"github.com/ethereum/go-ethereum/rlp"
	"math/big"
)

type Transaction struct {
	To    [20]byte
	Value uint64
	Nonce uint64
	V, R, S *big.Int
}

func (tx *Transaction) Hash() [32]byte {
	data, _ := rlp.EncodeToBytes(tx)
	hash := sha256.Sum256(data)
	return hash
}