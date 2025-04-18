package core

import (
	"crypto/sha256"
	"log"

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

type Block struct {
	Header       Header
	Transactions []Transaction
}

func (h *Header) EncodeHeader() [32]byte {
	data, err := rlp.EncodeToBytes(h)
	if err != nil {
		log.Fatal(err)
	}
	hash := sha256.Sum256(data)
	return hash
}
