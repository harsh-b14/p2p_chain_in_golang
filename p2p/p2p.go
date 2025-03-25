package p2p

import (
	"context"
	"fmt"
	"log"

	libp2p "github.com/libp2p/go-libp2p"
	p2phost "github.com/libp2p/go-libp2p/core/host"
	net "github.com/libp2p/go-libp2p/core/network"
	peer "github.com/libp2p/go-libp2p/core/peer"
	ping "github.com/libp2p/go-libp2p/p2p/protocol/ping"
	multiaddr "github.com/multiformats/go-multiaddr"
)

func StartP2P(port int) (p2phost.Host, error) {
	context.Background()

	host, err := libp2p.New(libp2p.ListenAddrStrings(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port)))
	if err != nil {
		return nil, err
	}
	fmt.Printf("P2P node started. Listening on: /ip4/127.0.0.1/tcp/%d\n", port)
	fmt.Println("Node ID: ", host.ID().String())

	pingService := &ping.PingService{Host: host}
	host.SetStreamHandler("/ping/1.0.0", func(s net.Stream) {
		pingService.PingHandler(s)
	})

	host.SetStreamHandler("/block/1.0.0", handleBlock)
	host.SetStreamHandler("/tx/1.0.0", handleTransaction)

	return host, nil
}

func ConnectToPeer(host p2phost.Host, peerAddr string) error {
	addr, err := multiaddr.NewMultiaddr(peerAddr)
	if err != nil {
		return err
	}
	info, err := peer.AddrInfoFromP2pAddr(addr)
	if err != nil {
		return err
	}
	if err := host.Connect(context.Background(), *info); err != nil {
		return err
	}
	fmt.Println("Connected to peer:", peerAddr)
	return nil
}

func handleBlock(s net.Stream) {
	defer s.Close()
	buf := make([]byte, 1024)
	n, err := s.Read(buf)
	if err != nil {
		log.Println("Error reading block:", err)
		return
	}
	fmt.Println("Received new block:", string(buf[:n]))
}

func handleTransaction(s net.Stream) {
	defer s.Close()
	buf := make([]byte, 1024)
	n, err := s.Read(buf)
	if err != nil {
		log.Println("Error reading transaction:", err)
		return
	}
	fmt.Println("Received new transaction:", string(buf[:n]))
}