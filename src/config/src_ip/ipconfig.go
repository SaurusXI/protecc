package src_ip

import (
	"github.com/SaurusXI/protecc/src/filter"
	"net"
)

type Config struct {
	setting filter.Gate
	allowed []net.IP
}
