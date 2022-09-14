package internal

import (
	"sync/atomic"

	"github.com/jkstack/anet"
	"github.com/jkstack/jkframe/logging"
	"github.com/jkstack/jkframe/utils"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/v3/disk"
)

func getUsage(warnings *uint64) *anet.HMDynamicUsage {
	var ret anet.HMDynamicUsage
	getCpuUsage(warnings, &ret)
	getMemoryUsage(warnings, &ret)
	getPartitionUsage(warnings, &ret)
	getInterfaceUsage(warnings, &ret)
	return &ret
}

func getCpuUsage(warnings *uint64, ret *anet.HMDynamicUsage) {
	usage, err := cpu.Percent(-1, false)
	if err != nil {
		logging.Warning("get cpu percent: %v", err)
		atomic.AddUint64(warnings, 1)
	}
	if len(usage) > 0 {
		ret.Cpu.Usage = utils.Float64P2(usage[0])
	}
}

func getMemoryUsage(warnings *uint64, ret *anet.HMDynamicUsage) {
	memStat, err := mem.VirtualMemory()
	if err != nil {
		logging.Warning("get memory usage: %v", err)
		atomic.AddUint64(warnings, 1)
	}
	if memStat != nil {
		ret.Memory.Used = memStat.Used
		ret.Memory.Free = memStat.Free
		ret.Memory.Available = memStat.Available
		ret.Memory.Usage = utils.Float64P2(memStat.UsedPercent)
		ret.Memory.SwapUsed = memStat.SwapCached
		ret.Memory.SwapFree = memStat.SwapFree
	}
}

func getPartitionUsage(warnings *uint64, ret *anet.HMDynamicUsage) {
	parts, err := disk.Partitions(false)
	if err != nil {
		logging.Warning("get partitions: %v", err)
		atomic.AddUint64(warnings, 1)
	}
	for _, part := range parts {
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			logging.Warning("get partition usage(%s): %v", part.Mountpoint, err)
			atomic.AddUint64(warnings, 1)
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
}

func getInterfaceUsage(warnings *uint64, ret *anet.HMDynamicUsage) {
	stats, err := net.IOCounters(true)
	if err != nil {
		logging.Warning("get interface stat: %v", err)
		atomic.AddUint64(warnings, 1)
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
}
