package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"
)

func main() {
	// Define command-line flags
	port := flag.Int("port", 4001, "Port to listen on")
	peerId := flag.String("peerId", "", "Multiaddress of the peer to connect to")
	rpcPort := flag.Int("rpcPort", 8080, "Port for RPC server")
	flag.Parse()

	// Start the P2P node
	host, err := startNode(*port)
	if err != nil {
		log.Fatalf("Failed to start the node: %v", err)
	}

	// Display node info
	fmt.Printf("✅ P2P Node started. Listening on: /ip4/127.0.0.1/tcp/%d\n", *port)
	fmt.Println("Node ID: ", host.ID())

	// Check if peerId is provided
	if *peerId != "" {
		fmt.Println("🔗 Connecting to peer:", *peerId)
		err := connectToPeer(host, *peerId)
		if err != nil {
			log.Fatalf("❌ Error connecting to peer: %v\n", err)
		} else {
			fmt.Println("✅ Successfully connected to peer:", *peerId)
		}
	} else {
		fmt.Println("🌱 No peerId provided. Starting a new blockchain...")
		startGenesisBlock()
	}

	// Start RPC server (optional)
	go startRPCServer(*rpcPort)

	// Keep the node running
	select {}
}

// Start a new libp2p node
func startNode(port int) (host.Host, error) {
	addr := fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port)
	listenAddr, _ := multiaddr.NewMultiaddr(addr)

	return libp2p.New(libp2p.ListenAddrs(listenAddr))
}

// Connect to a peer using peerId
func connectToPeer(host host.Host, peerAddr string) error {
	addr, err := multiaddr.NewMultiaddr(peerAddr)
	if err != nil {
		return fmt.Errorf("invalid multiaddress: %v", err)
	}

	// Extract peer ID from address
	peerInfo, err := peer.AddrInfoFromP2pAddr(addr)
	if err != nil {
		return fmt.Errorf("failed to parse peer info: %v", err)
	}

	// Connect to the peer
	return host.Connect(context.Background(), *peerInfo)
}

// Start a new blockchain with a genesis block
func startGenesisBlock() {
	fmt.Println("⛏️  Creating the genesis block...")
	// Add logic for genesis block initialization
	genesisBlock := "Genesis Block Created!"
	fmt.Println(genesisBlock)
}

// Start RPC server
func startRPCServer(port int) {
	fmt.Printf("📡 RPC server running on port %d\n", port)
	// Implement RPC server logic here
}
