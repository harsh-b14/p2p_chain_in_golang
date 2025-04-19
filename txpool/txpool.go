package txpool

import (
	"reflect"
	"sync"

	"github.com/harsh-b14/p2p-chain/core"
)

type TxPool struct {
	mu          sync.Mutex
	Transactions []core.Transaction
}

var TransactionPool *TxPool = NewTxPool()

func NewTxPool() *TxPool {
	return &TxPool{Transactions: []core.Transaction{}}
}	

func (pool *TxPool) AddTransaction(tx core.Transaction) *TxPool {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	pool.Transactions = append(pool.Transactions, tx)
	return pool
}

// func (pool *TxPool) GetPending() []core.Transaction {
// 	pool.mu.Lock()
// 	defer pool.mu.Unlock()
// 	return pool.Transactions
// }

func (pool *TxPool) GetAllTxs() []core.Transaction {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	
	if len(pool.Transactions) == 0 {
		return nil
	}

	transactions := make([]core.Transaction, len(pool.Transactions))
	copy(transactions, pool.Transactions)
	return transactions
}

func (pool *TxPool) AddTx(tx []core.Transaction) *TxPool {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	pool.Transactions = append(pool.Transactions, tx...)
	return pool
}

func (pool *TxPool) ClearAndAddTx(tx []core.Transaction) *TxPool {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	pool.Transactions = nil
	pool.Transactions = append(pool.Transactions, tx...)
	return pool
}

func (pool *TxPool) RemoveTransaction(txs []core.Transaction) *TxPool {
	pool.mu.Lock()
	defer pool.mu.Unlock()

	txsFromPool := pool.GetAllTxs()
	var updatedTxs []core.Transaction
	for _, i := range txsFromPool {
		flag := false
		for _, j := range txs {
			if reflect.DeepEqual(i, j) {
				flag = true
				break
			}
		}
		if !flag {
			updatedTxs = append(updatedTxs, i)
		}
	}
	pool.ClearAndAddTx(updatedTxs)
	return pool
}

func (pool *TxPool) Print() {
	pool.mu.Lock()
	defer pool.mu.Unlock()
	for i, tx := range pool.Transactions {
		println("Transaction: ", i)
		// println("Hash: ", tx.Hash)
		println("Nonce: ", tx.Nonce)
		println("V: ", tx.V)
		println("R: ", tx.R)
		println("S: ", tx.S)
		println("------------------------------------------------")
	}
}