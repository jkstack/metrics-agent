package internal

import (
	"metrics/internal/user"
	"net"
	"runtime"
	"strconv"
	"strings"
	"sync/atomic"
	"time"

	"github.com/jackpal/gateway"
	"github.com/jaypipes/ghw"
	"github.com/jkstack/anet"
	"github.com/jkstack/jkframe/logging"
	"github.com/jkstack/jkframe/utils"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
)

func getStatic(warnings *atomic.Uint64) *anet.HMStaticPayload {
	var ret anet.HMStaticPayload
	ret.Time = time.Now()
	fillStaticHostInfo(warnings, &ret)
	fillStaticCpuInfo(warnings, &ret)
	fillStaticMemoryInfo(warnings, &ret)
	fillStaticDiskInfo(warnings, &ret)
	fillStaticNetworkInfo(warnings, &ret)
	fillStaticUserInfo(warnings, &ret)
	return &ret
}

func fillStaticHostInfo(warnings *atomic.Uint64, ret *anet.HMStaticPayload) {
	info, err := host.Info()
	if err != nil {
		logging.Warning("get host.info: %v", err)
		warnings.Add(1)
		return
	}
	ret.Host.Name = info.Hostname
	ret.Host.UpTime = time.Duration(info.Uptime) * time.Second
	ret.OS.Name = info.OS
	ret.OS.PlatformName = info.Platform
	ret.OS.PlatformVersion = info.PlatformVersion
	it, err := getInstallTime()
	if err != nil {
		logging.Warning("get install time: %v", err)
		warnings.Add(1)
	}
	ret.OS.Install = it
	ret.OS.Startup = time.Now().Add(-ret.Host.UpTime)
	ret.Kernel.Version = info.KernelVersion
	ret.Kernel.Arch = info.KernelArch
}

func fillStaticCpuInfo(warnings *atomic.Uint64, ret *anet.HMStaticPayload) {
	var err error
	ret.CPU.Physical, err = cpu.Counts(false)
	if err != nil {
		logging.Warning("get physical cpu count: %v", err)
		warnings.Add(1)
	}
	ret.CPU.Logical, err = cpu.Counts(true)
	if err != nil {
		logging.Warning("get logical cpu count: %v", err)
		warnings.Add(1)
	}
	cores, err := cpu.Info()
	if err != nil {
		logging.Warning("get cpu.info: %v", err)
		warnings.Add(1)
		return
	}
	for _, core := range cores {
		id, _ := strconv.ParseUint(core.CoreID, 10, 32)
		physical, _ := strconv.ParseUint(core.PhysicalID, 10, 32)
		ret.CPU.Cores = append(ret.CPU.Cores, anet.HMCore{
			Processor: core.CPU,
			Model:     core.ModelName,
			Core:      int32(id),
			Cores:     core.Cores,
			Physical:  int32(physical),
			Mhz:       utils.Float64P2(core.Mhz),
		})
	}
}

func fillStaticMemoryInfo(warnings *atomic.Uint64, ret *anet.HMStaticPayload) {
	vm, err := mem.VirtualMemory()
	if err != nil {
		logging.Warning("get memory.info: %v", err)
		warnings.Add(1)
	}
	if vm != nil {
		ret.Memory.Physical = vm.Total
	}
	swap, err := mem.SwapMemory()
	if err != nil {
		logging.Warning("get swap.info: %v", err)
		warnings.Add(1)
	}
	if swap != nil {
		ret.Memory.Swap = swap.Total
	}
}

func fillStaticDiskInfo(warnings *atomic.Uint64, ret *anet.HMStaticPayload) {
	block, err := ghw.Block()
	if err != nil {
		logging.Warning("get blocks: %v", err)
		warnings.Add(1)
	}
	for _, disk := range block.Disks {
		switch disk.DriveType {
		case ghw.DRIVE_TYPE_HDD, ghw.DRIVE_TYPE_SSD:
		default:
			continue
		}
		var parts []string
		for _, part := range disk.Partitions {
			if runtime.GOOS == "linux" {
				part.Name = "/dev/" + part.Name
			}
			parts = append(parts, part.Name)
		}
		ret.Disks = append(ret.Disks, anet.HMDisk{
			Model:      disk.Model,
			Total:      disk.SizeBytes,
			Type:       disk.DriveType.String(),
			Partitions: parts,
		})
	}
	parts, err := disk.Partitions(false)
	if err != nil {
		logging.Warning("get partitions: %v", err)
		warnings.Add(1)
	}
	for _, part := range parts {
		usage, err := disk.Usage(part.Mountpoint)
		if err != nil {
			logging.Warning("get partition usage(%s): %v", part.Mountpoint, err)
			warnings.Add(1)
		}
		info := anet.HMPartition{
			Name:   part.Mountpoint,
			FSType: part.Fstype,
			Opts:   part.Opts,
		}
		if usage != nil {
			info.Total = usage.Total
			info.INodes = usage.InodesTotal
		}
		ret.Partitions = append(ret.Partitions, info)
	}
}

func fillStaticNetworkInfo(warnings *atomic.Uint64, ret *anet.HMStaticPayload) {
	gw, err := gateway.DiscoverGateway()
	if err != nil {
		logging.Warning("get gateway: %v", err)
		warnings.Add(1)
	}
	ret.GateWay = gw.String()
	intfs, err := net.Interfaces()
	if err != nil {
		logging.Warning("get interfaces: %v", err)
		warnings.Add(1)
	}
	for _, intf := range intfs {
		addrs, _ := intf.Addrs()
		ips := make([]string, len(addrs))
		for i, addr := range addrs {
			ips[i] = addr.String()
		}
		ret.Interface = append(ret.Interface, anet.HMInterface{
			Index:   intf.Index,
			Name:    intf.Name,
			Mtu:     intf.MTU,
			Flags:   strings.Split(intf.Flags.String(), "|"),
			Mac:     intf.HardwareAddr.String(),
			Address: ips,
		})
	}
}

func fillStaticUserInfo(warnings *atomic.Uint64, ret *anet.HMStaticPayload) {
	users, err := user.List()
	if err != nil {
		logging.Warning("get user list: %v", err)
		warnings.Add(1)
	}
	for _, user := range users {
		ret.User = append(ret.User, anet.HMUser{
			Name: user.Name,
			ID:   user.ID,
			GID:  user.GID,
		})
	}
}
