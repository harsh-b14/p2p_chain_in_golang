package txpool

import (
	"sync"
	"github.com/harsh-b14/p2p-chain/core"
)

type TxPool struct {
	mu          sync.Mutex
	pendingTxns []core.Transaction
}

func (pool *TxPool) Add(tx core.Transaction) {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	pool.pendingTxns = append(pool.pendingTxns, tx)
}

func (pool *TxPool) GetPending() []core.Transaction {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	return pool.pendingTxns
}
