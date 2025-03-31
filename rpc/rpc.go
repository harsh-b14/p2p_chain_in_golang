package rpc

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/harsh-b14/p2p-chain/core"
)

var blockchain []*core.Block

func StartRPC(port int) {
	r := http.NewServeMux()
	r.HandleFunc("/blockNumber", getBlockNumber)
	r.HandleFunc("/block", getBlock)
	r.HandleFunc("/sendTx", sendTransaction)

	fmt.Printf("RPC server running on port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

func AddBlock(block *core.Block) {
	blockchain = append(blockchain, block)
}

func getBlockNumber(w http.ResponseWriter, r *http.Request) {
	blockNum := len(blockchain)
	json.NewEncoder(w).Encode(blockNum)
}

func getBlock(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["number"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "Missing block number", http.StatusBadRequest)
		return
	}

	blockNum := keys[0]
	for _, block := range blockchain {
		if fmt.Sprintf("%d", block.Header.Number) == blockNum {
			json.NewEncoder(w).Encode(block)
			return
		}
	}
	http.Error(w, "Block not found", http.StatusNotFound)
}

func sendTransaction(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["signed"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "Missing signed transaction", http.StatusBadRequest)
		return
	}
	signedTx := keys[0]
	fmt.Println("Received signed transaction:", signedTx)
	w.WriteHeader(http.StatusOK)
}
