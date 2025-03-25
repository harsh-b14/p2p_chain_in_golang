package core

import "sync"

type Account struct {
	Address [20]byte
	Nonce   uint64
	Balance uint64
}

type State struct {
	Accounts map[[20]byte]*Account
	mu       sync.RWMutex
}

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
