package internal

import (
	"time"

	"github.com/jkstack/anet"
	"github.com/jkstack/jkframe/logging"
	"github.com/shirou/gopsutil/v3/host"
)

func getStatic() *anet.HMStaticPayload {
	var ret anet.HMStaticPayload
	ret.Time = time.Now()
	info, err := host.Info()
	if err != nil {
		logging.Warning("get host.info: %v", err)
	}
	if info != nil {
		fillStaticHostInfo(&ret, info)
	}
	return &ret
}

func fillStaticHostInfo(ret *anet.HMStaticPayload, info *host.InfoStat) {
	ret.Host.Name = info.Hostname
	ret.Host.UpTime = time.Duration(info.Uptime) * time.Second
	ret.OS.Name = info.OS
	ret.OS.PlatformName = info.Platform
	ret.OS.PlatformVersion = info.PlatformVersion
	ret.OS.Install = time.Unix(0, 0) // TODO
	ret.Kernel.Version = info.KernelVersion
	ret.Kernel.Arch = info.KernelArch
}
