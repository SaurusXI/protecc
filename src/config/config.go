package config

import (
	"github.com/SaurusXI/protecc/src/config/src_ip"
	"github.com/SaurusXI/protecc/src/config/dest_ip"
	"github.com/SaurusXI/protecc/src/config/src_port"
	"github.com/SaurusXI/protecc/src/config/dest_port"
	"github.com/SaurusXI/protecc/src/config/window_tcp"
	"github.com/SaurusXI/protecc/src/config/checksum_tcp"
)

type Config struct {
	IPsrc 		src_ip.Config
	IPdest 		dest_ip.Config
	Portsrc 	src_port.Config
	Portdest 	dest_port.Config
	Window 		window_tcp.Config
	Checksum 	checksum_tcp.Config
}