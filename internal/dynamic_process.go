package internal

import (
	"context"
	"metrics/internal/conf"
	"sort"
	"sync/atomic"

	"github.com/jkstack/anet"
	"github.com/jkstack/jkframe/logging"
	"github.com/jkstack/jkframe/utils"
	"github.com/shirou/gopsutil/v3/process"
	"golang.org/x/time/rate"
)

func getProcessList(cfg conf.ProcessConfigure, warnings *uint64, top int) []anet.HMDynamicProcess {
	limit := rate.NewLimiter(rate.Inf, 1)
	if cfg.Limit > 0 {
		limit = rate.NewLimiter(rate.Limit(cfg.Limit), 1)
	}
	list := processList(cfg, limit, warnings, top)
	var ret []anet.HMDynamicProcess
	for _, dy := range list {
		limit.Wait(context.Background())
		p := process.Process{Pid: dy.ID}
		var err error
		dy.ParentID, err = p.Ppid()
		if err != nil {
			logging.Warning("process => get parent id(%d): %v", dy.ID, err)
			atomic.AddUint64(warnings, 1)
			continue
		}
		dy.User, err = p.Username()
		if err != nil {
			logging.Warning("process => get username(%d): %v", dy.ID, err)
			atomic.AddUint64(warnings, 1)
			continue
		}
		memInfo, err := p.MemoryInfo()
		if err != nil {
			logging.Warning("process => get memory_info(%d): %v", dy.ID, err)
			atomic.AddUint64(warnings, 1)
			continue
		}
		if memInfo != nil {
			dy.RssMemory = memInfo.RSS
			dy.VirtualMemory = memInfo.VMS
			dy.SwapMemory = memInfo.Swap
		}
		percent, err := p.MemoryPercent()
		if err != nil {
			logging.Warning("process => get memory_usage(%d): %v", dy.ID, err)
			atomic.AddUint64(warnings, 1)
			continue
		}
		dy.MemoryUsage = utils.Float64P2(percent)
		dy.Cmd, err = p.CmdlineSlice()
		if err != nil {
			logging.Warning("process => get cmd(%d): %v", dy.ID, err)
			atomic.AddUint64(warnings, 1)
			continue
		}
		conns, err := p.Connections()
		if err != nil {
			logging.Warning("process => get connections(%d): %v", dy.ID, err)
			atomic.AddUint64(warnings, 1)
			continue
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

func processList(cfg conf.ProcessConfigure, limit *rate.Limiter, warnings *uint64, top int) []anet.HMDynamicProcess {
	pids, err := process.Pids()
	if err != nil {
		logging.Warning("get process list: %v", err)
		atomic.AddUint64(warnings, 1)
	}

	var list []anet.HMDynamicProcess
	for _, pid := range pids {
		if pid == 0 {
			continue
		}
		limit.Wait(context.Background())
		p := process.Process{Pid: pid}
		var dy anet.HMDynamicProcess
		dy.ID = pid
		usage, err := p.CPUPercent()
		if err != nil {
			logging.Warning("process => get cpu_usage(%d): %v", pid, err)
			atomic.AddUint64(warnings, 1)
			continue
		}
		dy.CpuUsage = utils.Float64P2(usage)
		list = append(list, dy)
	}
	if top > 0 && len(list) > top {
		sort.Slice(list, func(i, j int) bool {
			return list[i].CpuUsage > list[j].CpuUsage
		})
		return list[:top]
	}
	return list
}
