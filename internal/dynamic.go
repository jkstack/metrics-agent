package internal

import (
	"context"
	"metrics/internal/conf"

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
	var ret anet.HMDynamicRep
	for _, req := range req.Req {
		switch req {
		case anet.HMReqUsage:
			ret.Usage = getUsage()
		case anet.HMReqProcess:
			ret.Process = getProcessList(cfg.Task.Process)
		case anet.HMReqConnections:
			ret.Connections = getConnectionList()
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

func getProcessList(cfg conf.ProcessConfigure) []anet.HMDynamicProcess {
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
		dy.ParentID, err = p.Ppid()
		if err != nil {
			logging.Warning("process => get parent id(%d): %v", pid, err)
		}
		dy.User, err = p.Username()
		if err != nil {
			logging.Warning("process => get username(%d): %v", pid, err)
		}
		usage, err := p.CPUPercent()
		if err != nil {
			logging.Warning("process => get cpu_usage(%d): %v", pid, err)
		}
		dy.CpuUsage = utils.Float64P2(usage)
		memInfo, err := p.MemoryInfo()
		if err != nil {
			logging.Warning("process => get memory_info(%d): %v", pid, err)
		}
		if memInfo != nil {
			dy.RssMemory = memInfo.RSS
			dy.VirtualMemory = memInfo.VMS
			dy.SwapMemory = memInfo.Swap
		}
		percent, err := p.MemoryPercent()
		if err != nil {
			logging.Warning("process => get memory_usage(%d): %v", pid, err)
		}
		dy.MemoryUsage = utils.Float64P2(percent)
		dy.Cmd, err = p.CmdlineSlice()
		if err != nil {
			logging.Warning("process => get cmd(%d): %v", pid, err)
		}
		conns, err := p.Connections()
		if err != nil {
			logging.Warning("process => get connections(%d): %v", pid, err)
		}
		for _, conn := range conns {
			if conn.Status == "LISTEN" {
				dy.Listen = append(dy.Listen, conn.Laddr.Port)
			}
		}
		dy.Connections = len(conns)
		ret = append(ret, dy)
	}
	return ret
}

func getConnectionList() []anet.HMDynamicConnection {
	return nil
}
