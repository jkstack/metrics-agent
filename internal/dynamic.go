package internal

import (
	"github.com/jkstack/anet"
	"github.com/jkstack/jkframe/logging"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

func getDynamic(req *anet.HMDynamicReq) *anet.HMDynamicRep {
	var ret anet.HMDynamicRep
	for _, req := range req.Req {
		switch req {
		case anet.HMReqUsage:
			ret.Usage = getUsage()
		case anet.HMReqProcess:
			ret.Process = getProcessList()
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
		ret.Cpu.Usage = usage[0]
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
		ret.Memory.Usage = memStat.UsedPercent
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
			Usage:      usage.UsedPercent,
			InodeUsed:  usage.InodesUsed,
			InodeFree:  usage.InodesFree,
			InodeUsage: usage.InodesUsedPercent,
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

func getProcessList() []anet.HMDynamicProcess {
	return nil
}

func getConnectionList() []anet.HMDynamicConnection {
	return nil
}
