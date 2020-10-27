package tcpfilter

import (
	"github.com/SaurusXI/protecc/src/config"
	"github.com/SaurusXI/protecc/src/gate"
	"errors"
	"strconv"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// Implement IP filtration logic

func Allowed(c config.Config, packet gopacket.Packet) (bool, error) {
	
	if tcpLayer := packet.Layer(layers.LayerTypeTCP); tcpLayer != nil {
		// fmt.Println("IP Layer found")
		tcp, _ := tcpLayer.(*layers.TCP)
		blocked := false

		if c.Portsrc.Setting == gate.CLOSE {
			blocked = c.Portsrc.Blocked[key(tcp.SrcPort)] || blocked
		}
		if c.Portdest.Setting == gate.CLOSE {
			blocked = c.IPdest.Blocked[key(tcp.DstPort)] || blocked
		}
		if c.Checksum.Setting == gate.CLOSE {
			blocked = c.Checksum.Blocked[strconv.Itoa(int(tcp.Checksum))]
		}
		if c.Window.Setting == gate.CLOSE {
			blocked = c.Window.Blocked[strconv.Itoa(int(tcp.Window))]
		}

		return !blocked, nil
	}

	return false, errors.New("TCP Layer not found")
}

func key(port layers.TCPPort) string {
	return port.String()
}