package internal

import (
	"metrics/internal/conf"
	"sync/atomic"

	"github.com/jkstack/anet"
	"github.com/jkstack/jkframe/logging"
	"github.com/jkstack/jkframe/utils"
	"github.com/shirou/gopsutil/v3/host"
)

func getSensorsTemperatures(cfg conf.TempsConfigure, warnings *uint64) []anet.HMSensorTemperature {
	temps, err := host.SensorsTemperatures()
	if err != nil {
		logging.Warning("get sensors temperatures: %v", err)
		atomic.AddUint64(warnings, 1)
		return nil
	}
	var ret []anet.HMSensorTemperature
	for _, temp := range temps {
		ret = append(ret, anet.HMSensorTemperature{
			Name:        temp.SensorKey,
			Temperature: utils.Float64P2(temp.Temperature),
		})
	}
	return ret
}
