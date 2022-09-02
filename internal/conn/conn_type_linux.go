//go:build linux
// +build linux

package conn

import (
	"sync/atomic"
	"syscall"

	"github.com/jkstack/jkframe/logging"
	"github.com/shirou/gopsutil/v3/net"
)

func Type(warnings *uint64, conn net.ConnectionStat) string {
	switch conn.Type {
	case syscall.SOCK_STREAM:
		if conn.Family == syscall.AF_INET {
			return "tcp4"
		} else if conn.Family == syscall.AF_INET6 {
			return "tcp6"
		} else if conn.Family == syscall.AF_FILE {
			return "file"
		}
	case syscall.SOCK_DGRAM:
		if conn.Family == syscall.AF_INET {
			return "udp4"
		} else if conn.Family == syscall.AF_INET6 {
			return "udp6"
		}
	default:
		if conn.Family == syscall.AF_UNIX {
			return "unix"
		}
	}
	logging.Warning("connection unknown: type=%d, family=%d", conn.Type, conn.Family)
	atomic.AddUint64(warnings, 1)
	return "UNKNOWN"
}
