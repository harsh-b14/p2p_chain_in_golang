package main

import (
	"log"
	"github.com/harsh-b14/p2p-chain/p2p"
	"github.com/harsh-b14/p2p-chain/rpc"
)

func main() {
	// Start P2P node on port 4001
	host, err := p2p.StartP2P(4001)
	if err != nil {
		log.Fatal(err)
	}

	// Start RPC server on port 8080
	go func() {
		rpc.StartRPC(8080)
	}()

	// Optionally connect to a peer (replace with actual peer address)
	peerAddr := "/ip4/127.0.0.1/tcp/4002/p2p/12D3KooWDraX45FfP9kmdggy1AuguYQpuxDa7TdURj1stg1N1cGW"
	err = p2p.ConnectToPeer(host, peerAddr)
	if err != nil {
		log.Println("Error connecting to peer:", err)
	}

	select {}
}
