package main

import (
	"fmt"
	"log"
	"github.com/AkihiroSuda/go-netfilter-queue"
)

func main() {
	nfq, err := netfilter.NewNFQueue(0, 100, netfilter.NF_DEFAULT_PACKET_SIZE)
	if err != nil {
		log.Fatalln(err)
	}
	defer nfq.Close()

	packets := nfq.GetPackets()

	for true {
		select {
		case p := <-packets:
			fmt.Println(p.Packet)
			// getPacketInfo(p.Packet)
			p.SetVerdict(netfilter.NF_ACCEPT)
		}
	}
}