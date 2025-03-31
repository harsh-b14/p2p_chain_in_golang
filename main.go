package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	// "github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
	"github.com/harsh-b14/p2p-chain/rpc"
	"github.com/harsh-b14/p2p-chain/miner"
)

func main() {
	// port := flag.Int("port", 4001, "Port to listen on")
	// peerId := flag.String("peerId", "", "Multiaddress of the peer to connect to")
	// rpcPort := flag.Int("rpcPort", 8080, "Port for RPC server")
	// flag.Parse()

	// Start the P2P node
	host1, err := startNode(4001)
	host2, err := startNode(4002)

	if err != nil {
		log.Fatalf("Failed to start the node: %v", err)
	}

	// Display node info
	fmt.Printf("✅ Multiple P2P Node started. Listening on: /ip4/127.0.0.1/tcp/%d and /ip4/127.0.0.1/tcp/%d \n", 4001, 4002)
	fmt.Println("First node Id ", host1.ID())
	fmt.Println("Second node Id ", host2.ID())
	fmt.Println()

	fmt.Println("Mining the genesis block...")
	miner.MineGenesisBlock(host1.ID())
	fmt.Println("Mining the genesis block...")

	// Start RPC server (optional)
	go rpc.StartRPC(800)

	// Keep the node running
	select {}
}

// Start a new libp2p node
func startNode(port int) (host.Host, error) {
	addr := fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port)
	listenAddr, _ := multiaddr.NewMultiaddr(addr)

	return libp2p.New(libp2p.ListenAddrs(listenAddr))
}

// func connectToPeer(host host.Host, peerAddr string) error {
// 	addr, err := multiaddr.NewMultiaddr(peerAddr)
// 	if err != nil {
// 		return fmt.Errorf("invalid multiaddress: %v", err)
// 	}
// 	Extract peer ID from address
// 	peerInfo, err := peer.AddrInfoFromP2pAddr(addr)
// 	if err != nil {
// 		return fmt.Errorf("failed to parse peer info: %v", err)
// 	}
// 	Connect to the peer
// 	return host.Connect(context.Background(), *peerInfo)
// }

// Start a new blockchain with a genesis block

// Start RPC server
// func startRPCServer(port int) {
// 	fmt.Printf("📡 RPC server running on port %d\n", port)
// 	// Implement RPC server logic here
// }
