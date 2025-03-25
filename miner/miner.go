package miner

import (
	"time"
	"github.com/harsh-b14/p2p-chain/core"
	"github.com/harsh-b14/p2p-chain/txpool"
)

func MineBlock(txPool *txpool.TxPool, prevBlock *core.Block, minerAddr [20]byte) *core.Block {
	txs := txPool.GetPending()
	header := core.Header{
		ParentHash: prevBlock.Header.Hash(),
		Miner:      minerAddr,
		Number:     prevBlock.Header.Number + 1,
		Timestamp:  uint64(time.Now().Unix()),
	}

	block := core.Block{
		Header:       header,
		Transactions: txs,
	}
	return &block
}
