package utils

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rlp"
)

func GenerateKeysAndAddress() (*ecdsa.PrivateKey, string, common.Address, error) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		return nil, "", common.Address{}, err
	}
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	privateKeyHex := hex.EncodeToString(crypto.FromECDSA(privateKey))

	return privateKey, privateKeyHex, address, nil
}

// DecodeData decodes the data into the entity
// Entity should be a pointer to the struct to which the data should be decoded
func DecodeData(data []byte, entity interface{}) (any, error) {
	err := rlp.DecodeBytes(data, entity)
	if err != nil {
		return nil, fmt.Errorf("error decoding: %v", err)
	}
	return entity, nil
}

func EncodeData(data interface{}, isJson bool) ([]byte, error) {
	var err error
	if isJson {
		data, err = json.Marshal(data)
		if err != nil {
			return nil, fmt.Errorf("error json marshaling: %v", err)
		}
	}
	encodedData, err := rlp.EncodeToBytes(data)
	if err != nil {
		return nil, fmt.Errorf("error encoding: %v", err)
	}
	return encodedData, nil
}

func DecodeSignature(sig []byte) (r, s, v *big.Int) {
	if len(sig) != crypto.SignatureLength {
		panic(fmt.Sprintf("wrong size for signature: got %d, want %d", len(sig), crypto.SignatureLength))
	}
	r = new(big.Int).SetBytes(sig[:32])
	s = new(big.Int).SetBytes(sig[32:64])
	v = new(big.Int).SetBytes([]byte{sig[64] + 27})
	return r, s, v
}