package main

import (
	// "context"
	// "flag"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"

	// "github.com/libp2p/go-libp2p/core/peer"
	"github.com/harsh-b14/p2p-chain/miner"
	"github.com/harsh-b14/p2p-chain/rpc"
	"github.com/harsh-b14/p2p-chain/storage"
	"github.com/harsh-b14/p2p-chain/txpool"
	"github.com/harsh-b14/p2p-chain/utils"
	"github.com/multiformats/go-multiaddr"
)

var NodeIds []common.Address

func main() {
	// port := flag.Int("port", 4001, "Port to listen on")
	// peerId := flag.String("peerId", "", "Multiaddress of the peer to connect to")
	// rpcPort := flag.Int("rpcPort", 8080, "Port for RPC server")
	// flag.Parse()

	// Start the P2P node
	host1, err := startNode(4001)
	host2, err := startNode(4002)

	_, _, addr1, err := utils.GenerateKeysAndAddress()
	_, _, addr2, err := utils.GenerateKeysAndAddress()

	if err != nil {
		log.Fatalf("Failed to start the node: %v", err)
	}

	NodeIds = append(NodeIds, addr1, addr2)

	// Display node info
	fmt.Printf("✅ Multiple P2P Node started. Listening on: /ip4/127.0.0.1/tcp/%d and /ip4/127.0.0.1/tcp/%d \n", 4001, 4002)
	fmt.Println()
	fmt.Println("First node Id ", host1.ID())
	fmt.Println("First node adderss", addr1)
	fmt.Println()	
	fmt.Println("Second node Id ", host2.ID())
	fmt.Println("Second node adderss", addr2)
	fmt.Println()

	genesisBlock := miner.MineGenesisBlock(addr1)
	fmt.Println("Genesis block mined!!!")
	rpc.Blockchain = append(rpc.Blockchain, genesisBlock)
	fmt.Println()

	storage.StartDataBase()
	fmt.Println()

	go StartMining()

	// Start RPC server (optional)
	go rpc.StartRPC(8000)

	// Keep the node running
	select {}
}

// Start a new libp2p node
func startNode(port int) (host.Host, error) {
	addr := fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port)
	listenAddr, _ := multiaddr.NewMultiaddr(addr)

	return libp2p.New(libp2p.ListenAddrs(listenAddr))
}

func StartMining(){
	ticker := time.NewTicker(2 * time.Second)
	var cnt int32 = 0
	for range ticker.C {
		minerID := NodeIds[cnt%2]
		cnt++
		// currentMiner = (currentMiner + 1) % len(NodeIds)

		// var transactions []core.Transaction
		// if len(txPool) > 0 {
		// 	transactions = append(transactions, txPool...)
		// 	txPool = []core.Transaction{}
		// } else {
		// 	fmt.Printf("[%s] No transactions to include. Mining empty block.\n", minerID)
		// }

		// block := core.Block{
		// 	Miner:       minerID,
		// 	Transactions: transactions,
		// 	Timestamp:   time.Now().Unix(),
		// }
		// blockchain = append(blockchain, block)

		// Simulate mining process
		time.Sleep(2 * time.Second)
		// Simulate mining a block
		block := miner.MineBlock(txpool.TransactionPool, rpc.Blockchain[len(rpc.Blockchain)-1], minerID)
		if block == nil {
			fmt.Printf("[%s] No transactions to include. Mining empty block.\n\n", minerID)
			continue
		}
		rpc.Blockchain = append(rpc.Blockchain, block)
		fmt.Printf("[%s] Mined block with %d transaction(s)\n\n", minerID, len(txpool.TransactionPool.Transactions))
	}
}
