package rpc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harsh-b14/p2p-chain/core"
	"github.com/harsh-b14/p2p-chain/txpool"
)

var blockchain []*core.Block
var currentNonce uint64 = 0

func StartRPC(port int) {
	r := mux.NewRouter()
	r.HandleFunc("/blockNumber", getBlockNumber).Methods("GET")
	r.HandleFunc("/block", getBlock).Methods("GET")
	r.HandleFunc("/sendTx", sendTransaction).Methods("POST")

	fmt.Printf("RPC server runningon port %d\n", port)
	http.ListenAndServe(fmt.Sprintf (":%d", port), r)
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
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
		return
	}

	var tx core.UnSignedTx
	_ = json.NewDecoder(r.Body).Decode(&tx)
	if tx.To == [20]byte{} || tx.Value == 0 {
		http.Error(w, "Invalid transaction", http.StatusBadRequest)
		return
	}

	tx.Nonce = currentNonce
	currentNonce++

	// Here assuming sender's address is also part of UnSignedTx or a query param
	senderAddress := r.URL.Query().Get("from")
	if senderAddress == "" {
		http.Error(w, "Missing sender address", http.StatusBadRequest)
		return
	}

	privateKey, err := core.GetPrivateKeyByAddress(senderAddress)
	if err != nil {
		http.Error(w, "Sender wallet not found", http.StatusBadRequest)
		return
	}

	signedTx, err := core.SignTransaction(&tx, *privateKey)
	if err != nil {
		http.Error(w, "Failed to sign transaction", http.StatusInternalServerError)
		return
	}

	// TODO: Add to tx pool
	txpool.TransactionPool.AddTx([]core.Transaction{*signedTx})
	fmt.Println("Transaction sent successfully")
	fmt.Printf("Tx Hash: %x\n", signedTx.EncodeTx())
	json.NewEncoder(w).Encode(signedTx)
}
