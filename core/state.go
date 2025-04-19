package core

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"github.com/ethereum/go-ethereum/crypto"
)

type Account struct {
	Address [20]byte
	Nonce   uint64
	Balance uint64	
}

type State struct {
	Accounts map[[20]byte]*Account // map[key]value
	mu       sync.RWMutex
}

type Wallet struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privateKey"`
}

type KeyStore struct {
	Wallets []Wallet `json:"wallets"`
	mu      sync.Mutex
}

var (
	keystoreFile = "keystore.json"
	keyStore     = KeyStore{}
	currentNonce = uint64(0)
)

func (s *State) GetAccount(addr [20]byte) *Account {
	s.mu.RLock()
	defer s.mu.RUnlock()
	account, ok := s.Accounts[addr]
	if !ok {
		account = &Account{Address: addr}
		s.Accounts[addr] = account
	}
	return account
}

func LoadKeyStore() {
	file, err := os.Open(keystoreFile)
	if err != nil {
		fmt.Println("No keystore found. Creating new one.")
		keyStore = KeyStore{Wallets: []Wallet{}}
		return
	}
	defer file.Close()
	data, _ := ioutil.ReadAll(file)
	_ = json.Unmarshal(data, &keyStore.Wallets)
}

func SaveKeyStore() {
	keyStore.mu.Lock()
	defer keyStore.mu.Unlock()
	data, _ := json.MarshalIndent(keyStore.Wallets, "", "  ")
	_ = ioutil.WriteFile(keystoreFile, data, 0644)
}

func GenerateWallet() Wallet {
	privateKey, _ := crypto.GenerateKey()
	privateKeyHex := fmt.Sprintf("%x", crypto.FromECDSA(privateKey))
	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()

	fmt.Printf("Generated Address: %s\n\n", address)
	wallet := Wallet{Address: address, PrivateKey: privateKeyHex}
	keyStore.Wallets = append(keyStore.Wallets, wallet)
	SaveKeyStore()
	return wallet
}

func GetKeyStore() *KeyStore {
	keyStore.mu.Lock()
	defer keyStore.mu.Unlock()
	return &keyStore
}

func GetPrivateKeyByAddress(address string) (*ecdsa.PrivateKey, error) {
	for _, wallet := range GetKeyStore().Wallets {
		if wallet.Address == address {
			privateKeyBytes, _ := hex.DecodeString(wallet.PrivateKey)
			return crypto.ToECDSA(privateKeyBytes)
		}
	}
	return nil, fmt.Errorf("wallet not found")
}

