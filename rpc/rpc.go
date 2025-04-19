package rpc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harsh-b14/p2p-chain/core"
	"github.com/harsh-b14/p2p-chain/txpool"
)

var Blockchain []*core.Block
var currentNonce uint64 = 0

func StartRPC(port int) {
	r := mux.NewRouter()
	r.HandleFunc("/blockNumber", getBlockNumber).Methods("GET")
	r.HandleFunc("/block", getBlock).Methods("GET")
	r.HandleFunc("/sendTx", sendTransaction).Methods("POST")
	r.HandleFunc("/getBalance", getBalance).Methods("GET")

	fmt.Printf("RPC server runningon port %d\n", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}

func AddBlock(block *core.Block) {
	Blockchain = append(Blockchain, block)
}

func getBlockNumber(w http.ResponseWriter, r *http.Request) {
	blockNum := len(Blockchain)
	json.NewEncoder(w).Encode(blockNum)
}

func getBlock(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["number"]
	if !ok || len(keys[0]) < 1 {
		http.Error(w, "Missing block number", http.StatusBadRequest)
		return
	}

	blockNum := keys[0]
	for _, block := range Blockchain {
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
	var txs []core.Transaction
	txs = append(txs, *signedTx)
	txpool.TransactionPool.AddTx(txs)
	fmt.Println("Transaction sent successfully")
	fmt.Printf("Tx Hash: %x\n", signedTx.EncodeTx())
	json.NewEncoder(w).Encode(signedTx)
}

func getBalance(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("address") == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	address := r.URL.Query().Get("address")
	fmt.Println("Value of address:", address)

	// Convert address to [20]byte

	// Get the balance of the address
	// account, err := types.GetAccount(toAddress)
	// if err != nil {
	// 	http.Error(w, "Error getting account", http.StatusInternalServerError)
	// 	return
	// }

	// Encode the balance to JSON
	// balanceJSON, err := json.Marshal(account.Balance)
	// if err != nil {
	// 	http.Error(w, "Error encoding balance to JSON", http.StatusInternalServerError)
	// 	return
	// }

	// Write the JSON response
	// w.Header().Set("Content-Type", "application/json")
	// _, err = w.Write(balanceJSON)
	// if err != nil {
	// 	fmt.Println("Error writing response:", err)
	// }

	w.WriteHeader(http.StatusOK)
}
