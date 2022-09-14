package internal

import (
	"metrics/internal/conf"
	"time"

	"github.com/jkstack/anet"
)

func getDynamic(req *anet.HMDynamicReq, cfg *conf.Configure, warnings *uint64) *anet.HMDynamicRep {
	top := req.Top
	allowConns := req.AllowConns
	var ret anet.HMDynamicRep
	ret.Begin = time.Now()
	for _, req := range req.Req {
		switch req {
		case anet.HMReqUsage:
			ret.Usage = getUsage(warnings)
		case anet.HMReqProcess:
			ret.Process = getProcessList(cfg.Task.Process, warnings, top)
		case anet.HMReqConnections:
			ret.Connections = getConnectionList(cfg.Task.Conns, warnings, allowConns)
		case anet.HMReqSensorsTemperatures:
			ret.SensorsTemperatures = getSensorsTemperatures(cfg.Task.Temps, warnings)
		}
	}
	ret.End = time.Now()
	return &ret
}
