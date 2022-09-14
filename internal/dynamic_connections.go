package internal

import (
	"context"
	"fmt"
	"metrics/internal/conf"
	"metrics/internal/conn"
	"runtime"
	"sync/atomic"

	"github.com/jkstack/anet"
	"github.com/jkstack/jkframe/logging"
	"github.com/shirou/gopsutil/v3/net"
	"golang.org/x/time/rate"
)

func getConnectionList(cfg conf.ConnsConfigure, warnings *uint64, allow []string) []anet.HMDynamicConnection {
	allows := allowConns(cfg, allow)
	limit := rate.NewLimiter(rate.Inf, 1)
	if cfg.Limit > 0 {
		limit = rate.NewLimiter(rate.Limit(cfg.Limit), 1)
	}
	addr := func(addr net.Addr) string {
		if addr.Port > 0 {
			return fmt.Sprintf("%s:%d", addr.IP, addr.Port)
		}
		return addr.IP
	}
	var ret []anet.HMDynamicConnection
	for kind := range allows {
		var conns []net.ConnectionStat
		var err error
		if runtime.GOOS == "windows" {
			conns, err = net.Connections(kind)
		} else {
			conns, err = net.ConnectionsWithoutUids(kind)
		}
		if err != nil {
			logging.Warning("get connections of kind %s: %v", kind, err)
			atomic.AddUint64(warnings, 1)
			continue
		}
		for _, c := range conns {
			limit.Wait(context.Background())
			ret = append(ret, anet.HMDynamicConnection{
				Fd:     c.Fd,
				Pid:    c.Pid,
				Type:   conn.Type(warnings, c),
				Local:  addr(c.Laddr),
				Remote: addr(c.Raddr),
				Status: c.Status,
			})
		}
	}
	return ret
}

func allowConns(cfg conf.ConnsConfigure, allow []string) map[string]bool {
	allows := make(map[string]bool)
	parseAllow := func(allow []string) {
		for _, allow := range allow {
			switch allow {
			case "tcp", "tcp4", "tcp6":
				allows[allow] = true
			case "udp", "udp4", "udp6":
				allows[allow] = true
			case "unix":
				allows[allow] = true
			}
		}
		if allows["tcp4"] && allows["tcp6"] {
			allows["tcp"] = true
			delete(allows, "tcp4")
			delete(allows, "tcp6")
		}
		if allows["udp4"] && allows["udp6"] {
			allows["udp"] = true
			delete(allows, "udp4")
			delete(allows, "udp6")
		}
	}
	if len(allow) > 0 {
		parseAllow(allow)
	} else if len(cfg.Allow) > 0 {
		parseAllow(cfg.Allow)
	} else {
		allows["tcp"] = true
		allows["udp"] = true
		allows["unix"] = true
	}
	if allows["tcp"] && allows["udp"] && allows["unix"] {
		allows["all"] = true
		delete(allows, "tcp")
		delete(allows, "udp")
		delete(allows, "unix")
	}
	return allows
}
