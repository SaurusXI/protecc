package main

import (
	"log"
	"strconv"

	"github.com/AkihiroSuda/go-netfilter-queue"
	"github.com/SaurusXI/protecc/src/tui"
	"github.com/google/gopacket/layers"
)

func main() {
	nfq, err := netfilter.NewNFQueue(0, 100, netfilter.NF_DEFAULT_PACKET_SIZE)
	if err != nil {
		log.Fatalln(err)
	}
	defer nfq.Close()
	packetChannel := make(chan []string, 100)
	packets := nfq.GetPackets()
	go tui.Start(packetChannel)
	for true {
		select {
		case p := <-packets:
			ipLayer := p.Packet.Layer(layers.LayerTypeIPv4)
			tcpLayer := p.Packet.Layer(layers.LayerTypeTCP)
			tcp, _ := tcpLayer.(*layers.TCP)
			ip, _ := ipLayer.(*layers.IPv4)
			if tcp == nil || ip == nil {
				p.SetVerdict(netfilter.NF_ACCEPT)
				break
			}

			cols := []string{ip.SrcIP.String(), tcp.SrcPort.String(), ip.DstIP.String(), tcp.DstPort.String(), strconv.Itoa(int(tcp.Window)), strconv.Itoa(int(tcp.Checksum))}
			packetChannel <- cols
			// getPacketInfo(p.`Packet)
			p.SetVerdict(netfilter.NF_ACCEPT)
		}
	}
}
