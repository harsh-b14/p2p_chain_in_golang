package core

import (
	"crypto/sha256"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/rlp"
)

type Header struct {
	ParentHash       [32]byte
	Miner            [20]byte
	StateRoot        [32]byte
	TransactionsRoot [32]byte
	Number           uint64
	Timestamp        uint64
	ExtraData        []byte
}

type Transaction struct {
	To      [20]byte
	Value   uint64
	Nonce   uint64
	V, R, S *big.Int
}

type Block struct {
	Header       Header
	Transactions []Transaction
}

func (h *Header) Hash() [32]byte {
	data, err := rlp.EncodeToBytes(h)
	if err != nil {
		log.Fatal(err)
	}
	hash := sha256.Sum256(data)
	return hash
}

func (tx *Transaction) Hash() [32]byte {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		log.Fatal(err)
	}
	hash := sha256.Sum256(data)
	return hash
}
