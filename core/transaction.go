package core

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"errors"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/harsh-b14/p2p-chain/utils"
)

type Transaction struct {
	To      [20]byte
	Value   uint64
	Nonce   uint64
	V, R, S *big.Int
}

type UnSignedTx struct {
	To    [20]byte
	Value uint64
	Nonce uint64
}

func (tx *Transaction) EncodeTx() [32]byte {
	data, err := rlp.EncodeToBytes(tx)
	if err != nil {
		log.Fatal(err)
	}
	hash := sha256.Sum256(data)
	return hash
}

func CreateUnsignedTransaction(to [20]byte, value uint64, nonce uint64) *Transaction {
	return &Transaction{
		To:    to,
		Value: value,
		Nonce: nonce,
	}
}

func CreateTransaction(to [20]byte, value uint64, nonce uint64, v, r, s *big.Int) *Transaction {
	return &Transaction{
		To:    to,
		Value: value,
		Nonce: nonce,
		V:     v,
		R:     r,
		S:     s,
	}
}

func SignTransaction(tx *UnSignedTx, privateKey ecdsa.PrivateKey) (*Transaction, error) {
	hash, _ := utils.EncodeData(tx, false)
	sig, err := crypto.Sign(hash[:], &privateKey)
	if err != nil {
		return nil, err
	}
	R, S, V := utils.DecodeSignature(sig)
	return &Transaction{
		To:    tx.To,
		Value: tx.Value,
		Nonce: tx.Nonce,
		V:     V,
		R:     R,
		S:     S,
	}, nil
}

// copied from playground of manav
func RecoverPlain(sighash common.Hash, R, S, Vb *big.Int, homestead bool) (common.Address, error) {
	if Vb.BitLen() > 8 {
		return common.Address{}, errors.New("invalid signature")
	}
	V := byte(Vb.Uint64() - 27)
	if !crypto.ValidateSignatureValues(V, R, S, homestead) {
		return common.Address{}, errors.New("invalid signature")
	}
	// encode the signature in uncompressed format
	r, s := R.Bytes(), S.Bytes()
	sig := make([]byte, crypto.SignatureLength)
	copy(sig[32-len(r):32], r)
	copy(sig[64-len(s):64], s)
	sig[64] = V
	// recover the public key from the signature
	pub, err := crypto.Ecrecover(sighash[:], sig)
	if err != nil {
		return common.Address{}, err
	}
	if len(pub) == 0 || pub[0] != 4 {
		return common.Address{}, errors.New("invalid public key")
	}
	var addr common.Address
	copy(addr[:], crypto.Keccak256(pub[1:])[12:])
	return addr, nil
}
