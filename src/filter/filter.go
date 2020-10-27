package filter

import (
	"github.com/google/gopacket"
	"github.com/SaurusXI/protecc/src/filter/ipfilter"
	"github.com/SaurusXI/protecc/src/filter/tcpfilter"
	"github.com/SaurusXI/protecc/src/config"
)

type Filter struct {}

func (f Filter) Allowed(c config.Config, packet gopacket.Packet) bool {
	ipAllowed, _ := ipfilter.Allowed(c, packet) 
	tcpAllowed, _ := tcpfilter.Allowed(c, packet)
	return ipAllowed && tcpAllowed
}