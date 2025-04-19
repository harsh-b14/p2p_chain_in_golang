package miner

import (
	"time"

	"github.com/harsh-b14/p2p-chain/core"
	"github.com/harsh-b14/p2p-chain/txpool"
	t "github.com/chad-chain/chadChain/core/types"
)

func MineGenesisBlock(minerAddr [20]byte) *core.Block {
	header := core.Header{
		ParentHash: [32]byte{},
		Miner:      minerAddr,
		Number:     0,
		Timestamp:  uint64(time.Now().Unix()),
	}

	block := core.Block{
		Header: header,
	}
	
	return &block
}

func MineBlock(txPool *txpool.TxPool, prevBlock *core.Block, minerAddr [20]byte) *core.Block {
	txs := txPool.GetAllTxs()
	header := core.Header{
		// ParentHash: prevBlock.Header.Hash(),
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

func CreateEmptyBlock() t.Block {
	emptyBlock := new(t.Block)
	emptyBlock.Header.ParentHash = t.LatestBlockHash
	emptyBlock.Header.StateRoot = [32]byte(t.StateRootHash)
	emptyBlock.Header.TransactionsRoot = [32]byte(t.LatestBlock.Header.TransactionsRoot)
	emptyBlock.Header.Number = t.LatestBlock.Header.Number + 1
	emptyBlock.Header.Timestamp = uint64(time.Now().Unix())
	// sig, err := .SignHeader(&emptyBlock.Header)
	// if err != nil {
	// 	log.Fatalln("Failed to sign header: ", err)
	// }
	// emptyBlock.Header.ExtraData = sig
	return *emptyBlock
}
