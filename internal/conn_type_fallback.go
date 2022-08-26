//go:build !linux
// +build !linux

package internal

import (
	"syscall"

	"github.com/jkstack/jkframe/logging"
	"github.com/shirou/gopsutil/v3/net"
)

func connType(conn net.ConnectionStat) string {
	switch conn.Type {
	case syscall.SOCK_STREAM:
		if conn.Family == syscall.AF_INET {
			return "tcp4"
		} else if conn.Family == syscall.AF_INET6 {
			return "tcp6"
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
	return "UNKNOWN"
}
