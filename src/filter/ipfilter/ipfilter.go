package ipfilter

import (
	"github.com/SaurusXI/protecc/src/config"
	"github.com/SaurusXI/protecc/src/gate"
	"errors"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"net"
)

func Allowed(c config.Config, packet gopacket.Packet) (bool, error) {
	
	if ipLayer := packet.Layer(layers.LayerTypeIPv4); ipLayer != nil {
		// fmt.Println("IP Layer found")
		ip, _ := ipLayer.(*layers.IPv4)
		blocked := false

		if c.IPsrc.Setting == gate.CLOSE {
			blocked = c.IPsrc.Blocked[key(ip.SrcIP)] || blocked
		}
		if c.IPdest.Setting == gate.CLOSE {
			blocked = c.IPdest.Blocked[key(ip.DstIP)] || blocked
		}

		return !blocked, nil
	}

	return false, errors.New("IP Layer not found")
}

func key(ip net.IP) string {
	return string(ip.To16())
}