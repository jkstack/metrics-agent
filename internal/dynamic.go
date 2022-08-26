package internal

import (
	"context"
	"fmt"
	"metrics/internal/conf"
	"sort"

	"github.com/jkstack/anet"
	"github.com/jkstack/jkframe/logging"
	"github.com/jkstack/jkframe/utils"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
	"github.com/shirou/gopsutil/v3/process"
	"golang.org/x/time/rate"
)

func getDynamic(req *anet.HMDynamicReq, cfg *conf.Configure) *anet.HMDynamicRep {
	top := req.Top
	allowConns := req.AllowConns
	var ret anet.HMDynamicRep
	for _, req := range req.Req {
		switch req {
		case anet.HMReqUsage:
			ret.Usage = getUsage()
		case anet.HMReqProcess:
			ret.Process = getProcessList(cfg.Task.Process, top)
		case anet.HMReqConnections:
			ret.Connections = getConnectionList(cfg.Task.Conns, allowConns)
		}
	}
	return &ret
}

func getUsage() *anet.HMDynamicUsage {
	var ret anet.HMDynamicUsage
	usage, err := cpu.Percent(-1, false)
	if err != nil {
		logging.Warning("get cpu percent: %v", err)
	}
	if len(usage) > 0 {
		ret.Cpu.Usage = utils.Float64P2(usage[0])
	}
	memStat, err := mem.VirtualMemory()
	if err != nil {
		logging.Warning("get memory usage: %v", err)
	}
	if memStat != nil {
		ret.Memory.Used = memStat.Used
		ret.Memory.Free = memStat.Free
		ret.Memory.Available = memStat.Available
		ret.Memory.Total = memStat.Total
		ret.Memory.Usage = utils.Float64P2(memStat.UsedPercent)
		ret.Memory.SwapUsed = memStat.SwapCached
		ret.Memory.SwapFree = memStat.SwapFree
		ret.Memory.SwapTotal = memStat.Total
	}
	parts, err := disk.Partitions(false)
	if err != nil {
		logging.Warning("get partitions: %v", err)
	}
	for _, part := range parts {
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			logging.Warning("get partition usage(%s): %v", part.Mountpoint, err)
			continue
		}
		ret.Partitions = append(ret.Partitions, anet.HMDynamicPartition{
			Name:       part.Mountpoint,
			Used:       usage.Used,
			Free:       usage.Free,
			Usage:      utils.Float64P2(usage.UsedPercent),
			InodeUsed:  usage.InodesUsed,
			InodeFree:  usage.InodesFree,
			InodeUsage: utils.Float64P2(usage.InodesUsedPercent),
		})
	}
	stats, err := net.IOCounters(true)
	if err != nil {
		logging.Warning("get interface stat: %v", err)
	}
	for _, stat := range stats {
		ret.Interface = append(ret.Interface, anet.HMDynamicInterface{
			Name:        stat.Name,
			BytesSent:   stat.BytesSent,
			BytesRecv:   stat.BytesRecv,
			PacketsSent: stat.PacketsSent,
			PacketsRecv: stat.PacketsRecv,
		})
	}
	return &ret
}

func getProcessList(cfg conf.ProcessConfigure, top int) []anet.HMDynamicProcess {
	pids, err := process.Pids()
	if err != nil {
		logging.Warning("get process list: %v", err)
	}
	limit := rate.NewLimiter(rate.Inf, 1)
	if cfg.Limit > 0 {
		limit = rate.NewLimiter(rate.Limit(cfg.Limit), 1)
	}
	var ret []anet.HMDynamicProcess
	for _, pid := range pids {
		limit.Wait(context.Background())
		p := process.Process{Pid: pid}
		var dy anet.HMDynamicProcess
		dy.ID = pid
		usage, err := p.CPUPercent()
		if err != nil {
			logging.Warning("process => get cpu_usage(%d): %v", pid, err)
		}
		dy.CpuUsage = utils.Float64P2(usage)
		ret = append(ret, dy)
	}
	if top > 0 && len(ret) > top {
		sort.Slice(ret, func(i, j int) bool {
			return ret[i].CpuUsage > ret[j].CpuUsage
		})
		ret = ret[:top]
	}
	for i, dy := range ret {
		limit.Wait(context.Background())
		p := process.Process{Pid: dy.ID}
		dy.ParentID, err = p.Ppid()
		if err != nil {
			logging.Warning("process => get parent id(%d): %v", dy.ID, err)
		}
		dy.User, err = p.Username()
		if err != nil {
			logging.Warning("process => get username(%d): %v", dy.ID, err)
		}
		memInfo, err := p.MemoryInfo()
		if err != nil {
			logging.Warning("process => get memory_info(%d): %v", dy.ID, err)
		}
		if memInfo != nil {
			dy.RssMemory = memInfo.RSS
			dy.VirtualMemory = memInfo.VMS
			dy.SwapMemory = memInfo.Swap
		}
		percent, err := p.MemoryPercent()
		if err != nil {
			logging.Warning("process => get memory_usage(%d): %v", dy.ID, err)
		}
		dy.MemoryUsage = utils.Float64P2(percent)
		dy.Cmd, err = p.CmdlineSlice()
		if err != nil {
			logging.Warning("process => get cmd(%d): %v", dy.ID, err)
		}
		conns, err := p.Connections()
		if err != nil {
			logging.Warning("process => get connections(%d): %v", dy.ID, err)
		}
		for _, conn := range conns {
			if conn.Status == "LISTEN" {
				dy.Listen = append(dy.Listen, conn.Laddr.Port)
			}
		}
		dy.Connections = len(conns)
		ret[i] = dy
	}
	return ret
}

func getConnectionList(cfg conf.ConnsConfigure, allow []string) []anet.HMDynamicConnection {
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
		allows["all"] = true
	}
	if allows["tcp"] && allows["udp"] && allows["unix"] {
		allows["all"] = true
		delete(allows, "tcp")
		delete(allows, "udp")
		delete(allows, "unix")
	}
	pids, err := process.Pids()
	if err != nil {
		logging.Warning("get process list in connection.list: %v", err)
		return nil
	}
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
	for _, pid := range pids {
		for kind := range allows {
			conns, err := net.ConnectionsPidWithoutUids(kind, pid)
			if err != nil {
				logging.Warning("get connections of kind %s: %v", kind, err)
				continue
			}
			for _, conn := range conns {
				limit.Wait(context.Background())
				ret = append(ret, anet.HMDynamicConnection{
					Fd:     conn.Fd,
					Pid:    pid,
					Type:   connType(conn),
					Local:  addr(conn.Laddr),
					Remote: addr(conn.Raddr),
					Status: conn.Status,
				})
			}
		}
	}
	return ret
}
